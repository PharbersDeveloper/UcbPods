package UcbModel

import "gopkg.in/mgo.v2/bson"

// ManagerConfig Info
type ManagerConfig struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	TotalBusinessIndicators int32 `json:"total-business-indicators" bson:"total-business-indicators"`
	TotalBudgets             int32 `json:"total-budgets" bson:"total-budgets"`
	TotalMeetingPlaces      int32 `json:"total-meeting-places" bson:"total-meeting-places"`
	ManagerKPI              int32 `json:"manager-kpi" bson:"manager-kpi"`
	ManagerTime             int32 `json:"manager-time" bson:"manager-time"`
	VisitTotalTime          int32 `json:"visit-total-time" bson:"visit-total-time"`
	TeamBusinessExperience 	string `json:"team-business-experience" bson:"team-business-experience"`
	TeamDescribe 			string `json:"team-describe" bson:"team-describe"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ManagerConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ManagerConfig) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *ManagerConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
