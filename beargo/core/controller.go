/*
* @Author: bear
* @Date:   2018-09-22 16:29:58
* @Last Modified by:   bear
* @Last Modified time: 2018-09-22 16:30:02
 */
package core


type CoreController struct {

}


type ControllerInterface interface {

	Get()
	Post()
	Put()
	Delete()

}


func (c *CoreController) Get() {

}

func (c *CoreController) Post() {

}

func (c *CoreController) Put() {

}

func (c *CoreController) Delete() {

}