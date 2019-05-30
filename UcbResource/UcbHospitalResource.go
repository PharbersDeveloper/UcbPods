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

type UcbHospitalResource struct {
	UcbHospitalStorage *UcbDataStorage.UcbHospitalStorage
	UcbImageStorage    *UcbDataStorage.UcbImageStorage
	UcbHospitalConfigStorage *UcbDataStorage.UcbHospitalConfigStorage
}

func (s UcbHospitalResource) NewHospitalResource (args []BmDataStorage.BmStorage) *UcbHospitalResource {
	var is *UcbDataStorage.UcbImageStorage
	var hs *UcbDataStorage.UcbHospitalStorage
	var hcs *UcbDataStorage.UcbHospitalConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbImageStorage" {
			is = arg.(*UcbDataStorage.UcbImageStorage)
		} else if tp.Name() == "UcbHospitalStorage" {
			hs = arg.(*UcbDataStorage.UcbHospitalStorage)
		} else if tp.Name() == "UcbHospitalConfigStorage" {
			hcs = arg.(*UcbDataStorage.UcbHospitalConfigStorage)
		}
	}
	return &UcbHospitalResource{
		UcbImageStorage: is,
		UcbHospitalStorage: hs,
		UcbHospitalConfigStorage: hcs,
	}
}

func (s UcbHospitalResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	hospitalConfigsID, pciok := r.QueryParams["hospitalConfigsID"]

	if pciok {
		modelRootID := hospitalConfigsID[0]

		modelRoot, err := s.UcbHospitalConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}

		model, err := s.UcbHospitalStorage.GetOne(modelRoot.HospitalID)
		if err != nil {
			return &Response{}, err
		}
		return &Response{Res: model}, nil
	}

	var result []*UcbModel.Hospital

	models := s.UcbHospitalStorage.GetAll(r, -1, -1)
	for _, model := range models {
		result = append(result, model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbHospitalResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Hospital
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
		for _, iter := range s.UcbHospitalStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbHospitalStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Hospital{}
	count := s.UcbHospitalStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbHospitalResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbHospitalStorage.GetOne(ID)
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
func (s UcbHospitalResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Hospital)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbHospitalStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbHospitalResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbHospitalStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbHospitalResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Hospital)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbHospitalStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
