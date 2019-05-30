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

type UcbRepresentativeAbilityResource struct {
	UcbRepresentativeAbilityStorage *UcbDataStorage.UcbRepresentativeAbilityStorage
	UcbPersonnelAssessmentStorage	*UcbDataStorage.UcbPersonnelAssessmentStorage
}

func (s UcbRepresentativeAbilityResource) NewRepresentativeAbilityResource(args []BmDataStorage.BmStorage) *UcbRepresentativeAbilityResource {
	var ras *UcbDataStorage.UcbRepresentativeAbilityStorage
	var pas *UcbDataStorage.UcbPersonnelAssessmentStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbRepresentativeAbilityStorage" {
			ras = arg.(*UcbDataStorage.UcbRepresentativeAbilityStorage)
		} else if tp.Name() == "UcbPersonnelAssessmentStorage" {
			pas = arg.(*UcbDataStorage.UcbPersonnelAssessmentStorage)
		}
	}
	return &UcbRepresentativeAbilityResource{
		UcbRepresentativeAbilityStorage: ras,
		UcbPersonnelAssessmentStorage:   pas,
	}
}

func (s UcbRepresentativeAbilityResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.RepresentativeAbility

	personnelAssessmentsID, pasok := r.QueryParams["personnelAssessmentsID"]

	if pasok {
		modelRootID := personnelAssessmentsID[0]

		modelRoot, err := s.UcbPersonnelAssessmentStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		r.QueryParams["ids"] = modelRoot.RepresentativeAbilityIDs

		result := s.UcbRepresentativeAbilityStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := s.UcbRepresentativeAbilityStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbRepresentativeAbilityResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.UcbRepresentativeAbilityStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeAbilityResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.RepresentativeAbility)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbRepresentativeAbilityStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeAbilityResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbRepresentativeAbilityStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbRepresentativeAbilityResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.RepresentativeAbility)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbRepresentativeAbilityStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
