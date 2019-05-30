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

// UcbAssessmentReportDescribeStorage stores all of the tasty modelleaf, needs to be injected into
// AssessmentReportDescribe and AssessmentReportDescribe Resource. In the real world, you would use a database for that.
type UcbAssessmentReportDescribeStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbAssessmentReportDescribeStorage) NewAssessmentReportDescribeStorage(args []BmDaemons.BmDaemon) *UcbAssessmentReportDescribeStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbAssessmentReportDescribeStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbAssessmentReportDescribeStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.AssessmentReportDescribe {
	in := UcbModel.AssessmentReportDescribe{}
	var out []UcbModel.AssessmentReportDescribe
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
func (s UcbAssessmentReportDescribeStorage) GetOne(id string) (UcbModel.AssessmentReportDescribe, error) {
	in := UcbModel.AssessmentReportDescribe{ID: id}
	out := UcbModel.AssessmentReportDescribe{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("AssessmentReportDescribe for id %s not found", id)
	return UcbModel.AssessmentReportDescribe{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbAssessmentReportDescribeStorage) Insert(c UcbModel.AssessmentReportDescribe) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbAssessmentReportDescribeStorage) Delete(id string) error {
	in := UcbModel.AssessmentReportDescribe{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("AssessmentReportDescribe with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbAssessmentReportDescribeStorage) Update(c UcbModel.AssessmentReportDescribe) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("AssessmentReportDescribe with id does not exist")
	}

	return nil
}
