package UcbModel

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
	"errors"
	"strconv"
)

// ResourceConfig Info
type ResourceConfig struct {
	ID           string        `json:"-"`
	Id_          bson.ObjectId `json:"-" bson:"_id"`
	ScenarioId   string        `json:"scenario-id" bson:"scenario-id"`
	ResourceType float64       `json:"resource-type" bson:"resource-type"`
	// 0 => ManagerConfig ; 1 => RepresentativeConfig
	ResourceID string `json:"resource-id" bson:"resource-id"`

	ManagerConfig        *ManagerConfig        `json:"-"`
	RepresentativeConfig *RepresentativeConfig `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ResourceConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ResourceConfig) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u ResourceConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "managerConfigs",
			Name: "managerConfig",
		},
		{
			Type: "representativeConfigs",
			Name: "representativeConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u ResourceConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.ResourceType == 0 {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ResourceID,
			Type: "managerConfigs",
			Name: "managerConfig",
		})
	} else if u.ResourceType == 1 {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ResourceID,
			Type: "representativeConfigs",
			Name: "representativeConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u ResourceConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.ResourceType == 0 && u.ManagerConfig != nil {
		result = append(result, u.ManagerConfig)
	} else if u.ResourceType == 1 && u.RepresentativeConfig != nil {
		result = append(result, u.RepresentativeConfig)
	}

	return result
}

func (u *ResourceConfig) SetToOneReferenceID(name, ID string) error {
	if name == "managerConfig" {
		u.ResourceID = ID
		return nil
	}
	if name == "representativeConfig" {
		u.ResourceID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *ResourceConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "resource-type":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			rst[k] = val
		case "status":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			rst[k] = val
		case "course-type":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			rst[k] = val
		case "lt[create-time]":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			r := make(map[string]interface{})
			r["$lt"] = val
			rst["create-time"] = r
		case "lte[create-time]":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			r := make(map[string]interface{})
			r["$lte"] = val
			rst["create-time"] = r
		case "gt[apply-time]":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			r := make(map[string]interface{})
			r["$gt"] = val
			rst["apply-time"] = r
		case "gte[apply-time]":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			r := make(map[string]interface{})
			r["$gte"] = val
			rst["apply-time"] = r
		case "ne[course-type]":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			r := make(map[string]interface{})
			r["$ne"] = val
			rst["course-type"] = r
		}
	}

	return rst
}
