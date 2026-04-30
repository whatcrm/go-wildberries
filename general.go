package wildberries

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/whatcrm/go-wildberries/models"
	"github.com/whatcrm/go-wildberries/utils/common"
	"github.com/whatcrm/go-wildberries/utils/feedbacks"
)

func (c *Client) GetPing(ctx context.Context) (*models.PingResponse, error) {
	requestURL := c.CommonBaseURL + common.PingEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.PingResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetNews(ctx context.Context, query models.NewsQuery) (*models.NewsResponse, error) {
	if query.From == "" && query.FromID == 0 {
		return nil, fmt.Errorf("one of parameters must be set: from or fromID")
	}

	requestURL, err := url.Parse(c.CommonBaseURL + common.NewsEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	if query.From != "" {
		params.Set("from", query.From)
	}
	if query.FromID != 0 {
		params.Set("fromID", strconv.FormatUint(query.FromID, 10))
	}
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.NewsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetSellerInfo(ctx context.Context) (*models.SellerInfo, error) {
	requestURL := c.CommonBaseURL + common.SellerInfoEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.SellerInfo
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetRating(ctx context.Context) (*models.SupplierRating, error) {
	requestURL := c.FeedbacksBaseURL + feedbacks.RatingEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.SupplierRating
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetSubscriptions(ctx context.Context) (*models.SubscriptionsJamInfo, error) {
	requestURL := c.CommonBaseURL + common.SubscriptionsEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.SubscriptionsJamInfo
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
