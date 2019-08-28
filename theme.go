package theme

import (
	"github.com/aeridya/core"
	"github.com/aeridya/page"
)

//Option is the function type for initialization of Themes
type Option func()

//Themer is the interface to satisfy for Themes
type themer interface {
	Init(options ...Option) error
	Serve(resp *core.Response)
	Error(resp *core.Response)
}

//Theme is a basic struct with helper functions for themes
type Theme struct {
	themer
}

//ParseOpts runs through the available Option functions and sets the values
func (t *Theme) ParseOpts(opts []Option) {
	for i := range opts {
		opts[i]()
	}
}

//Register will make Aeridya aware of the theme functions to be used
func Register(t themer) {
	core.Serve = t.Serve
	core.Error = t.Error
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
