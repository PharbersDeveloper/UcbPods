package UcbDataStorage

import (
	"errors"
	"fmt"
	"Ucb/UcbModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/manyminds/api2go"
	"net/http"
)

type UcbUseableProposalStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbUseableProposalStorage) NewUseableProposalStorage(args []BmDaemons.BmDaemon) *UcbUseableProposalStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbUseableProposalStorage{mdb}
}

func (s UcbUseableProposalStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.UseableProposal {
	in := UcbModel.UseableProposal{}
	var out []UcbModel.UseableProposal
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		var tmp []*UcbModel.UseableProposal
		for i := 0; i < len(out); i++ {
			ptr := out[i]
			s.db.ResetIdWithId_(&ptr)
			tmp = append(tmp, &ptr)
		}
		return tmp
	} else {
		return nil
	}
}

func (s UcbUseableProposalStorage) GetOne(id string) (UcbModel.UseableProposal, error) {
	in := UcbModel.UseableProposal{ID: id}
	out := UcbModel.UseableProposal{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("UseableProposal for id %s not found", id)
	return UcbModel.UseableProposal{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

func (s *UcbUseableProposalStorage) Insert(c UcbModel.UseableProposal) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbUseableProposalStorage) Delete(id string) error {
	in := UcbModel.UseableProposal{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("UseableProposal with id %s does not exist", id)
	}

	return nil
}

// Update a model
func (s *UcbUseableProposalStorage) Update(c UcbModel.UseableProposal) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("UseableProposal with id does not exist")
	}

	return nil
}

func (s *UcbUseableProposalStorage) Count(req api2go.Request, c UcbModel.UseableProposal) int {
	r, _ := s.db.Count(req, &c)
	return r
}
