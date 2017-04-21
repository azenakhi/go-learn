package extension

import (
	"github.com/azenakhi/go-learn/plugin/plug"
)

type DataBasePlugin struct{}

func (DataBasePlugin) Message() string {
	return "DataBasePlugin"
}

func init() {
	plug.Registry = append(plug.Registry, &DataBasePlugin{})
}
