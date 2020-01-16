/**
 * @Time  : 2020-01-16 10:12
 * @Author: Lynn
 * @File  : gogin-graphql-boilerplate_schema
 * @Description:
 * @History:
 *  1.[2020-01-16-10:12] new created
 */
package gql

import "github.com/graphql-go/graphql"

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
