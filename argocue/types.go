package argocue

type Command string

type CommandsType struct {
	Generate Command
}

var Commands CommandsType = CommandsType{
	Generate: "generate",
}
