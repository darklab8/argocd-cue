package argocue

type Command string

type CommandsType struct {
	Generate      Command
	GetParameters Command
}

var Commands CommandsType = CommandsType{
	Generate:      "generate",
	GetParameters: "get_parameters",
}

type ApplicationParams struct {
	Name           string `json:"name"`
	Title          string `json:"title"`
	CollectionType string `json:"collectionType"`
	Map            map[string]interface{}
}
