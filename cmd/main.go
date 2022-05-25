package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	cmd := cobra.Command{
		Use: "yaml-readme",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var items []map[string]string

			// find YAML files
			var files []string
			var data []byte
			if files, err = filepath.Glob("items/*.yaml"); err == nil {
				for _, metaFile := range files {
					if data, err = ioutil.ReadFile(metaFile); err != nil {
						fmt.Println("failed to read file", metaFile, err)
						continue
					}

					metaMap := make(map[string]string)
					if err = yaml.Unmarshal(data, metaMap); err != nil {
						fmt.Printf("failed to parse file [%s] as a YAML, error: %v\n", metaFile, err)
						continue
					}
					items = append(items, metaMap)
				}
			}

			// load readme template
			var readmeTpl string
			if data, err = ioutil.ReadFile("README.tpl"); err != nil {
				fmt.Printf("failed to load README template, error: %v\n", err)
				readmeTpl = `
|中文名称|英文名称|JD|
|---|---|---|
{{- range $val := .}}
|{{$val.zh}}|{{$val.en}}|{{$val.jd}}|
{{end}}
`
			}
			readmeTpl = string(data)

			// generate readme file
			var tpl *template.Template
			if tpl, err = template.New("readme").Parse(readmeTpl); err != nil {
				return
			}
			err = tpl.Execute(os.Stdout, items)
			return
		},
	}
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
