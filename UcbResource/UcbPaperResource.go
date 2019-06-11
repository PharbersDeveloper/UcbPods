package UcbResource

import (
	"Ucb/UcbDataStorage"
	"Ucb/UcbModel"
	"Ucb/Util/array"
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

	UcbCitySalesReportStorage		*UcbDataStorage.UcbCitySalesReportStorage
	UcbCityStorage					*UcbDataStorage.UcbCityStorage

	UcbRepresentativeSalesReportStorage	*UcbDataStorage.UcbRepresentativeSalesReportStorage
	UcbResourceConfigStorage			*UcbDataStorage.UcbResourceConfigStorage
	UcbRepresentativeConfigStorage		*UcbDataStorage.UcbRepresentativeConfigStorage
	UcbRepresentativeStorage			*UcbDataStorage.UcbRepresentativeStorage

	UcbHospitalSalesReportStorage		*UcbDataStorage.UcbHospitalSalesReportStorage
	UcbDestConfigStorage				*UcbDataStorage.UcbDestConfigStorage
	UcbHospitalConfigStorage			*UcbDataStorage.UcbHospitalConfigStorage
	UcbHospitalStorage					*UcbDataStorage.UcbHospitalStorage
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
	var prods  *UcbDataStorage.UcbProductStorage

	var csr *UcbDataStorage.UcbCitySalesReportStorage
	var cs *UcbDataStorage.UcbCityStorage


	var repsrs *UcbDataStorage.UcbRepresentativeSalesReportStorage
	var rcs *UcbDataStorage.UcbResourceConfigStorage
	var repcs *UcbDataStorage.UcbRepresentativeConfigStorage
	var rs *UcbDataStorage.UcbRepresentativeStorage

	var hsrs *UcbDataStorage.UcbHospitalSalesReportStorage
	var dcs *UcbDataStorage.UcbDestConfigStorage
	var hcs *UcbDataStorage.UcbHospitalConfigStorage
	var hs *UcbDataStorage.UcbHospitalStorage


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
		} else if tp.Name() == "UcbCitySalesReportStorage" {
			csr = arg.(*UcbDataStorage.UcbCitySalesReportStorage)
		} else if tp.Name() == "UcbCityStorage" {
			cs = arg.(*UcbDataStorage.UcbCityStorage)
		} else if tp.Name() == "UcbRepresentativeSalesReportStorage" {
			repsrs = arg.(*UcbDataStorage.UcbRepresentativeSalesReportStorage)
		} else if tp.Name() == "UcbResourceConfigStorage" {
			rcs = arg.(*UcbDataStorage.UcbResourceConfigStorage)
		} else if tp.Name() == "UcbRepresentativeConfigStorage" {
			repcs = arg.(*UcbDataStorage.UcbRepresentativeConfigStorage)
		} else if tp.Name() == "UcbRepresentativeStorage" {
			rs = arg.(*UcbDataStorage.UcbRepresentativeStorage)
		} else if tp.Name() == "UcbHospitalSalesReportStorage" {
			hsrs = arg.(*UcbDataStorage.UcbHospitalSalesReportStorage)
		} else if tp.Name() == "UcbDestConfigStorage" {
			dcs = arg.(*UcbDataStorage.UcbDestConfigStorage)
		} else if tp.Name() == "UcbHospitalConfigStorage" {
			hcs = arg.(*UcbDataStorage.UcbHospitalConfigStorage)
		} else if tp.Name() == "UcbHospitalStorage" {
			hs = arg.(*UcbDataStorage.UcbHospitalStorage)
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

		UcbCitySalesReportStorage: csr,
		UcbCityStorage: cs,

		UcbRepresentativeSalesReportStorage: repsrs,
		UcbResourceConfigStorage: rcs,
		UcbRepresentativeConfigStorage: repcs,
		UcbRepresentativeStorage: rs,

		UcbHospitalSalesReportStorage: hsrs,
		UcbDestConfigStorage: dcs,
		UcbHospitalConfigStorage: hcs,
		UcbHospitalStorage: hs,

	}
}

func (s UcbPaperResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	chartType, ctOk := r.QueryParams["chart-type"]

	if ctOk && chartType[0] == "product-compete-line" { // 前端应该有以前写好的function，直接返回数据前端展示
		return s.productCompeteLine(r)
	} else if ctOk && chartType[0] == "product-sales-report-summary" { // 处理产品销售报告饼状图数据
		return s.productSalesReportSummary(r)
	} else if ctOk && chartType[0] == "city-sales-report-summary" { // 处理城市销售报告饼状图数据
		return s.citySalesReportSummary(r)
	} else if ctOk && chartType[0] == "representative-sales-report-summary" { // 处理代表报告饼状图数据
		return s.representativeSalesReportSummary(r)
	} else if ctOk && chartType[0] == "hospital-sales-report-summary" { // 处理医院报告饼状图数据
		return s.representativeSalesReportSummary(r)
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

func (s UcbPaperResource) productSalesReportSummary(r api2go.Request) (api2go.Responder, error) {
	r.QueryParams["orderby"] = []string{"time"}
	result := s.UcbPaperStorage.GetAll(r, -1, -1)
	curr := result[len(result)-1:]

	r.QueryParams["ids"] = curr[0].SalesReportIDs
	salesReportModels := s.UcbSalesReportStorage.GetAll(r, -1,-1)

	srms := salesReportModels[len(salesReportModels)-2:]

	for _, salesReportModel := range srms {
		var (
			goodsIds []string
			goodsConfigMapProductConfig map[string]string
			summary map[string]interface{}
			detail [] interface{}
		)

		r.QueryParams = map[string][]string{}
		goodsConfigMapProductConfig = make(map[string]string)
		summary = map[string]interface{}{}

		scenarioModel, _ := s.UcbScenarioStorage.GetOne(salesReportModel.ScenarioID)

		summary["scenarioName"] = scenarioModel.Name

		r.QueryParams["scenario-id"] = []string{salesReportModel.ScenarioID}

		goodsConfigModels := s.UcbGoodsConfigStorage.GetAll(r, -1, -1)
		for _, goodsConfig := range goodsConfigModels {
			goodsIds = append(goodsIds, goodsConfig.GoodsID)
			goodsConfigMapProductConfig[goodsConfig.GoodsID] = goodsConfig.ID
		}

		r.QueryParams = map[string][]string{}
		r.QueryParams["ids"] = salesReportModel.ProductSalesReportIDs
		prodReps := s.UcbProductSalesReportStorage.GetAll(r, -1, -1)

		r.QueryParams = map[string][]string{}

		r.QueryParams["ids"] = goodsIds
		r.QueryParams["product-type"] = []string{"0"}

		for _, productConfigModel := range s.UcbProductConfigStorage.GetAll(r, -1, -1) {
			productModel, _ := s.UcbProductStorage.GetOne(productConfigModel.ProductID)
			if v, ok := goodsConfigMapProductConfig[productConfigModel.ID]; ok {
				tempMap := map[string]interface{}{}
				for _, prodRep := range prodReps {
					if v == prodRep.GoodsConfigID {
						tempMap["goodsName"] = productModel.Name
						tempMap["sales"] = prodRep.Sales
						tempMap["contribution"] = prodRep.SalesContribute
					}
				}
				detail = append(detail, tempMap)
			}
		}

		summary["values"] = detail

		salesReportModel.ProductSalesReportSummary = summary

	}

	curr[0].SalesReports = salesReportModels
	return &Response{Res: curr}, nil
}

func (s UcbPaperResource) citySalesReportSummary(r api2go.Request) (api2go.Responder, error) {
	r.QueryParams["orderby"] = []string{"time"}
	result := s.UcbPaperStorage.GetAll(r, -1, -1)
	curr := result[len(result)-1:]

	r.QueryParams["ids"] = curr[0].SalesReportIDs
	salesReportModels := s.UcbSalesReportStorage.GetAll(r, -1,-1)

	srms := salesReportModels[len(salesReportModels) - 2:]

	for _, salesReportModel := range srms {
		var (
			summary map[string]interface{}
			detail [] interface{}
			distinctCityIds []string
			sales float64
			contribution float64
		)

		r.QueryParams = map[string][]string{}
		summary = map[string]interface{}{}

		scenarioModel, _ := s.UcbScenarioStorage.GetOne(salesReportModel.ScenarioID)

		summary["scenarioName"] = scenarioModel.Name

		r.QueryParams["ids"] = salesReportModel.CitySalesReportIDs

		citySalesReports := s.UcbCitySalesReportStorage.GetAll(r, -1,-1)
		for _, citySalesReport := range citySalesReports {
			distinctCityIds = append(distinctCityIds, citySalesReport.CityId)
		}
		distinctCityIds = array.Distinct(distinctCityIds)

		for _, cityId := range distinctCityIds {
			city, _ := s.UcbCityStorage.GetOne(cityId)
			for _, citySalesReport := range citySalesReports {
				if cityId == citySalesReport.CityId {
					sales = sales + citySalesReport.Sales
					contribution = contribution + citySalesReport.SalesContribute
				}
			}

			detail = append(detail, map[string]interface{}{
				"cityName": city.Name,
				"sales": sales,
				"contribution": contribution,
			})
		}
		summary["values"] = detail
		salesReportModel.CitySalesReportSummary = summary
	}

	curr[0].SalesReports = salesReportModels
	return &Response{Res: curr}, nil
}

func (s UcbPaperResource) representativeSalesReportSummary(r api2go.Request) (api2go.Responder, error){
	r.QueryParams["orderby"] = []string{"time"}
	result := s.UcbPaperStorage.GetAll(r, -1, -1)
	curr := result[len(result)-1:]

	r.QueryParams["ids"] = curr[0].SalesReportIDs
	salesReportModels := s.UcbSalesReportStorage.GetAll(r, -1,-1)

	srms := salesReportModels[len(salesReportModels) - 2:]

	for _, salesReportModel := range srms {
		var (
			summary map[string]interface{}
			salesTo3 float64
			salesTo2 float64
			salesTo1 float64
			salesOut float64
			contributionTo3 float64
			contributionTo2 float64
			contributionTo1 float64
			contributionOut float64
			detail []interface{}
		)
		summary = map[string]interface{}{}

		r.QueryParams = map[string][]string{}

		scenarioModel, _ := s.UcbScenarioStorage.GetOne(salesReportModel.ScenarioID)

		summary["scenarioName"] = scenarioModel.Name

		r.QueryParams["ids"] = salesReportModel.HospitalSalesReportIDs

		hospitalSalesReports := s.UcbHospitalSalesReportStorage.GetAll(r, -1,-1)

		r.QueryParams = map[string][]string{}
		r.QueryParams["scenario-id"] = []string{salesReportModel.ScenarioID}
		r.QueryParams["resource-type"] = []string{"1"}

		destConfigs := s.UcbDestConfigStorage.GetAll(r, -1, -1)

		for _, destConfig := range  destConfigs {
			hospitalConfig, _:= s.UcbHospitalConfigStorage.GetOne(destConfig.DestID)
			hospital, _ := s.UcbHospitalStorage.GetOne(hospitalConfig.HospitalID)
			for _, hospitalSalesReport := range hospitalSalesReports {
				if destConfig.ID == hospitalSalesReport.DestConfigID && hospital.HospitalLevel == "三级"{
					salesTo3 = salesTo3 + hospitalSalesReport.Sales
					contributionTo3 = contributionTo3 + hospitalSalesReport.SalesContribute
				} else if destConfig.ID == hospitalSalesReport.DestConfigID && hospital.HospitalLevel == "二级" {
					salesTo2 = salesTo2 + hospitalSalesReport.Sales
					contributionTo2 = contributionTo2 + hospitalSalesReport.SalesContribute
				} else if destConfig.ID == hospitalSalesReport.DestConfigID && hospital.HospitalLevel == "一级" {
					salesTo1 = salesTo1 + hospitalSalesReport.Sales
					contributionTo1 = contributionTo1 + hospitalSalesReport.SalesContribute
				} else if hospitalSalesReport.DestConfigID == "-1" {
					salesOut = salesOut + hospitalSalesReport.Sales
					contributionOut = contributionOut + hospitalSalesReport.SalesContribute
				}
			}
		}

		detail = append(detail, map[string]interface{}{
			"hospitalLevel": "院外",
			"sales": salesOut,
			"contribution": contributionOut,
		})

		detail = append(detail, map[string]interface{}{
			"hospitalLevel": "三级",
			"sales": salesTo3,
			"contribution": contributionTo3,
		})

		detail = append(detail, map[string]interface{}{
			"hospitalLevel": "二级",
			"sales": salesTo2,
			"contribution": contributionTo2,
		})

		detail = append(detail, map[string]interface{}{
			"hospitalLevel": "一级",
			"sales": salesTo1,
			"contribution": contributionTo1,
		})

		summary["values"] = detail

		salesReportModel.HospitalSalesReportSummary = summary
	}

	curr[0].SalesReports = salesReportModels

	return &Response{Res: curr}, nil
}