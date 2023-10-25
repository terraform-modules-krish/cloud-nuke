package main

import (
	"github.com/gruntwork-io/aws-nuke/commands"
	gruntwork-cli "github.com/terraform-modules-krish/go-commons/entrypoint"
)

// VERSION - Set at build time
var VERSION string

func main() {
	app := commands.CreateCli(VERSION)
	entrypoint.RunApp(app)
}
