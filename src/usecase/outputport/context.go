package outputport

type Context interface {
	JSON(code int, i interface{}) error
	Param(name string) string
	ParamNames() []string
	QueryParam(name string) string
	FormValue(name string) string
	Bind(i interface{}) error
	Redirect(code int, url string) error
	Get(key string) interface{}
	Set(key string, val interface{})
	QueryString() string
	String(code int, s string) error
}
