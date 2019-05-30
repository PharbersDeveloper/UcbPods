package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Title struct {
	ID               string        `json:"-"`
	Id_              bson.ObjectId `json:"-" bson:"_id"`
	Value            string        `json:"value" bson:"value"`
	Describe         string        `json:"describe" bson:"describe"`
	ImagesID 		 string 	   `json:"-" bson:"image-id"`
	Img      		 *Image 	   `json:"-"`
}

func (c Title) GetID() string {
	return c.ID
}

func (c Title) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Title) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "images",
			Name: "image",
		},
	}
}

func (c Title) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if c.ImagesID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   c.ImagesID,
			Type: "images",
			Name: "image",
		})
	}

	return result
}

func (c *Title) SetToOneReferenceID(name, ID string) error {
	if name == "image" {
		c.ImagesID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (c Title) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if c.ImagesID != "" && c.Img != nil {
		result = append(result, c.Img)
	}
	return result
}



func (c *Title) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
