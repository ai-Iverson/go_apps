package utility

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)
// 权限中间件
type middlewareService struct{}

var middleware = middlewareService{}

func Middleware() *middlewareService {
	return &middleware
}

//func (s *middlewareService) Auth(r *ghttp.Request) {
//	// GfJWTMiddleware gf jwt集成的中间件
//	// Auth是权限service中配置的gf jwt
//	Auth().MiddlewareFunc()(r)
//	r.Middleware.Next()
//}

type DefaultHandlerRes struct {
	ResultCode int         `json:"resultCode"    dc:"Error code"`
	Message    string      `json:"message" dc:"Error message" d:""`
	Data       interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

func (s *middlewareService) CustomResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		default:
			code = gcode.CodeUnknown
		}
	} else {
		code = gcode.New(200, "success", "")
	}
	r.Response.WriteJson(DefaultHandlerRes{
		ResultCode: code.Code(),
		Message:    msg,
		Data:       res,
	})
}