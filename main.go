package main

import (
	"github.com/terraform-modules-krish/cloud-nuke/commands"
	"github.com/terraform-modules-krish/go-commons/entrypoint"
)

// VERSION - Set at build time
var VERSION string
var MixPanelClientId string

func main() {
	app := commands.CreateCli(VERSION, MixPanelClientId)
	entrypoint.RunApp(app)
}
