package http

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/things-go/dyn/encoding"
)

func RequestWithUri(req *http.Request, params gin.Params) *http.Request {
	vars := make(url.Values, len(params))
	for _, p := range params {
		vars.Set(p.Key, p.Value)
	}
	return encoding.RequestWithUri(req, vars)
}
