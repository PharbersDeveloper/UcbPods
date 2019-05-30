package UcbResource

import (
	"errors"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type UcbManageTeamResultResource struct {
	UcbManageTeamResultStorage	*UcbDataStorage.UcbManageTeamResultStorage
	UcbAssessmentReportStorage			*UcbDataStorage.UcbAssessmentReportStorage
}

func (c UcbManageTeamResultResource) NewManageTeamResultResource(args []BmDataStorage.BmStorage) *UcbManageTeamResultResource {
	var rdr *UcbDataStorage.UcbManageTeamResultStorage
	var ar *UcbDataStorage.UcbAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbManageTeamResultStorage" {
			rdr = arg.(*UcbDataStorage.UcbManageTeamResultStorage)
		} else if tp.Name() == "UcbAssessmentReportStorage" {
			ar = arg.(*UcbDataStorage.UcbAssessmentReportStorage)
		}
	}
	return &UcbManageTeamResultResource{
		UcbManageTeamResultStorage: rdr,
		UcbAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c UcbManageTeamResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.UcbAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbManageTeamResultStorage.GetOne(modelRoot.ManageTeamResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.ManageTeamResult
	result = c.UcbManageTeamResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbManageTeamResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbManageTeamResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbManageTeamResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManageTeamResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbManageTeamResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbManageTeamResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbManageTeamResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbManageTeamResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManageTeamResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbManageTeamResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
