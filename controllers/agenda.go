package controllers

import (
	//"encoding/json"
	"skripsih/models"
	//"skripsih/modules/utils"
	//"strconv"
	"github.com/astaxie/beego"
)

// oprations for Kv
type AgendaController struct {
	beego.Controller
}
type resultAgenda struct {
	Success bool
	Rows    interface{}
	Message string
}

// @Title Get All
// @Description get Kv
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Kv
// @Failure 403
// @router / [get]
func (c *AgendaController) GetAgenda() {
	var limit int64 = 10
	var offset int64 = 0
	res := struct {
		Success bool
		Rows    []interface{}
		Total   int64
		Message string
	}{}

	fields := c.GetString("fields")
	limit, err := c.GetInt64("limit")
	offset, err = c.GetInt64("offset")
	sortby := c.GetString("sortby")
	order := c.GetString("order")
	query := c.GetString("query")


	if offset == 0 {
		total, err := models.CountGetAllAgenda(query)
		if err != nil {
			res.Message = err.Error()
		}
		res.Total = total
	} else {
		res.Total = 0
	}
	l, err := models.GetAllAgenda(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
		res.Success = false
		res.Rows = l
		res.Message = err.Error()
	} else {
		res.Success = true
		res.Rows = l
	}
	c.Data["json"] = res
	c.ServeJSON()
}
