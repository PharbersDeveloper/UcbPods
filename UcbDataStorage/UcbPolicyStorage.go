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

// UcbPolicyStorage stores all of the tasty modelleaf, needs to be injected into
// Policy and Policy Resource. In the real world, you would use a database for that.
type UcbPolicyStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbPolicyStorage) NewPolicyStorage(args []BmDaemons.BmDaemon) *UcbPolicyStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbPolicyStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbPolicyStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Policy {
	in := UcbModel.Policy{}
	var out []*UcbModel.Policy
	err := s.db.FindMulti(r, &in, &out, skip, take)
	if err == nil {
		for i, iter := range out {
			s.db.ResetIdWithId_(iter)
			out[i] = iter
		}
		return out
	} else {
		return nil
	}
}

// GetOne tasty modelleaf
func (s UcbPolicyStorage) GetOne(id string) (UcbModel.Policy, error) {
	in := UcbModel.Policy{ID: id}
	out := UcbModel.Policy{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Policy for id %s not found", id)
	return UcbModel.Policy{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbPolicyStorage) Insert(c UcbModel.Policy) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbPolicyStorage) Delete(id string) error {
	in := UcbModel.Policy{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Policy with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbPolicyStorage) Update(c UcbModel.Policy) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Policy with id does not exist")
	}

	return nil
}
