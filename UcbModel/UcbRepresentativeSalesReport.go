package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// RepresentativeSalesReport Info
type RepresentativeSalesReport struct {
	ID         		string        `json:"-"`
	Id_        		bson.ObjectId `json:"-" bson:"_id"`
	ResourceConfigID	string	`json:"-" bson:"resource-config-id"`
	GoodsConfigID	string  `json:"-" bson:"goods-config-id"`

	ResourceConfig		*ResourceConfig	`json:"-"`
	GoodsConfig 		*GoodsConfig `json:"-"`

	//RepresentativeName string `json:"representative-name" bson:"representative-name"`
	//ProductName		string `json:"product-name" bson:"product-name"`

	Potential		float64	`json:"potential" bson:"potential"`
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
	Contribution	float64 `json:"contribution" bson:"contribution"`
	PatientCount		int		`json:"patient-count" bson:"patient-count"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c RepresentativeSalesReport) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *RepresentativeSalesReport) SetID(id string) error {
	c.ID = id
	return nil
}


// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u RepresentativeSalesReport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "resourceConfigs",
			Name: "resourceConfig",
		},
		{
			Type: "goodsConfigs",
			Name: "goodsConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u RepresentativeSalesReport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	if u.ResourceConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ResourceConfigID,
			Type: "resourceConfigs",
			Name: "resourceConfig",
		})
	}

	if u.GoodsConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.GoodsConfigID,
			Type: "goodsConfigs",
			Name: "goodsConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u RepresentativeSalesReport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.ResourceConfigID != "" && u.ResourceConfig != nil {
		result = append(result, u.ResourceConfig)
	}

	if u.GoodsConfigID != "" && u.GoodsConfig != nil {
		result = append(result, u.GoodsConfig)
	}

	return result
}

func (u *RepresentativeSalesReport) SetToOneReferenceID(name, ID string) error {
	if name == "resourceConfig" {
		u.ResourceConfigID = ID
		return nil
	}
	if name == "goodsConfig" {
		u.GoodsConfigID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *RepresentativeSalesReport) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
