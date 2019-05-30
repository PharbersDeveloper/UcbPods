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

type UcbTargetAssignsResultResource struct {
	UcbTargetAssignsResultStorage	*UcbDataStorage.UcbTargetAssignsResultStorage
	UcbAssessmentReportStorage			*UcbDataStorage.UcbAssessmentReportStorage
}

func (c UcbTargetAssignsResultResource) NewTargetAssignsResultResource(args []BmDataStorage.BmStorage) *UcbTargetAssignsResultResource {
	var rdr *UcbDataStorage.UcbTargetAssignsResultStorage
	var ar *UcbDataStorage.UcbAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbTargetAssignsResultStorage" {
			rdr = arg.(*UcbDataStorage.UcbTargetAssignsResultStorage)
		} else if tp.Name() == "UcbAssessmentReportStorage" {
			ar = arg.(*UcbDataStorage.UcbAssessmentReportStorage)
		}
	}
	return &UcbTargetAssignsResultResource{
		UcbTargetAssignsResultStorage: rdr,
		UcbAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c UcbTargetAssignsResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.UcbAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbTargetAssignsResultStorage.GetOne(modelRoot.TargetAssignsResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.TargetAssignsResult
	result = c.UcbTargetAssignsResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbTargetAssignsResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbTargetAssignsResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbTargetAssignsResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.TargetAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbTargetAssignsResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbTargetAssignsResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbTargetAssignsResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbTargetAssignsResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.TargetAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbTargetAssignsResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
