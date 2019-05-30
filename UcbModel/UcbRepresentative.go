package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Representative struct {
	ID     string        `json:"-"`
	Id_    bson.ObjectId `json:"-" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender float64       `json:"gender" bson:"gender"`
	// 0 => 女 ; 1 => 男

	ImagesIDs []string `json:"-" bson:"image-ids"`
	Imgs      []*Image `json:"-"`
}

func (c Representative) GetID() string {
	return c.ID
}

func (c Representative) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Representative) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "images",
			Name: "images",
		},
	}
}

func (c Representative) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.ImagesIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "images",
			Name: "images",
		})
	}
	return result
}

func (c Representative) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.Imgs {
		result = append(result, c.Imgs[key])
	}
	return result
}

func (c *Representative) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImagesIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Representative) AddToManyIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImagesIDs = append(c.ImagesIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Representative) DeleteToManyIDs(name string, IDs []string) error {
	if name == "images" {
		for _, ID := range IDs {
			for pos, oldID := range c.ImagesIDs {
				if ID == oldID {
					c.ImagesIDs = append(c.ImagesIDs[:pos], c.ImagesIDs[pos+1:]...)
				}
			}
		}
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Representative) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
