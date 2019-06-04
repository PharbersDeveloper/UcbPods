package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// SalesReport Info
type SalesReport struct {
	ID         						string        `json:"-"`
	Id_        						bson.ObjectId `json:"-" bson:"_id"`
	ScenarioID						string `json:"-" bson:"scenario-id"`
	HospitalSalesReportIDs			[]string	`json:"-" bson:"hospital-sales-report-ids"`
	RepresentativeSalesReportIDs	[]string  	`json:"-" bson:"representative-sales-report-ids"`
	ProductSalesReportIDs			[]string	`json:"-" bson:"product-sales-report-ids"`
	CitySalesReportIDs				[]string	`json:"-" bson:"city-sales-report-ids"`

	HospitalSalesReport 		[]*HospitalSalesReport			`json:"-"`
	RepresentativeSalesReport	[]*RepresentativeSalesReport	`json:"-"`
	ProductSalesReport			[]*ProductSalesReport			`json:"-"`
	CitySalesReport				[]*CitySalesReport				`json:"-"`
	Scenario					*Scenario `json:"-"`

	Time 						float64 `json:"time" bson:"time"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c SalesReport) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *SalesReport) SetID(id string) error {
	c.ID = id
	return nil
}


// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u SalesReport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "hospitalSalesReports",
			Name: "hospitalSalesReports",
		},
		{
			Type: "representativeSalesReports",
			Name: "representativeSalesReports",
		},
		{
			Type: "productSalesReports",
			Name: "productSalesReports",
		},
		{
			Type: "citySalesReports",
			Name: "citySalesReports",
		},
		{
			Type: "scenarios",
			Name: "scenario",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u SalesReport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range u.HospitalSalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "hospitalSalesReports",
			Name: "hospitalSalesReports",
		})
	}


	for _, kID := range u.RepresentativeSalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "representativeSalesReports",
			Name: "representativeSalesReports",
		})
	}

	for _, kID := range u.ProductSalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "productSalesReports",
			Name: "productSalesReports",
		})
	}

	for _, kID := range u.CitySalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "citySalesReports",
			Name: "citySalesReports",
		})
	}

	if u.ScenarioID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.ScenarioID,
			Type: "scenarios",
			Name: "scenario",
		})
	}


	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u SalesReport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range u.HospitalSalesReport {
		result = append(result, u.HospitalSalesReport[key])
	}

	for key := range u.RepresentativeSalesReport {
		result = append(result, u.RepresentativeSalesReport[key])
	}

	for key := range u.ProductSalesReport {
		result = append(result, u.ProductSalesReport[key])
	}

	for key := range u.CitySalesReport {
		result = append(result, u.CitySalesReport[key])
	}

	if u.ScenarioID != "" && u.Scenario != nil {
		result = append(result, u.Scenario)
	}

	return result
}

func (c *SalesReport) SetToOneReferenceID(name, ID string) error {
	if name == "scenario" {
		c.ScenarioID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *SalesReport) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "hospitalSalesReports" {
		u.HospitalSalesReportIDs = IDs
		return nil
	} else if name == "representativeSalesReports" {
		u.RepresentativeSalesReportIDs = IDs
		return nil
	} else if name == "productSalesReports" {
		u.ProductSalesReportIDs = IDs
		return nil
	} else if name == "citySalesReports" {
		u.CitySalesReportIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *SalesReport) AddToManyIDs(name string, IDs []string) error {
	if name == "hospitalSalesReports" {
		u.HospitalSalesReportIDs = append(u.HospitalSalesReportIDs, IDs...)
		return nil
	} else if name == "representativeSalesReports" {
		u.RepresentativeSalesReportIDs = append(u.RepresentativeSalesReportIDs, IDs...)
		return nil
	} else if name == "productSalesReports" {
		u.ProductSalesReportIDs = append(u.ProductSalesReportIDs, IDs...)
		return nil
	} else if name == "citySalesReports" {
		u.CitySalesReportIDs = append(u.CitySalesReportIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}


func (u *SalesReport) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "scenario-id":
			rst[k] = v[0]
		}
	}

	return rst
}
