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

type UcbResourceConfigResource struct {
	UcbResourceConfigStorage       	*UcbDataStorage.UcbResourceConfigStorage
	UcbManagerConfigStorage        	*UcbDataStorage.UcbManagerConfigStorage
	UcbRepresentativeConfigStorage 	*UcbDataStorage.UcbRepresentativeConfigStorage
	UcbTeamConfigStorage		   	*UcbDataStorage.UcbTeamConfigStorage
	UcbRepresentativeinputStorage	*UcbDataStorage.UcbRepresentativeinputStorage
	UcbBusinessinputStorage			*UcbDataStorage.UcbBusinessinputStorage
	UcbRepresentativeSalesReportStorage	*UcbDataStorage.UcbRepresentativeSalesReportStorage
	UcbHospitalSalesReportStorage		*UcbDataStorage.UcbHospitalSalesReportStorage
}

func (s UcbResourceConfigResource) NewResourceConfigResource(args []BmDataStorage.BmStorage) *UcbResourceConfigResource {
	var rcs *UcbDataStorage.UcbResourceConfigStorage
	var mcs *UcbDataStorage.UcbManagerConfigStorage
	var repcs *UcbDataStorage.UcbRepresentativeConfigStorage
	var tcs *UcbDataStorage.UcbTeamConfigStorage
	var ris *UcbDataStorage.UcbRepresentativeinputStorage
	var bis *UcbDataStorage.UcbBusinessinputStorage
	var rsr *UcbDataStorage.UcbRepresentativeSalesReportStorage
	var hsrs *UcbDataStorage.UcbHospitalSalesReportStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbResourceConfigStorage" {
			rcs = arg.(*UcbDataStorage.UcbResourceConfigStorage)
		} else if tp.Name() == "UcbManagerConfigStorage" {
			mcs = arg.(*UcbDataStorage.UcbManagerConfigStorage)
		} else if tp.Name() == "UcbRepresentativeConfigStorage" {
			repcs = arg.(*UcbDataStorage.UcbRepresentativeConfigStorage)
		} else if tp.Name() == "UcbTeamConfigStorage" {
			tcs = arg.(*UcbDataStorage.UcbTeamConfigStorage)
		} else if tp.Name() == "UcbRepresentativeinputStorage" {
			ris = arg.(*UcbDataStorage.UcbRepresentativeinputStorage)
		} else if tp.Name() == "UcbBusinessinputStorage" {
			bis = arg.(*UcbDataStorage.UcbBusinessinputStorage)
		} else if tp.Name() == "UcbRepresentativeSalesReportStorage" {
			rsr = arg.(*UcbDataStorage.UcbRepresentativeSalesReportStorage)
		} else if tp.Name() == "UcbHospitalSalesReportStorage" {
			hsrs = arg.(*UcbDataStorage.UcbHospitalSalesReportStorage)
		}
	}
	return &UcbResourceConfigResource{
		UcbResourceConfigStorage:       rcs,
		UcbManagerConfigStorage:        mcs,
		UcbRepresentativeConfigStorage: repcs,
		UcbTeamConfigStorage: tcs,
		UcbRepresentativeinputStorage: ris,
		UcbBusinessinputStorage: bis,
		UcbRepresentativeSalesReportStorage: rsr,
		UcbHospitalSalesReportStorage: hsrs,
	}
}

func (s UcbResourceConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.ResourceConfig
	teamConfigsID, tcok := r.QueryParams["teamConfigsID"]
	representativeinputsID, tok := r.QueryParams["representativeinputsID"]
	businessinputsID, bok := r.QueryParams["businessinputsID"]
	representativeSalesReportsID, rsrok := r.QueryParams["representativeSalesReportsID"]
	hospitalSalesReportsID, hsrok := r.QueryParams["hospitalSalesReportsID"]

	if tcok {
		modelRootID := teamConfigsID[0]
		modelRoot, err := s.UcbTeamConfigStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		r.QueryParams["ids"] = modelRoot.ResourceConfigIDs

		result := s.UcbResourceConfigStorage.GetAll(r, -1,-1)

		return &Response{Res: result}, nil
	}

	if tok {
		modelRootID := representativeinputsID[0]
		modelRoot, err := s.UcbRepresentativeinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		result, err := s.UcbResourceConfigStorage.GetOne(modelRoot.ResourceConfigId)

		if err != nil {
			return &Response{}, nil
		}

		return &Response{Res: result}, nil
	}

	if bok {
		modelRootID := businessinputsID[0]
		modelRoot, err := s.UcbBusinessinputStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		result, err := s.UcbResourceConfigStorage.GetOne(modelRoot.ResourceConfigId)

		if err != nil {
			return &Response{}, nil
		}

		return &Response{Res: result}, nil
	}

	if rsrok {
		modelRootID := representativeSalesReportsID[0]
		modelRoot, err := s.UcbRepresentativeSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		result, err := s.UcbResourceConfigStorage.GetOne(modelRoot.ResourceConfigID)

		if err != nil {
			return &Response{}, nil
		}

		return &Response{Res: result}, nil
	}

	if hsrok {
		modelRootID := hospitalSalesReportsID[0]
		modelRoot, err := s.UcbHospitalSalesReportStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		result, err := s.UcbResourceConfigStorage.GetOne(modelRoot.ResourceConfigID)

		if err != nil {
			return &Response{}, nil
		}

		return &Response{Res: result}, nil

	}

	models := s.UcbResourceConfigStorage.GetAll(r, -1, -1)

	if len(models) == 1 {
		return &Response{Res: models[0]}, nil
	}

	for _, model := range models {
		result = append(result, *model)
	}

	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbResourceConfigResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.ResourceConfig
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
		for _, iter := range s.UcbResourceConfigStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbResourceConfigStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.ResourceConfig{}
	count := s.UcbResourceConfigStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbResourceConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbResourceConfigStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	if model.ResourceType == 0 {
		mcr, err := s.UcbManagerConfigStorage.GetOne(model.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		model.ManagerConfig = &mcr
	} else if model.ResourceType == 1 {
		rcr, err := s.UcbRepresentativeConfigStorage.GetOne(model.ResourceID)
		if err != nil {
			return &Response{}, err
		}
		model.RepresentativeConfig = &rcr
	}

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbResourceConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.ResourceConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := s.UcbResourceConfigStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbResourceConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbResourceConfigStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbResourceConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.ResourceConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbResourceConfigStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}
