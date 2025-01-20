package main

import (
	"github.com/webolton/prat/internal/bootstrap"
	"github.com/webolton/prat/internal/commands"
)

var config bootstrap.Config

func init() {
	config = bootstrap.Execute()
}

func main() {
	commands.Execute(config)
}
