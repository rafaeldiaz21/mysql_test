package main

import (
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"mysql_test/generators"
)

const Template string = `
package queries
// Automatically generated queries 
// ALL EDITS WILL BE LOST
const (
{{- range .Queries }}
	{{ printf "%s = ` + "`%s`" + `" .Name .Sql }}
{{- end }} 
)
`

type query struct {
	Name string
	Sql  string
}

type queryHandler struct {
	template            *template.Template
	multipleSpacesRegex *regexp.Regexp
	operatorSpacesRegex *regexp.Regexp
	queries             []query
	sourcePath          string
	targetPath          string
}

func (handler *queryHandler) SetSourcePath(sourcePath string) {
	handler.sourcePath = sourcePath
}

func (handler *queryHandler) GetSourcePath() string {
	return handler.sourcePath
}

func (handler *queryHandler) GetTargetPath() string {
	return handler.targetPath
}

func (handler *queryHandler) Prepare() {
}

func (handler *queryHandler) Write() error {
	var file *os.File
	var tpl = handler.template
	var err error
	file, err = os.Create(path.Join(handler.GetTargetPath(), "main.go"))
	if err != nil {
		return err
	}
	return tpl.Execute(file, struct {
		Queries []query
	}{
		Queries: handler.queries,
	})
}

func (handler *queryHandler) HandleFile(file os.FileInfo) error {
	// Just to avoid ugly syntax
	var multipleSpacesRegex = handler.multipleSpacesRegex
	var operatorSpacesRegex = handler.operatorSpacesRegex
	// Declare local stuff
	var err error
	var sql string
	var content string
	var name string

	name, err = generators.GetTypeName(file.Name())
	if err != nil {
		return err
	}
	content, err = generators.GetContent(handler.sourcePath, file.Name())
	if err != nil {
		return err
	}

	sql = strings.Replace(content, "\t\n\r", " ", -1)
	sql = multipleSpacesRegex.ReplaceAllString(sql, " ")
	sql = operatorSpacesRegex.ReplaceAllString(sql, "$1")
	sql = strings.Trim(sql, " \t\n\r")

	handler.queries = append(handler.queries, query{Name: name, Sql: sql})
	return nil
}

func main() {
	var sourcePath = "./sql"
	var targetPath = "./queries"
	var err error
	var parser = template.New("")
	var handler = &queryHandler{
		sourcePath:          sourcePath,
		targetPath:          targetPath,
		template:            template.Must(parser.Parse(Template)),
		multipleSpacesRegex: regexp.MustCompile(`\s+`),
		operatorSpacesRegex: regexp.MustCompile(`(?:([()])\s+)`),
	}
	// Let the handler prepare itself
	handler.Prepare()
	// Generate
	err = generators.Generate(handler)
	if err != nil {
		log.Println(err)
	}
	err = handler.Write()
	if err != nil {
		log.Fatalln(err)
	}
}
