package UcbDataStorage

import (
	"fmt"
	"errors"
	"Ucb/UcbModel"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
)

// UcbProposalStorage stores all of the tasty modelleaf, needs to be injected into
// Proposal and Proposal Resource. In the real world, you would use a database for that.
type UcbProposalStorage struct {
	Proposal  map[string]*UcbModel.Proposal
	idCount int

	db *BmMongodb.BmMongodb
}

func (s UcbProposalStorage) NewProposalStorage(args []BmDaemons.BmDaemon) *UcbProposalStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbProposalStorage{make(map[string]*UcbModel.Proposal), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbProposalStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.Proposal {
	in := UcbModel.Proposal{}
	var out []UcbModel.Proposal
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(&iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s UcbProposalStorage) GetOne(id string) (UcbModel.Proposal, error) {
	in := UcbModel.Proposal{ID: id}
	out := UcbModel.Proposal{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Proposal for id %s not found", id)
	return UcbModel.Proposal{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbProposalStorage) Insert(c UcbModel.Proposal) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbProposalStorage) Delete(id string) error {
	in := UcbModel.Proposal{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Proposal with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbProposalStorage) Update(c UcbModel.Proposal) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Proposal with id does not exist")
	}

	return nil
}
