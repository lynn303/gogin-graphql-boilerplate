/**
 * @Time  : 2020-01-16 18:10
 * @Author: Lynn
 * @File  : user_schema
 * @Description:
 * @History:
 *  1.[2020-01-16-18:10] new created
 */
package schema

import (
	"github.com/graphql-go/graphql"
	"gogin-graphql-boilerplate/app/gql"
	"gogin-graphql-boilerplate/app/service"
)

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

var UserAop = graphql.Field{
	Description: "新增|编辑用户",
	Type:        gql.ResultType,
	Args:        graphql.FieldConfigArgument{"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(userInputType)}},
	Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
		input, _ := p.Args["input"].(map[string]interface{})
		return service.AopUser(input)
	},
}
var UserDel = graphql.Field{
	Description: "删除用户",
	Type:        gql.ResultType,
	Args:        graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)}},
	Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
		id, _ := p.Args["id"].(int)
		return service.DelUser(uint(id))
	},
}
var UserList = graphql.Field{
	Description: "获取用户列表",
	Type:        graphql.NewList(userRstType),
	Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
		return service.GetUserList()
	},
}
