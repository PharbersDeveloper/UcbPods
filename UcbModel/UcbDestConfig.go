package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

// DestConfig Info
type DestConfig struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	ScenarioId string        `json:"scenario-id" bson:"scenario-id"`
	DestType   float64       `json:"dest-type" bson:"dest-type"`
	// 0 => RegionConfig; 1 => HospitalConfig
	DestID     string        `json:"dest-id" bson:"dest-id"`

	RegionConfig   *RegionConfig   `json:"-"`
	HospitalConfig *HospitalConfig `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c DestConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *DestConfig) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u DestConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "regionConfigs",
			Name: "regionConfig",
		},
		{
			Type: "hospitalConfigs",
			Name: "hospitalConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u DestConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.DestType == 0 {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.DestID,
			Type: "regionConfigs",
			Name: "regionConfig",
		})
	} else if u.DestType == 1 {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.DestID,
			Type: "hospitalConfigs",
			Name: "hospitalConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u DestConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.DestType == 0 && u.RegionConfig != nil{
		result = append(result, u.RegionConfig)
	} else if u.DestType == 1 && u.HospitalConfig != nil {
		result = append(result, u.HospitalConfig)
	}

	return result
}

func (u *DestConfig) SetToOneReferenceID(name, ID string) error {
	if name == "regionConfig" {
		u.DestID = ID
		return nil
	}
	if name == "hospitalConfig" {
		u.DestID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *DestConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "dest-type":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			rst[k] = val
		}
	}

	return rst
}
