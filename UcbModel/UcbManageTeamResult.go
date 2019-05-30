package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type ManageTeamResult struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`

	LevelConfigID    string        `json:"-" bson:"level-config-id"`
	LevelConfig 	 *LevelConfig  `json:"-"`

	AssessmentReportDescribeIDs		[]string	`json:"-" bson:"assessment-report-describe-ids"`
	AssessmentReportDescribes		[]*AssessmentReportDescribe		`json:"-"`
}

func (c ManageTeamResult) GetID() string {
	return c.ID
}

func (c ManageTeamResult) SetID(id string) error {
	c.ID = id
	return nil
}

func (c ManageTeamResult) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "levelConfigs",
			Name: "levelConfig",
		},
		{
			Type: "assessmentReportDescribes",
			Name: "assessmentReportDescribes",
		},
	}
}

func (c ManageTeamResult) GetReferencedIDs() []jsonapi.ReferenceID {
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

	return result
}

func (u *ManageTeamResult) SetToOneReferenceID(name, ID string) error {
	if name == "levelConfig" {
		u.LevelConfigID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (c ManageTeamResult) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if c.LevelConfigID != "" && c.LevelConfig != nil {
		result = append(result, c.LevelConfig)
	}

	for key := range c.AssessmentReportDescribes {
		result = append(result, c.AssessmentReportDescribes[key])
	}

	return result
}

func (c *ManageTeamResult) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "assessmentReportDescribes" {
		c.AssessmentReportDescribeIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *ManageTeamResult) AddToManyIDs(name string, IDs []string) error {
	if name == "assessmentReportDescribes" {
		c.AssessmentReportDescribeIDs = append(c.AssessmentReportDescribeIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *ManageTeamResult) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
