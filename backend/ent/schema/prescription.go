package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Prescription holds the schema definition for the Prescription entity.
type Prescription struct {
	ent.Schema
}

// Fields of the Prescription.
func (Prescription) Fields() []ent.Field {

	return []ent.Field{

		field.String("Prescrip_Note"),

		field.Time("Prescrip_DateTime").Default(time.Now),
	}
}

// Edges of the Prescription.
func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("doctor", Doctor.Type).Ref("doctor_prescription").Unique(),
		edge.From("patient", Patient.Type).Ref("patient_prescription").Unique(),
		edge.From("nurse", Nurse.Type).Ref("nurse_prescription").Unique(),
		edge.From("drug", Drug.Type).Ref("drug_prescription").Unique(),
	}
}
