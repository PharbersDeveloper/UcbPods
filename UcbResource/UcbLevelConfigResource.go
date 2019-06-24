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

type UcbLevelConfigResource struct {
	UcbLevelConfigStorage          		*UcbDataStorage.UcbLevelConfigStorage
	UcbRegionalDivisionResultStorage	*UcbDataStorage.UcbRegionalDivisionResultStorage
	UcbTargetAssignsResultStorage		*UcbDataStorage.UcbTargetAssignsResultStorage
	UcbResourceAssignsResultStorage		*UcbDataStorage.UcbResourceAssignsResultStorage
	UcbManageTimeResultStorage			*UcbDataStorage.UcbManageTimeResultStorage
	UcbManageTeamResultStorage			*UcbDataStorage.UcbManageTeamResultStorage
	UcbGeneralPerformanceResultStorage	*UcbDataStorage.UcbGeneralPerformanceResultStorage
	UcbSimplifyResultStorage			*UcbDataStorage.UcbSimplifyResultStorage
}

func (c UcbLevelConfigResource) NewLevelConfigResource(args []BmDataStorage.BmStorage) *UcbLevelConfigResource {
	var lcs *UcbDataStorage.UcbLevelConfigStorage
	var rdr	*UcbDataStorage.UcbRegionalDivisionResultStorage
	var tar	*UcbDataStorage.UcbTargetAssignsResultStorage
	var rar *UcbDataStorage.UcbResourceAssignsResultStorage
	var mtr *UcbDataStorage.UcbManageTimeResultStorage
	var mtrs *UcbDataStorage.UcbManageTeamResultStorage
	var gpr *UcbDataStorage.UcbGeneralPerformanceResultStorage
	var srs *UcbDataStorage.UcbSimplifyResultStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbLevelConfigStorage" {
			lcs = arg.(*UcbDataStorage.UcbLevelConfigStorage)
		} else if tp.Name() == "UcbRegionalDivisionResultStorage" {
			rdr = arg.(*UcbDataStorage.UcbRegionalDivisionResultStorage)
		} else if tp.Name() == "UcbTargetAssignsResultStorage" {
			tar = arg.(*UcbDataStorage.UcbTargetAssignsResultStorage)
		} else if tp.Name() == "UcbResourceAssignsResultStorage" {
			rar = arg.(*UcbDataStorage.UcbResourceAssignsResultStorage)
		} else if tp.Name() == "UcbManageTimeResultStorage" {
			mtr = arg.(*UcbDataStorage.UcbManageTimeResultStorage)
		} else if tp.Name() == "UcbManageTeamResultStorage" {
			mtrs = arg.(*UcbDataStorage.UcbManageTeamResultStorage)
		} else if tp.Name() == "UcbGeneralPerformanceResultStorage" {
			gpr = arg.(*UcbDataStorage.UcbGeneralPerformanceResultStorage)
		} else if tp.Name() == "UcbSimplifyResultStorage" {
			srs = arg.(*UcbDataStorage.UcbSimplifyResultStorage)
		}
	}
	return &UcbLevelConfigResource{
		UcbLevelConfigStorage:	lcs,
		UcbRegionalDivisionResultStorage: rdr,
		UcbTargetAssignsResultStorage: tar,
		UcbResourceAssignsResultStorage: rar,
		UcbManageTimeResultStorage: mtr,
		UcbManageTeamResultStorage: mtrs,
		UcbGeneralPerformanceResultStorage: gpr,
		UcbSimplifyResultStorage: srs,
	}
}

// FindAll images
func (c UcbLevelConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []UcbModel.LevelConfig
	regionalDivisionResultsID, rdrOk := r.QueryParams["regionalDivisionResultsID"]
	targetAssignsResultsID, tarOk := r.QueryParams["targetAssignsResultsID"]
	resourceAssignsResultsID, rarOk := r.QueryParams["resourceAssignsResultsID"]
	manageTimeResultsID, mtrOk := r.QueryParams["manageTimeResultsID"]
	manageTeamResultsID, mtrsOk := r.QueryParams["manageTeamResultsID"]
	generalPerformanceResultsID, gprOk := r.QueryParams["generalPerformanceResultsID"]
	simplifyResultsID, srOk := r.QueryParams["simplifyResultsID"]

	if rdrOk {
		modelRootID := regionalDivisionResultsID[0]
		modelRoot, err := c.UcbRegionalDivisionResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if tarOk {
		modelRootID := targetAssignsResultsID[0]
		modelRoot, err := c.UcbTargetAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if rarOk {
		modelRootID := resourceAssignsResultsID[0]
		modelRoot, err := c.UcbResourceAssignsResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if mtrOk {
		modelRootID := manageTimeResultsID[0]
		modelRoot, err := c.UcbManageTimeResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if mtrsOk {
		modelRootID := manageTeamResultsID[0]
		modelRoot, err := c.UcbManageTeamResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if gprOk {
		modelRootID := generalPerformanceResultsID[0]
		modelRoot, err := c.UcbGeneralPerformanceResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	if srOk {
		modelRootID := simplifyResultsID[0]
		modelRoot, err := c.UcbSimplifyResultStorage.GetOne(modelRootID)
		if err != nil {
			return &Response{}, nil
		}

		model, err:= c.UcbLevelConfigStorage.GetOne(modelRoot.LevelConfigID)

		if err != nil {
			return &Response{}, nil
		}
		return &Response{Res: model}, nil
	}

	result = c.UcbLevelConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbLevelConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbLevelConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbLevelConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.LevelConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbLevelConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbLevelConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbLevelConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbLevelConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.LevelConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbLevelConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
