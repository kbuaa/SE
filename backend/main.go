package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/itclub/app/controllers"
	"github.com/itclub/app/ent"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Doctors struct
type Doctors struct {
	Doctor []Doctor
}

//Doctor struct
type Doctor struct {
	Name  string
	Email string
	Tel   string
}

//Nurses struct
type Nurses struct {
	Nurse []Nurse
}

//Nurse struct
type Nurse struct {
	Name     string
	Email    string
	Password string
	Tel      string
}

//Patients struct
type Patients struct {
	Patient []Patient
}

//Patient struct
type Patient struct {
	Name string
}

//Drugs struct
type Drugs struct {
	Drug []Drug
}

//Drug struct
type Drug struct {
	Name string
}

// @title SUT SA Example API Prescription
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:se.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDoctorController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewDrugController(v1, client)
	controllers.NewNurseController(v1, client)
	controllers.NewPrescriptionController(v1, client)

	// Set Doctors Data
	doctors := Doctors{
		Doctor: []Doctor{
			Doctor{"นายแพทย์ สมศักดิ์ ชูใจ", "Gop.krongsin@hotmail.com", "0698742651"},
			Doctor{"แพทย์หญิง สมหญิง สมัครใจ", "Gopla7789@gmail.com", "0983549871"},
		},
	}

	for _, d := range doctors.Doctor {
		client.Doctor.
			Create().
			SetDoctorName(d.Name).
			SetDoctorEmail(d.Email).
			SetDoctorTel(d.Tel).
			Save(context.Background())
	}

	// Set Nurses Data
	nurses := Nurses{
		Nurse: []Nurse{
			Nurse{"นางสาว สายสมร   เพลินตา", "Saisamon_123@hotmail.com", "Gop069874", "0698742651"},
			Nurse{"นาง ตาหวาน   ตากลม", "tawan123456@gmail.com", "Gopla9871", "0983549871"},
		},
	}

	for _, n := range nurses.Nurse {
		client.Nurse.
			Create().
			SetNurseName(n.Name).
			SetNurseEmail(n.Email).
			SetNursePassword(n.Password).
			SetNurseTel(n.Tel).
			Save(context.Background())
	}

	// Set Patient Data
	patients := Patients{
		Patient: []Patient{
			Patient{"นาย สมชาย ใจดี"},
			Patient{"นางสาว สมหญิง สวยจัง"},
			Patient{"นาง สมร ดีใจ"},
		},
	}

	for _, p := range patients.Patient {

		client.Patient.
			Create().
			SetPatientName(p.Name).
			Save(context.Background())
	}

	// Set Drugs Data
	drugs := Drugs{
		Drug: []Drug{
			Drug{"Paracetamol"},
			Drug{"ยาธาตุน้ำขาว"},
			Drug{"Aspirin"},
			Drug{"ยาธาตุน้ำแดง"},
			Drug{"ยาแก้ไอน้ำดำ"},
			Drug{"ENO"},
			Drug{"Decolgen"},
		},
	}

	for _, dr := range drugs.Drug {

		client.Drug.
			Create().
			SetDrugName(dr.Name).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
