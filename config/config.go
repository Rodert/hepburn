package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/sirupsen/logrus"

	// "github.com/hashicorp/vault/api"
	"github.com/hashicorp/consul/api"
)

var Configure Configuration

type Configuration struct {
	// Log             Log             `hcl:"log,block"`
	Web   Web   `hcl:"web,block"`
	MySQL MySQL `hcl:"mysql,block"`
}

type Web struct {
	Address string `hcl:"address"`
}

type MySQL struct {
	DNS string `hcl:"dns"`
}

func NewConfigure(configPaths ...string) error {
	var data []byte
	var config Configuration
	if len(configPaths) > 0 {
		configPath := configPaths[0]
		src, err := ioutil.ReadFile(configPath)
		if err != nil {
			return fmt.Errorf("new configure err %w", err)
		}
		data = src
	} else {
		client, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			panic(err)
		}
		kv := client.KV()
		pair, _, err := kv.Get("s_graph/config", nil)
		if err != nil {
			panic(err)
		}
		_ = pair
	}

	// file, diags := hclsyntax.ParseConfig(data, "config", hcl.Pos{Line: 1, Column: 1})
	// if diags.HasErrors() {
	// 	return diags
	// }
	// diags = gohcl.DecodeBody(file.Body, nil, &config)
	// if diags.HasErrors() {
	// 	log.Fatal(diags)
	// return diags
	// }

	if err := readConfig(data, &config); err != nil {
		logrus.Error(err)
	}

	Configure = config
	return nil
}

func readConfig[T any](data []byte, t *T) error {
	if len(data) == 0 {
		panic("no configuration was read")
	}
	file, diags := hclsyntax.ParseConfig(data, "config", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return diags
	}
	diags = gohcl.DecodeBody(file.Body, nil, t)
	if diags.HasErrors() {
		return diags
	}
	return nil
}
