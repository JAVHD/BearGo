/*
* @Author: bear
* @Date:   2018-09-19 13:56:15
* @Last Modified by:   bear
* @Last Modified time: 2018-09-22 15:51:14
 */

package route

import (
	"beargo/beargo/core"
	"beargo/application/index/controller"
	"fmt"
)

func init(){
	fmt.Println("this is a index module route")
	core.AddRoute("public/index", &controller.PublicController{})
}
