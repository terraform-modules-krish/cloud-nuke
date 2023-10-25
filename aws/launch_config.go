package aws

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/terraform-modules-krish/cloud-nuke/logging"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/errors"
)

// Returns a formatted string of Launch config Names
func getAllLaunchConfigurations(session *session.Session, region string, excludeAfter time.Time) ([]*string, error) {
	svc := autoscaling.New(session)
	result, err := svc.DescribeLaunchConfigurations(&autoscaling.DescribeLaunchConfigurationsInput{})

	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	var configNames []*string
	for _, config := range result.LaunchConfigurations {
		if excludeAfter.After(*config.CreatedTime) {
			configNames = append(configNames, config.LaunchConfigurationName)
		}
	}

	return configNames, nil
}

// Deletes all Launch configurations
func nukeAllLaunchConfigurations(session *session.Session, configNames []*string) error {
	svc := autoscaling.New(session)

	if len(configNames) == 0 {
		logging.Logger.Infof("No Launch Configurations to nuke in region %s", *session.Config.Region)
		return nil
	}

	logging.Logger.Infof("Deleting all Launch Configurations in region %s", *session.Config.Region)
	var deletedConfigNames []*string

	for _, configName := range configNames {
		params := &autoscaling.DeleteLaunchConfigurationInput{
			LaunchConfigurationName: configName,
		}

		_, err := svc.DeleteLaunchConfiguration(params)
		if err != nil {
			logging.Logger.Errorf("[Failed] %s", err)
		} else {
			deletedConfigNames = append(deletedConfigNames, configName)
			logging.Logger.Infof("Deleted Launch configuration: %s", *configName)
		}
	}

	logging.Logger.Infof("[OK] %d Launch Configuration(s) deleted in %s", len(deletedConfigNames), *session.Config.Region)
	return nil
}
