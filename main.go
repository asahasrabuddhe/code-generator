package main

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"log"
	"os"
	"strings"
	"text/template"
)

type JsonMarshalerMeta struct {
	PackageName     string
	ObjectName      string
	ObjectType      string
	PointerReceiver bool
}

func main() {
	var ptrReceiver bool
	var pkgName, objName, objType string

	ptrReceiver = true

	flag.StringVar(&pkgName, "package", "", "name of the package for which the code is to be generated")
	flag.StringVar(&objName, "name", "", "name of object")
	flag.StringVar(&objType, "type", "", "type of object")

	flag.BoolVar(&ptrReceiver, "pointer", true, "toggle pointer receiver")

	flag.Parse()

	if pkgName == "" {
		log.Fatal("package name cannot be blank")
	}

	if objName == "" {
		log.Fatal("object name cannot be blank")
	}

	if objType == "" {
		log.Fatal("type name cannot be blank")
	}

	box := packr.New("templates", "./templates")

	funcMap := map[string]interface{}{
		"printReceiver": func(receiver string, pointer bool) string {
			if pointer {
				return fmt.Sprintf("*%v", receiver)
			}

			return receiver
		},
	}

	t, err := box.FindString("marshal-json.gotpl")
	if err != nil {
		log.Fatal(err)
	}

	tpl, err := template.New("jsonmarshaler").Funcs(funcMap).Parse(t)
	if err != nil {
		log.Fatalln(err)
	}

	metadata := JsonMarshalerMeta{
		PackageName:     pkgName,
		ObjectName:      objName,
		ObjectType:      objType,
		PointerReceiver: ptrReceiver,
	}

	file, err := os.Create(fmt.Sprintf("%v_json.go", strings.ToLower(metadata.ObjectType)))

	err = tpl.Execute(file, metadata)

	if err != nil {
		log.Fatal(err)
	}
}
