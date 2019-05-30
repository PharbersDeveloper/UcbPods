package UcbModel

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

// Scenario Info
type Scenario struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	ProposalID string        `json:"proposal-id" bson:"proposal-id"`
	Phase      int           `json:"phase" bson:"phase"`
	Name	   string		 `json:"name" bson:"name"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Scenario) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Scenario) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Scenario) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "proposal-id":
			rst[k] = v[0]
		//case "account-id":
		//	rst[k] = v[0]
		case "phase":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			rst[k] = val
		}
	}

	return rst
}
