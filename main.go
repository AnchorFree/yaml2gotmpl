package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Obj interface{}

func readYAML(filename string) (Obj, error) {
	var o Obj

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &o)
	if err != nil {
		return nil, err
	}

	return o, err
}

func main() {

	inFile := flag.String("in", "", "Input file name, default is stdin")
	tmpl := flag.String("tmpl", "", "Go template file")
	outFile := flag.String("out", "", "Output file, default is stdout")

	flag.Parse()

	var out *os.File
	var in []byte
	var o Obj
	var err error

	if *tmpl == "" {
		panic("Template is mandatory")
	}

	if *inFile == "" {
		in, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

	} else {
		f, err := os.Open(*inFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		in, err = ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
	}

	if *outFile == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(*outFile)
		if err != nil {
			panic(err)
		}
	}

	err = yaml.Unmarshal(in, &o)
	if err != nil {
		panic(err)
	}

	t := template.Must(template.ParseFiles(*tmpl))
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(out, o)
	if err != nil {
		panic(err)
	}
}
