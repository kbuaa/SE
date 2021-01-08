package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/itclub/app/ent"
	"github.com/itclub/app/ent/nurse"
)

type NurseController struct {
	client *ent.Client
	router gin.IRouter
}

type Nurse struct {
	Email    string
	Name     string
	Tel      string
	Password string
}

// CreateNurse handles POST requests for adding nurse entities
// @Summary Create nurse
// @Description Create nurse
// @ID create-nurse
// @Accept   json
// @Produce  json
// @Param nurse body ent.Nurse true "Nurse entity"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses [post]
func (ctl *NurseController) CreateNurse(c *gin.Context) {
	obj := Nurse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nurse binding failed",
		})
		return
	}

	n, err := ctl.client.Nurse.
		Create().
		SetNurseEmail(obj.Email).
		SetNurseName(obj.Name).
		SetNursePassword(obj.Password).
		SetNurseTel(obj.Tel).
		Save(context.Background())

	fmt.Println(err)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, n)
}

// GetNurse handles GET requests to retrieve a nurse entity
// @Summary Get a nurse entity by ID
// @Description get nurse by ID
// @ID get-nurse
// @Produce  json
// @Param id path int true "Nurse ID"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses/{id} [get]
func (ctl *NurseController) GetNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, n)
}

// ListNurse handles request to get a list of nurse entities
// @Summary List nurse entities
// @Description list nurse entities
// @ID list-nurse
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurses [get]
func (ctl *NurseController) ListNurse(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	nurses, err := ctl.client.Nurse.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, nurses)
}

// NewNurseController creates and registers handles for the nurse controller
func NewNurseController(router gin.IRouter, client *ent.Client) *NurseController {
	nc := &NurseController{
		client: client,
		router: router,
	}

	nc.register()

	return nc

}

func (ctl *NurseController) register() {
	nurses := ctl.router.Group("/nurses")

	nurses.POST("", ctl.CreateNurse)
	nurses.GET(":id", ctl.GetNurse)
	nurses.GET("", ctl.ListNurse)

}
