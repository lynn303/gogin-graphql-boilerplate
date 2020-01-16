/**
 * @Time  : 2020-01-16 18:10
 * @Author: Lynn
 * @File  : user_schema
 * @Description:
 * @History:
 *  1.[2020-01-16-18:10] new created
 */
package schema

import "github.com/graphql-go/graphql"

var userInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "userInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"userName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "用戶名"},
		"headUrl":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "头像"},
	},
})
var userRstType = graphql.NewObject(graphql.ObjectConfig{
	Name: "userRst",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.NewNonNull(graphql.Int), Description: "Id"},
		"userName": &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "用户名"},
		"headUrl":  &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "头像"},
		"isActive": &graphql.Field{Type: graphql.NewNonNull(graphql.Boolean), Description: "是否激活"},
	},
})
