package UcbModel

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/manyminds/api2go/jsonapi"
	"errors"
	"strconv"
)

// GoodsConfig Info
type GoodsConfig struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	ScenarioId string        `json:"scenario-id" bson:"scenario-id"`
	GoodsType  float64       `json:"goods-type" bson:"goods-type"`
	// 0 => ProductConfig ;
	GoodsID string `json:"goods-id" bson:"goods-id"`

	ProductConfig *ProductConfig `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c GoodsConfig) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *GoodsConfig) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u GoodsConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "productConfigs",
			Name: "productConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u GoodsConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	if u.GoodsType == 0 {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.GoodsID,
			Type: "productConfigs",
			Name: "productConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u GoodsConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var result []jsonapi.MarshalIdentifier

	if u.GoodsType == 0 && u.ProductConfig != nil {
		result = append(result, u.ProductConfig)
	}

	return result
}

func (u *GoodsConfig) SetToOneReferenceID(name, ID string) error {
	if name == "productConfig" {
		u.GoodsID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *GoodsConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "goods-type":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			rst[k] = val
		}
	}

	return rst
}
