package UcbResource

import (
	"Ucb/UcbDataStorage"
	"errors"
	"Ucb/UcbModel"
	"reflect"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type UcbActionKpiResource struct {
	UcbActionKpiStorage 			*UcbDataStorage.UcbActionKpiStorage
	UcbPersonnelAssessmentStorage	*UcbDataStorage.UcbPersonnelAssessmentStorage
}

func (c UcbActionKpiResource) NewActionKpiResource(args []BmDataStorage.BmStorage) *UcbActionKpiResource {
	var cs *UcbDataStorage.UcbActionKpiStorage
	var pas *UcbDataStorage.UcbPersonnelAssessmentStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbActionKpiStorage" {
			cs = arg.(*UcbDataStorage.UcbActionKpiStorage)
		} else if tp.Name() == "UcbPersonnelAssessmentStorage" {
			pas = arg.(*UcbDataStorage.UcbPersonnelAssessmentStorage)
		}
	}
	return &UcbActionKpiResource{
		UcbActionKpiStorage: cs,
		UcbPersonnelAssessmentStorage: pas,
	}
}

// FindAll ActionKpis
func (c UcbActionKpiResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*UcbModel.ActionKpi
	personnelAssessmentsID, pasok := r.QueryParams["personnelAssessmentsID"]

	if pasok {
		modelRootID := personnelAssessmentsID[0]

		modelRoot, err := c.UcbPersonnelAssessmentStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		r.QueryParams["ids"] = modelRoot.ActionKpiIDs

		result = c.UcbActionKpiStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	result = c.UcbActionKpiStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbActionKpiResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbActionKpiStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbActionKpiResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ActionKpi)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbActionKpiStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbActionKpiResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbActionKpiStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbActionKpiResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ActionKpi)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbActionKpiStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
