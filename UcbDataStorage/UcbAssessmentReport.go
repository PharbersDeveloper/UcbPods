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

// UcbAssessmentReportStorage stores all of the tasty modelleaf, needs to be injected into
// AssessmentReport and AssessmentReport Resource. In the real world, you would use a database for that.
type UcbAssessmentReportStorage struct {
	db *BmMongodb.BmMongodb
}

func (s UcbAssessmentReportStorage) NewAssessmentReportStorage(args []BmDaemons.BmDaemon) *UcbAssessmentReportStorage {
	mdb := args[0].(*BmMongodb.BmMongodb)
	return &UcbAssessmentReportStorage{mdb}
}

// GetAll of the modelleaf
func (s UcbAssessmentReportStorage) GetAll(r api2go.Request, skip int, take int) []UcbModel.AssessmentReport {
	in := UcbModel.AssessmentReport{}
	var out []UcbModel.AssessmentReport
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
func (s UcbAssessmentReportStorage) GetOne(id string) (UcbModel.AssessmentReport, error) {
	in := UcbModel.AssessmentReport{ID: id}
	out := UcbModel.AssessmentReport{ID: id}
	err := s.db.FindOne(&in, &out)
	if err == nil {
		return out, nil
	}
	errMessage := fmt.Sprintf("AssessmentReport for id %s not found", id)
	return UcbModel.AssessmentReport{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
}

// Insert a fresh one
func (s *UcbAssessmentReportStorage) Insert(c UcbModel.AssessmentReport) string {
	tmp, err := s.db.InsertBmObject(&c)
	if err != nil {
		fmt.Println(err)
	}

	return tmp
}

// Delete one :(
func (s *UcbAssessmentReportStorage) Delete(id string) error {
	in := UcbModel.AssessmentReport{ID: id}
	err := s.db.Delete(&in)
	if err != nil {
		return fmt.Errorf("AssessmentReport with id %s does not exist", id)
	}

	return nil
}

// Update updates an existing modelleaf
func (s *UcbAssessmentReportStorage) Update(c UcbModel.AssessmentReport) error {
	err := s.db.Update(&c)
	if err != nil {
		return fmt.Errorf("AssessmentReport with id does not exist")
	}

	return nil
}
