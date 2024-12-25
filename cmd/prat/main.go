package main

import (
	"github.com/webolton/prat/internal/bootstrap"
	"github.com/webolton/prat/internal/commands"
)

func main() {
	bootstrap.Execute()
	commands.Execute()
}
