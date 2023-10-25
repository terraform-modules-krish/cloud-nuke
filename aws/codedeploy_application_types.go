package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/terraform-modules-krish/go-commons/errors"
)

// CodeDeployApplications - represents all codedeploy applications
type CodeDeployApplications struct {
	AppNames []string
}

// ResourceName - the simple name of the aws resource
func (c CodeDeployApplications) ResourceName() string {
	return "codedeploy-application"
}

// ResourceIdentifiers - The instance ids of the code deploy applications
func (c CodeDeployApplications) ResourceIdentifiers() []string {
	return c.AppNames
}

func (c CodeDeployApplications) MaxBatchSize() int {
	// Tentative batch size to ensure AWS doesn't throttle.
	return 100
}

// Nuke - nuke 'em all!!!
func (c CodeDeployApplications) Nuke(session *session.Session, identifiers []string) error {
	if err := nukeAllCodeDeployApplications(session, identifiers); err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}
