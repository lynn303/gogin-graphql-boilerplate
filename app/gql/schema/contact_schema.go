/**
 * @Time  : 2020-01-16 18:10
 * @Author: Lynn
 * @File  : contact_schema
 * @Description:
 * @History:
 *  1.[2020-01-16-18:10] new created
 */
package schema

import "github.com/graphql-go/graphql"

var contactInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "userInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"userId":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int), Description: "用戶名"},
		"name":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "姓名"},
		"phone":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "手机号码"},
		"address":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "地址"},
		"position": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "职位"},
		"qq":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "QQ号"},
		"wechat":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "微信号"},
	},
})
var contactRstType = graphql.NewObject(graphql.ObjectConfig{
	Name: "userRst",
	Fields: graphql.Fields{
		"userId":   &graphql.Field{Type: graphql.NewNonNull(graphql.Int), Description: "用戶名"},
		"name":     &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "姓名"},
		"phone":    &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "手机号码"},
		"address":  &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "地址"},
		"position": &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "职位"},
		"qq":       &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "QQ号"},
		"wechat":   &graphql.Field{Type: graphql.NewNonNull(graphql.String), Description: "微信号"},
	},
})
