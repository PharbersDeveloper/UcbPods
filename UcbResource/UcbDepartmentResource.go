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

type UcbDepartmentResource struct {
	UcbDepartmentStorage *UcbDataStorage.UcbDepartmentStorage
	UcbHospitalConfigStorage *UcbDataStorage.UcbHospitalConfigStorage
}

func (c UcbDepartmentResource) NewDepartmentResource(args []BmDataStorage.BmStorage) *UcbDepartmentResource {
	var cs *UcbDataStorage.UcbDepartmentStorage
	var hcs *UcbDataStorage.UcbHospitalConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbDepartmentStorage" {
			cs = arg.(*UcbDataStorage.UcbDepartmentStorage)
		}else if tp.Name() == "UcbHospitalConfigStorage" {
	 		hcs = arg.(*UcbDataStorage.UcbHospitalConfigStorage)
		}
	}
	return &UcbDepartmentResource{
		UcbDepartmentStorage: cs,
		UcbHospitalConfigStorage: hcs,
	}
}

// FindAll Departments
func (c UcbDepartmentResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*UcbModel.Department
	hospitalConfigsID, pciok := r.QueryParams["hospitalConfigsID"]

	if pciok {
		modelRootID := hospitalConfigsID[0]

		modelRoot, err := c.UcbHospitalConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.DepartmentIDs

		result = c.UcbDepartmentStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	result = c.UcbDepartmentStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbDepartmentResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbDepartmentStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbDepartmentResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Department)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbDepartmentStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbDepartmentResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbDepartmentStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbDepartmentResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Department)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbDepartmentStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
