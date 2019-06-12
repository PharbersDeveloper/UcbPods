package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type SimplifyResult struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`

	LevelConfigID    string        `json:"-" bson:"level-config-id"`
	LevelConfig 	 *LevelConfig  `json:"-"`

	AssessmentReportDescribeIDs		[]string	`json:"-" bson:"assessment-report-describe-ids"`
	AssessmentReportDescribes		[]*AssessmentReportDescribe		`json:"-"`

	ScenarioResultsIDs		[]string	`json:"-" bson:"scenario-result-ids"`
	ScenarioResults			[]*ScenarioResult		`json:"-"`

	TotalQuotaAchievement	float64 	`json:"total-quota-achievement" bson:"total-quota-achievement"`
}

func (c SimplifyResult) GetID() string {
	return c.ID
}

func (c SimplifyResult) SetID(id string) error {
	c.ID = id
	return nil
}

func (c SimplifyResult) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "levelConfigs",
			Name: "levelConfig",
		},
		{
			Type: "assessmentReportDescribes",
			Name: "assessmentReportDescribes",
		},
		{
			Type: "scenarioResults",
			Name: "scenarioResults",
		},
	}
}

func (c SimplifyResult) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if c.LevelConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   c.LevelConfigID,
			Type: "levelConfigs",
			Name: "levelConfig",
		})
	}

	for _, kID := range c.AssessmentReportDescribeIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "assessmentReportDescribes",
			Name: "assessmentReportDescribes",
		})
	}

	for _, kID := range c.ScenarioResultsIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "scenarioResults",
			Name: "scenarioResults",
		})
	}

	return result
}

func (u *SimplifyResult) SetToOneReferenceID(name, ID string) error {
	if name == "levelConfig" {
		u.LevelConfigID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (c SimplifyResult) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if c.LevelConfigID != "" && c.LevelConfig != nil {
		result = append(result, c.LevelConfig)
	}

	for key := range c.AssessmentReportDescribes {
		result = append(result, c.AssessmentReportDescribes[key])
	}

	for key := range c.ScenarioResults {
		result = append(result, c.ScenarioResults[key])
	}

	return result
}

func (c *SimplifyResult) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "assessmentReportDescribes" {
		c.AssessmentReportDescribeIDs = IDs
		return nil
	}

	if name == "scenarioResults" {
		c.ScenarioResultsIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *SimplifyResult) AddToManyIDs(name string, IDs []string) error {
	if name == "assessmentReportDescribes" {
		c.AssessmentReportDescribeIDs = append(c.AssessmentReportDescribeIDs, IDs...)
		return nil
	}

	if name == "scenarioResults" {
		c.ScenarioResultsIDs = append(c.ScenarioResultsIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *SimplifyResult) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
