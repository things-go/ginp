package ginp

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/things-go/dyn/encoding"
	"github.com/things-go/dyn/encoding/jsonpb"
	"github.com/things-go/dyn/errors"
	transportHttp "github.com/things-go/dyn/transport/http"
)

func Response(c *gin.Context, data ...any) {
	var obj any

	if len(data) > 0 {
		obj = data[0]
	} else {
		obj = struct{}{}
	}
	c.JSON(http.StatusOK, obj)
}

func Abort(c *gin.Context, err error) {
	e := errors.FromError(err)

	status := 599
	switch {
	case e.Code == -1:
		status = http.StatusInternalServerError
	case e.Code < 1000:
		status = int(e.Code)
	}
	c.AbortWithStatusJSON(status, e)
}

type Implemented struct {
	Encoding              *encoding.Encoding
	Validation            *validator.Validate
	DisableBindValidation bool
}

func NewDefaultImplemented() *Implemented {
	e := encoding.New()
	err := e.Register(encoding.MIMEJSON, &Codec{
		Marshaler: &jsonpb.Codec{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:  true,
				UseEnumNumbers: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return &Implemented{
		Encoding:              e,
		Validation:            transportHttp.Validator(),
		DisableBindValidation: false,
	}
}

func (i *Implemented) Validate(ctx context.Context, v any) error {
	if i.DisableBindValidation {
		return nil
	}
	return i.Validation.StructCtx(ctx, v)
}

func (*Implemented) ErrorEncoder(c *gin.Context, err error, isBadRequest bool) {
	if isBadRequest {
		err = errors.ErrBadRequest(err.Error())
	}
	Abort(c, err)
}

func (i *Implemented) Bind(c *gin.Context, v any) error {
	return i.Encoding.Bind(c.Request, v)
}
func (i *Implemented) BindQuery(c *gin.Context, v any) error {
	return i.Encoding.BindQuery(c.Request, v)
}
func (i *Implemented) BindUri(c *gin.Context, v any) error {
	return i.Encoding.BindUri(c.Request, v)
}
func (i *Implemented) RequestWithUri(req *http.Request, params gin.Params) *http.Request {
	return transportHttp.RequestWithUri(req, params)
}
func (i *Implemented) Render(c *gin.Context, statusCode int, v any) {
	c.Writer.WriteHeader(statusCode)
	err := i.Encoding.Render(c.Writer, c.Request, v)
	if err != nil {
		c.String(500, "Render failed cause by %v", err)
	}
}
