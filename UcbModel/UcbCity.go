package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type City struct {
	ID       				string        `json:"-"`
	Id_      				bson.ObjectId `json:"-" bson:"_id"`
	Name     				string        `json:"name" bson:"name"`
	Level	 				string        `json:"level" bson:"level"`
	Type	 				string        `json:"type" bson:"type"`
	LocalPatientRatio 		float64       `json:"local-patient-ratio" bson:"local-patient-ratio"`
	NonLocalPatientRatio 	float64       `json:"nonlocal-patient-ratio" bson:"nonlocal-patient-ratio"`



	HospitalConfigIds 		[]string 			`json:"-" bson:"City-config-ids"`
	HospitalConfigs			[]*HospitalConfig 	`json:"-"`
	ImageIds 				[]string			`json:"-" bson:"image-ids"`
	Images 					[]*Image 			`json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c City) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *City) SetID(id string) error {
	c.ID = id
	return nil
}

func (c City) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "images",
			Name: "images",
		},
		{
			Type: "hospitalConfigs",
			Name: "hospitalConfigs",
		},
	}
}

func (c City) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.ImageIds {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "images",
			Name: "images",
		})
	}

	for _, kID := range c.HospitalConfigIds {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "hospitalConfigs",
			Name: "hospitalConfigs",
		})
	}
	return result
}

func (c City) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.Images {
		result = append(result, c.Images[key])
	}

	for key := range c.HospitalConfigs {
		result = append(result, c.HospitalConfigs[key])
	}
	return result
}

func (c *City) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImageIds = IDs
		return nil
	}

	if name == "hospitalConfigs" {
		c.HospitalConfigIds = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *City) AddToManyIDs(name string, IDs []string) error {
	if name == "images" {
		c.ImageIds = append(c.ImageIds, IDs...)
		return nil
	}

	if name == "hospitalConfigs" {
		c.HospitalConfigIds = append(c.HospitalConfigIds, IDs...)
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}


func (u *City) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
