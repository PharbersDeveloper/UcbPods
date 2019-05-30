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

type UcbDestConfigResource struct {
	UcbDestConfigStorage    			*UcbDataStorage.UcbDestConfigStorage
	UcbHospitalConfigStorage			*UcbDataStorage.UcbHospitalConfigStorage
	UcbRegionConfigStorage				*UcbDataStorage.UcbRegionConfigStorage

	UcbHospitalSalesReportStorage		*UcbDataStorage.UcbHospitalSalesReportStorage
	UcbSalesConfigStorage 				*UcbDataStorage.UcbSalesConfigStorage
	UcbBusinessinputStorage				*UcbDataStorage.UcbBusinessinputStorage
}

func (s UcbDestConfigResource) NewDestConfigResource(args []BmDataStorage.BmStorage) *UcbDestConfigResource {
	var dcs *UcbDataStorage.UcbDestConfigStorage
	var hcs *UcbDataStorage.UcbHospitalConfigStorage
	var rcs *UcbDataStorage.UcbRegionConfigStorage
	var hsr *UcbDataStorage.UcbHospitalSalesReportStorage
	var sc *UcbDataStorage.UcbSalesConfigStorage
	var bis *UcbDataStorage.UcbBusinessinputStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbDestConfigStorage" {
			dcs = arg.(*UcbDataStorage.UcbDestConfigStorage)
		} else if tp.Name() == "UcbHospitalConfigStorage" {
			hcs = arg.(*UcbDataStorage.UcbHospitalConfigStorage)
		} else if tp.Name() == "UcbRegionConfigStorage" {
			rcs = arg.(*UcbDataStorage.UcbRegionConfigStorage)
		} else if tp.Name() == "UcbHospitalSalesReportStorage" {
			hsr = arg.(*UcbDataStorage.UcbHospitalSalesReportStorage)
		} else if tp.Name() == "UcbSalesConfigStorage" {
			sc = arg.(*UcbDataStorage.UcbSalesConfigStorage)
		} else if tp.Name() == "UcbBusinessinputStorage" {
			bis = arg.(*UcbDataStorage.UcbBusinessinputStorage)
		}
	}
	return &UcbDestConfigResource{
		UcbDestConfigStorage:    	dcs,
		UcbHospitalConfigStorage: hcs,
		UcbRegionConfigStorage: rcs,
		UcbHospitalSalesReportStorage : hsr,
		UcbSalesConfigStorage: sc,
		UcbBusinessinputStorage: bis,
	}
}

func (s UcbDestConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	hospitalSalesReportsID, hsrok := r.QueryParams["hospitalSalesReportsID"]
	salesConfigsID, scok := r.QueryParams["salesConfigsID"]
	businessinputsID, bok := r.QueryParams["businessinputsID"]

	if hsrok {
		modelRootID := hospitalSalesReportsID[0]
		modelRoot, err := s.UcbHospitalSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= s.UcbDestConfigStorage.GetOne(modelRoot.DestConfigID)


		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if scok {
		modelRootID := salesConfigsID[0]
		modelRoot, err := s.UcbSalesConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= s.UcbDestConfigStorage.GetOne(modelRoot.DestConfigID)


		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if bok {
		modelRootID := businessinputsID[0]
		modelRoot, err := s.UcbBusinessinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		result, err := s.UcbDestConfigStorage.GetOne(modelRoot.DestConfigId)

		if err != nil {
			return &Response{}, nil
		}

		return &Response{Res: result}, nil
	}

	var result []UcbModel.DestConfig
	models := s.UcbDestConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbDestConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.DestConfig
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
		for _, iter := range s.UcbDestConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbDestConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.DestConfig{}
	count := s.UcbDestConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbDestConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	modelRoot, err := s.UcbDestConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if modelRoot.DestType == 0 {
		model, err := s.UcbRegionConfigStorage.GetOne(modelRoot.DestID)
		if err != nil {
			return &Response{}, err
		}
		modelRoot.RegionConfig = &model
	}

	if modelRoot.DestType == 1 {
		model, err := s.UcbHospitalConfigStorage.GetOne(modelRoot.DestID)
		if err != nil {
			return &Response{}, err
		}
		modelRoot.HospitalConfig = &model
	}

	return &Response{Res: modelRoot}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbDestConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.DestConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbDestConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbDestConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbDestConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbDestConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.DestConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbDestConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
