package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type RegionConfig struct {
	ID  				string        `json:"-"`
	Id_ 				bson.ObjectId `json:"-" bson:"_id"`

	Population			float64 `json:"population" bson:"population"`
 	HealthSpending      float64  `json:"health-spending" bson:"health-spending"`

	RegionID 			string `json:"-" bson:"region-id"`
	Region   			*Region `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u RegionConfig) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *RegionConfig) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u RegionConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "regions",
			Name: "region",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u RegionConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.RegionID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.RegionID,
			Type: "regions",
			Name: "region",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u RegionConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.RegionID != "" && u.Region != nil {
		result = append(result, u.Region)
	}

	return result
}

func (u *RegionConfig) SetToOneReferenceID(name, ID string) error {
	if name == "region" {
		u.RegionID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *RegionConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
