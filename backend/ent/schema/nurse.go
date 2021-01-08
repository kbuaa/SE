package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Nurse holds the schema definition for the Nurse entity.
type Nurse struct {
	ent.Schema
}

// Fields of the Nurse.
func (Nurse) Fields() []ent.Field {

	return []ent.Field{

		field.String("Nurse_Email").NotEmpty(),

		field.String("Nurse_Name").NotEmpty(),

		field.String("Nurse_Tel").NotEmpty(),

		field.String("Nurse_Password").NotEmpty(),
	}
}

// Edges of the Nurse.
func (Nurse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nurse_prescription", Prescription.Type).StorageKey(edge.Column("nurse_id")),
	}
}
