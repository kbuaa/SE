package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itclub/app/ent"
	"github.com/itclub/app/ent/doctor"
	"github.com/itclub/app/ent/drug"
	"github.com/itclub/app/ent/nurse"
	"github.com/itclub/app/ent/patient"
	"github.com/itclub/app/ent/prescription"
)

// PrescriptionController defines the struct for the prescription controller
type PrescriptionController struct {
	client *ent.Client
	router gin.IRouter
}

//Prescription struct
type Prescription struct {
	Doctor  int
	Patient int
	Nurse   int
	Drug    int
	Note    string
	Added   string
}

// CreatePrescription handles POST requests for adding prescription entities
// @Summary Create prescription
// @Description Create prescription
// @ID create-prescription
// @Accept   json
// @Produce  json
// @Param prescription body Prescription true "Prescription entity"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions [post]
func (ctl *PrescriptionController) CreatePrescription(c *gin.Context) {
	obj := Prescription{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Prescription binding failed",
		})
		return
	}
	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}
	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}
	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Nurse not found",
		})
		return
	}
	dr, err := ctl.client.Drug.
		Query().
		Where(drug.IDEQ(int(obj.Drug))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Drug not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.Added)
	pr, err := ctl.client.Prescription.
		Create().
		SetDoctor(d).
		SetPatient(p).
		SetNurse(n).
		SetDrug(dr).
		SetPrescripDateTime(time).
		SetPrescripNote(obj.Note).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   pr,
	})
}

// GetPrescription handles GET requests to retrieve a prescription entity
// @Summary Get a prescription entity by ID
// @Description get prescription by ID
// @ID get-prescription
// @Produce  json
// @Param id path int true "Prescription ID"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions/{id} [get]
func (ctl *PrescriptionController) GetPrescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.Prescription.
		Query().
		Where(prescription.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, b)
}

// ListPrescription handles request to get a list of prescription entities
// @Summary List prescription entities
// @Description list prescription entities
// @ID list-prescription
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions [get]
func (ctl *PrescriptionController) ListPrescription(c *gin.Context) {
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
	prescriptions, err := ctl.client.Prescription.
		Query().
		WithDoctor().
		WithPatient().
		WithNurse().
		WithDrug().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, prescriptions)
}

// DeletePrescription handles DELETE requests to delete a prescription entity
// @Summary Delete a prescription entity by ID
// @Description get prescription by ID
// @ID delete-prescription
// @Produce  json
// @Param id path int true "Prescription ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions/{id} [delete]
func (ctl *PrescriptionController) DeletePrescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Prescription.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdatePrescription handles PUT requests to update a Prescription entity
// @Summary Update a prescription entity by ID
// @Description update prescription by ID
// @ID update-prescription
// @Accept   json
// @Produce  json
// @Param id path int true "Prescription ID"
// @Param drugAllergy body ent.Prescription true "Prescription entity"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prescriptions/{id} [put]
func (ctl *PrescriptionController) UpdatePrescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.Prescription{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "booking binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.Prescription.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, b)
}

// NewPrescriptionController creates and registers handles for the prescription controller
func NewPrescriptionController(router gin.IRouter, client *ent.Client) *PrescriptionController {
	da := &PrescriptionController{
		client: client,
		router: router,
	}
	da.register()
	return da
}
func (ctl *PrescriptionController) register() {
	prescriptions := ctl.router.Group("/prescriptions")
	prescriptions.POST("", ctl.CreatePrescription)
	prescriptions.GET("", ctl.ListPrescription)
	prescriptions.DELETE(":id", ctl.DeletePrescription)
}
