/**
 * @Time  : 2020-01-16 10:12
 * @Author: Lynn
 * @File  : mutation
 * @Description:
 * @History:
 *  1.[2020-01-16-10:12] new created
 */
package gql

import "github.com/graphql-go/graphql"

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutation",
	Fields: graphql.Fields{},
})
