package aws

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/terraform-modules-krish/cloud-nuke/config"
	"github.com/terraform-modules-krish/cloud-nuke/logging"
	"github.com/terraform-modules-krish/cloud-nuke/telemetry"
	commonTelemetry "github.com/terraform-modules-krish/go-commons/telemetry"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/errors"
	"github.com/hashicorp/go-multierror"
)

// getAllEc2KeyPairs extracts the list of existing ec2 key pairs.
func (k EC2KeyPairs) getAll(configObj config.Config) ([]*string, error) {
	result, err := k.Client.DescribeKeyPairs(&ec2.DescribeKeyPairsInput{})
	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	var ids []*string
	for _, keyPair := range result.KeyPairs {
		if configObj.EC2KeyPairs.ShouldInclude(config.ResourceValue{
			Name: keyPair.KeyName,
			Time: keyPair.CreateTime,
		}) {
			ids = append(ids, keyPair.KeyPairId)
		}
	}

	return ids, nil
}

// deleteKeyPair is a helper method that deletes the given ec2 key pair.
func (k EC2KeyPairs) deleteKeyPair(keyPairId *string) error {
	params := &ec2.DeleteKeyPairInput{
		KeyPairId: keyPairId,
	}

	_, err := k.Client.DeleteKeyPair(params)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

// nukeAllEc2KeyPairs attempts to delete given ec2 key pair IDs.
func (k EC2KeyPairs) nukeAll(keypairIds []*string) error {
	if len(keypairIds) == 0 {
		logging.Logger.Infof("No EC2 key pairs to nuke in region %s", k.Region)
		return nil
	}

	logging.Logger.Infof("Terminating all EC2 key pairs in region %s", k.Region)

	deletedKeyPairs := 0
	var multiErr *multierror.Error
	for _, keypair := range keypairIds {
		if err := k.deleteKeyPair(keypair); err != nil {
			telemetry.TrackEvent(commonTelemetry.EventContext{
				EventName: "Error Nuking EC2 Key Pair",
			}, map[string]interface{}{
				"region": k.Region,
			})
			logging.Logger.Errorf("[Failed] %s", err)
			multiErr = multierror.Append(multiErr, err)
		} else {
			deletedKeyPairs++
			logging.Logger.Infof("Deleted EC2 KeyPair: %s", *keypair)
		}
	}

	logging.Logger.Infof("[OK] %d EC2 KeyPair(s) terminated", deletedKeyPairs)
	return multiErr.ErrorOrNil()
}
