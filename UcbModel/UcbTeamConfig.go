package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// TeamConfig Info
type TeamConfig struct {
	ID         					string        	`json:"-"`
	Id_        					bson.ObjectId	`json:"-" bson:"_id"`
	ScenarioId 					string 		  	`json:"scenario-id" bson:"scenario-id"`
	BusinessExperience 			string			`json:"business-experience" bson:"business-experience"`
	Describe 					string			`json:"describe" bson:"describe"`
	ProductKnowledge 			float64			`json:"product-knowledge" bson:"product-knowledge"`
	SalesAbility 				float64			`json:"sales-ability" bson:"sales-ability"`
	RegionalManagementAbility	float64			`json:"regional-management-ability" bson:"regional-management-ability"`
	JobEnthusiasm 				float64			`json:"job-enthusiasm" bson:"job-enthusiasm"`
	BehaviorValidity 			float64			`json:"behavior-validity" bson:"behavior-validity"`

	ResourceConfigIDs			[]string		`json:"-" bson:"resource-config-ids"`
	ResourceConfig				[]*ResourceConfig `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c TeamConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *TeamConfig) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u TeamConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "resourceConfigs",
			Name: "resourceConfigs",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u TeamConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range u.ResourceConfigIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "resourceConfigs",
			Name: "resourceConfigs",
		})
	}
	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u TeamConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range u.ResourceConfig {
		result = append(result, u.ResourceConfig[key])
	}

	return result
}

func (u *TeamConfig) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "resourceConfigs" {
		u.ResourceConfigIDs = IDs
		return nil
	}
	return errors.New("There is no to-one relationship with the name " + name)
}

func (c *TeamConfig) AddToManyIDs(name string, IDs []string) error {
	if name == "resourceConfigs" {
		c.ResourceConfigIDs = append(c.ResourceConfigIDs, IDs...)
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *TeamConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
