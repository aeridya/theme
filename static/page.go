package static

import (
	"html/template"

	"github.com/aeridya/core"
	"github.com/aeridya/page"
)

type Page struct {
	page.Page

	Templates []string
	PageTemp  *template.Template
}

func (p *Page) PageInit(title, route, tmpldir string, tmpls ...string) error {
	p.Route = route
	p.Title = title
	p.Templates = AddDir(tmpldir, tmpls)
	err := p.LoadPage()
	return err
}

func AddDir(tmpldir string, tmpls []string) []string {
	t := make([]string, len(tmpls))
	for i := range tmpls {
		t[i] = tmpldir + "/" + tmpls[i]
	}
	return t
}

func (p *Page) LoadPage() error {
	var err error
	p.PageTemp, err = template.ParseFiles(p.Templates...)
	return err
}

func (p *Page) Get(resp *core.Response) {
	resp.Good(200)
	if core.Development {
		p.LoadPage()
	}
	p.PageTemp.Execute(resp.W, p)
}
