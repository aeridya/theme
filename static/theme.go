package static

import (
	"html/template"

	"github.com/aeridya/core"
)

type Theme struct {
	Pages  map[string]core.Paging
	Errors map[int]*template.Template

	TemplateDir string
}

func (t *Theme) StaticInit(base string) error {
	t.Pages = make(map[string]core.Paging)
	t.Errors = make(map[int]*template.Template)

	a := template.New("default")
	t.Errors[0], _ = a.Parse("An error has occurred: {{.Status}}\n{{.Err}}\n")

	if s, err := core.Config.GetString("", "Template"); err != nil {
		return err
	} else {
		t.TemplateDir = s
	}
	return nil
}

func (t Theme) Parse(input string, resp *core.Response) core.Paging {
	if p, ok := t.Pages[input]; !ok {
		resp.Error("Page " + input + " not found")
		return nil
	} else {
		return p
	}
}

func (t Theme) Serve(resp *core.Response) {
	o := t.Parse(resp.R.URL.Path, resp)
	if o == nil {
		resp.Bad(404, resp.Err.Error())
		//core.ThemeError(resp, t)
		return
	}
	core.ServePage(resp, o)
	return
}

func (t Theme) Error(resp *core.Response) {
	if resp.Data == nil {
		resp.Data = resp
	}
	if s, ok := t.Errors[resp.Status]; ok {
		s.Execute(resp.W, resp.Data)
	} else {
		t.Errors[0].Execute(resp.W, resp)
	}
	return
}
