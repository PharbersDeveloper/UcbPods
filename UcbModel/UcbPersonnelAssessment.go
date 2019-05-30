package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type PersonnelAssessment struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	Time		float64	`json:"time" bson:"time"`

	ScenarioID	string	`json:"-" bson:"scenario-id"`
	Scenario	*Scenario `json:"-"`

	RepresentativeAbilityIDs    []string      `json:"-" bson:"representative-ability-ids"`
	RepresentativeAbility 		[]*RepresentativeAbility `json:"-"`

	ActionKpiIDs		[]string			`json:"-" bson:"action-kpi-ids"`
	ActionKpi 			[]*ActionKpi		`json:"-"`
}

func (c PersonnelAssessment) GetID() string {
	return c.ID
}

func (c PersonnelAssessment) SetID(id string) error {
	c.ID = id
	return nil
}

func (c PersonnelAssessment) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "representativeAbilities",
			Name: "representativeAbilities",
		},
		{
			Type: "actionKpis",
			Name: "actionKpis",
		},
		{
			Type: "scenarios",
			Name: "scenario",
		},
	}
}

func (c PersonnelAssessment) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.RepresentativeAbilityIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "representativeAbilities",
			Name: "representativeAbilities",
		})
	}

	for _, kID := range  c.ActionKpiIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "actionKpis",
			Name: "actionKpis",
		})
	}

	if c.ScenarioID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: c.ScenarioID,
			Type: "scenarios",
			Name: "scenario",
		})
	}

	return result
}

func (c PersonnelAssessment) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.RepresentativeAbility {
		result = append(result, c.RepresentativeAbility[key])
	}


	for key := range c.ActionKpi {
		result = append(result, c.ActionKpi[key])
	}

	if c.ScenarioID != "" && c.Scenario != nil {
		result = append(result, c.Scenario)
	}

	return result
}

func (c *PersonnelAssessment) SetToOneReferenceID(name, ID string) error {
	if name == "scenario" {
		c.ScenarioID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (c *PersonnelAssessment) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "representativeAbilities" {
		c.RepresentativeAbilityIDs = IDs
		return nil
	}

	if name == "actionKpis" {
		c. ActionKpiIDs= IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *PersonnelAssessment) AddToManyIDs(name string, IDs []string) error {
	if name == "representativeAbilities" {
		c.RepresentativeAbilityIDs = append(c.RepresentativeAbilityIDs, IDs...)
		return nil
	}

	if name == "actionKpis" {
		c.ActionKpiIDs = append(c.ActionKpiIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *PersonnelAssessment) DeleteToManyIDs(name string, IDs []string) error {
	if name == "representativeAbilities" {
		for _, ID := range IDs {
			for pos, oldID := range c.RepresentativeAbilityIDs {
				if ID == oldID {
					c.RepresentativeAbilityIDs = append(c.RepresentativeAbilityIDs[:pos], c.RepresentativeAbilityIDs[pos+1:]...)
				}
			}
		}
	}

	if name == "actionKpis" {
		for _, ID := range IDs {
			for pos, oldID := range c.ActionKpiIDs {
				if ID == oldID {
					c.ActionKpiIDs = append(c.ActionKpiIDs[:pos], c.ActionKpiIDs[pos+1:]...)
				}
			}
		}
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *PersonnelAssessment) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "proposal-id":
			rst[k] = v[0]
		}
	}
	return rst
}
