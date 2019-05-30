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

type UcbManageTimeResultResource struct {
	UcbManageTimeResultStorage	*UcbDataStorage.UcbManageTimeResultStorage
	UcbAssessmentReportStorage			*UcbDataStorage.UcbAssessmentReportStorage
}

func (c UcbManageTimeResultResource) NewManageTimeResultResource(args []BmDataStorage.BmStorage) *UcbManageTimeResultResource {
	var rdr *UcbDataStorage.UcbManageTimeResultStorage
	var ar *UcbDataStorage.UcbAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbManageTimeResultStorage" {
			rdr = arg.(*UcbDataStorage.UcbManageTimeResultStorage)
		} else if tp.Name() == "UcbAssessmentReportStorage" {
			ar = arg.(*UcbDataStorage.UcbAssessmentReportStorage)
		}
	}
	return &UcbManageTimeResultResource{
		UcbManageTimeResultStorage: rdr,
		UcbAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c UcbManageTimeResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.UcbAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbManageTimeResultStorage.GetOne(modelRoot.ManageTimeResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.ManageTimeResult
	result = c.UcbManageTimeResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbManageTimeResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbManageTimeResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbManageTimeResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManageTimeResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbManageTimeResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbManageTimeResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbManageTimeResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbManageTimeResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ManageTimeResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbManageTimeResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
