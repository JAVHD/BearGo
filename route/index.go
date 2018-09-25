/*
* @Author: bear
* @Date:   2018-09-19 13:56:15
* @Last Modified by:   bear
* @Last Modified time: 2018-09-25 16:18:40
 */

package route

import (
	"beargo/application/index/controller"
	"beargo/beargo/core"
	"fmt"
)

func init() {
	fmt.Println("this is a index module route")
	core.AddRoute("public/index", &controller.PublicController{})
}
