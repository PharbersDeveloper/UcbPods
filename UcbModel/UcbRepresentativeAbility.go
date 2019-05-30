package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type RepresentativeAbility struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	ProductKnowledge          float64 `json:"product-knowledge" bson:"product-knowledge"`
	SalesAbility              float64 `json:"sales-ability" bson:"sales-ability"`
	RegionalManagementAbility float64 `json:"regional-management-ability" bson:"regional-management-ability"`
	JobEnthusiasm             float64 `json:"job-enthusiasm" bson:"job-enthusiasm"`
	BehaviorValidity          float64 `json:"behavior-validity" bson:"behavior-validity"`

	RepresentativeID 	string          `json:"-" bson:"representative-id"`
	Representative   	*Representative `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u RepresentativeAbility) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *RepresentativeAbility) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u RepresentativeAbility) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "representatives",
			Name: "representative",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u RepresentativeAbility) GetReferencedIDs() []jsonapi.ReferenceID {
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
func (u RepresentativeAbility) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.RepresentativeID != "" && u.Representative != nil {
		result = append(result, u.Representative)
	}

	return result
}

func (u *RepresentativeAbility) SetToOneReferenceID(name, ID string) error {
	if name == "representative" {
		u.RepresentativeID = ID
		return nil
	}
	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *RepresentativeAbility) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		}
	}
	return rst
}
