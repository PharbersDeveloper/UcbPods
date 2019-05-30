package UcbModel

import (
	"gopkg.in/mgo.v2/bson"
)

type AssessmentReportDescribe struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	Text	         string        `json:"text" bson:"text"`
	Code             int           `json:"code" bson:"code"`
}

func (c AssessmentReportDescribe) GetID() string {
	return c.ID
}

func (c AssessmentReportDescribe) SetID(id string) error {
	c.ID = id
	return nil
}

func (c *AssessmentReportDescribe) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
