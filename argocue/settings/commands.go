package settings

type Command string

type CommandsType struct {
	Generate      Command
	GetParameters Command
}

var Commands CommandsType = CommandsType{
	Generate:      "generate",
	GetParameters: "get_parameters",
}
