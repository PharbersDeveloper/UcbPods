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

type UcbRepresentativeResource struct {
	UcbRepresentativeStorage       *UcbDataStorage.UcbRepresentativeStorage
	UcbRepresentativeConfigStorage *UcbDataStorage.UcbRepresentativeConfigStorage
	UcbImageStorage                *UcbDataStorage.UcbImageStorage
	UcbActionKpiStorage			   *UcbDataStorage.UcbActionKpiStorage
	UcbRepresentativeAbilityStorage *UcbDataStorage.UcbRepresentativeAbilityStorage
}

func (s UcbRepresentativeResource) NewRepresentativeResource(args []BmDataStorage.BmStorage) *UcbRepresentativeResource {
	var is *UcbDataStorage.UcbImageStorage
	var reps *UcbDataStorage.UcbRepresentativeStorage
	var repcs *UcbDataStorage.UcbRepresentativeConfigStorage
	var aks	*UcbDataStorage.UcbActionKpiStorage
	var ras *UcbDataStorage.UcbRepresentativeAbilityStorage
	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbImageStorage" {
			is = arg.(*UcbDataStorage.UcbImageStorage)
		} else if tp.Name() == "UcbRepresentativeStorage" {
			reps = arg.(*UcbDataStorage.UcbRepresentativeStorage)
		} else if tp.Name() == "UcbRepresentativeConfigStorage" {
			repcs = arg.(*UcbDataStorage.UcbRepresentativeConfigStorage)
		} else if tp.Name() == "UcbActionKpiStorage" {
			aks = arg.(*UcbDataStorage.UcbActionKpiStorage)
		} else if tp.Name() == "UcbRepresentativeAbilityStorage" {
			ras = arg.(*UcbDataStorage.UcbRepresentativeAbilityStorage)
		}
	}
	return &UcbRepresentativeResource{
		UcbImageStorage:                is,
		UcbRepresentativeStorage:       reps,
		UcbRepresentativeConfigStorage: repcs,
		UcbActionKpiStorage: 			aks,
		UcbRepresentativeAbilityStorage: ras,
	}
}

func (s UcbRepresentativeResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	representativeConfigsID, rcok := r.QueryParams["representativeConfigsID"]
	actionKpisID, akok := r.QueryParams["actionKpisID"]
	representativeAbilitiesID, raok := r.QueryParams["representativeAbilitiesID"]
	var result []UcbModel.Representative

	if rcok {
		modelRootID := representativeConfigsID[0]
		modelRoot, err := s.UcbRepresentativeConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.UcbRepresentativeStorage.GetOne(modelRoot.RepresentativeID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	if akok {
		modelRootID := actionKpisID[0]
		modelRoot, err := s.UcbActionKpiStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.UcbRepresentativeStorage.GetOne(modelRoot.RepresentativeID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	if raok {
		modelRootID := representativeAbilitiesID[0]
		modelRoot, err := s.UcbRepresentativeAbilityStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := s.UcbRepresentativeStorage.GetOne(modelRoot.RepresentativeID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	models := s.UcbRepresentativeStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbRepresentativeResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Representative
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
		for _, iter := range s.UcbRepresentativeStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbRepresentativeStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Representative{}
	count := s.UcbRepresentativeStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbRepresentativeResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbRepresentativeStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	r.QueryParams["ids"] = model.ImagesIDs
	images := s.UcbImageStorage.GetAll(r, -1, -1)
	for _, image := range images {
		model.Imgs = append(model.Imgs, &image)
	}
	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Representative)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbRepresentativeStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbRepresentativeResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbRepresentativeStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbRepresentativeResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Representative)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbRepresentativeStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
