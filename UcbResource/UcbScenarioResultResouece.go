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

type UcbScenarioResultResource struct {
	UcbScenarioResultStorage 		*UcbDataStorage.UcbScenarioResultStorage
	UcbSimplifyResultStorage 		*UcbDataStorage.UcbSimplifyResultStorage
}

func (c UcbScenarioResultResource) NewScenarioResultResource(args []BmDataStorage.BmStorage) *UcbScenarioResultResource {
	var cs *UcbDataStorage.UcbScenarioResultStorage
	var srs *UcbDataStorage.UcbSimplifyResultStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbScenarioResultStorage" {
			cs = arg.(*UcbDataStorage.UcbScenarioResultStorage)
		} else if tp.Name() == "UcbSimplifyResultStorage" {
			srs = arg.(*UcbDataStorage.UcbSimplifyResultStorage)
		}
	}
	return &UcbScenarioResultResource{
		UcbScenarioResultStorage: cs,
		UcbSimplifyResultStorage: srs,
	}
}

// FindAll ScenarioResults
func (c UcbScenarioResultResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	simplifyResultsID, sOk := r.QueryParams["simplifyResultsID"]

	if sOk {
		modelRootID := simplifyResultsID[0]
		modelRoot, err := c.UcbSimplifyResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		r.QueryParams["ids"] = modelRoot.ScenarioResultsIDs
		model := c.UcbScenarioResultStorage.GetAll(r, -1,-1)

		return &Response{Res: model}, nil
	}

	result := c.UcbScenarioResultStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbScenarioResultResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbScenarioResultStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbScenarioResultResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ScenarioResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbScenarioResultStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbScenarioResultResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbScenarioResultStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbScenarioResultResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.ScenarioResult)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbScenarioResultStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
