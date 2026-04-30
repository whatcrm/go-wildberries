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
	"github.com/whatcrm/go-wildberries/utils/documents"
	"github.com/whatcrm/go-wildberries/utils/finance"
)

// Finance: Balance
func (c *Client) GetFinanceBalance(ctx context.Context) (*models.BalanceResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.FinanceBaseURL+finance.BalanceEndpoint, nil)
	if err != nil {
		return nil, err
	}

	var response models.BalanceResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Finance: Sales reports
func (c *Client) GetSalesReportsList(ctx context.Context, request models.SalesReportListReq) ([]models.SalesReportListRes, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.FinanceBaseURL+finance.SalesReportsListEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response []models.SalesReportListRes
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetSalesReportsDetailedByReportID(ctx context.Context, reportID int64, request models.FinancialReportsDetailedReportIdReq) ([]models.SalesReportsDetailedRes, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqURL := c.FinanceBaseURL + fmt.Sprintf( /* endpoint includes {reportId} */ finance.SalesReportsDetailedByIDEndpoint, reportID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response []models.SalesReportsDetailedRes
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetSalesReportsDetailed(ctx context.Context, request models.SalesReportsDetailedReq) ([]models.SalesReportsDetailedRes, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.FinanceBaseURL+finance.SalesReportsDetailedEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response []models.SalesReportsDetailedRes
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Deprecated: GET /api/v5/supplier/reportDetailByPeriod
func (c *Client) GetSupplierReportDetailByPeriod(ctx context.Context, query models.SupplierReportDetailByPeriodQuery) ([]models.DetailReportItem, error) {
	reqURL, err := url.Parse(c.StatisticsBaseURL + finance.SupplierReportDetailByPeriodEndpoint)
	if err != nil {
		return nil, err
	}

	params := reqURL.Query()
	params.Set("dateFrom", query.DateFrom)
	params.Set("dateTo", query.DateTo)
	if query.Limit > 0 {
		params.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.RrdID > 0 {
		params.Set("rrdid", strconv.Itoa(query.RrdID))
	}
	if query.Period != "" {
		params.Set("period", query.Period)
	}
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response []models.DetailReportItem
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Finance: Acquiring reports
func (c *Client) GetAcquiringReportsList(ctx context.Context, request models.AcquiringReportListReq) ([]models.AcquiringReportListRes, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.FinanceBaseURL+finance.AcquiringListEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response []models.AcquiringReportListRes
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetAcquiringReportsDetailedByReportID(ctx context.Context, reportID int64, request models.FinancialReportsDetailedReportIdReq) ([]models.AcquiringReportsDetailedRes, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqURL := c.FinanceBaseURL + fmt.Sprintf( /* endpoint includes {reportId} */ finance.AcquiringDetailedByIDEndpoint, reportID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response []models.AcquiringReportsDetailedRes
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetAcquiringReportsDetailed(ctx context.Context, request models.AcquiringReportsDetailedReq) ([]models.AcquiringReportsDetailedRes, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.FinanceBaseURL+finance.AcquiringDetailedEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response []models.AcquiringReportsDetailedRes
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Documents: categories
func (c *Client) GetDocumentsCategories(ctx context.Context, locale string) (*models.GetCategoriesResponse, error) {
	reqURL, err := url.Parse(c.DocumentsBaseURL + documents.CategoriesEndpoint)
	if err != nil {
		return nil, err
	}
	params := reqURL.Query()
	if locale == "" {
		locale = "en"
	}
	params.Set("locale", locale)
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.GetCategoriesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Documents: list
func (c *Client) GetDocumentsList(ctx context.Context, query models.DocumentsListQuery) (*models.GetListResponse, error) {
	reqURL, err := url.Parse(c.DocumentsBaseURL + documents.ListEndpoint)
	if err != nil {
		return nil, err
	}
	params := reqURL.Query()

	locale := query.Locale
	if locale == "" {
		locale = "en"
	}
	params.Set("locale", locale)

	if query.BeginTime != "" {
		params.Set("beginTime", query.BeginTime)
	}
	if query.EndTime != "" {
		params.Set("endTime", query.EndTime)
	}
	if query.Sort != "" {
		params.Set("sort", query.Sort)
	}
	if query.Order != "" {
		params.Set("order", query.Order)
	}
	if query.Category != "" {
		params.Set("category", query.Category)
	}
	if query.ServiceName != "" {
		params.Set("serviceName", query.ServiceName)
	}
	if query.Limit > 0 {
		params.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.Offset > 0 {
		params.Set("offset", strconv.Itoa(query.Offset))
	}

	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.GetListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Documents: download one
func (c *Client) DownloadDocument(ctx context.Context, serviceName, extension string) (*models.GetDocResponse, error) {
	reqURL, err := url.Parse(c.DocumentsBaseURL + documents.DownloadEndpoint)
	if err != nil {
		return nil, err
	}
	params := reqURL.Query()
	params.Set("serviceName", serviceName)
	params.Set("extension", extension)
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.GetDocResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Documents: download multiple
func (c *Client) DownloadDocumentsAll(ctx context.Context, request models.RequestDownload) (*models.GetDocsResponse, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.DocumentsBaseURL+documents.DownloadAllEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetDocsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
