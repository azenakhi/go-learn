package ext

import (
	"github.com/azenakhi/go-learn/pluggo/plugin"
)

type FilePlugin struct{}

func (FilePlugin) Message() string {
	return "FilePlugin"
}

func init() {
	plugin.Registry = append(plugin.Registry, &FilePlugin{})

}
