package extension

import "github.com/azenakhi/go-learn/plugin/plug"

type FilePlugin struct{}

func (FilePlugin) Message() string {
	return "FilePlugin"
}

func init() {
	plug.Registry = append(plug.Registry, &FilePlugin{})
}
