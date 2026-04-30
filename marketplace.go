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
	"github.com/whatcrm/go-wildberries/utils/marketplace"
)

func (c *Client) UpdateStocks(ctx context.Context, warehouseID int64, request models.UpdateStocksRequest) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.StocksByWarehouseEndpoint, warehouseID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) DeleteStocks(ctx context.Context, warehouseID int64, request models.DeleteStocksRequest) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.StocksByWarehouseEndpoint, warehouseID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) GetStocks(ctx context.Context, warehouseID int64, request models.GetStocksRequest) (*models.GetStocksResponse, error) {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.StocksByWarehouseEndpoint, warehouseID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.GetStocksResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetOffices(ctx context.Context) ([]models.Office, error) {
	requestURL := c.MarketplaceBaseURL + marketplace.OfficesEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response []models.Office
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetWarehouses(ctx context.Context) ([]models.Warehouse, error) {
	requestURL := c.MarketplaceBaseURL + marketplace.WarehousesEndpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response []models.Warehouse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateWarehouse(ctx context.Context, request models.CreateWarehouseRequest) (*models.CreateWarehouseResponse, error) {
	requestURL := c.MarketplaceBaseURL + marketplace.WarehousesEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.CreateWarehouseResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateWarehouse(ctx context.Context, warehouseID int64, request models.UpdateWarehouseRequest) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.WarehouseByIDEndpoint, warehouseID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) DeleteWarehouse(ctx context.Context, warehouseID int64) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.WarehouseByIDEndpoint, warehouseID)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL, nil)
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) GetWarehouseContacts(ctx context.Context, warehouseID int64) (*models.WarehouseContactsResponse, error) {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.DBWWarehouseContactsByIDEndpoint, warehouseID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.WarehouseContactsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateWarehouseContacts(ctx context.Context, warehouseID int64, request models.WarehouseContactsRequest) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.DBWWarehouseContactsByIDEndpoint, warehouseID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) GetPassOffices(ctx context.Context) ([]models.PassOffice, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+marketplace.PassesOfficesEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response []models.PassOffice
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetPasses(ctx context.Context) ([]models.Pass, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+marketplace.PassesEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response []models.Pass
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CreatePass(ctx context.Context, request models.UpsertPassRequest) (*models.CreatePassResponse, error) {
	var response models.CreatePassResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.PassesEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) UpdatePass(ctx context.Context, passID int64, request models.UpsertPassRequest) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.PassByIDEndpoint, passID)
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) DeletePass(ctx context.Context, passID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.PassByIDEndpoint, passID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetNewOrders(ctx context.Context) (*models.OrdersNewResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+marketplace.OrdersNewEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response models.OrdersNewResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetOrders(ctx context.Context, query models.OrdersQuery) (*models.OrdersListResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.OrdersEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	if query.Limit > 0 {
		params.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	if query.Next >= 0 {
		params.Set("next", strconv.FormatInt(query.Next, 10))
	}
	if query.DateFrom > 0 {
		params.Set("dateFrom", strconv.FormatInt(query.DateFrom, 10))
	}
	if query.DateTo > 0 {
		params.Set("dateTo", strconv.FormatInt(query.DateTo, 10))
	}
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.OrdersListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetOrdersStatus(ctx context.Context, request models.OrdersIDsRequest) (*models.OrdersStatusResponse, error) {
	var response models.OrdersStatusResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.OrdersStatusEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetReshipmentOrders(ctx context.Context) (*models.ReshipmentOrdersResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+marketplace.OrdersReshipmentEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response models.ReshipmentOrdersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CancelOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.OrderCancelEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetOrdersStickers(ctx context.Context, query models.StickerQuery, request models.OrdersStickersRequest) (*models.OrdersStickersResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.OrdersStickersEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("type", query.Type)
	params.Set("width", strconv.FormatInt(query.Width, 10))
	params.Set("height", strconv.FormatInt(query.Height, 10))
	requestURL.RawQuery = params.Encode()
	var response models.OrdersStickersResponse
	if err = c.postMarketplaceJSON(ctx, requestURL.String(), request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetOrdersMeta(ctx context.Context, request models.OrdersMetaRequest) (*models.OrdersMetaResponse, error) {
	var response models.OrdersMetaResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.OrdersMetaBatchEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteOrderMeta(ctx context.Context, orderID int64, key string) error {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + fmt.Sprintf(marketplace.OrderMetaEndpoint, orderID))
	if err != nil {
		return err
	}
	params := requestURL.Query()
	params.Set("key", key)
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL.String(), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) SetOrderMetaSgtin(ctx context.Context, orderID int64, request models.MetaSgtinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.OrderMetaSgtinEndpoint, orderID), request)
}

func (c *Client) SetOrderMetaUin(ctx context.Context, orderID int64, request models.MetaUinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.OrderMetaUinEndpoint, orderID), request)
}

func (c *Client) SetOrderMetaImei(ctx context.Context, orderID int64, request models.MetaImeiRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.OrderMetaImeiEndpoint, orderID), request)
}

func (c *Client) SetOrderMetaGtin(ctx context.Context, orderID int64, request models.MetaGtinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.OrderMetaGtinEndpoint, orderID), request)
}

func (c *Client) SetOrderMetaExpiration(ctx context.Context, orderID int64, request models.MetaExpirationRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.OrderMetaExpirationEndpoint, orderID), request)
}

func (c *Client) SetOrderMetaCustomsDeclaration(ctx context.Context, orderID int64, request models.MetaCustomsDeclarationRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.OrderMetaCustomsEndpoint, orderID), request)
}

func (c *Client) GetCrossBorderStickers(ctx context.Context, request models.OrdersStickersRequest) (*models.OrdersStickersResponse, error) {
	var response models.OrdersStickersResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.OrdersStickersCrossBorder, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetOrdersStatusHistory(ctx context.Context, request models.OrdersIDsRequest) (*models.OrdersStatusHistoryResponse, error) {
	var response models.OrdersStatusHistoryResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.OrdersStatusHistoryEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetOrdersClientInfo(ctx context.Context, request models.OrdersIDsRequest) (*models.OrdersClientInfoResponse, error) {
	var response models.OrdersClientInfoResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.OrdersClientInfoEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CreateSupply(ctx context.Context, request models.CreateSupplyRequest) (*models.CreateSupplyResponse, error) {
	var response models.CreateSupplyResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.SuppliesEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetSupplies(ctx context.Context, query models.PaginationQuery) (*models.SuppliesResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.SuppliesEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("limit", strconv.FormatInt(query.Limit, 10))
	params.Set("next", strconv.FormatInt(query.Next, 10))
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.SuppliesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) AddOrdersToSupply(ctx context.Context, supplyID string, request models.SupplyOrdersRequest) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.SupplyOrdersPatchEndpoint, supplyID)
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetSupply(ctx context.Context, supplyID string) (*models.Supply, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.SupplyByIDEndpoint, supplyID), nil)
	if err != nil {
		return nil, err
	}
	var response models.Supply
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteSupply(ctx context.Context, supplyID string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.SupplyByIDEndpoint, supplyID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetSupplyOrderIDs(ctx context.Context, supplyID string) (*models.SupplyOrderIDsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.SupplyOrderIDsEndpoint, supplyID), nil)
	if err != nil {
		return nil, err
	}
	var response models.SupplyOrderIDsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeliverSupply(ctx context.Context, supplyID string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.SupplyDeliverEndpoint, supplyID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetSupplyBarcode(ctx context.Context, supplyID, stickerType string) (*models.SupplyBarcodeResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + fmt.Sprintf(marketplace.SupplyBarcodeEndpoint, supplyID))
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("type", stickerType)
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.SupplyBarcodeResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetSupplyTrbx(ctx context.Context, supplyID string) (*models.SupplyTrbxResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.SupplyTrbxEndpoint, supplyID), nil)
	if err != nil {
		return nil, err
	}
	var response models.SupplyTrbxResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) AddSupplyTrbx(ctx context.Context, supplyID string, request models.SupplyTrbxAddRequest) (*models.SupplyTrbxAddResponse, error) {
	var response models.SupplyTrbxAddResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.SupplyTrbxEndpoint, supplyID), request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteSupplyTrbx(ctx context.Context, supplyID string, request models.SupplyTrbxDeleteRequest) error {
	requestURL := c.MarketplaceBaseURL + fmt.Sprintf(marketplace.SupplyTrbxEndpoint, supplyID)
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetSupplyTrbxStickers(ctx context.Context, supplyID, stickerType string, request models.SupplyTrbxStickersRequest) (*models.SupplyTrbxStickersResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + fmt.Sprintf(marketplace.SupplyTrbxStickersEndpoint, supplyID))
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("type", stickerType)
	requestURL.RawQuery = params.Encode()
	var response models.SupplyTrbxStickersResponse
	if err = c.postMarketplaceJSON(ctx, requestURL.String(), request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBWNewOrders(ctx context.Context) (*models.DBWOrdersNewResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+marketplace.DBWOrdersNewEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response models.DBWOrdersNewResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBWOrders(ctx context.Context, query models.DBWOrdersQuery) (*models.DBWOrdersListResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.DBWOrdersEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("limit", strconv.FormatInt(query.Limit, 10))
	params.Set("next", strconv.FormatInt(query.Next, 10))
	params.Set("dateFrom", strconv.FormatInt(query.DateFrom, 10))
	params.Set("dateTo", strconv.FormatInt(query.DateTo, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.DBWOrdersListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBWDeliveryDates(ctx context.Context, request models.DBWDeliveryDatesRequest) (*models.DBWDeliveryDatesResponse, error) {
	var response models.DBWDeliveryDatesResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBWOrdersDeliveryDateEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBWOrdersClientInfo(ctx context.Context, request models.OrdersIDsRequest) (*models.DBWOrdersClientInfoResponse, error) {
	var response models.DBWOrdersClientInfoResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBWOrdersClientEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBWOrdersStatus(ctx context.Context, request models.OrdersIDsRequest) (*models.OrdersStatusResponse, error) {
	var response models.OrdersStatusResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBWOrdersStatusEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ConfirmDBWOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderConfirmEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetDBWOrdersStickers(ctx context.Context, query models.StickerQuery, request models.OrdersStickersRequest) (*models.OrdersStickersResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.DBWOrdersStickersEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("type", query.Type)
	params.Set("width", strconv.FormatInt(query.Width, 10))
	params.Set("height", strconv.FormatInt(query.Height, 10))
	requestURL.RawQuery = params.Encode()

	var response models.OrdersStickersResponse
	if err = c.postMarketplaceJSON(ctx, requestURL.String(), request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) AssembleDBWOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderAssembleEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetDBWOrdersCourierInfo(ctx context.Context, request models.OrdersIDsRequest) (*models.DBWOrdersCourierInfoResponse, error) {
	var response models.DBWOrdersCourierInfoResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBWOrdersCourierEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CancelDBWOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderCancelEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetDBWOrderMeta(ctx context.Context, orderID int64) (*models.DBWOrderMetaResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderMetaEndpoint, orderID), nil)
	if err != nil {
		return nil, err
	}
	var response models.DBWOrderMetaResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteDBWOrderMeta(ctx context.Context, orderID int64, key string) error {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + fmt.Sprintf(marketplace.DBWOrderMetaEndpoint, orderID))
	if err != nil {
		return err
	}
	params := requestURL.Query()
	params.Set("key", key)
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL.String(), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) SetDBWOrderMetaSgtin(ctx context.Context, orderID int64, request models.MetaSgtinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderMetaSgtinEndpoint, orderID), request)
}

func (c *Client) SetDBWOrderMetaUin(ctx context.Context, orderID int64, request models.MetaUinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderMetaUinEndpoint, orderID), request)
}

func (c *Client) SetDBWOrderMetaImei(ctx context.Context, orderID int64, request models.MetaImeiRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderMetaImeiEndpoint, orderID), request)
}

func (c *Client) SetDBWOrderMetaGtin(ctx context.Context, orderID int64, request models.MetaGtinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.DBWOrderMetaGtinEndpoint, orderID), request)
}

func (c *Client) GetDBSNewOrders(ctx context.Context) (*models.DBSOrdersNewResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+marketplace.DBSOrdersNewEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response models.DBSOrdersNewResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBSOrders(ctx context.Context, query models.DBWOrdersQuery) (*models.DBSOrdersListResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.DBSOrdersEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("limit", strconv.FormatInt(query.Limit, 10))
	params.Set("next", strconv.FormatInt(query.Next, 10))
	params.Set("dateFrom", strconv.FormatInt(query.DateFrom, 10))
	params.Set("dateTo", strconv.FormatInt(query.DateTo, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.DBSOrdersListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBSGroupsInfo(ctx context.Context, request models.DBSOrderGroupsRequest) (*models.DBSOrderGroupsResponse, error) {
	var response models.DBSOrderGroupsResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSGroupsInfoEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBSClientInfo(ctx context.Context, request models.OrdersIDsRequest) (*models.DBSClientInfoResponse, error) {
	var response models.DBSClientInfoResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSOrdersClientEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBSB2BClientInfo(ctx context.Context, request models.DBSOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	var response models.DBSStatusSetResponses
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSB2BClientInfoEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBSDeliveryDates(ctx context.Context, request models.DBWDeliveryDatesRequest) (*models.DBWDeliveryDatesResponse, error) {
	var response models.DBWDeliveryDatesResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSOrdersDeliveryDateEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBSStatusesInfo(ctx context.Context, request models.DBSOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	var response models.DBSStatusSetResponses
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSOrdersStatusInfoEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SetDBSStatusCancel(ctx context.Context, request models.DBSOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersStatusCancelEndpoint, request)
}

func (c *Client) SetDBSStatusConfirm(ctx context.Context, request models.DBSOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersStatusConfirmEndpoint, request)
}

func (c *Client) GetDBSStickers(ctx context.Context, stickerType string, width, height int64, request models.OrdersStickersRequest) (*models.OrdersStickersResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.DBSOrdersStickersEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("type", stickerType)
	params.Set("width", strconv.FormatInt(width, 10))
	params.Set("height", strconv.FormatInt(height, 10))
	requestURL.RawQuery = params.Encode()

	var response models.OrdersStickersResponse
	if err = c.postMarketplaceJSON(ctx, requestURL.String(), request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SetDBSStatusDeliver(ctx context.Context, request models.DBSOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersStatusDeliverEndpoint, request)
}

func (c *Client) SetDBSStatusReceive(ctx context.Context, request models.DBSOrdersCodeRequest) (*models.DBSStatusSetResponses, error) {
	var response models.DBSStatusSetResponses
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSOrdersStatusReceiveEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SetDBSStatusReject(ctx context.Context, request models.DBSOrdersCodeRequest) (*models.DBSStatusSetResponses, error) {
	var response models.DBSStatusSetResponses
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSOrdersStatusRejectEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetDBSMetaInfo(ctx context.Context, request models.DBSOrdersRequestV2) (*models.DBSOrdersMetaInfoResponse, error) {
	var response models.DBSOrdersMetaInfoResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.DBSOrdersMetaInfoEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteDBSMeta(ctx context.Context, request models.DBSOrdersMetaDeleteRequest) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersMetaDeleteEndpoint, request)
}

func (c *Client) SetDBSMetaSgtin(ctx context.Context, request models.DBSOrdersSGTINSetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersMetaSgtinEndpoint, request)
}

func (c *Client) SetDBSMetaUin(ctx context.Context, request models.DBSOrdersUINSetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersMetaUinEndpoint, request)
}

func (c *Client) SetDBSMetaImei(ctx context.Context, request models.DBSOrdersIMEISetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersMetaImeiEndpoint, request)
}

func (c *Client) SetDBSMetaGtin(ctx context.Context, request models.DBSOrdersGTINSetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postDBSStatusSetter(ctx, marketplace.DBSOrdersMetaGtinEndpoint, request)
}

func (c *Client) SetDBSMetaCustomsDeclaration(ctx context.Context, request models.DBSOrdersCustomsSetRequest) error {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.MarketplaceBaseURL+marketplace.DBSOrdersMetaCustomsEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetClickCollectNewOrders(ctx context.Context) (*models.OrdersNewResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+marketplace.ClickCollectOrdersNewEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response models.OrdersNewResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SetClickCollectStatusConfirm(ctx context.Context, request models.ClickCollectOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectStatusConfirmV3, request)
}

func (c *Client) SetClickCollectStatusPrepare(ctx context.Context, request models.ClickCollectOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectStatusPrepareV3, request)
}

func (c *Client) ConfirmClickCollectOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderConfirmEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) PrepareClickCollectOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderPrepareEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetClickCollectClientInfo(ctx context.Context, request models.OrdersIDsRequest) (*models.ClickCollectOrderClientInfoResponse, error) {
	var response models.ClickCollectOrderClientInfoResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.ClickCollectClientInfoEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CheckClickCollectClientIdentity(ctx context.Context, request models.ClickCollectCheckIdentityRequest) (*models.ClickCollectCheckedIdentity, error) {
	var response models.ClickCollectCheckedIdentity
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.ClickCollectClientIdentity, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SetClickCollectStatusReceive(ctx context.Context, request models.ClickCollectOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectStatusReceiveV3, request)
}

func (c *Client) SetClickCollectStatusReject(ctx context.Context, request models.ClickCollectOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectStatusRejectV3, request)
}

func (c *Client) ReceiveClickCollectOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderReceiveEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) RejectClickCollectOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderRejectEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetClickCollectStatusesInfo(ctx context.Context, request models.ClickCollectOrdersRequestV2) (*models.ClickCollectOrderStatusesV2, error) {
	var response models.ClickCollectOrderStatusesV2
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.ClickCollectStatusInfoV3, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetClickCollectStatuses(ctx context.Context, request models.OrdersIDsRequest) (*models.ClickCollectOrderStatusResponse, error) {
	var response models.ClickCollectOrderStatusResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.ClickCollectStatusEndpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetClickCollectOrders(ctx context.Context, query models.OrdersQuery) (*models.ClickCollectOrdersListResponse, error) {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + marketplace.ClickCollectOrdersEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	if query.Limit > 0 {
		params.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	if query.Next >= 0 {
		params.Set("next", strconv.FormatInt(query.Next, 10))
	}
	if query.DateFrom > 0 {
		params.Set("dateFrom", strconv.FormatInt(query.DateFrom, 10))
	}
	if query.DateTo > 0 {
		params.Set("dateTo", strconv.FormatInt(query.DateTo, 10))
	}
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.ClickCollectOrdersListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) SetClickCollectStatusCancel(ctx context.Context, request models.ClickCollectOrdersRequestV2) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectStatusCancelV3, request)
}

func (c *Client) CancelClickCollectOrder(ctx context.Context, orderID int64) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderCancelEndpoint, orderID), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) GetClickCollectMetaInfo(ctx context.Context, request models.ClickCollectOrdersRequestV2) (*models.ClickCollectOrdersMetaResponse, error) {
	var response models.ClickCollectOrdersMetaResponse
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+marketplace.ClickCollectMetaInfoV3, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteClickCollectMeta(ctx context.Context, request models.ClickCollectOrdersMetaDeleteRequest) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectMetaDeleteV3, request)
}

func (c *Client) SetClickCollectMetaSgtin(ctx context.Context, request models.DBSOrdersSGTINSetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectMetaSgtinV3, request)
}

func (c *Client) SetClickCollectMetaUin(ctx context.Context, request models.DBSOrdersUINSetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectMetaUinV3, request)
}

func (c *Client) SetClickCollectMetaImei(ctx context.Context, request models.DBSOrdersIMEISetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectMetaImeiV3, request)
}

func (c *Client) SetClickCollectMetaGtin(ctx context.Context, request models.DBSOrdersGTINSetRequest) (*models.DBSStatusSetResponses, error) {
	return c.postClickCollectStatusSetter(ctx, marketplace.ClickCollectMetaGtinV3, request)
}

func (c *Client) GetClickCollectOrderMeta(ctx context.Context, orderID int64) (*models.DBWOrderMetaResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderMetaEndpoint, orderID), nil)
	if err != nil {
		return nil, err
	}
	var response models.DBWOrderMetaResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) DeleteClickCollectOrderMeta(ctx context.Context, orderID int64, key string) error {
	requestURL, err := url.Parse(c.MarketplaceBaseURL + fmt.Sprintf(marketplace.ClickCollectOrderMetaEndpoint, orderID))
	if err != nil {
		return err
	}
	params := requestURL.Query()
	params.Set("key", key)
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL.String(), nil)
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) SetClickCollectOrderMetaSgtin(ctx context.Context, orderID int64, request models.MetaSgtinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderMetaSgtin, orderID), request)
}

func (c *Client) SetClickCollectOrderMetaUin(ctx context.Context, orderID int64, request models.MetaUinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderMetaUin, orderID), request)
}

func (c *Client) SetClickCollectOrderMetaImei(ctx context.Context, orderID int64, request models.MetaImeiRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderMetaImei, orderID), request)
}

func (c *Client) SetClickCollectOrderMetaGtin(ctx context.Context, orderID int64, request models.MetaGtinRequest) error {
	return c.putMarketplaceNoContent(ctx, c.MarketplaceBaseURL+fmt.Sprintf(marketplace.ClickCollectOrderMetaGtin, orderID), request)
}

func (c *Client) postClickCollectStatusSetter(ctx context.Context, endpoint string, payload interface{}) (*models.DBSStatusSetResponses, error) {
	var response models.DBSStatusSetResponses
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+endpoint, payload, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) postDBSStatusSetter(ctx context.Context, endpoint string, payload interface{}) (*models.DBSStatusSetResponses, error) {
	var response models.DBSStatusSetResponses
	if err := c.postMarketplaceJSON(ctx, c.MarketplaceBaseURL+endpoint, payload, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) postMarketplaceJSON(ctx context.Context, endpoint string, payload interface{}, out interface{}) error {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	if err = c.Send(req, out); err != nil {
		return err
	}
	return nil
}

func (c *Client) putMarketplaceNoContent(ctx context.Context, endpoint string, payload interface{}) error {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}
