package UcbResource

import (
	"errors"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type UcbPaperinputResource struct {
	UcbPaperinputStorage          *UcbDataStorage.UcbPaperinputStorage
	UcbBusinessinputStorage       *UcbDataStorage.UcbBusinessinputStorage
	UcbRepresentativeinputStorage *UcbDataStorage.UcbRepresentativeinputStorage
	UcbManagerinputStorage        *UcbDataStorage.UcbManagerinputStorage
	UcbPaperStorage				  *UcbDataStorage.UcbPaperStorage
	UcbScenarioStorage			  *UcbDataStorage.UcbScenarioStorage
}

func (s UcbPaperinputResource) NewPaperinputResource(args []BmDataStorage.BmStorage) *UcbPaperinputResource {
	var pis *UcbDataStorage.UcbPaperinputStorage
	var bis *UcbDataStorage.UcbBusinessinputStorage
	var ris *UcbDataStorage.UcbRepresentativeinputStorage
	var mis *UcbDataStorage.UcbManagerinputStorage
	var ps *UcbDataStorage.UcbPaperStorage
	var ss *UcbDataStorage.UcbScenarioStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbPaperinputStorage" {
			pis = arg.(*UcbDataStorage.UcbPaperinputStorage)
		} else if tp.Name() == "UcbBusinessinputStorage" {
			bis = arg.(*UcbDataStorage.UcbBusinessinputStorage)
		} else if tp.Name() == "UcbRepresentativeinputStorage" {
			ris = arg.(*UcbDataStorage.UcbRepresentativeinputStorage)
		} else if tp.Name() == "UcbManagerinputStorage" {
			mis = arg.(*UcbDataStorage.UcbManagerinputStorage)
		} else if tp.Name() == "UcbPaperStorage" {
			ps = arg.(*UcbDataStorage.UcbPaperStorage)
		} else if tp.Name() == "UcbScenarioStorage" {
			ss = arg.(*UcbDataStorage.UcbScenarioStorage)
		}
	}
	return &UcbPaperinputResource{
		UcbPaperinputStorage:          	pis,
		UcbBusinessinputStorage:       	bis,
		UcbRepresentativeinputStorage: 	ris,
		UcbManagerinputStorage:        	mis,
		UcbPaperStorage: 				ps,
		UcbScenarioStorage: 			ss,
	}
}

func (s UcbPaperinputResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	papersID, dcok := r.QueryParams["papersID"]
	var result []UcbModel.Paperinput

	if dcok {
		modelRootID := papersID[0]
		modelRoot, err := s.UcbPaperStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.InputIDs

		result := s.UcbPaperinputStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	models := s.UcbPaperinputStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbPaperinputResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Paperinput
		number, size, offset, limit string
	)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}
	offsetQuery, ok := r.QueryParams["page[offset]"]
	if ok {
		offset = offsetQuery[0]
	}
	limitQuery, ok := r.QueryParams["page[limit]"]
	if ok {
		limit = limitQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for _, iter := range s.UcbPaperinputStorage.GetAll(r, int(start), int(sizeI)) {
			result = append(result, *iter)
		}

	} else {
		limitI, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		offsetI, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		for _, iter := range s.UcbPaperinputStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Paperinput{}
	count := s.UcbPaperinputStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbPaperinputResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbPaperinputStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	err = s.ResetReferencedModel(&model, &r)

	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbPaperinputResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Paperinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbPaperinputStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbPaperinputResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbPaperinputStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbPaperinputResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Paperinput)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbPaperinputStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}

func (s UcbPaperinputResource) ResetReferencedModel(model *UcbModel.Paperinput, r *api2go.Request) error {
	model.Businessinputs = []*UcbModel.Businessinput{}
	r.QueryParams["ids"] = model.BusinessinputIDs
	for _, Businessinput := range s.UcbBusinessinputStorage.GetAll(*r, -1, -1) {
		model.Businessinputs = append(model.Businessinputs, Businessinput)
	}

	model.Representativeinputs = []*UcbModel.Representativeinput{}
	r.QueryParams["ids"] = model.RepresentativeinputIDs
	for _, Representativeinput := range s.UcbRepresentativeinputStorage.GetAll(*r, -1, -1) {
		model.Representativeinputs = append(model.Representativeinputs, Representativeinput)
	}

	model.Managerinputs = []*UcbModel.Managerinput{}
	r.QueryParams["ids"] = model.ManagerinputIDs
	for _, manageInput := range s.UcbManagerinputStorage.GetAll(*r, -1, -1) {
		model.Managerinputs = append(model.Managerinputs, manageInput)
	}

	result, _ := s.UcbScenarioStorage.GetOne(model.ScenarioID)
	model.Scenario = &result

	return nil
}
