package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type LevelConfig struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	Text	         string        `json:"text" bson:"text"`
	Code             int        `json:"code" bson:"code"`
	LevelID          string        `json:"level-id" bson:"level-id"`
	Level 			*Level			`json:"-"`
	EvaluationID			string	`json:"-" bson:"evaluation-id"`
	Evaluation			*Evaluation	`json:"-"`
}

func (c LevelConfig) GetID() string {
	return c.ID
}

func (c LevelConfig) SetID(id string) error {
	c.ID = id
	return nil
}

func (c LevelConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "levels",
			Name: "level",
		},
		{
			Type: "evaluations",
			Name: "evaluation",
		},
	}
}

func (c LevelConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if c.LevelID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   c.LevelID,
			Type: "levels",
			Name: "level",
		})
	}

	if c.EvaluationID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   c.EvaluationID,
			Type: "evaluations",
			Name: "evaluation",
		})
	}

	return result
}

func (u *LevelConfig) SetToOneReferenceID(name, ID string) error {
	if name == "level" {
		u.LevelID = ID
		return nil
	}

	if name == "evaluation" {
		u.EvaluationID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (c LevelConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if c.LevelID != "" && c.Level != nil {
		result = append(result, c.Level)
	}

	if c.EvaluationID != "" && c.Evaluation != nil {
		result = append(result, c.Evaluation)
	}

	return result
}

func (c *LevelConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "level-id":
			rst[k] = v[0]
		case "code":
			v, _ := strconv.Atoi(v[0])
			rst[k] = v
		}
	}
	return rst
}
