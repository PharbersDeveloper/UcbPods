package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// HospitalSalesReport Info
type HospitalSalesReport struct {
	ID         		string        `json:"-"`
	Id_        		bson.ObjectId `json:"-" bson:"_id"`
	DestConfigID	string	`json:"dest-config-id" bson:"dest-config-id" mapstructure:"hospital-id"`
	ResourceConfigID	string	`json:"-" bson:"resource-config-id" mapstructure:"representative-id"`
	GoodsConfigID	string  `json:"-" bson:"goods-config-id" mapstructure:"product-id"`

	//HospitalID	string	`json:"-" bson:"dest-config-id"`
	//ProductID	string	`json:"-" bson:"resource-config-id"`
	//RepresentativeID	string  `json:"-" bson:"goods-config-id"`

	DestConfig		*DestConfig	`json:"-"`
	GoodsConfig 	*GoodsConfig `json:"-"`
	ResourceConfig	*ResourceConfig	`json:"-"`

	//HospitalName 	string `json:"hospital-name" bson:"hospital-name"`
	//ProductName		string `json:"product-name" bson:"product-name"`

	Potential		float64	`json:"potential" bson:"potential" mapstructure:""`
	Sales			float64 `json:"sales" bson:"sales" mapstructure:"sales"`
	SalesQuota 		float64	`json:"sales-quota" bson:"sales-quota" mapstructure:"sales-quota"`
	Share 			float64 `json:"share" bson:"share" mapstructure:"share"`
	QuotaAchievement float64 `json:"quota-achievement" bson:"quota-achievement" mapstructure:"quota-achievement"`
	SalesGrowth		float64	`json:"sales-growth" bson:"sales-growth" mapstructure:"sales-growth"`

	QuotaContribute float64 `json:"quota-contribute" bson:"quota-contribute" mapstructure:"quota-contribute"`
	QuotaGrowth		float64	`json:"quota-growth" bson:"quota-growth" mapstructure:"quota-growth"`
	YTDSales		float64	`json:"ytd-sales" bson:"ytd-sales" mapstructure:"ytd-sales"`
	SalesContribute	float64	`json:"sales-contribute" bson:"sales-contribute" mapstructure:"sales-contribute"`
	SalesYearOnYear	float64	`json:"sales-year-on-year" bson:"sales-year-on-year" mapstructure:"sales-year-on-year"`
	SalesMonthOnMonth float64	`json:"sales-month-on-month" bson:"sales-month-on-month" mapstructure:"sales-month-on-month"`
	DrugEntranceInfo  string	`json:"drug-entrance-info" bson:"drug-entrance-info" mapstructure:"drug-entrance-info"`
	PatientCount		int		`json:"patient-count" bson:"patient-count" mapstructure:"patient-count"`

}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c HospitalSalesReport) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *HospitalSalesReport) SetID(id string) error {
	c.ID = id
	return nil
}


// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u HospitalSalesReport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "destConfigs",
			Name: "destConfig",
		},
		{
			Type: "goodsConfigs",
			Name: "goodsConfig",
		},
		{
			Type: "resourceConfigs",
			Name: "resourceConfig",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u HospitalSalesReport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	if u.DestConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.DestConfigID,
			Type: "destConfigs",
			Name: "destConfig",
		})
	}

	if u.GoodsConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.GoodsConfigID,
			Type: "goodsConfigs",
			Name: "goodsConfig",
		})
	}

	if u.ResourceConfigID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ResourceConfigID,
			Type: "resourceConfigs",
			Name: "resourceConfig",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u HospitalSalesReport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.DestConfigID != "" && u.DestConfig != nil {
		result = append(result, u.DestConfig)
	}

	if u.GoodsConfigID != "" && u.GoodsConfig != nil {
		result = append(result, u.GoodsConfig)
	}

	if u.ResourceConfigID != "" && u.ResourceConfig != nil {
		result = append(result, u.ResourceConfig)
	}

	return result
}

func (u *HospitalSalesReport) SetToOneReferenceID(name, ID string) error {
	if name == "DestConfig" {
		u.DestConfigID = ID
		return nil
	}
	if name == "goodsConfig" {
		u.GoodsConfigID = ID
		return nil
	}
	if name == "resourceConfig" {
		u.ResourceConfigID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *HospitalSalesReport) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "notEq[destConfigId]":
			r := map[string]interface{}{}
			r["$ne"] = v[0]
			rst["dest-config-id"] = r
		}
	}

	return rst
}
