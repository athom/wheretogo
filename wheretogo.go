package wheretogo

import (
	"github.com/paulbellamy/mango"
	"net/http"
)

var DefaultRouter = NewRouter()
var DefaultPath = "/"

func Go(key string) {
	DefaultRouter.Go(key)
}

func Where(key string) (r string) {
	return DefaultRouter.Where(key)
}

func AddRoute(key string, path string) {
	DefaultRouter.AddRoute(key, path)
}

type Router struct {
	Pool map[string]string
}

func NewRouter() (r *Router) {
	r = &Router{
		Pool: make(map[string]string),
	}
	return
}

func (this *Router) Where(key string) (r string) {
	path, ok := this.Pool[key]
	if ok {
		delete(this.Pool, key)
		r = path
		return
	}

	r = DefaultPath
	return
}

func (this *Router) Go(key string) {
	path := Where(key)
	mango.Redirect(http.StatusFound, path)
}

func (this *Router) AddRoute(key string, path string) {
	if this.Pool == nil {
		this.Pool = make(map[string]string)
	}
	this.Pool[key] = path
}
