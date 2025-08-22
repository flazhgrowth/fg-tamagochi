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
				return router.NewRouter("")
			},
		},
	})
}
	`

	EntityTemplate ProjectTemplate = `package {{.entity}}`

	EntityAPIInterfaceTemplate ProjectTemplate = `package {{.entity}}

type {{.entity_title}}API interface {}
	`

	EntityUsecaseInterfaceTemplate ProjectTemplate = `package {{.entity}}

type {{.entity_title}}Usecase interface {}
	`

	EntityRepositoryInterfaceTemplate ProjectTemplate = `package {{.entity}}

type {{.entity_title}}Repository interface {}
	`

	TransportImplTemplate ProjectTemplate = `package {{.entity}}api

import "{{.packagename}}/internal/entity/{{.entity}}"

type API struct {
	{{.entity}}Usecase {{.entity}}.{{.entity_title}}Usecase
}

func New({{.entity}}usecase {{.entity}}.{{.entity_title}}Usecase) {{.entity}}.{{.entity_title}}API {
	return &API{
		{{.entity}}Usecase: {{.entity}}usecase,
	}
}
	`

	TransportImplEmptyTemplate ProjectTemplate = `package {{.entity}}api`

	UsecaseImplTemplate ProjectTemplate = `package {{.entity}}uc

import "{{.packagename}}/internal/entity/{{.entity}}"

type Usecase struct {
	{{.entity}}Repo {{.entity}}.{{.entity_title}}Repository
}

func New({{.entity}}repo {{.entity}}.{{.entity_title}}Repository) {{.entity}}.{{.entity_title}}Usecase {
	return &Usecase{
		{{.entity}}Repo: {{.entity}}repo,
	}
}
	`

	UsecaseImplEmptyTemplate ProjectTemplate = `package {{.entity}}uc`

	DBRepositoryImplTemplate ProjectTemplate = `package {{.entity}}repo

import (
	"github.com/flazhgrowth/fg-tamagochi/pkg/db/sqlator"
	"{{.packagename}}/internal/entity/{{.entity}}"
)

type Repository struct {
	actuator sqlator.SQLator
}

func New(actuator sqlator.SQLator) {{.entity}}.{{.entity_title}}Repository {
	return &Repository{
		actuator: actuator,
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
  server: '11011'`
	VaultTemplate ProjectTemplate = `{
	"database": {
		"driver": "postgres",
		"reader_dsn": "dsn",
		"writer_dsn": "dsn"
	}
}`
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
