package aws

import (
	awsgo "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/terraform-modules-krish/go-commons/errors"
)

// LaunchConfigs - represents all launch configurations
type LaunchConfigs struct {
	Client                   autoscalingiface.AutoScalingAPI
	Region                   string
	LaunchConfigurationNames []string
}

// ResourceName - the simple name of the aws resource
func (lc LaunchConfigs) ResourceName() string {
	return "lc"
}

func (lc LaunchConfigs) MaxBatchSize() int {
	// Tentative batch size to ensure AWS doesn't throttle
	return 49
}

// ResourceIdentifiers - The names of the launch configurations
func (lc LaunchConfigs) ResourceIdentifiers() []string {
	return lc.LaunchConfigurationNames
}

// Nuke - nuke 'em all!!!
func (lc LaunchConfigs) Nuke(session *session.Session, identifiers []string) error {
	if err := lc.nukeAll(awsgo.StringSlice(identifiers)); err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}
