package commands

import (
	"strings"

	gruntwork-cli "github.com/terraform-modules-krish/go-commons/collections"

	"github.com/fatih/color"
	"github.com/gruntwork-io/aws-nuke/aws"
	"github.com/gruntwork-io/aws-nuke/logging"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/errors"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/shell"
	"github.com/urfave/cli"
)

// CreateCli - Create the CLI app with all commands, flags, and usage text configured.
func CreateCli(version string) *cli.App {
	app := cli.NewApp()

	app.Name = "aws-nuke"
	app.HelpName = app.Name
	app.Author = "Gruntwork <www.gruntwork.io>"
	app.Version = version
	app.Usage = "A CLI tool to cleanup AWS resources (ASG, ELB, ELBv2, EBS, EC2). THIS TOOL WILL COMPLETELY REMOVE ALL RESOURCES AND ITS EFFECTS ARE IRREVERSIBLE!!!"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "exclude-region",
			Usage: "regions to exclude",
		},
	}
	app.Action = errors.WithPanicHandling(awsNuke)

	return app
}

// Nuke it all!!!
func awsNuke(c *cli.Context) error {
	regions := aws.GetAllRegions()
	excludedRegions := c.StringSlice("exclude-region")

	for _, excludedRegion := range excludedRegions {
		if !collections.ListContainsElement(regions, excludedRegion) {
			return InvalidFlagError{
				Name: "exclude-regions",
				Value: excludedRegion,
			}
		}
	}

	logging.Logger.Infoln("Retrieving all active AWS resources")
	account, err := aws.GetAllResources(regions, excludedRegions)

	if err != nil {
		return errors.WithStackTrace(err)
	}

	if len(account.Resources) == 0 {
		logging.Logger.Infoln("Nothing to nuke, you're all good!")
		return nil
	}

	logging.Logger.Infoln("The following AWS resources are going to be nuked: ")

	for region, resourcesInRegion := range account.Resources {
		for _, resources := range resourcesInRegion.Resources {
			for _, identifier := range resources.ResourceIdentifiers() {
				logging.Logger.Infof("* %s-%s-%s\n", resources.ResourceName(), identifier, region)
			}
		}
	}

	color := color.New(color.FgHiRed, color.Bold)
	color.Println("\nTHE NEXT STEPS ARE DESTRUCTIVE AND COMPLETELY IRREVERSIBLE, PROCEED WITH CAUTION!!!")

	prompt := "\nAre you sure you want to nuke all listed resources? Enter 'nuke' to confirm: "
	shellOptions := shell.ShellOptions{Logger: logging.Logger}
	input, err := shell.PromptUserForInput(prompt, &shellOptions)

	if err != nil {
		return errors.WithStackTrace(err)
	}

	if strings.ToLower(input) == "nuke" {
		aws.NukeAllResources(account, regions)
	}

	return nil
}
