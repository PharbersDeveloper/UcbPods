package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Goodsinput struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	SalesTarget      float64       `json:"sales-target" bson:"sales-target"`
	Budget           float64       `json:"budget" bson:"budget"`

	GoodsConfigId 	 string		   `json:"goods-config-id" bson:"goods-config-id"`
	GoodsConfig 	 *GoodsConfig	`json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Goodsinput) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Goodsinput) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Goodsinput) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{

		{
			Type: "goodsConfigs",
			Name: "goodsConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Goodsinput) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.GoodsConfigId != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.GoodsConfigId,
			Type: "goodsConfigs",
			Name: "goodsConfig",
		})
	}


	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Goodsinput) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.GoodsConfigId != "" && u.GoodsConfig != nil {
		result = append(result, u.GoodsConfig)
	}


	return result
}

func (u *Goodsinput) SetToOneReferenceID(name, ID string) error {

	if name == "goodsConfig" {
		u.GoodsConfigId = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *Goodsinput) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		}
	}
	return rst
}
