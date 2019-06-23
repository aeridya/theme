package sqltheme

import (
	"errors"
	"github.com/aeridya/core"
	"database/sql"
	"html/template"
)

type Theme struct {
	DB *sql.DB
}

func (t *Theme)SQLInit() error {
	return errors.New("SQL Theme Requires a custom function for SQLInit to connect to the Database.")
}

func (t Theme) Parse(input string, resp *core.Response) *template.HTML {
	out := "hello"
	return &(template.HTML)out
}

func (t Theme) Serve(resp *core.Response) {
	o := t.Parse(resp.R.URL.Path, resp)
	if o == nil {
		resp.Bad(404, resp.Err.Error())
		return
	}
	core.ServePage(resp, o)
}

func (t Theme) Error(resp *core.Response) {

}