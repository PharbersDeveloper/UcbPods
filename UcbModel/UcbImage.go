package UcbModel

import "gopkg.in/mgo.v2/bson"

// Image Info
type Image struct {
	ID   string        `json:"-"`
	Id_  bson.ObjectId `json:"-" bson:"_id"`
	Img  string        `json:"img" bson:"img"`
	Tag  string        `json:"tag" bson:"tag"`
	Flag float64       `json:"flag" bson:"flag"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Image) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Image) SetID(id string) error {
	c.ID = id
	return nil
}

func (u *Image) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
