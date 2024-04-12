package app

import "kang-blogging/internal/blogging/app/command"

type Application struct {
	Command Command
}

type Command struct {
	// This is a sample command
	DoSomething command.DoSomethingHandler
}
