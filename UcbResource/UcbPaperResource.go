package UcbResource

import (
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"errors"
	"github.com/alfredyang1986/BmServiceDef/BmDataStorage"
	"github.com/manyminds/api2go"
	"net/http"
	"reflect"
	"strconv"
)

type UcbPaperResource struct {
	UcbPaperStorage			*UcbDataStorage.UcbPaperStorage
	UcbPaperinputStorage	*UcbDataStorage.UcbPaperinputStorage
	UcbSalesReportStorage	*UcbDataStorage.UcbSalesReportStorage
	UcbPersonnelAssessmentStorage	*UcbDataStorage.UcbPersonnelAssessmentStorage

	UcbScenarioStorage				*UcbDataStorage.UcbScenarioStorage
	UcbProductSalesReportStorage	*UcbDataStorage.UcbProductSalesReportStorage
	UcbGoodsConfigStorage 			*UcbDataStorage.UcbGoodsConfigStorage
	UcbProductConfigStorage			*UcbDataStorage.UcbProductConfigStorage
	UcbProductStorage				*UcbDataStorage.UcbProductStorage
}

func (s UcbPaperResource) NewPaperResource (args []BmDataStorage.BmStorage) *UcbPaperResource {
	var ps *UcbDataStorage.UcbPaperStorage
	var pis *UcbDataStorage.UcbPaperinputStorage
	var srs *UcbDataStorage.UcbSalesReportStorage
	var pas *UcbDataStorage.UcbPersonnelAssessmentStorage

	var ss *UcbDataStorage.UcbScenarioStorage
	var psrs *UcbDataStorage.UcbProductSalesReportStorage
	var gcs  *UcbDataStorage.UcbGoodsConfigStorage
	var pcs  *UcbDataStorage.UcbProductConfigStorage
	var prods   *UcbDataStorage.UcbProductStorage

	for _, arg := range args {
		tp := reflect.ValueOf(arg).Elem().Type()
		if tp.Name() == "UcbPaperStorage" {
			ps = arg.(*UcbDataStorage.UcbPaperStorage)
		} else if tp.Name() == "UcbPaperinputStorage" {
			pis = arg.(*UcbDataStorage.UcbPaperinputStorage)
		} else if tp.Name() == "UcbSalesReportStorage" {
			srs = arg.(*UcbDataStorage.UcbSalesReportStorage)
		} else if tp.Name() == "UcbPersonnelAssessmentStorage" {
			pas = arg.(*UcbDataStorage.UcbPersonnelAssessmentStorage)
		} else if tp.Name() == "UcbProductSalesReportStorage" {
			psrs = arg.(*UcbDataStorage.UcbProductSalesReportStorage)
		} else if tp.Name() == "UcbGoodsConfigStorage" {
			gcs = arg.(*UcbDataStorage.UcbGoodsConfigStorage)
		} else if tp.Name() == "UcbProductConfigStorage" {
			pcs = arg.(*UcbDataStorage.UcbProductConfigStorage)
		} else if tp.Name() == "UcbProductStorage" {
			prods = arg.(*UcbDataStorage.UcbProductStorage)
		} else if tp.Name() == "UcbScenarioStorage" {
			ss = arg.(*UcbDataStorage.UcbScenarioStorage)
		}
	}
	return &UcbPaperResource{
		UcbPaperinputStorage: pis,
		UcbPaperStorage: ps,
		UcbSalesReportStorage: srs,
		UcbPersonnelAssessmentStorage: pas,

		UcbScenarioStorage: ss,
		UcbProductSalesReportStorage: psrs,
		UcbGoodsConfigStorage: gcs,
		UcbProductConfigStorage: pcs,
		UcbProductStorage: prods,
	}
}

func (s UcbPaperResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	chartType, ctOk := r.QueryParams["chart-type"]

	if ctOk && chartType[0] == "product-compete-line" { // 前端应该有以前写好的function，直接返回数据前端展示
		return s.productCompeteLine(r)
	} else if ctOk && chartType[0] == "product-sales-report-summary" { // 处理产品销售报告饼状图数据
		return s.productSalesReportPie(r)
	}

	result := s.UcbPaperStorage.GetAll(r, -1, -1)
	return &Response{Res: result}, nil
}

// PaginatedFindAll can be used to load models in chunks
func (s UcbPaperResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	var (
		result                      []UcbModel.Paper
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
		for _, iter := range s.UcbPaperStorage.GetAll(r, int(start), int(sizeI)) {
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

		for _, iter := range s.UcbPaperStorage.GetAll(r, int(offsetI), int(limitI)) {
			result = append(result, *iter)
		}
	}

	in := UcbModel.Paper{}
	count := s.UcbPaperStorage.Count(r, in)

	return uint(count), &Response{Res: result}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the model with the given ID, otherwise an error
func (s UcbPaperResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	model, err := s.UcbPaperStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	model.Paperinputs = []*UcbModel.Paperinput{}
	r.QueryParams["ids"] = model.InputIDs
	paperInputModels := s.UcbPaperinputStorage.GetAll(r, -1,-1)
	model.Paperinputs = paperInputModels

	model.SalesReports = []*UcbModel.SalesReport{}
	r.QueryParams["ids"] = model.SalesReportIDs
	salesReportModels := s.UcbSalesReportStorage.GetAll(r, -1,-1)
	model.SalesReports = salesReportModels

	model.PersonnelAssessments = []*UcbModel.PersonnelAssessment{}
	r.QueryParams["ids"] = model.PersonnelAssessmentIDs
	personnelAssessmentModels := s.UcbPersonnelAssessmentStorage.GetAll(r, -1,-1)
	model.PersonnelAssessments = personnelAssessmentModels

	return &Response{Res: model}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UcbPaperResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Paper)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id := s.UcbPaperStorage.Insert(model)
	model.ID = id

	return &Response{Res: model, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UcbPaperResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UcbPaperStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the model
func (s UcbPaperResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	model, ok := obj.(UcbModel.Paper)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UcbPaperStorage.Update(model)
	return &Response{Res: model, Code: http.StatusNoContent}, err
}

func (s UcbPaperResource) productCompeteLine(r api2go.Request) (api2go.Responder, error) {
	r.QueryParams["orderby"] = []string{"time"}
	result := s.UcbPaperStorage.GetAll(r, -1, -1)
	curr := result[len(result)-1:]

	for _, model := range curr {
		salesReportIds := model.SalesReportIDs[len(model.SalesReportIDs)-4:] // 写死了
		model.SalesReportIDs = salesReportIds
		r.QueryParams["ids"] = salesReportIds
		salesReportModels := s.UcbSalesReportStorage.GetAll(r, -1,-1)
		model.SalesReports = salesReportModels
	}

	return &Response{Res: curr}, nil
}

func (s UcbPaperResource) productSalesReportPie(r api2go.Request) (api2go.Responder, error) {
	r.QueryParams["orderby"] = []string{"time"}
	result := s.UcbPaperStorage.GetAll(r, -1, -1)
	curr := result[len(result)-1:]

	r.QueryParams["ids"] = curr[0].SalesReportIDs
	salesReportModels := s.UcbSalesReportStorage.GetAll(r, -1,-1)

	srms := salesReportModels[len(salesReportModels)-2:]
	var (
		goodsConfigIds []string
		goodsIds []string
		goodsConfigMapProductConfig map[string]string
		summary []map[string]interface{}
	)

	for _, salesReportModel := range srms {
		r.QueryParams = map[string][]string{}
		goodsConfigMapProductConfig = make(map[string]string)
		summary = []map[string]interface{}{}

		scenarioModel, _ := s.UcbScenarioStorage.GetOne(salesReportModel.ScenarioID)

		r.QueryParams["scenario-id"] = []string{salesReportModel.ScenarioID}

		goodsConfigModels := s.UcbGoodsConfigStorage.GetAll(r, -1, -1)
		for _, goodsConfig := range goodsConfigModels {
			goodsIds = append(goodsIds, goodsConfig.GoodsID)
			goodsConfigMapProductConfig[goodsConfig.GoodsID] = goodsConfig.ID
		}

		r.QueryParams = map[string][]string{}

		r.QueryParams["ids"] = goodsIds
		r.QueryParams["product-type"] = []string{"0"}
		for _, productConfigModel := range s.UcbProductConfigStorage.GetAll(r, -1, -1) {
			productModel, _ := s.UcbProductStorage.GetOne(productConfigModel.ProductID)
			if _, ok := goodsConfigMapProductConfig[productConfigModel.ID]; ok {
				detail := map[string]interface{}{}
				detail["goodsConfigId"] = goodsConfigMapProductConfig[productConfigModel.ID]
				detail["goodsName"] = productModel.Name
				detail["scenarioName"] = scenarioModel.Name
				summary = append(summary, detail)
			}
		}

		for _, v := range summary {
			goodsConfigIds = append(goodsConfigIds, v["goodsConfigId"].(string))
		}

		r.QueryParams = map[string][]string{}

		r.QueryParams["ids"] = salesReportModel.ProductSalesReportIDs
		r.QueryParams["goodsConfigIds"] = goodsConfigIds
		prodReps := s.UcbProductSalesReportStorage.GetAll(r, -1, -1)
		for _, prodRep := range prodReps  {
			for _, v := range summary {
				if v["goodsConfigId"].(string) == prodRep.GoodsConfigID {
					v["sales"] = prodRep.Sales
					v["contribution"] = prodRep.Contribution
				}
			}
		}

		salesReportModel.ProductSalesReportSummary = summary

	}

	curr[0].SalesReports = salesReportModels
	return &Response{Res: curr}, nil
}
