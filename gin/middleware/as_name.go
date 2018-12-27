/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-27
 * Time: 21:42
 */
package ginmiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterAsName(routerAsName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//apiG := api.Gin{C: c}
		//apiG.C.rea
		fmt.Println(c.Request.RequestURI,c.HandlerName(),c.Keys,"234234234")
		c.Set("router_as_name",routerAsName)
		c.Next()
	}
}

