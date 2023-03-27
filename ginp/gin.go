package ginp

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/things-go/dyn/genproto/errors"
	transportHttp "github.com/things-go/dyn/transport/http"
)

var _ transportHttp.Carrier = (*GinCarry)(nil)

type GinCarry struct {
	validation *validator.Validate
}

func NewCarryForGin() *GinCarry {
	return &GinCarry{
		validation: func() *validator.Validate {
			v := validator.New()
			v.SetTagName("binding")
			return v
		}(),
	}
}
func (*GinCarry) WithValueUri(req *http.Request, params gin.Params) *http.Request {
	return transportHttp.WithValueUri(req, params)
}
func (*GinCarry) Bind(cg *gin.Context, v any) error {
	return cg.ShouldBind(v)
}
func (*GinCarry) BindQuery(cg *gin.Context, v any) error {
	return cg.ShouldBindQuery(v)
}
func (*GinCarry) BindUri(cg *gin.Context, v any) error {
	return cg.ShouldBindUri(v)
}
func (*GinCarry) ErrorBadRequest(cg *gin.Context, err error) {
	Abort(cg, errors.ErrBadRequest(err.Error()))
}
func (*GinCarry) Error(cg *gin.Context, err error) {
	Abort(cg, err)
}
func (*GinCarry) Render(cg *gin.Context, v any) {
	Response(cg, v)
}
func (cg *GinCarry) Validator() *validator.Validate {
	return cg.validation
}
func (cg *GinCarry) Validate(ctx context.Context, v any) error {
	return cg.validation.StructCtx(ctx, v)
}
func (cg *GinCarry) StructCtx(ctx context.Context, v any) error {
	return cg.validation.StructCtx(ctx, v)
}
func (cg *GinCarry) Struct(v any) error {
	return cg.validation.Struct(v)
}
func (cg *GinCarry) VarCtx(ctx context.Context, v any, tag string) error {
	return cg.validation.VarCtx(ctx, v, tag)
}
func (cg *GinCarry) Var(v any, tag string) error {
	return cg.validation.Var(v, tag)
}
