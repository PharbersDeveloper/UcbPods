package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type HospitalConfig struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	DoctorNumber  int     `json:"doctor-number" bson:"doctor-number"`
	BedNumber     int     `json:"bed-number" bson:"bed-number"`
	Income        float64 `json:"income" bson:"income"`
	SpaceBelongs  string  `json:"space-belongs" bson:"space-belongs"`
	Ability2Pay   string  `json:"ability-to-pay" bson:"ability-to-pay"`

	Hospital   *Hospital `json:"-"`
	HospitalID string    `json:"-" bson:"hospital-id"`

	DepartmentIDs []string      `json:"-" bson:"department-ids"`
	Departments   []*Department `json:"-"`

	PolicyIDs []string  `json:"-" bson:"policy-ids"`
	Policies  []*Policy `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u HospitalConfig) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *HospitalConfig) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u HospitalConfig) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "hospitals",
			Name: "hospital",
		},
		{
			Type: "policies",
			Name: "policies",
		},
		{
			Type: "departments",
			Name: "departments",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u HospitalConfig) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	for _, kID := range u.PolicyIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "policies",
			Name: "policies",
		})
	}

	for _, kID := range u.DepartmentIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "departments",
			Name: "departments",
		})
	}

	if u.HospitalID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.HospitalID,
			Type: "hospitals",
			Name: "hospital",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u HospitalConfig) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	for key := range u.Policies {
		result = append(result, u.Policies[key])
	}

	for key := range u.Departments {
		result = append(result, u.Departments[key])
	}

	if u.HospitalID != "" && u.Hospital != nil {
		result = append(result, u.Hospital)
	}

	return result
}

func (u *HospitalConfig) SetToOneReferenceID(name, ID string) error {
	if name == "hospital" {
		u.HospitalID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *HospitalConfig) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "policies" {
		u.PolicyIDs = IDs
		return nil
	}
	if name == "departments" {
		u.DepartmentIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new leafs that a users loves so much
func (u *HospitalConfig) AddToManyIDs(name string, IDs []string) error {
	if name == "policies" {
		u.PolicyIDs = append(u.PolicyIDs, IDs...)
		return nil
	}
	if name == "departments" {
		u.DepartmentIDs = append(u.DepartmentIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some leafs from a users because they made him very sick
func (u *HospitalConfig) DeleteToManyIDs(name string, IDs []string) error {
	if name == "policies" {
		for _, ID := range IDs {
			for pos, oldID := range u.PolicyIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.PolicyIDs = append(u.PolicyIDs[:pos], u.PolicyIDs[pos+1:]...)
				}
			}
		}
	}

	if name == "departments" {
		for _, ID := range IDs {
			for pos, oldID := range u.DepartmentIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.DepartmentIDs = append(u.DepartmentIDs[:pos], u.DepartmentIDs[pos+1:]...)
				}
			}
		}
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *HospitalConfig) GetConditionsBsonM(parameters map[string][]string) bson.M {
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
