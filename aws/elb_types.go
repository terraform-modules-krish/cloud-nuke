package aws

import (
	awsgo "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/errors"
)

// LoadBalancers - represents all load balancers
type LoadBalancers struct {
	Names []string
}

// ResourceName - the simple name of the aws resource
func (balancer LoadBalancers) ResourceName() string {
	return "elb"
}

// ResourceIdentifiers - The names of the load balancers
func (balancer LoadBalancers) ResourceIdentifiers() []string {
	return balancer.Names
}

// Nuke - nuke 'em all!!!
func (balancer LoadBalancers) Nuke(session *session.Session) error {
	if err := nukeAllElbInstances(session, awsgo.StringSlice(balancer.Names)); err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

type ElbDeleteError struct{}

func (e ElbDeleteError) Error() string {
	return "ELB was not deleted"
}
