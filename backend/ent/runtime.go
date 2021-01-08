// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/itclub/app/ent/doctor"
	"github.com/itclub/app/ent/drug"
	"github.com/itclub/app/ent/nurse"
	"github.com/itclub/app/ent/patient"
	"github.com/itclub/app/ent/prescription"
	"github.com/itclub/app/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	doctorFields := schema.Doctor{}.Fields()
	_ = doctorFields
	// doctorDescDoctorEmail is the schema descriptor for Doctor_Email field.
	doctorDescDoctorEmail := doctorFields[0].Descriptor()
	// doctor.DoctorEmailValidator is a validator for the "Doctor_Email" field. It is called by the builders before save.
	doctor.DoctorEmailValidator = doctorDescDoctorEmail.Validators[0].(func(string) error)
	// doctorDescDoctorName is the schema descriptor for Doctor_Name field.
	doctorDescDoctorName := doctorFields[1].Descriptor()
	// doctor.DoctorNameValidator is a validator for the "Doctor_Name" field. It is called by the builders before save.
	doctor.DoctorNameValidator = doctorDescDoctorName.Validators[0].(func(string) error)
	// doctorDescDoctorTel is the schema descriptor for Doctor_Tel field.
	doctorDescDoctorTel := doctorFields[2].Descriptor()
	// doctor.DoctorTelValidator is a validator for the "Doctor_Tel" field. It is called by the builders before save.
	doctor.DoctorTelValidator = doctorDescDoctorTel.Validators[0].(func(string) error)
	drugFields := schema.Drug{}.Fields()
	_ = drugFields
	// drugDescDrugName is the schema descriptor for Drug_Name field.
	drugDescDrugName := drugFields[0].Descriptor()
	// drug.DrugNameValidator is a validator for the "Drug_Name" field. It is called by the builders before save.
	drug.DrugNameValidator = drugDescDrugName.Validators[0].(func(string) error)
	nurseFields := schema.Nurse{}.Fields()
	_ = nurseFields
	// nurseDescNurseEmail is the schema descriptor for Nurse_Email field.
	nurseDescNurseEmail := nurseFields[0].Descriptor()
	// nurse.NurseEmailValidator is a validator for the "Nurse_Email" field. It is called by the builders before save.
	nurse.NurseEmailValidator = nurseDescNurseEmail.Validators[0].(func(string) error)
	// nurseDescNurseName is the schema descriptor for Nurse_Name field.
	nurseDescNurseName := nurseFields[1].Descriptor()
	// nurse.NurseNameValidator is a validator for the "Nurse_Name" field. It is called by the builders before save.
	nurse.NurseNameValidator = nurseDescNurseName.Validators[0].(func(string) error)
	// nurseDescNurseTel is the schema descriptor for Nurse_Tel field.
	nurseDescNurseTel := nurseFields[2].Descriptor()
	// nurse.NurseTelValidator is a validator for the "Nurse_Tel" field. It is called by the builders before save.
	nurse.NurseTelValidator = nurseDescNurseTel.Validators[0].(func(string) error)
	// nurseDescNursePassword is the schema descriptor for Nurse_Password field.
	nurseDescNursePassword := nurseFields[3].Descriptor()
	// nurse.NursePasswordValidator is a validator for the "Nurse_Password" field. It is called by the builders before save.
	nurse.NursePasswordValidator = nurseDescNursePassword.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientName is the schema descriptor for Patient_Name field.
	patientDescPatientName := patientFields[0].Descriptor()
	// patient.PatientNameValidator is a validator for the "Patient_Name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	prescriptionFields := schema.Prescription{}.Fields()
	_ = prescriptionFields
	// prescriptionDescPrescripDateTime is the schema descriptor for Prescrip_DateTime field.
	prescriptionDescPrescripDateTime := prescriptionFields[1].Descriptor()
	// prescription.DefaultPrescripDateTime holds the default value on creation for the Prescrip_DateTime field.
	prescription.DefaultPrescripDateTime = prescriptionDescPrescripDateTime.Default.(func() time.Time)
}