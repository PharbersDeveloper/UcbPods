package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type ActionKpi struct {
	ID						string        `json:"-"`
	Id_						bson.ObjectId `json:"-" bson:"_id"`
	TargetNumber			float64	`json:"target-number" bson:"target-number"`
	TargetCoverage			float64	`json:"target-coverage" bson:"target-coverage"`
	HighLevelFrequency		float64	`json:"high-level-frequency" bson:"high-level-frequency"`
	MiddleLevelFrequency	float64	`json:"middle-level-frequency" bson:"middle-level-frequency"`
	LowLevelFrequency		float64	`json:"low-level-frequency" bson:"low-level-frequency"`
	RepresentativeID 	string          `json:"-" bson:"representative-id"`
	Representative   	*Representative `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (a ActionKpi) GetID() string {
	return a.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (a *ActionKpi) SetID(id string) error {
	a.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u ActionKpi) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "representatives",
			Name: "representative",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u ActionKpi) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.RepresentativeID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.RepresentativeID,
			Type: "representatives",
			Name: "representative",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u ActionKpi) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.RepresentativeID != "" && u.Representative != nil {
		result = append(result, u.Representative)
	}

	return result
}

func (u *ActionKpi) SetToOneReferenceID(name, ID string) error {
	if name == "representative" {
		u.RepresentativeID = ID
		return nil
	}
	return errors.New("There is no to-one relationship with the name " + name)
}

func (a *ActionKpi) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "ids":
			r := make(map[string]interface{})
			var ids []bson.ObjectId
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "scenario-id":
			rst[k] = v[0]
		}
	}
	return rst
}
