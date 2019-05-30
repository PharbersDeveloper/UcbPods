package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Hospital struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	Name             string        `json:"name" bson:"name"`
	Describe         string        `json:"describe" bson:"describe"`
	Code             string        `json:"code" bson:"code"`
	Regtime          string        `json:"regtime" bson:"regtime"`
	HospitalCategory string        `json:"hospital-category" bson:"hospital-category"`
	HospitalLevel    string        `json:"hospital-level" bson:"hospital-level"`
	Position         string        `json:"position" bson:"position"`

	ImagesIDs []string `json:"-" bson:"image-ids"`
	Imgs      []*Image `json:"-"`
}

func (c Hospital) GetID() string {
	return c.ID
}

func (c Hospital) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Hospital) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "images",
			Name: "images",
		},
	}
}

func (c Hospital) GetReferencedIDs() []jsonapi.ReferenceID {
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

func (c Hospital) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.Imgs {
		result = append(result, c.Imgs[key])
	}
	return result
}

func (c *Hospital) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImagesIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Hospital) AddToManyIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImagesIDs = append(c.ImagesIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Hospital) DeleteToManyIDs(name string, IDs []string) error {
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

func (c *Hospital) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
