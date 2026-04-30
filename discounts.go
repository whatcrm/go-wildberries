package wildberries

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/whatcrm/go-wildberries/models"
	"github.com/whatcrm/go-wildberries/utils/discounts"
)

func (c *Client) GetBufferTask(ctx context.Context, uploadID int64) (*models.BufferTaskResponse, error) {
	requestURL, err := url.Parse(c.DiscountsBaseURL + discounts.BufferTasksEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	params.Set("uploadID", strconv.FormatInt(uploadID, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.BufferTaskResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetBufferGoodsTask(ctx context.Context, uploadID, limit, offset int64) (*models.BufferGoodsTaskResponse, error) {
	requestURL, err := url.Parse(c.DiscountsBaseURL + discounts.BufferGoodsTaskEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	params.Set("uploadID", strconv.FormatInt(uploadID, 10))
	params.Set("limit", strconv.FormatInt(limit, 10))
	params.Set("offset", strconv.FormatInt(offset, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.BufferGoodsTaskResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetGoodsFilter(ctx context.Context, query models.GoodsFilterQuery) (*models.GoodsListResponse, error) {
	requestURL, err := url.Parse(c.DiscountsBaseURL + discounts.ListGoodsFilterEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	if query.Limit > 0 {
		params.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	if query.Offset >= 0 {
		params.Set("offset", strconv.FormatInt(query.Offset, 10))
	}
	if query.FilterNmID > 0 {
		params.Set("filterNmID", strconv.FormatInt(query.FilterNmID, 10))
	}
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.GoodsListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetGoodsFilterByNmList(ctx context.Context, request models.GoodsFilterRequest) (*models.GoodsListResponse, error) {
	requestURL := c.DiscountsBaseURL + discounts.ListGoodsFilterEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GoodsListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetGoodsSizeByNm(ctx context.Context, query models.SizeGoodsQuery) (*models.SizeGoodsResponse, error) {
	requestURL, err := url.Parse(c.DiscountsBaseURL + discounts.ListGoodsSizeNMEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	if query.Limit > 0 {
		params.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	if query.Offset >= 0 {
		params.Set("offset", strconv.FormatInt(query.Offset, 10))
	}
	params.Set("nmID", strconv.FormatInt(query.NmID, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.SizeGoodsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetQuarantineGoods(ctx context.Context, query models.QuarantineGoodsQuery) (*models.QuarantineGoodsResponse, error) {
	requestURL, err := url.Parse(c.DiscountsBaseURL + discounts.QuarantineGoodsEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	if query.Limit > 0 {
		params.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	if query.Offset >= 0 {
		params.Set("offset", strconv.FormatInt(query.Offset, 10))
	}
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.QuarantineGoodsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UploadGoodsPricesAndDiscounts(ctx context.Context, request models.GoodsUploadRequest) (*models.TaskCreatedResponse, error) {
	return c.postDiscountsTask(ctx, discounts.UploadTaskEndpoint, request)
}

func (c *Client) UploadSizesPrices(ctx context.Context, request models.SizeUploadRequest) (*models.TaskCreatedResponse, error) {
	return c.postDiscountsTask(ctx, discounts.UploadTaskSizeEndpoint, request)
}

func (c *Client) UploadClubDiscounts(ctx context.Context, request models.ClubDiscountUploadRequest) (*models.TaskCreatedResponse, error) {
	return c.postDiscountsTask(ctx, discounts.UploadTaskClubEndpoint, request)
}

func (c *Client) GetHistoryTask(ctx context.Context, uploadID int64) (*models.TaskHistoryResponse, error) {
	requestURL, err := url.Parse(c.DiscountsBaseURL + discounts.HistoryTasksEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("uploadID", strconv.FormatInt(uploadID, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.TaskHistoryResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetHistoryGoodsTask(ctx context.Context, uploadID, limit, offset int64) (*models.GoodsTaskHistoryResponse, error) {
	requestURL, err := url.Parse(c.DiscountsBaseURL + discounts.HistoryGoodsEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("uploadID", strconv.FormatInt(uploadID, 10))
	params.Set("limit", strconv.FormatInt(limit, 10))
	params.Set("offset", strconv.FormatInt(offset, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.GoodsTaskHistoryResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) postDiscountsTask(ctx context.Context, endpoint string, payload interface{}) (*models.TaskCreatedResponse, error) {
	requestURL := c.DiscountsBaseURL + endpoint
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.TaskCreatedResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
