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

// UcbDepartmentStorage stores all of the tasty modelleaf, needs to be injected into
// Department and Department Resource. In the real world, you would use a database for that.
type UcbDepartmentStorage struct {
	Departments map[string]*UcbModel.Department
	idCount     int

	db *BmMongodb.BmMongodb
}

func (s UcbDepartmentStorage) NewDepartmentStorage(args []BmDaemons.BmDaemon) *UcbDepartmentStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbDepartmentStorage{make(map[string]*UcbModel.Department), 1, mdb}
}

// GetAll of the modelleaf
func (s UcbDepartmentStorage) GetAll(r api2go.Request, skip int, take int) []*UcbModel.Department {
	in := UcbModel.Department{}
	var out []*UcbModel.Department
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
func (s UcbDepartmentStorage) GetOne(id string) (UcbModel.Department, error) {
	in := UcbModel.Department{ID: id}
	out := UcbModel.Department{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("Department for id %s not found", id)
	return UcbModel.Department{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbDepartmentStorage) Insert(c UcbModel.Department) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbDepartmentStorage) Delete(id string) error {
	in := UcbModel.Department{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("Department with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbDepartmentStorage) Update(c UcbModel.Department) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("Department with id does not exist")
	}

	return nil
}
