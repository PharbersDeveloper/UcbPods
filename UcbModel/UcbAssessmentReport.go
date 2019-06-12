package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// AssessmentReport Info
type AssessmentReport struct {
	ID         						string        `json:"-"`
	Id_        						bson.ObjectId `json:"-" bson:"_id"`

	RegionalDivisionResultID		string	`json:"-" bson:"regional-division-result-id"`
	RegionalDivisionResult			*RegionalDivisionResult `json:"-"`

	TargetAssignsResultID			string  `json:"-" bson:"target-assigns-result-id"`
	TargetAssignsResult				*TargetAssignsResult	`json:"-"`

	ResourceAssignsResultID			string	`json:"-" bson:"resource-assigns-result-id"`
	ResourceAssignsResult			*ResourceAssignsResult	`json:"-"`

	ManageTimeResultID				string	`json:"-" bson:"manage-time-result-id"`
	ManageTimeResult				*ManageTimeResult	`json:"-"`

	ManageTeamResultID				string	`json:"-" bson:"manage-team-result-id"`
	ManageTeamResult				*ManageTeamResult	`json:"-"`

	GeneralPerformanceResultID		string	`json:"-" bson:"general-performance-id"`
	GeneralPerformanceResult		*GeneralPerformanceResult	`json:"-"`

	SimplifyResultID				string 	`json:"-" bson:"simplify-result-id"`
	SimplifyResult					*SimplifyResult	`json:"-"`

	ScenarioID						string 	`json:"-" bson:"scenario-id"`
	Scenario						*Scenario `json:"-"`

	PaperInputID					string `json:"-" bson:"paper-input-id"`

	Time 							int64 `json:"time" bson:"time"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c AssessmentReport) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *AssessmentReport) SetID(id string) error {
	c.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u AssessmentReport) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "regionalDivisionResults",
			Name: "regionalDivisionResult",
		},
		{
			Type: "targetAssignsResults",
			Name: "targetAssignsResult",
		},
		{
			Type: "resourceAssignsResults",
			Name: "resourceAssignsResult",
		},
		{
			Type: "manageTimeResults",
			Name: "manageTimeResult",
		},
		{
			Type: "manageTeamResults",
			Name: "manageTeamResult",
		},
		{
			Type: "generalPerformanceResults",
			Name: "generalPerformanceResult",
		},
		{
			Type: "simplifyResults",
			Name: "simplifyResult",
		},
		{
			Type: "scenarios",
			Name: "scenario",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u AssessmentReport) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.RegionalDivisionResultID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.RegionalDivisionResultID,
			Type: "regionalDivisionResults",
			Name: "regionalDivisionResult",
		})
	}

	if u.TargetAssignsResultID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.TargetAssignsResultID,
			Type: "targetAssignsResults",
			Name: "targetAssignsResult",
		})
	}

	if u.ResourceAssignsResultID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.ResourceAssignsResultID,
			Type: "resourceAssignsResults",
			Name: "resourceAssignsResult",
		})
	}

	if u.ManageTimeResultID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.ManageTimeResultID,
			Type: "manageTimeResults",
			Name: "manageTimeResult",
		})
	}

	if u.ManageTeamResultID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.ManageTeamResultID,
			Type: "manageTeamResults",
			Name: "manageTeamResult",
		})
	}

	if u.GeneralPerformanceResultID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.GeneralPerformanceResultID,
			Type: "generalPerformanceResults",
			Name: "generalPerformanceResult",
		})
	}
	if u.SimplifyResultID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.SimplifyResultID,
			Type: "simplifyResults",
			Name: "simplifyResult",
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
func (u AssessmentReport) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.RegionalDivisionResultID != "" && u.RegionalDivisionResult != nil {
		result = append(result, u.RegionalDivisionResult)
	}

	if u.TargetAssignsResultID != "" && u.TargetAssignsResult != nil {
		result = append(result, u.TargetAssignsResult)
	}

	if u.ResourceAssignsResultID != "" && u.ResourceAssignsResult != nil {
		result = append(result, u.ResourceAssignsResult)
	}

	if u.ManageTimeResultID != "" && u.ManageTimeResult != nil {
		result = append(result, u.ManageTimeResult)
	}

	if u.ManageTeamResultID != "" && u.ManageTeamResult != nil {
		result = append(result, u.ManageTeamResult)
	}

	if u.GeneralPerformanceResultID != "" && u.GeneralPerformanceResult != nil {
		result = append(result, u.GeneralPerformanceResult)
	}

	if u.SimplifyResultID != "" && u.SimplifyResult != nil {
		result = append(result, u.SimplifyResult)
	}

	if u.ScenarioID != "" && u.Scenario != nil {
		result = append(result, u.Scenario)
	}

	return result
}

func (c *AssessmentReport) SetToOneReferenceID(name, ID string) error {
	if name == "regionalDivisionResult" {
		c.ResourceAssignsResultID = ID
		return nil
	}

	if name == "targetAssignsResult" {
		c.TargetAssignsResultID = ID
		return nil
	}

	if name == "resourceAssignsResult" {
		c.ResourceAssignsResultID = ID
		return nil
	}

	if name == "manageTimeResult" {
		c.ManageTimeResultID = ID
		return nil
	}

	if name == "manageTeamResult" {
		c.ManageTimeResultID = ID
		return nil
	}

	if name == "generalPerformanceResult" {
		c.GeneralPerformanceResultID = ID
		return nil
	}

	if name == "simplifyResult" {
		c.SimplifyResultID = ID
		return nil
	}

	if name == "scenario" {
		c.ScenarioID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *AssessmentReport) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
