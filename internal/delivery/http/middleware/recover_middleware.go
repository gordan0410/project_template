package middleware

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	common_response "github.com/gordan0410/common/enum/response"
	common_wh_error "github.com/gordan0410/common/enum/wh_error"
	common_log "github.com/gordan0410/common/log"
)

func (m *middleware) RecoverMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}

				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					_ = ctx.Error(err.(error))
					ctx.Abort()
				} else {
					common_log.WithCtx(ctx).Error().
						Interface("error", err).
						Interface("header", headers).
						Msg("panic recovered")

					resp := common_response.NewResponse(ctx)
					resp.Data = struct{}{}
					resp.Code = common_wh_error.InternalServerError.GetCode()
					resp.Msg = common_wh_error.InternalServerError.GetMessage()
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
				}

			}
		}()
		ctx.Next()
	}
}
