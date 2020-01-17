/**
 * @Time  : 2020-01-16 10:12
 * @Author: Lynn
 * @File  : query
 * @Description:
 * @History:
 *  1.[2020-01-16-10:12] new created
 */
package gql

import (
	"github.com/graphql-go/graphql"
	"gogin-graphql-boilerplate/app/gql/schema"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"userList": &schema.UserList,
	},
})
