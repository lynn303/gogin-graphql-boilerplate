/**
 * @Time  : 2020-01-16 18:34
 * @Author: Lynn
 * @File  : router
 * @Description:
 * @History:
 *  1.[2020-01-16-18:34] new created
 */
package router

import (
	"github.com/gin-gonic/gin"
	"gogin-graphql-boilerplate/app/gql"
	"gogin-graphql-boilerplate/middleware/cors"
)

var Router *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
}

func SetRouter() {
	Router.Use(cors.CORS())
	Router.POST("/graphql", gql.GraphqlHandler())
	Router.GET("/graphql", gql.GraphqlHandler())
}
