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

type UcbRegionalDivisionResultResource struct {
	UcbRegionalDivisionResultStorage	*UcbDataStorage.UcbRegionalDivisionResultStorage
	UcbAssessmentReportStorage			*UcbDataStorage.UcbAssessmentReportStorage
}

func (c UcbRegionalDivisionResultResource) NewRegionalDivisionResultResource(args []BmDataStorage.BmStorage) *UcbRegionalDivisionResultResource {
	var rdr *UcbDataStorage.UcbRegionalDivisionResultStorage
	var ar *UcbDataStorage.UcbAssessmentReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbRegionalDivisionResultStorage" {
			rdr = arg.(*UcbDataStorage.UcbRegionalDivisionResultStorage)
		} else if tp.Name() == "UcbAssessmentReportStorage" {
			ar = arg.(*UcbDataStorage.UcbAssessmentReportStorage)
		}
	}
	return &UcbRegionalDivisionResultResource{
		UcbRegionalDivisionResultStorage: rdr,
		UcbAssessmentReportStorage: ar,
	}
}

// FindAll images
func (c UcbRegionalDivisionResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	assessmentReportsID, arOk := r.QueryParams["assessmentReportsID"]

	if arOk {
		modelRootID := assessmentReportsID[0]
		modelRoot, err := c.UcbAssessmentReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbRegionalDivisionResultStorage.GetOne(modelRoot.RegionalDivisionResultID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.RegionalDivisionResult
	result = c.UcbRegionalDivisionResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbRegionalDivisionResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbRegionalDivisionResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbRegionalDivisionResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.RegionalDivisionResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbRegionalDivisionResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbRegionalDivisionResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbRegionalDivisionResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbRegionalDivisionResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.RegionalDivisionResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbRegionalDivisionResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
