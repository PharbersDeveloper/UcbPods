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

type UcbResourceAssignsResultResource struct {
	UcbResourceAssignsResultStorage	*UcbDataStorage.UcbResourceAssignsResultStorage
	UcbAssessmentReportStorage			*UcbDataStorage.UcbAssessmentReportStorage
}

func (c UcbResourceAssignsResultResource) NewResourceAssignsResultResource(args []BmDataStorage.BmStorage) *UcbResourceAssignsResultResource {
	var rdr *UcbDataStorage.UcbResourceAssignsResultStorage
	var ar *UcbDataStorage.UcbAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbResourceAssignsResultStorage" {
			rdr = arg.(*UcbDataStorage.UcbResourceAssignsResultStorage)
		} else if tp.Name() == "UcbAssessmentReportStorage" {
			ar = arg.(*UcbDataStorage.UcbAssessmentReportStorage)
		}
	}
	return &UcbResourceAssignsResultResource{
		UcbResourceAssignsResultStorage: rdr,
		UcbAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c UcbResourceAssignsResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.UcbAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbResourceAssignsResultStorage.GetOne(modelRoot.ResourceAssignsResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.ResourceAssignsResult
	result = c.UcbResourceAssignsResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbResourceAssignsResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbResourceAssignsResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbResourceAssignsResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ResourceAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbResourceAssignsResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbResourceAssignsResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbResourceAssignsResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbResourceAssignsResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ResourceAssignsResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbResourceAssignsResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
