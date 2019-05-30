package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// Proposal Info
type Proposal struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Describe   string        `json:"describe" bson:"describe"`
	TotalPhase int           `json:"total-phase" bson:"total-phase"`
	InputIDs   []string      `json:"-" bson:"input-ids"`
	SalesReportIDs  []string  `json:"-" bson:"sales-report-ids"`
	SalesReports 		[]*SalesReport `json:"-"`
	PersonnelAssessmentIDs  []string  `json:"-" bson:"personnel-assessment-ids"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Proposal) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Proposal) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Proposal) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "salesReports",
			Name: "salesReports",
		},
	}
}

func (c Proposal) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.SalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "salesReports",
			Name: "salesReports",
		})
	}

	return result
}

func (c Proposal) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.SalesReports {
		result = append(result, c.SalesReports[key])
	}
	return result
}

func (c *Proposal) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "salesReports" {
		c.SalesReportIDs = IDs
		return nil
	}
	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Proposal) AddToManyIDs(name string, IDs []string) error {

	if name == "salesReports" {
		c.SalesReportIDs = append(c.SalesReportIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *Proposal) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
