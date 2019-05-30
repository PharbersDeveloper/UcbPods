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

type UcbRegionConfigResource struct {
	UcbRegionConfigStorage		*UcbDataStorage.UcbRegionConfigStorage
	UcbRegionStorage			*UcbDataStorage.UcbRegionStorage
	UcbDestConfigStorage 		*UcbDataStorage.UcbDestConfigStorage
}

func (s UcbRegionConfigResource) NewRegionConfigResource(args []BmDataStorage.BmStorage) *UcbRegionConfigResource {
	var rcs *UcbDataStorage.UcbRegionConfigStorage
	var rs *UcbDataStorage.UcbRegionStorage
	var dcs *UcbDataStorage.UcbDestConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbRegionConfigStorage" {
			rcs = arg.(*UcbDataStorage.UcbRegionConfigStorage)
		} else if tp.Name() == "UcbRegionStorage" {
			rs = arg.(*UcbDataStorage.UcbRegionStorage)
		} else if tp.Name() == "UcbDestConfigStorage" {
			dcs = arg.(*UcbDataStorage.UcbDestConfigStorage)
		}
	}
	return &UcbRegionConfigResource{
		UcbRegionStorage: rs,
		UcbRegionConfigStorage: rcs,
		UcbDestConfigStorage: dcs,
	}
}

func (s UcbRegionConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	destConfigsID, dcok := r.QueryParams["destConfigsID"]


	if dcok {
		modelRootID := destConfigsID[0]
		modelRoot, err := s.UcbDestConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err := s.UcbRegionConfigStorage.GetOne(modelRoot.DestID)
		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.RegionConfig

	models := s.UcbRegionConfigStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbRegionConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.RegionConfig
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
		for _, iter := range s.UcbRegionConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbRegionConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.RegionConfig{}
	count := s.UcbRegionConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbRegionConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbRegionConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.RegionID != "" {
		regionModel, err := s.UcbRegionStorage.GetOne(model.RegionID)
		if err != nil {
			return &Response{}, err
		}
		model.Region = &regionModel
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbRegionConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.RegionConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbRegionConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbRegionConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbRegionConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbRegionConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.RegionConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbRegionConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
