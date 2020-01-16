/**
 * @Time  : 2020-01-16 10:12
 * @Author: Lynn
 * @File  : graphql
 * @Description:
 * @History:
 *  1.[2020-01-16-10:12] new created
 */
package gql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
