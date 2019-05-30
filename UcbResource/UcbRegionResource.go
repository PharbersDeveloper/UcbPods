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

type UcbRegionResource struct {
	UcbRegionStorage *UcbDataStorage.UcbRegionStorage
	UcbImageStorage  *UcbDataStorage.UcbImageStorage
	UcbRegionConfigStorage *UcbDataStorage.UcbRegionConfigStorage
}

func (s UcbRegionResource) NewRegionResource(args []BmDataStorage.BmStorage) *UcbRegionResource {
	var is *UcbDataStorage.UcbImageStorage
	var hs *UcbDataStorage.UcbRegionStorage
	var rcs *UcbDataStorage.UcbRegionConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbImageStorage" {
			is = arg.(*UcbDataStorage.UcbImageStorage)
		} else if tp.Name() == "UcbRegionStorage" {
			hs = arg.(*UcbDataStorage.UcbRegionStorage)
		} else if tp.Name() == "UcbRegionConfigStorage" {
			rcs = arg.(*UcbDataStorage.UcbRegionConfigStorage)
		}
	}
	return &UcbRegionResource{
		UcbImageStorage: is,
		UcbRegionStorage: hs,
		UcbRegionConfigStorage: rcs,
	}
}

func (s UcbRegionResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	regionConfigsID, pciok := r.QueryParams["regionConfigsID"]

	if pciok {
		modelRootID := regionConfigsID[0]
		modelRoot, err := s.UcbRegionConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err := s.UcbRegionConfigStorage.GetOne(modelRoot.RegionID)
		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.Region
	models := s.UcbRegionStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbRegionResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Region
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
		for _, iter := range s.UcbRegionStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbRegionStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Region{}
	count := s.UcbRegionStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbRegionResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbRegionStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Imgs = []*UcbModel.Image{}
	for _, kID := range model.ImagesIDs {
		choc, err := s.UcbImageStorage.GetOne(kID)
		if err != nil {
			return &Response{}, err
		}
		model.Imgs = append(model.Imgs, &choc)
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbRegionResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Region)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbRegionStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbRegionResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbRegionStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbRegionResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Region)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbRegionStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
