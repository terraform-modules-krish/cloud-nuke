package resources

import (
	"context"
	"strings"
	"time"

	awsgo "github.com/aws/aws-sdk-go/aws"
	commonTelemetry "github.com/terraform-modules-krish/go-commons/telemetry"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/terraform-modules-krish/cloud-nuke/config"
	"github.com/terraform-modules-krish/cloud-nuke/logging"
	"github.com/terraform-modules-krish/cloud-nuke/report"
	"github.com/terraform-modules-krish/cloud-nuke/telemetry"
	"github.com/terraform-modules-krish/go-commons/errors"
)

// Returns a formatted string of AMI Image ids
func (ami *AMIs) getAll(c context.Context, configObj config.Config) ([]*string, error) {
	params := &ec2.DescribeImagesInput{
		Owners: []*string{awsgo.String("self")},
	}

	output, err := ami.Client.DescribeImages(params)
	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	var imageIds []*string
	for _, image := range output.Images {
		layout := "2006-01-02T15:04:05.000Z"
		createdTime, err := time.Parse(layout, *image.CreationDate)
		if err != nil {
			return nil, err
		}

		// Check if the image has a tag that indicates AWS management
		isAWSManaged := false
		for _, tag := range image.Tags {
			if *tag.Key == "aws-managed" && *tag.Value == "true" {
				isAWSManaged = true
				break
			}
		}

		// Skip AWS managed images and images created by AWS Backup
		if isAWSManaged || strings.HasPrefix(*image.Name, "AwsBackup") {
			continue
		}

		if configObj.AMI.ShouldInclude(config.ResourceValue{
			Name: image.Name,
			Time: &createdTime,
		}) {
			imageIds = append(imageIds, image.ImageId)
		}
	}

	return imageIds, nil
}

// Deletes all AMI
func (ami *AMIs) nukeAll(imageIds []*string) error {
	if len(imageIds) == 0 {
		logging.Debugf("No AMI to nuke in region %s", ami.Region)
		return nil
	}

	logging.Debugf("Deleting all AMI in region %s", ami.Region)

	deletedCount := 0
	for _, imageID := range imageIds {
		params := &ec2.DeregisterImageInput{
			ImageId: imageID,
		}

		_, err := ami.Client.DeregisterImage(params)

		// Record status of this resource
		e := report.Entry{
			Identifier:   aws.StringValue(imageID),
			ResourceType: "Amazon Machine Image (AMI)",
			Error:        err,
		}
		report.Record(e)

		if err != nil {
			logging.Debugf("[Failed] %s", err)
			telemetry.TrackEvent(commonTelemetry.EventContext{
				EventName: "Error Nuking AMI",
			}, map[string]interface{}{
				"region": ami.Region,
			})
		} else {
			deletedCount++
			logging.Debugf("Deleted AMI: %s", *imageID)
		}
	}

	logging.Debugf("[OK] %d AMI(s) terminated in %s", deletedCount, ami.Region)
	return nil
}
