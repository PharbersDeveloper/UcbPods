package UcbModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type UseableProposal struct {
	ID        string        `json:"-"`
	Id_       bson.ObjectId `json:"-" bson:"_id"`
	AccountID string        `json:"account-id" bson:"account-id"`

	ProposalID string   `json:"-" bson:"proposal-id"`
	Proposal   *Proposal `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u UseableProposal) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *UseableProposal) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u UseableProposal) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "proposals",
			Name: "proposal",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u UseableProposal) GetReferencedIDs() []jsonapi.ReferenceID {
	var result []jsonapi.ReferenceID

	if u.ProposalID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.ProposalID,
			Type: "proposals",
			Name: "proposal",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u UseableProposal) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	var result []jsonapi.MarshalIdentifier

	if u.ProposalID != "" && u.Proposal != nil {
		result = append(result, u.Proposal)
	}

	return result
}

func (u *UseableProposal) SetToOneReferenceID(name, ID string) error {
	if name == "proposal" {
		u.ProposalID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}

func (u *UseableProposal) GetConditionsBsonM(parameters map[string][]string) bson.M {
	rst := make(map[string]interface{})
	for k, v := range parameters {
		switch k {
		case "account-id":
			rst[k] = v[0]
		}
	}

	return rst
}
