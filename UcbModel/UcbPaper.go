package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Paper struct {
	ID         string        `json:"-"`
	Id_        bson.ObjectId `json:"-" bson:"_id"`
	AccountID  string        `json:"account-id" bson:"account-id"`
	ProposalID string        `json:"proposal-id" bson:"proposal-id"`
	Name       string        `json:"name" bson:"name"`
	Describe   string        `json:"describe" bson:"describe"`
	StartTime  int64         `json:"start-time" bson:"start-time"`
	EndTime    int64         `json:"end-time" bson:"end-time"`
	TotalPhase	int 		 `json:"total-phase" bson:"total-phase"`
	Time		int64			`json:"time" bson:"time"`

	// 0 关卡未开始 1 关卡内正在做周期，未执最后行计算 2 关卡部分周期做完，新的周期未开始 3 关卡内的所有周期结束
	InputState int       `json:"state" bson:"input-state"`

	InputIDs    []string      `json:"-" bson:"input-ids"`
	Paperinputs []*Paperinput `json:"-"`

	SalesReportIDs    	[]string       `json:"-" bson:"sales-report-ids"`
	SalesReports 		[]*SalesReport `json:"-"`

	PersonnelAssessmentIDs		[]string					`json:"-" bson:"personnel-assessment-ids"`
	PersonnelAssessments 		[]*PersonnelAssessment		`json:"-"`

	AssessmentReportIDs		[]string	`json:"-" bson:"assessment-report-ids"`
	AssessmentReports		[]*AssessmentReport	`json:"-"`
}

func (c Paper) GetID() string {
	return c.ID
}

func (c Paper) SetID(id string) error {
	c.ID = id
	return nil
}

func (c Paper) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "paperinputs",
			Name: "paperinputs",
		},
		{
			Type: "salesReports",
			Name: "salesReports",
		},
		{
			Type: "personnelAssessments",
			Name: "personnelAssessments",
		},
		{
			Type: "assessmentReports",
			Name: "assessmentReports",
		},
	}
}

func (c Paper) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range c.InputIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "paperinputs",
			Name: "paperinputs",
		})
	}

	for _, kID := range c.SalesReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "salesReports",
			Name: "salesReports",
		})
	}

	for _, kID := range  c.PersonnelAssessmentIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "personnelAssessments",
			Name: "personnelAssessments",
		})
	}

	for _, kID := range  c.AssessmentReportIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "assessmentReports",
			Name: "assessmentReports",
		})
	}

	return result
}

func (c Paper) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range c.Paperinputs {
		result = append(result, c.Paperinputs[key])
	}

	for key := range c.SalesReports {
		result = append(result, c.SalesReports[key])
	}

	for key := range c.PersonnelAssessments {
		result = append(result, c.PersonnelAssessments[key])
	}

	for key := range c.AssessmentReports {
		result = append(result, c.AssessmentReports[key])
	}

	return result
}

func (c *Paper) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "paperinputs" {
		c.InputIDs = IDs
		return nil
	}

	if name == "salesReports" {
		c.SalesReportIDs = IDs
		return nil
	}

	if name == "personnelAssessments" {
		c. PersonnelAssessmentIDs= IDs
		return nil
	}

	if name == "assessmentReports" {
		c. AssessmentReportIDs= IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Paper) AddToManyIDs(name string, IDs []string) error {
	if name == "paperinputs" {
		c.InputIDs = append(c.InputIDs, IDs...)
		return nil
	}

	if name == "salesReports" {
		c.SalesReportIDs = append(c.SalesReportIDs, IDs...)
		return nil
	}

	if name == "personnelAssessments" {
		c.PersonnelAssessmentIDs = append(c.PersonnelAssessmentIDs, IDs...)
		return nil
	}

	if name == "assessmentReports" {
		c.AssessmentReportIDs = append(c.AssessmentReportIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (c *Paper) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
		case "proposal-id":
			rst[k] = v[0]
		case "account-id":
			rst[k] = v[0]
		}
	}
	return rst
}
