package UcbResource

import (
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"errors"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type UcbSimplifyResultResource struct {
	UcbSimplifyResultStorage 		*UcbDataStorage.UcbSimplifyResultStorage
	UcbAssessmentReportStorage		*UcbDataStorage.UcbAssessmentReportStorage
}

func (c UcbSimplifyResultResource) NewSimplifyResultResource(args []BmDataStorage.BmStorage) *UcbSimplifyResultResource {
	var cs *UcbDataStorage.UcbSimplifyResultStorage
	var ars *UcbDataStorage.UcbAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbSimplifyResultStorage" {
			cs = arg.(*UcbDataStorage.UcbSimplifyResultStorage)
		} else if tp.Name() == "UcbAssessmentReportStorage" {
			ars = arg.(*UcbDataStorage.UcbAssessmentReportStorage)
		}
	}
	return &UcbSimplifyResultResource{
		UcbSimplifyResultStorage: cs,
		UcbAssessmentReportStorage: ars,
	}
}

// FindAll SimplifyResults
func (c UcbSimplifyResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	assessmentReportsID, aOk := r.QueryParams["assessmentReportsID"]

	if aOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.UcbAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, _ := c.UcbSimplifyResultStorage.GetOne(modelRoot.SimplifyResultID)
		return &Response{Res: model}, nil
	}

	result := c.UcbSimplifyResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbSimplifyResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbSimplifyResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbSimplifyResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.SimplifyResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbSimplifyResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbSimplifyResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbSimplifyResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbSimplifyResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.SimplifyResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbSimplifyResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
