/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-27
 * Time: 21:53
 */
package ginmiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		v,res := c.Get("router_as_name")
		fmt.Println(v,res,"路由别名")
	}
}

