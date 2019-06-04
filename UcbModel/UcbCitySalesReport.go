package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type CitySalesReport struct {
	ID       string        `json:"-"`
	Id_      bson.ObjectId `json:"-" bson:"_id"`
	Sales			float64 `json:"sales" bson:"sales"`
	SalesQuota 		float64	`json:"sales-quota" bson:"sales-quota"`
	Share 			float64 `json:"share" bson:"share"`
	QuotaAchievement float64 `json:"quota-achievement" bson:"quota-achievement"`
	SalesGrowth		float64	`json:"sales-growth" bson:"sales-growth"`

	QuotaContribute float64 `json:"quota-contribute" bson:"quota-contribute"`
	QuotaGrowth		float64	`json:"quota-growth" bson:"quota-growth"`
	YTDSales		float64	`json:"ytd-sales" bson:"ytd-sales"`
	SalesContribute	float64	`json:"sales-contribute" bson:"sales-contribute"`
	SalesYearOnYear	float64	`json:"sales-year-on-year" bson:"sales-year-on-year"`
	SalesMonthOnMonth float64	`json:"sales-month-on-month" bson:"sales-month-on-month"`
	PatientCount	int		`json:"patient-count" bson:"patient-count"`

	CityId		string `json:"-" bson:"city-id"`
	City 		*City  `json:"-"`
}

func (c CitySalesReport) GetID() string {
	return c.ID
}

func (c CitySalesReport) SetID(id string) error {
	c.ID = id
	return nil
}

func (c CitySalesReport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "cities",
			Name: "city",
		},
	}
}

func (c CitySalesReport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if c.CityId != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   c.CityId,
			Type: "cities",
			Name: "city",
		})
	}

	return result
}

func (c CitySalesReport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if c.CityId != "" && c.City != nil {
		result = append(result, c.City)
	}
	return result
}

func (c *CitySalesReport) SetToOneReferenceID(name, ID string) error {
	if name == "city" {
		c.CityId = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (c *CitySalesReport) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
