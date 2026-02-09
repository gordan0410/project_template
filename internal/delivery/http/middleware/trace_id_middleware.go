package middleware

import (
	"github.com/gin-gonic/gin"
	common_enum "github.com/gordan0410/common/enum"
	common_helper "github.com/gordan0410/common/helper"
)

func (m *middleware) AddTraceID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if xTraceID := ctx.GetHeader(common_enum.TraceId.ToString()); xTraceID == "" {
			newTraceID := common_helper.GetUUID()
			ctx.Set(common_enum.TraceId.ToString(), newTraceID) // set to context
		} else {
			ctx.Set(common_enum.TraceId.ToString(), xTraceID) // set to context
		}
		newSubTraceID := common_helper.GetUUID()
		ctx.Set(common_enum.SubTraceId.ToString(), newSubTraceID)
		ctx.Next()
	}
}
