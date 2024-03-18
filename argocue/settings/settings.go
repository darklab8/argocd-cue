package settings

type ApplicationParams struct {
	Name           string                 `json:"name"`
	Title          string                 `json:"title"`
	CollectionType string                 `json:"collectionType"`
	Map            map[string]interface{} `json:"map"`
}

type ParameterGroup struct {
	Name  string   `json:"name"`
	Array []string `json:"array"`
}
