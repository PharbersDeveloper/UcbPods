package UcbResource

import (
	"errors"
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"net/http"
	"reflect"

	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
)

type UcbSalesConfigResource struct {
	UcbSalesConfigStorage       *UcbDataStorage.UcbSalesConfigStorage
	UcbPaperStorage				*UcbDataStorage.UcbPaperStorage
	UcbSalesReportStorage 		*UcbDataStorage.UcbSalesReportStorage
	UcbHospitalSalesReportStorage *UcbDataStorage.UcbHospitalSalesReportStorage
	UcbSalesReportResource 			*UcbSalesReportResource
	UcbDestConfigStorage 		*UcbDataStorage.UcbDestConfigStorage
	UcbHospitalConfigStorage 	*UcbDataStorage.UcbHospitalConfigStorage
	UcbGoodsConfigStorage 		*UcbDataStorage.UcbGoodsConfigStorage
}

func (c UcbSalesConfigResource) NewSalesConfigResource(args []BmDataStorage.BmStorage) *UcbSalesConfigResource {
	var sc *UcbDataStorage.UcbSalesConfigStorage
	var ps *UcbDataStorage.UcbPaperStorage
	var srr	*UcbSalesReportResource


	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbSalesConfigStorage" {
			sc = arg.(*UcbDataStorage.UcbSalesConfigStorage)
		} else if tp.Name() == "UcbPaperStorage" {
			ps = arg.(*UcbDataStorage.UcbPaperStorage)
		} else if tp.Name() == "UcbSalesReportResource" {
			srr = arg.(interface{}).(*UcbSalesReportResource)
		}
	}
	return &UcbSalesConfigResource{
		UcbSalesConfigStorage:	sc,
		UcbPaperStorage: ps,
		UcbSalesReportResource: srr,
	}
}

// FindAll SalesConfigs
func (c UcbSalesConfigResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	var result []*UcbModel.SalesConfig
	_, acok := r.QueryParams["account-id"]
	_, psok := r.QueryParams["proposal-id"]

	if acok && psok{
		result = c.UcbSalesConfigStorage.GetAll(r, -1, -1)

		r.QueryParams["orderby"] = []string{"time"}
		paperModel := c.UcbPaperStorage.GetAll(r, -1, -1)


		if len(paperModel) > 0 {
			r.QueryParams = map[string][]string{}
			// 获取这个用户在关卡下最新的报告
			SalesReportIDs := paperModel[0].SalesReportIDs
			LastSalesReportID := SalesReportIDs[len(SalesReportIDs)-1]
			SalesReportResponse, err := c.UcbSalesReportResource.FindOne(LastSalesReportID, r)
			if err != nil {
				return &Response{}, err
			}
			response := SalesReportResponse.Result()
			item := response.(UcbModel.SalesReport)

			for _, salesConfigModel := range result {
				salesConfigModel.SalesReportID = item.ID
				salesConfigModel.SalesReport = &item
			}
		}
		return &Response{Res: result}, nil
	}


	result = c.UcbSalesConfigStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// FindOne choc
func (c UcbSalesConfigResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	res, err := c.UcbSalesConfigStorage.GetOne(ID)
	return &Response{Res: res}, err
}

// Create a new choc
func (c UcbSalesConfigResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.SalesConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	id := c.UcbSalesConfigStorage.Insert(choc)
	choc.ID = id
	return &Response{Res: choc, Code: http.StatusCreated}, nil
}

// Delete a choc :(
func (c UcbSalesConfigResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := c.UcbSalesConfigStorage.Delete(id)
	return &Response{Code: http.StatusOK}, err
}

// Update a choc
func (c UcbSalesConfigResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	choc, ok := obj.(UcbModel.SalesConfig)
	if !ok {
		return &Response{}, api2go.NewHTTPError(
			errors.New("Invalid instance given"),
			"Invalid instance given",
			http.StatusBadRequest,
		)
	}

	err := c.UcbSalesConfigStorage.Update(choc)
	return &Response{Res: choc, Code: http.StatusNoContent}, err
}
