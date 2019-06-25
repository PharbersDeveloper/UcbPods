package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type ScenarioResult struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	QuotaAchievement float64        `json:"quota-achievement" bson:"quota-achievement" mapstructure:"quota-achievement"`

	ScenarioID          string        `json:"scenario-id" bson:"scenario-id" mapstructure:"scenario-id"`
	Scenario 			*Scenario			`json:"-"`
}

func (c ScenarioResult) GetID() string {
	return c.ID
}

func (c ScenarioResult) SetID(id string) error {
	c.ID = id
	return nil
}

func (c ScenarioResult) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "scenarios",
			Name: "scenario",
		},
	}
}

func (c ScenarioResult) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if c.ScenarioID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   c.ScenarioID,
			Type: "scenarios",
			Name: "scenario",
		})
	}

	return result
}

func (u *ScenarioResult) SetToOneReferenceID(name, ID string) error {
	if name == "scenario" {
		u.ScenarioID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (c ScenarioResult) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if c.ScenarioID != "" && c.Scenario != nil {
		result = append(result, c.Scenario)
	}
	return result
}

func (c *ScenarioResult) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
