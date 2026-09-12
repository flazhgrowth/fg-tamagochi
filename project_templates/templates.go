package projecttemplates

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/flazhgrowth/fg-gotools/random"
)

type (
	ProjectTemplate string
)

const (
	MaingoTemplate ProjectTemplate = `package main

import (
	"github.com/flazhgrowth/fg-tamagochi/app"
	"github.com/flazhgrowth/fg-tamagochi/cmd"
	"github.com/flazhgrowth/fg-tamagochi/cmd/serve"
	"github.com/flazhgrowth/fg-tamagochi/pkg/http/router"
)

func main() {
	cmd.Conjure(cmd.CmdArgs{
		ServeCmdArgs: serve.ServeCmdArgs{
			GetRoutesFn: func(app *app.App) router.Router {
				return router.NewRouter(router.OpenAPISpecInfo{
					Title:   "tamagochi service",
					Version: "0.0.1",
					Desc:    "tamagochi service api",
				})
			},
		},
	})
}
	`

	EntityEmptyTemplate      ProjectTemplate = `package {{.entity}}`
	EntityInterfacesTemplate ProjectTemplate = `package {{.entity}}

type API interface {}
type Service interface {}
type Repository interface {}
	`
	EntityDatabaseTemplate ProjectTemplate = `package {{.entity}}
type (
	{{.entity_title}} struct {}
)
	`
	EntityFilterSorterFieldsTemplate ProjectTemplate = `package {{.entity}}
type (
	{{.entity_title}}Filter struct {}
	{{.entity_title}}Sorter struct {}
	{{.entity_title}}UpdateFields struct {}
)
	`

	APIImplTemplate ProjectTemplate = `package {{.entity}}api

import "{{.packagename}}/internal/entity/{{.entity}}"

type api struct {
	{{.entity}}Svc {{.entity}}.Service
}

func New({{.entity}}svc {{.entity}}.Service) {{.entity}}.API {
	return &api{
		{{.entity}}Svc: {{.entity}}svc,
	}
}
	`

	APIImplEmptyTemplate ProjectTemplate = `package {{.entity}}api`

	ServiceImplTemplate ProjectTemplate = `package {{.entity}}svc

import "{{.packagename}}/internal/entity/{{.entity}}"

type service struct {
	{{.entity}}Repo {{.entity}}.Repository
}

func New({{.entity}}repo {{.entity}}.Repository) {{.entity}}.Service {
	return &service{
		{{.entity}}Repo: {{.entity}}repo,
	}
}
	`

	ServiceImplEmptyTemplate ProjectTemplate = `package {{.entity}}uc`

	DBRepositoryImplTemplate ProjectTemplate = `package {{.entity}}repo

import (
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator"
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator/sqltx"
	"{{.packagename}}/internal/entity/{{.entity}}"
)

type repository struct {
	actuator 	sqlator.SQLator
	tx 			sqltx.SQLTx
}

func New(actuator sqlator.SQLator, tx sqltx.SQLTx) {{.entity}}.Repository {
	return &repository{
		actuator: 	actuator,
		tx:			tx,
	}
}
	`

	DBRepositoryImplEmptyTemplate ProjectTemplate = `package {{.entity}}repo`

	GitignoreTemplate ProjectTemplate = `/etc/*`

	ConfigTemplate ProjectTemplate = `env: 'local'
http:
  timeout:
    unit: 'second'
    write: '30'
    read: '30'
    idle: '30'
  server: '11011'
  `
	VaultTemplate ProjectTemplate = `{
	"database": {
		"driver": "postgres",
		"reader_dsn": "dsn",
		"writer_dsn": "dsn"
	}
}
	`
	FeatureflagTemplate ProjectTemplate = `
example_feature_enabled: true
	`
)

func (templ ProjectTemplate) WriteTo(path string, binding map[string]any) error {
	if binding == nil {
		return os.WriteFile(path, []byte(templ.val()), 0644)
	}
	id := random.GenerateRandomNumber(3)
	contents := bindData(int64(id), templ.val(), binding)
	return os.WriteFile(path, contents, 0644)
}

func (templ ProjectTemplate) val() string {
	return string(templ)
}

func bindData(id int64, templ string, data any) []byte {
	if data == nil {
		return []byte(templ)
	}
	var templateBuffer bytes.Buffer

	name := fmt.Sprintf("%d-%s", id, templ)
	htmlTemplate := template.Must(template.New(name).Parse(templ))
	if err := htmlTemplate.ExecuteTemplate(&templateBuffer, name, data); err != nil {
		return []byte(templ)
	}

	return templateBuffer.Bytes()
}
