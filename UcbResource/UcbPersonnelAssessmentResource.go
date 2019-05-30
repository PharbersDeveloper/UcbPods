package UcbResource

import (
	"errors"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
)

type UcbPersonnelAssessmentResource struct {
	UcbPersonnelAssessmentStorage 			*UcbDataStorage.UcbPersonnelAssessmentStorage
	UcbPaperStorage							*UcbDataStorage.UcbPaperStorage
}

func (s UcbPersonnelAssessmentResource) NewPersonnelAssessmentResource(args []BmDataStorage.BmStorage) *UcbPersonnelAssessmentResource {
	var pas *UcbDataStorage.UcbPersonnelAssessmentStorage
	var ps  *UcbDataStorage.UcbPaperStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbPersonnelAssessmentStorage" {
			pas = arg.(*UcbDataStorage.UcbPersonnelAssessmentStorage)
		} else if tp.Name() == "UcbPaperStorage" {
			ps = arg.(*UcbDataStorage.UcbPaperStorage)
		}
	}
	return &UcbPersonnelAssessmentResource{
		UcbPersonnelAssessmentStorage:		pas,
		UcbPaperStorage: 					ps,
	}
}

func (s UcbPersonnelAssessmentResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.PersonnelAssessment
	papersID, pok := r.QueryParams["papersID"]

	if pok {
		modelRootID := papersID[0]
		modelRoot, err := s.UcbPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.PersonnelAssessmentIDs

		result := s.UcbPersonnelAssessmentStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := s.UcbPersonnelAssessmentStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbPersonnelAssessmentResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.UcbPersonnelAssessmentStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbPersonnelAssessmentResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.PersonnelAssessment)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbPersonnelAssessmentStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbPersonnelAssessmentResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbPersonnelAssessmentStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbPersonnelAssessmentResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.PersonnelAssessment)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbPersonnelAssessmentStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
