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

type UcbEvaluationResource struct {
	UcbEvaluationStorage	*UcbDataStorage.UcbEvaluationStorage
	UcbLevelConfigStorage 		*UcbDataStorage.UcbLevelConfigStorage
}

func (c UcbEvaluationResource) NewEvaluationResource(args []BmDataStorage.BmStorage) *UcbEvaluationResource {
	var rdr *UcbDataStorage.UcbEvaluationStorage
	var lcs *UcbDataStorage.UcbLevelConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbEvaluationStorage" {
			rdr = arg.(*UcbDataStorage.UcbEvaluationStorage)
		} else if tp.Name() == "UcbLevelConfigStorage" {
			lcs = arg.(*UcbDataStorage.UcbLevelConfigStorage)
		}
	}
	return &UcbEvaluationResource{
		UcbEvaluationStorage: rdr,
		UcbLevelConfigStorage: lcs,
	}
}

// FindAll images
func (c UcbEvaluationResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	levelConfigsID, lcsOk := r.QueryParams["levelConfigsID"]

	if lcsOk {
		modelRootID := levelConfigsID[0]
		modelRoot, err := c.UcbLevelConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbEvaluationStorage.GetOne(modelRoot.EvaluationID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.Evaluation
	result = c.UcbEvaluationStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbEvaluationResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbEvaluationStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbEvaluationResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Evaluation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbEvaluationStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbEvaluationResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbEvaluationStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbEvaluationResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Evaluation)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbEvaluationStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
