package argocue

type Command string

type CommandsType struct {
	Generate  Command
	GetParams Command
}

var Commands CommandsType = CommandsType{
	Generate:  "generate",
	GetParams: "get_params",
}

type ApplicationParams struct {
	Name           string `json:"name"`
	Title          string `json:"title"`
	CollectionType string `json:"collectionType"`
	Map            map[string]interface{}
}
