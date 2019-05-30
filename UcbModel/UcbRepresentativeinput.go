package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Representativeinput struct {
	ID                       string        `json:"-"`
	Id_                      bson.ObjectId `json:"-" bson:"_id"`
	ProductKnowledgeTraining float64       `json:"product-knowledge-training" bson:"product-knowledge-training"`
	SalesAbilityTraining     float64       `json:"sales-ability-training" bson:"sales-ability-training"`
	RegionTraining           float64       `json:"region-training" bson:"region-training"`
	PerformanceTraining      float64       `json:"performance-training" bson:"performance-training"`
	VocationalDevelopment    float64       `json:"vocational-development" bson:"vocational-development"`
	AssistAccessTime         float64       `json:"assist-access-time" bson:"assist-access-time"`
	AbilityCoach             float64       `json:"ability-coach" bson:"ability-coach"`
	ResourceConfigId         string        `json:"resource-config-id" bson:"resource-config-id"`
	ResourceConfig 			 *ResourceConfig `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Representativeinput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Representativeinput) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Representativeinput) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "resourceConfigs",
			Name: "resourceConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Representativeinput) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	if u.ResourceConfigId != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ResourceConfigId,
			Type: "resourceConfigs",
			Name: "resourceConfig",
		})
	}


	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Representativeinput) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.ResourceConfigId != "" && u.ResourceConfig != nil {
		result = append(result, u.ResourceConfig)
	}


	return result
}

func (u *Representativeinput) SetToOneReferenceID(name, ID string) error {
	if name == "resourceConfig" {
		u.ResourceConfigId = ID
		return nil
	}
	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *Representativeinput) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	r := make(map[string]interface{})
	var ids []bson.ObjectId
	for k, v := range parameters {
		switch k {
		case "ids":
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		}
	}
	return rst
}
