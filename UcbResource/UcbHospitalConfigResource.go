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

type UcbHospitalConfigResource struct {
	UcbHospitalConfigStorage 	*UcbDataStorage.UcbHospitalConfigStorage
	UcbHospitalStorage			*UcbDataStorage.UcbHospitalStorage
	UcbPolicyStorage 			*UcbDataStorage.UcbPolicyStorage
	UcbDepartmentStorage 		*UcbDataStorage.UcbDepartmentStorage
	UcbDestConfigStorage 		*UcbDataStorage.UcbDestConfigStorage
}

func (s UcbHospitalConfigResource) NewHospitalConfigResource(args []BmDataStorage.BmStorage) *UcbHospitalConfigResource {
	var hcs *UcbDataStorage.UcbHospitalConfigStorage
	var hs *UcbDataStorage.UcbHospitalStorage
	var ps *UcbDataStorage.UcbPolicyStorage
	var ds *UcbDataStorage.UcbDepartmentStorage
	var dcs *UcbDataStorage.UcbDestConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()

		if tp.Name() == "UcbHospitalConfigStorage" {
			hcs = arg.(*UcbDataStorage.UcbHospitalConfigStorage)
		} else if tp.Name() == "UcbHospitalStorage" {
			hs = arg.(*UcbDataStorage.UcbHospitalStorage)
		} else if tp.Name() == "UcbPolicyStorage" {
			ps = arg.(*UcbDataStorage.UcbPolicyStorage)
		} else if tp.Name() == "UcbDepartmentStorage" {
			ds = arg.(*UcbDataStorage.UcbDepartmentStorage)
		} else if tp.Name() == "UcbDestConfigStorage" {
			dcs = arg.(*UcbDataStorage.UcbDestConfigStorage)
		}
	}
	return &UcbHospitalConfigResource{
		UcbHospitalConfigStorage: hcs,
		UcbHospitalStorage: hs,
		UcbPolicyStorage: ps,
		UcbDepartmentStorage: ds,
		UcbDestConfigStorage: dcs,
	}
}

func (s UcbHospitalConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	destConfigsID, dcok := r.QueryParams["destConfigsID"]

	if dcok {
		modelRootID := destConfigsID[0]
		modelRoot, err := s.UcbDestConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}
		model, err := s.UcbHospitalConfigStorage.GetOne(modelRoot.DestID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	var result []UcbModel.HospitalConfig

	models := s.UcbHospitalConfigStorage.GetAll(r, -1, -1)

	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbHospitalConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.HospitalConfig
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
		for _, iter := range s.UcbHospitalConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbHospitalConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.HospitalConfig{}
	count := s.UcbHospitalConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbHospitalConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbHospitalConfigStorage.GetOne(ID)
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
func (s UcbHospitalConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.HospitalConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbHospitalConfigStorage.Insert(model)
	model.ID = id

	s.ResetReferencedModel(&model, &r)

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbHospitalConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbHospitalConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbHospitalConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.HospitalConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbHospitalConfigStorage.Update(model)

	s.ResetReferencedModel(&model, &r)

	return &Response{Res: model, Code: http.StatusNoContent}, err
}

func (s UcbHospitalConfigResource) ResetReferencedModel(model *UcbModel.HospitalConfig, r *api2go.Request) error {
	model.Policies = []*UcbModel.Policy{}
	r.QueryParams["ids"] = model.PolicyIDs
	for _, policy := range s.UcbPolicyStorage.GetAll(*r, -1,-1) {
		model.Policies = append(model.Policies, policy)
	}

	model.Departments = []*UcbModel.Department{}
	r.QueryParams["ids"] = model.DepartmentIDs
	for _, department := range s.UcbDepartmentStorage.GetAll(*r, -1, -1) {
		model.Departments = append(model.Departments, department)
	}

	if model.HospitalID != "" {
		hospital, err := s.UcbHospitalStorage.GetOne(model.HospitalID)
		if err != nil {
			return err
		}
		model.Hospital = &hospital
	}

	return nil
}
