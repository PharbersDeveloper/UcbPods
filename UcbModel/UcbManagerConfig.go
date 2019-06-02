package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

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

	ManagerGoodsConfigIds	[]string `json:"-" bson:"manager-goods-config-ids"`
	ManagerGoodsConfigs		[]*ManagerGoodsConfig `json:"-"`
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

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u *ManagerConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "managerGoodsConfigs",
			Name: "managerGoodsConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (c *ManagerConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.ManagerGoodsConfigIds {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "managerGoodsConfigs",
			Name: "managerGoodsConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (c *ManagerConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.ManagerGoodsConfigs {
		result = append(result, c.ManagerGoodsConfigs[key])
	}

	return result
}

func (c *ManagerConfig) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "managerGoodsConfigs" {
		c.ManagerGoodsConfigIds = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new leafs that a users loves so much
func (c *ManagerConfig) AddToManyIDs(name string, IDs []string) error {
	if name == "managerGoodsConfigs" {
		c.ManagerGoodsConfigIds = append(c.ManagerGoodsConfigIds, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
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
