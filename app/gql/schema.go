/**
 * @Time  : 2020-01-16 10:16
 * @Author: Lynn
 * @File  : schema
 * @Description:
 * @History:
 *  1.[2020-01-16-10:16] new created
 */
package gql

import "github.com/graphql-go/graphql"

var ResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "result",
	Fields: graphql.Fields{
		"success": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "success:true;failed:false",
		},
	},
})

func GeneratePageType(otype graphql.Type, name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"count": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"items": &graphql.Field{Type: graphql.NewList(otype)},
		},
	})
}

var PageInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "pageInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id":       &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"page":     &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"pageSize": &graphql.InputObjectFieldConfig{Type: graphql.Int},
		"filter":   &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})
