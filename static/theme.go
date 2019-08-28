package static

import (
	"html/template"

	"github.com/aeridya/core"
	"github.com/aeridya/page"
	"github.com/aeridya/theme"

	"github.com/aeridya/core/configurit"
	"github.com/aeridya/core/logit"
)

type Theme struct {
	theme.Theme
	Pages  map[string]page.Paging
	Errors map[int]*template.Template

	TemplateDir string
}

func (t *Theme) StaticInit() error {
	t.Pages = make(map[string]page.Paging)
	t.Errors = make(map[int]*template.Template)

	a := template.New("default")
	t.Errors[0], _ = a.Parse("An error has occurred: {{.Status}}\n{{.Err}}\n")
	t.Errors[404], _ = a.Parse("Page Not found: {{.Status}}\n{{.Err}}\n")

	if s, err := configurit.Config.GetString("static", "Template"); err != nil {
		return err
	} else {
		t.TemplateDir = s
	}

	logit.Log(0, "hello from theme init: ", t.TemplateDir)

	return nil
}

func (t Theme) Parse(input string, resp *core.Response) page.Paging {
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
		return
	}
	ServePage(resp, o)
	return
}

func (t Theme) Error(resp *core.Response) {
	if resp.Data == nil {
		resp.Data = resp
	}
	if s, ok := t.Errors[resp.Status]; ok {
		s.Execute(resp.W, resp.Data)
	} else {
		resp.Bad(400, resp.Err.Error())
		t.Errors[0].Execute(resp.W, resp)
	}
	return
}

func ServePage(resp *core.Response, p page.Paging) {
	switch resp.R.Method {
	case "GET":
		p.Get(resp)
	case "PUT":
		p.Put(resp)
	case "POST":
		p.Post(resp)
	case "DELETE":
		p.Delete(resp)
	case "OPTIONS":
		p.Options(resp)
	case "HEAD":
		p.Head(resp)
	default:
		p.Unsupported(resp)
	}
}
