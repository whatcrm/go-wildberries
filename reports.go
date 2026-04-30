package wildberries

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/whatcrm/go-wildberries/models"
	"github.com/whatcrm/go-wildberries/utils/reports"
)

func (c *Client) GetSupplierStocks(ctx context.Context, dateFrom string) ([]models.StocksItem, error) {
	reqURL, err := url.Parse(c.StatisticsBaseURL + reports.SupplierStocksEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", dateFrom)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out []models.StocksItem
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetSupplierOrders(ctx context.Context, query models.StatsDateFromQuery) ([]models.OrdersItem, error) {
	reqURL, err := url.Parse(c.StatisticsBaseURL + reports.SupplierOrdersEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", query.DateFrom)
	if query.Flag != nil {
		q.Set("flag", strconv.Itoa(*query.Flag))
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out []models.OrdersItem
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetSupplierSales(ctx context.Context, query models.StatsDateFromQuery) ([]models.SalesItem, error) {
	reqURL, err := url.Parse(c.StatisticsBaseURL + reports.SupplierSalesEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", query.DateFrom)
	if query.Flag != nil {
		q.Set("flag", strconv.Itoa(*query.Flag))
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out []models.SalesItem
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetExciseReport(ctx context.Context, dateFrom, dateTo string, request *models.ExciseReportRequest) (*models.ExciseReportResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.ExciseReportEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", dateFrom)
	q.Set("dateTo", dateTo)
	reqURL.RawQuery = q.Encode()

	var body []byte
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	var out models.ExciseReportResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CreateWarehouseRemainsTask(ctx context.Context, query models.WarehouseRemainsTaskQuery) (*models.CreateTaskResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.WarehouseRemainsCreateTaskEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	if query.Locale != "" {
		q.Set("locale", query.Locale)
	}
	setBool := func(k string, v *bool) {
		if v != nil {
			q.Set(k, strconv.FormatBool(*v))
		}
	}
	setBool("groupByBrand", query.GroupByBrand)
	setBool("groupBySubject", query.GroupBySubject)
	setBool("groupBySa", query.GroupBySa)
	setBool("groupByNm", query.GroupByNm)
	setBool("groupByBarcode", query.GroupByBarcode)
	setBool("groupBySize", query.GroupBySize)
	if query.FilterPics != nil {
		q.Set("filterPics", strconv.Itoa(*query.FilterPics))
	}
	if query.FilterVolume != nil {
		q.Set("filterVolume", strconv.Itoa(*query.FilterVolume))
	}
	reqURL.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.CreateTaskResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetWarehouseRemainsTaskStatus(ctx context.Context, taskID string) (*models.GetTaskStatusResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SellerAnalyticsBaseURL+fmt.Sprintf(reports.WarehouseRemainsTaskStatusEndpoint, taskID), nil)
	if err != nil {
		return nil, err
	}
	var out models.GetTaskStatusResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DownloadWarehouseRemainsReport(ctx context.Context, taskID string) ([]models.WarehouseRemainsItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SellerAnalyticsBaseURL+fmt.Sprintf(reports.WarehouseRemainsTaskDownloadEndpoint, taskID), nil)
	if err != nil {
		return nil, err
	}
	var out []models.WarehouseRemainsItem
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetMeasurementPenalties(ctx context.Context, query models.RetentionsQuery) (*models.MeasurementPenaltiesResponse, error) {
	return c.getRetentions(ctx, reports.MeasurementPenaltiesEndpoint, query)
}

func (c *Client) GetWarehouseMeasurements(ctx context.Context, query models.RetentionsQuery) (*models.WarehouseMeasurementsResponse, error) {
	return c.getRetentions(ctx, reports.WarehouseMeasurementsEndpoint, query)
}

func (c *Client) GetDeductions(ctx context.Context, query models.DeductionsQuery) (*models.DeductionsResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.DeductionsEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	if query.DateFrom != "" {
		q.Set("dateFrom", query.DateFrom)
	}
	q.Set("dateTo", query.DateTo)
	q.Set("limit", strconv.Itoa(query.Limit))
	if query.Offset > 0 {
		q.Set("offset", strconv.Itoa(query.Offset))
	}
	if query.Sort != "" {
		q.Set("sort", query.Sort)
	}
	if query.Order != "" {
		q.Set("order", query.Order)
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.DeductionsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetAntifraudDetails(ctx context.Context, date string) (*models.AntifraudDetailsResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.AntifraudDetailsEndpoint)
	if err != nil {
		return nil, err
	}
	if date != "" {
		q := reqURL.Query()
		q.Set("date", date)
		reqURL.RawQuery = q.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.AntifraudDetailsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetGoodsLabeling(ctx context.Context, dateFrom, dateTo string) (*models.GoodsLabelingResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.GoodsLabelingEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", dateFrom)
	q.Set("dateTo", dateTo)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.GoodsLabelingResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) CreateAcceptanceReportTask(ctx context.Context, dateFrom, dateTo string) (*models.CreateTaskResponse, error) {
	return c.createTaskWithDateRange(ctx, reports.AcceptanceReportCreateTaskEndpoint, dateFrom, dateTo)
}

func (c *Client) GetAcceptanceReportTaskStatus(ctx context.Context, taskID string) (*models.GetTaskStatusResponse, error) {
	return c.getTaskStatus(ctx, reports.AcceptanceReportTaskStatusEndpoint, taskID)
}

func (c *Client) DownloadAcceptanceReport(ctx context.Context, taskID string) ([]models.AcceptanceReportItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SellerAnalyticsBaseURL+fmt.Sprintf(reports.AcceptanceReportTaskDownloadEndpoint, taskID), nil)
	if err != nil {
		return nil, err
	}
	var out []models.AcceptanceReportItem
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) CreatePaidStorageTask(ctx context.Context, dateFrom, dateTo string) (*models.CreateTaskResponse, error) {
	return c.createTaskWithDateRange(ctx, reports.PaidStorageCreateTaskEndpoint, dateFrom, dateTo)
}

func (c *Client) GetPaidStorageTaskStatus(ctx context.Context, taskID string) (*models.GetTaskStatusResponse, error) {
	return c.getTaskStatus(ctx, reports.PaidStorageTaskStatusEndpoint, taskID)
}

func (c *Client) DownloadPaidStorageReport(ctx context.Context, taskID string) ([]models.PaidStorageItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SellerAnalyticsBaseURL+fmt.Sprintf(reports.PaidStorageTaskDownloadEndpoint, taskID), nil)
	if err != nil {
		return nil, err
	}
	var out []models.PaidStorageItem
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) GetRegionSale(ctx context.Context, dateFrom, dateTo string) (*models.RegionSaleResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.RegionSaleEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", dateFrom)
	q.Set("dateTo", dateTo)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.RegionSaleResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetBrandShareBrands(ctx context.Context) (*models.BrandShareBrandsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SellerAnalyticsBaseURL+reports.BrandShareBrandsEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var out models.BrandShareBrandsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetBrandShareParentSubjects(ctx context.Context, locale, brand, dateFrom, dateTo string) (*models.ParentSubjectsResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.BrandShareParentsEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	if locale != "" {
		q.Set("locale", locale)
	}
	q.Set("brand", brand)
	q.Set("dateFrom", dateFrom)
	q.Set("dateTo", dateTo)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.ParentSubjectsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetBrandShareReport(ctx context.Context, parentID int64, brand, dateFrom, dateTo string) (*models.BrandShareResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.BrandShareReportEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("parentId", strconv.FormatInt(parentID, 10))
	q.Set("brand", brand)
	q.Set("dateFrom", dateFrom)
	q.Set("dateTo", dateTo)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.BrandShareResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetBannedProductsBlocked(ctx context.Context, sort, order string) (*models.BannedProductsResponse, error) {
	return c.getBannedProducts(ctx, reports.BannedBlockedEndpoint, sort, order)
}

func (c *Client) GetBannedProductsShadowed(ctx context.Context, sort, order string) (*models.BannedProductsResponse, error) {
	return c.getBannedProducts(ctx, reports.BannedShadowedEndpoint, sort, order)
}

func (c *Client) GetGoodsReturn(ctx context.Context, dateFrom, dateTo string) (*models.GoodsReturnResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + reports.GoodsReturnEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", dateFrom)
	q.Set("dateTo", dateTo)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.GoodsReturnResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) createTaskWithDateRange(ctx context.Context, endpoint, dateFrom, dateTo string) (*models.CreateTaskResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("dateFrom", dateFrom)
	q.Set("dateTo", dateTo)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.CreateTaskResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) getTaskStatus(ctx context.Context, endpoint, taskID string) (*models.GetTaskStatusResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SellerAnalyticsBaseURL+fmt.Sprintf(endpoint, taskID), nil)
	if err != nil {
		return nil, err
	}
	var out models.GetTaskStatusResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) getRetentions(ctx context.Context, endpoint string, query models.RetentionsQuery) (*models.MeasurementPenaltiesResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	if query.DateFrom != "" {
		q.Set("dateFrom", query.DateFrom)
	}
	q.Set("dateTo", query.DateTo)
	q.Set("limit", strconv.Itoa(query.Limit))
	if query.Offset > 0 {
		q.Set("offset", strconv.Itoa(query.Offset))
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.MeasurementPenaltiesResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) getBannedProducts(ctx context.Context, endpoint, sort, order string) (*models.BannedProductsResponse, error) {
	reqURL, err := url.Parse(c.SellerAnalyticsBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("sort", sort)
	q.Set("order", order)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.BannedProductsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
