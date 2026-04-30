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
	"github.com/whatcrm/go-wildberries/utils/ordersfbw"
)

func (c *Client) GetFBWAcceptanceOptions(ctx context.Context, goods []models.FBWGood, warehouseID *int64) (*models.FBWAcceptanceOptionsResponse, error) {
	requestURL, err := url.Parse(c.SuppliesBaseURL + ordersfbw.AcceptanceOptionsEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	if warehouseID != nil {
		params.Set("warehouseID", strconv.FormatInt(*warehouseID, 10))
	}
	requestURL.RawQuery = params.Encode()

	jsonBody, err := json.Marshal(goods)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FBWAcceptanceOptionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetFBWWarehouses(ctx context.Context) ([]models.FBWWarehouse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SuppliesBaseURL+ordersfbw.WarehousesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	var response []models.FBWWarehouse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetFBWTransitTariffs(ctx context.Context) ([]models.FBWTransitTariff, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SuppliesBaseURL+ordersfbw.TransitTariffsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	var response []models.FBWTransitTariff
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetFBWSupplies(ctx context.Context, query models.FBWSuppliesQuery, request models.FBWSuppliesFiltersRequest) ([]models.FBWSupply, error) {
	requestURL, err := url.Parse(c.SuppliesBaseURL + ordersfbw.SuppliesEndpoint)
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

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response []models.FBWSupply
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetFBWSupplyDetails(ctx context.Context, id int64, query models.FBWSupplyDetailsQuery) (*models.FBWSupplyDetails, error) {
	requestURL, err := url.Parse(c.SuppliesBaseURL + fmt.Sprintf(ordersfbw.SupplyByIDEndpoint, id))
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	if query.IsPreorderID {
		params.Set("isPreorderID", strconv.FormatBool(query.IsPreorderID))
	}
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.FBWSupplyDetails
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetFBWSupplyGoods(ctx context.Context, id int64, query models.FBWSupplyGoodsQuery) ([]models.FBWGoodInSupply, error) {
	requestURL, err := url.Parse(c.SuppliesBaseURL + fmt.Sprintf(ordersfbw.SupplyGoodsByIDEndpoint, id))
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
	if query.IsPreorderID {
		params.Set("isPreorderID", strconv.FormatBool(query.IsPreorderID))
	}
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response []models.FBWGoodInSupply
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetFBWSupplyPackage(ctx context.Context, supplyID int64) ([]models.FBWBox, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.SuppliesBaseURL+fmt.Sprintf(ordersfbw.SupplyPackageByIDEndpoint, supplyID), nil)
	if err != nil {
		return nil, err
	}

	var response []models.FBWBox
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}
