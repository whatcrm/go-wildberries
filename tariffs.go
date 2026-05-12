package wildberries

import (
	"context"
	"net/http"
	"net/url"

	"github.com/whatcrm/go-wildberries/models"
	"github.com/whatcrm/go-wildberries/utils/common"
)

func (c *Client) GetTariffsCommission(ctx context.Context, locale string) (*models.CommissionResponse, error) {
	requestURL, err := url.Parse(c.CommonBaseURL + common.TariffsCommission)
	if err != nil {
		return nil, err
	}

	if locale != "" {
		params := requestURL.Query()
		params.Set("locale", locale)
		requestURL.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.CommissionResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetTariffsBox(ctx context.Context, date string) (*models.TariffsBoxResponse, error) {
	requestURL, err := url.Parse(c.CommonBaseURL + common.TariffsBox)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	params.Set("date", date)
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.TariffsBoxResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetTariffsPallet(ctx context.Context, date string) (*models.TariffsPalletResponse, error) {
	requestURL, err := url.Parse(c.CommonBaseURL + common.TariffsPallet)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	params.Set("date", date)
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.TariffsPalletResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetTariffsAcceptanceCoefficients(ctx context.Context, warehouseIDs string) ([]models.AcceptanceCoefficient, error) {
	requestURL, err := url.Parse(c.CommonBaseURL + common.TariffsAcceptance)
	if err != nil {
		return nil, err
	}

	if warehouseIDs != "" {
		params := requestURL.Query()
		params.Set("warehouseIDs", warehouseIDs)
		requestURL.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response []models.AcceptanceCoefficient
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetTariffsReturn(ctx context.Context, date string) (*models.ReturnTariffsResponse, error) {
	requestURL, err := url.Parse(c.CommonBaseURL + common.TariffsReturn)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	params.Set("date", date)
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.ReturnTariffsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
