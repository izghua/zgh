package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/izghua/zgh/conf"
	"github.com/izghua/zgh/utils"
	"strconv"
	"time"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	if data == nil {
		data = []string{}
	}
	msg := conf.GetMsg(errCode)
	beginTime, _ := strconv.ParseInt(g.C.Writer.Header().Get("X-Begin-Time"), 10, 64)

	duration := time.Now().Sub(time.Unix(0, beginTime))
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	roundedStr := fmt.Sprintf("%.3fms", rounded)
	g.C.Writer.Header().Set("X-Response-time", roundedStr)
	g.C.JSON(httpCode, gin.H{
		"code":    errCode,
		"message": msg,
		"data":    data,
	})
	utils.ZLog().Info("message", "API Response", "code", errCode, "errMsg", msg, "took", roundedStr)
	g.C.Abort()
	return
}
