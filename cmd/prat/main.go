package main

import (
	"github.com/webolton/prat/internal/commands"
)

func main() {
	// bootstrap.Execute()
	commands.Execute("this would be the config passed down to version")
}
