package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type ProductConfig struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	ProductType      int `json:"product-type" bson:"product-type"`
	PriceType        string  `json:"price-type" bson:"price-type"`
	ReferencePrice   float64 `json:"reference-price" bson:"reference-price"`
	CostPrice        float64 `json:"cost-price" bson:"cost-price"`
	LifeCycle        string  `json:"life-cycle" bson:"life-cycle"`
	LaunchTime       float64 `json:"launch-time" bson:"launch-time"`
	TreatmentArea    string  `json:"treatment-area" bson:"treatment-area"`
	ProductFeature   string  `json:"product-feature" bson:"product-feature"`
	CostEffective    string  `json:"cost-effective" bson:"cost-effective"`
	Safety           string  `json:"safety" bson:"safety"`
	Effectiveness    string  `json:"effectiveness" bson:"effectiveness"`
	Convenience      string  `json:"convenience" bson:"convenience"`
	TargetDepartment string  `json:"target-department" bson:"target-department"`
	PatentDescribe	 string	 `json:"patent-describe" bson:"patent-describe"`
	ProductID string   `json:"-" bson:"product-id"`
	Product   *Product `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u ProductConfig) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *ProductConfig) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u ProductConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "products",
			Name: "product",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u ProductConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	if u.ProductID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ProductID,
			Type: "products",
			Name: "product",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u ProductConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var result []jsonapi.MarshalIdentifier

	if u.ProductID != "" && u.Product != nil {
		result = append(result, u.Product)
	}

	return result
}

func (u *ProductConfig) SetToOneReferenceID(name, ID string) error {
	if name == "product" {
		u.ProductID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *ProductConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "product-type":
			val, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				panic(err.Error())
			}
			rst[k] = val
		}
	}
	return rst
}
