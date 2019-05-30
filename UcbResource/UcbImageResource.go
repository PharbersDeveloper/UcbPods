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

type UcbImageResource struct {
	UcbImageStorage          *UcbDataStorage.UcbImageStorage
	UcbProductStorage        *UcbDataStorage.UcbProductStorage
	UcbHospitalStorage       *UcbDataStorage.UcbHospitalStorage
	UcbRegionStorage         *UcbDataStorage.UcbRegionStorage
	UcbRepresentativeStorage *UcbDataStorage.UcbRepresentativeStorage
	UcbLevelStorage	 		 *UcbDataStorage.UcbLevelStorage
	UcbTitleStorage			 *UcbDataStorage.UcbTitleStorage
}

func (c UcbImageResource) NewImageResource(args []BmDataStorage.BmStorage) *UcbImageResource {
	var cs *UcbDataStorage.UcbImageStorage
	var ps *UcbDataStorage.UcbProductStorage
	var hs *UcbDataStorage.UcbHospitalStorage
	var rs *UcbDataStorage.UcbRegionStorage
	var rt *UcbDataStorage.UcbRepresentativeStorage
	var ls *UcbDataStorage.UcbLevelStorage
	var ts *UcbDataStorage.UcbTitleStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbImageStorage" {
			cs = arg.(*UcbDataStorage.UcbImageStorage)
		} else if tp.Name() == "UcbProductStorage" {
			ps = arg.(*UcbDataStorage.UcbProductStorage)
		} else if tp.Name() == "UcbHospitalStorage" {
			hs = arg.(*UcbDataStorage.UcbHospitalStorage)
		} else if tp.Name() == "UcbRegionStorage" {
			rs = arg.(*UcbDataStorage.UcbRegionStorage)
		} else if tp.Name() == "UcbRepresentativeStorage" {
			rt = arg.(*UcbDataStorage.UcbRepresentativeStorage)
		} else if tp.Name() == "UcbLevelStorage" {
			ls = arg.(*UcbDataStorage.UcbLevelStorage)
		} else if tp.Name() == "UcbTitleStorage" {
			ts = arg.(*UcbDataStorage.UcbTitleStorage)
		}
	}
	return &UcbImageResource{
		UcbImageStorage:          cs,
		UcbProductStorage:        ps,
		UcbHospitalStorage:       hs,
		UcbRegionStorage:         rs,
		UcbRepresentativeStorage: rt,
		UcbLevelStorage: 		  ls,
		UcbTitleStorage:      	  ts,
	}
}

// FindAll images
func (c UcbImageResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	productsID, pok := r.QueryParams["productsID"]
	hospitalsID, hsok := r.QueryParams["hospitalsID"]
	regionsID, rsok := r.QueryParams["regionsID"]
	representativeID, rtok := r.QueryParams["representativesID"]
	levelsID, lok := r.QueryParams["levelsID"]
	titlesID, tok := r.QueryParams["titlesID"]

	if pok {
		modelRootID := productsID[0]
		modelRoot, err := c.UcbProductStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.ImagesIDs
	} else if hsok {
		modelRootID := hospitalsID[0]
		modelRoot, err := c.UcbHospitalStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.ImagesIDs
	} else if rsok {
		modelRootID := regionsID[0]
		modelRoot, err := c.UcbRegionStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.ImagesIDs
	} else if rtok {
		modelRootID := representativeID[0]
		modelRoot, err := c.UcbRepresentativeStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		r.QueryParams["ids"] = modelRoot.ImagesIDs
	} else if lok {
		modelRootID := levelsID[0]
		modelRoot, err := c.UcbLevelStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbImageStorage.GetOne(modelRoot.ImagesID)

		if err != nil {
			return &Response{}, err
		}

		return  &Response{Res: model}, nil
	} else if tok {
		modelRootID := titlesID[0]
		modelRoot, err := c.UcbTitleStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, err
		}
		model, err := c.UcbImageStorage.GetOne(modelRoot.ImagesID)

		if err != nil {
			return &Response{}, err
		}

		return  &Response{Res: model}, nil
	}

	var result []UcbModel.Image
	result = c.UcbImageStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbImageResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbImageStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbImageResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbImageStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbImageResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbImageStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbImageResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.Image)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbImageStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
