package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	defaultOption = "default"
)

var filename = flag.String("filename", "", "filename")

func main() {
	flag.Parse()

	plugin, err := plugin.NewPlugin()
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"toUpper": strings.ToUpper,

		"contentType": func(route spec.Route) string {
			// 判断 Content-Type
			if defineStruct, ok := route.RequestType.(spec.DefineStruct); ok {
				// json
				if len(defineStruct.GetBodyMembers()) > 0 {
					return "application/json"
				} else if len(defineStruct.GetFormMembers()) > 0 {
					return "application/x-www-form-urlencoded"
				}
			}
			return ""
		},
		"genTypes": func(route spec.Route) string {
			structType, ok := route.RequestType.(spec.DefineStruct)
			if !ok {
				return ""
			}
			var builder strings.Builder
			fmt.Fprintf(&builder, "{\n")
			for index, member := range structType.Members {
				for _, tag := range member.Tags() {
					if tag.Key == "json" {
						fmt.Fprintf(&builder, "	%s : \"%s\"", tag.Name, getTagDefaultVaule(tag.Options))
					}
				}
				if index < len(structType.Members)-1 {
					fmt.Fprintf(&builder, ", \n")
				} else {
					fmt.Fprintf(&builder, "\n")
				}
			}
			fmt.Fprintf(&builder, "}")
			return builder.String()
		},
	}

	t := template.Must(template.New("service").Delims("<<", ">>").Funcs(funcMap).Parse(apiTemplate))

	output := plugin.Dir + "/" + *filename

	f, err := pathx.CreateIfNotExist(output)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(f, map[string]interface{}{
		"Service": plugin.Api.Service,
	})

	fmt.Println(err)
}

func getTagDefaultVaule(opts []string) (val string) {
	val = ""
	for _, option := range opts {
		if strings.HasPrefix(option, defaultOption) {
			val = strings.Replace(option, defaultOption+"=", "", -1)
			return
		}
	}
	return
}

//go:embed api.tpl
var apiTemplate string
