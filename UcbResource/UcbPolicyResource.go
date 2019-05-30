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

type UcbPolicyResource struct {
	UcbPolicyStorage *UcbDataStorage.UcbPolicyStorage
	UcbHospitalConfigStorage *UcbDataStorage.UcbHospitalConfigStorage
}

func (c UcbPolicyResource) NewPolicyResource(args []BmDataStorage.BmStorage) *UcbPolicyResource {
	var cs *UcbDataStorage.UcbPolicyStorage
	var hcs *UcbDataStorage.UcbHospitalConfigStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbPolicyStorage" {
			cs = arg.(*UcbDataStorage.UcbPolicyStorage)
	 	} else if tp.Name() == "UcbHospitalConfigStorage" {
	 		hcs = arg.(*UcbDataStorage.UcbHospitalConfigStorage)
		}
	}
	return &UcbPolicyResource{
		UcbPolicyStorage: cs,
		UcbHospitalConfigStorage: hcs,
	}
}

// FindAll Policys
func (c UcbPolicyResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*UcbModel.Policy
	hospitalConfigsID, pciok := r.QueryParams["hospitalConfigsID"]

	if pciok {
		modelRootID := hospitalConfigsID[0]

		modelRoot, err := c.UcbHospitalConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.PolicyIDs

		result = c.UcbPolicyStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	result = c.UcbPolicyStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbPolicyResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbPolicyStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbPolicyResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Policy)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbPolicyStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbPolicyResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbPolicyStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbPolicyResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Policy)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbPolicyStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
