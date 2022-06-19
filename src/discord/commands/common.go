package commands

import "github.com/vcokltfre/reviver/src/discord/impl"

type CommandError struct {
	Command string
	Error   error
}

const Ephemeral = 1 << 6

var (
	Errors chan CommandError
	Topics impl.TopicManager
)

func sendError(command string, err error) {
	Errors <- CommandError{command, err}
}
