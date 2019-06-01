package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Businessinput struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	MeetingPlaces    float64       `json:"meeting-places" bson:"meeting-places"`
	VisitTime        float64       `json:"visit-time" bson:"visit-time"`

	DestConfigId     string         `json:"dest-config-id" bson:"dest-config-id"`
	DestConfig 		 *DestConfig 	 `json:"-"`
	ResourceConfigId string          `json:"resource-config-id" bson:"resource-config-id"`
	ResourceConfig   *ResourceConfig `json:"-"`
	GoodsInputIds	 []string		 `json:"-" bson:"goods-input-id"`
	GoodsInputs		 []*Goodsinput	 `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Businessinput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Businessinput) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Businessinput) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "destConfigs",
			Name: "destConfig",
		},
		{
			Type: "resourceConfigs",
			Name: "resourceConfig",
		},
		{
			Type: "goodsinputs",
			Name: "goodsinputs",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Businessinput) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	if u.DestConfigId != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.DestConfigId,
			Type: "destConfigs",
			Name: "destConfig",
		})
	}

	if u.ResourceConfigId != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ResourceConfigId,
			Type: "resourceConfigs",
			Name: "resourceConfig",
		})
	}

	for _, kID := range u.GoodsInputIds {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "goodsinputs",
			Name: "goodsinputs",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Businessinput) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.DestConfigId != "" && u.DestConfig != nil {
		result = append(result, u.ResourceConfig)
	}

	if u.ResourceConfigId != "" && u.ResourceConfig != nil {
		result = append(result, u.ResourceConfig)
	}

	for key := range u.GoodsInputs {
		result = append(result, u.GoodsInputs[key])
	}

	return result
}

func (u *Businessinput) SetToOneReferenceID(name, ID string) error {
	if name == "resourceConfig" {
		u.ResourceConfigId = ID
		return nil
	}

	if name == "destConfig" {
		u.DestConfigId = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *Businessinput) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "goodsinputs" {
		u.GoodsInputIds = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *Businessinput) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	r := make(map[string]interface{})
	var ids []bson.ObjectId
	for k, v := range parameters {
		switch k {
		case "ids":
			for i := 0; i < len(v); i++ {
				ids = append(ids, bson.ObjectIdHex(v[i]))
			}
			r["$in"] = ids
			rst["_id"] = r
		case "dest-config-id":
			rst[k] = v[0]
		case "resource-config-id":
			rst[k] = v[0]
		}
	}
	return rst
}
