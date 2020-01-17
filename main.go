/**
 * @Time  : 2020-01-16 9:58
 * @Author: Lynn
 * @File  : main
 * @Description:
 * @History:
 *  1.[2020-01-16-9:58] new created
 */
package main

import "gogin-graphql-boilerplate/app/router"

func main() {
	r := router.Router
	router.SetRouter()
	r.Run(":2006")
}
