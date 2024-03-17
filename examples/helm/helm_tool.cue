package helm

import (
	"tool/file"
    "encoding/yaml"
    "github.com/darklab8/argocue/examples/helm/templates"

)

command: build: {
	task1: mkdir: file.Create & {
		filename: "values.yaml"
		contents: yaml.Marshal(values)
	}
	task2: mkdir: file.Create & {
		filename: "Chart.yaml"
		contents: yaml.Marshal(chart)
	}
    task3: mkdir: file.Create & {
		filename: "templates/build.yaml"
		contents: yaml.MarshalStream(templates.files)
	}
}