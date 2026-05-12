package wildberries

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/whatcrm/go-wildberries/utils"
)

const (
	DefaultCommonBaseURL          = "https://common-api.wildberries.ru"
	DefaultContentBaseURL         = "https://content-api.wildberries.ru"
	DefaultFeedbacksBaseURL       = "https://feedbacks-api.wildberries.ru"
	DefaultBuyerChatBaseURL       = "https://buyer-chat-api.wildberries.ru"
	DefaultDiscountsBaseURL       = "https://discounts-prices-api.wildberries.ru"
	DefaultMarketplaceBaseURL     = "https://marketplace-api.wildberries.ru"
	DefaultReturnsBaseURL         = "https://returns-api.wildberries.ru"
	DefaultSuppliesBaseURL        = "https://supplies-api.wildberries.ru"
	DefaultFinanceBaseURL         = "https://finance-api.wildberries.ru"
	DefaultStatisticsBaseURL      = "https://statistics-api.wildberries.ru"
	DefaultDocumentsBaseURL       = "https://documents-api.wildberries.ru"
	DefaultSellerAnalyticsBaseURL = "https://seller-analytics-api.wildberries.ru"
	DefaultUserManagementBaseURL  = "https://user-management-api.wildberries.ru"
)

type Client struct {
	HTTPClient *http.Client

	Token string

	CommonBaseURL          string
	ContentBaseURL         string
	FeedbacksBaseURL       string
	BuyerChatBaseURL       string
	DiscountsBaseURL       string
	MarketplaceBaseURL     string
	ReturnsBaseURL         string
	SuppliesBaseURL        string
	FinanceBaseURL         string
	StatisticsBaseURL      string
	DocumentsBaseURL       string
	SellerAnalyticsBaseURL string
	UserManagementBaseURL  string
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("token is required")
	}

	return &Client{
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Token:                  token,
		CommonBaseURL:          DefaultCommonBaseURL,
		ContentBaseURL:         DefaultContentBaseURL,
		FeedbacksBaseURL:       DefaultFeedbacksBaseURL,
		BuyerChatBaseURL:       DefaultBuyerChatBaseURL,
		DiscountsBaseURL:       DefaultDiscountsBaseURL,
		MarketplaceBaseURL:     DefaultMarketplaceBaseURL,
		ReturnsBaseURL:         DefaultReturnsBaseURL,
		SuppliesBaseURL:        DefaultSuppliesBaseURL,
		FinanceBaseURL:         DefaultFinanceBaseURL,
		StatisticsBaseURL:      DefaultStatisticsBaseURL,
		DocumentsBaseURL:       DefaultDocumentsBaseURL,
		SellerAnalyticsBaseURL: DefaultSellerAnalyticsBaseURL,
		UserManagementBaseURL:  DefaultUserManagementBaseURL,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	if req.Header.Get("Content-Type") == "" && req.Method != http.MethodDelete {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-wildberries-sdk")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return utils.ParseAPIError(resp)
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err = io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
