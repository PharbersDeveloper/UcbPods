package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Region struct {
	ID       string        `json:"-"`
	Id_      bson.ObjectId `json:"-" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Describe string        `json:"describe" bson:"describe"`

	ImagesIDs []string `json:"-" bson:"image-ids"`
	Imgs      []*Image `json:"-"`
}

func (c Region) GetID() string {
	return c.ID
}

func (c Region) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Region) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "images",
			Name: "images",
		},
	}
}

func (c Region) GetReferencedIDs() []jsonapi.ReferenceID {
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

func (c Region) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.Imgs {
		result = append(result, c.Imgs[key])
	}
	return result
}

func (c *Region) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImagesIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Region) AddToManyIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImagesIDs = append(c.ImagesIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Region) DeleteToManyIDs(name string, IDs []string) error {
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

func (c *Region) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
