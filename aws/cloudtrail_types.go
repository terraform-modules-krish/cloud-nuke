package aws

import (
	awsgo "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
	"github.com/terraform-modules-krish/go-commons/errors"
)

// CloudWatchLogGroup - represents all ec2 instances
type CloudtrailTrail struct {
	Client cloudtrailiface.CloudTrailAPI
	Region string
	Arns   []string
}

// ResourceName - the simple name of the aws resource
func (ct CloudtrailTrail) ResourceName() string {
	return "cloudtrail"
}

// ResourceIdentifiers - The instance ids of the ec2 instances
func (ct CloudtrailTrail) ResourceIdentifiers() []string {
	return ct.Arns
}

func (ct CloudtrailTrail) MaxBatchSize() int {
	return 50
}

// Nuke - nuke 'em all!!!
func (ct CloudtrailTrail) Nuke(session *session.Session, identifiers []string) error {
	if err := ct.nukeAll(awsgo.StringSlice(identifiers)); err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}
