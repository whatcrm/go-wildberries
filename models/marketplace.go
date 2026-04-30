package models

type UpdateStocksRequest struct {
	Stocks []StockItem `json:"stocks"`
}

type StockItem struct {
	ChrtID int64 `json:"chrtId"`
	Amount int   `json:"amount"`
}

type DeleteStocksRequest struct {
	ChrtIDs []int64 `json:"chrtIds"`
}

type GetStocksRequest struct {
	ChrtIDs []int64 `json:"chrtIds"`
}

type GetStocksResponse struct {
	Stocks []StockItem `json:"stocks"`
}

type Office struct {
	Address         string  `json:"address"`
	Name            string  `json:"name"`
	City            string  `json:"city"`
	ID              int64   `json:"id"`
	Longitude       float64 `json:"longitude"`
	Latitude        float64 `json:"latitude"`
	CargoType       int     `json:"cargoType"`
	DeliveryType    int     `json:"deliveryType"`
	FederalDistrict *string `json:"federalDistrict"`
	Selected        bool    `json:"selected"`
}

type Warehouse struct {
	Name         string `json:"name"`
	OfficeID     int64  `json:"officeId"`
	ID           int64  `json:"id"`
	CargoType    int    `json:"cargoType"`
	DeliveryType int    `json:"deliveryType"`
	IsDeleting   bool   `json:"isDeleting"`
	IsProcessing bool   `json:"isProcessing"`
}

type CreateWarehouseRequest struct {
	Name     string `json:"name"`
	OfficeID int64  `json:"officeId"`
}

type CreateWarehouseResponse struct {
	ID int64 `json:"id"`
}

type UpdateWarehouseRequest struct {
	Name     string `json:"name"`
	OfficeID int64  `json:"officeId"`
}

type WarehouseContactsResponse struct {
	Contacts []WarehouseContact `json:"contacts"`
}

type WarehouseContactsRequest struct {
	Contacts []WarehouseContact `json:"contacts"`
}

type WarehouseContact struct {
	Comment string `json:"comment"`
	Phone   string `json:"phone"`
}

type PaginationQuery struct {
	Limit int64 `json:"-"`
	Next  int64 `json:"-"`
}

type OrdersQuery struct {
	Limit    int64 `json:"-"`
	Next     int64 `json:"-"`
	DateFrom int64 `json:"-"`
	DateTo   int64 `json:"-"`
}

type PassOffice struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	ID      int64  `json:"id"`
}

type Pass struct {
	FirstName     string `json:"firstName"`
	DateEnd       string `json:"dateEnd"`
	LastName      string `json:"lastName"`
	CarModel      string `json:"carModel"`
	CarNumber     string `json:"carNumber"`
	OfficeName    string `json:"officeName"`
	OfficeAddress string `json:"officeAddress"`
	OfficeID      int64  `json:"officeId"`
	ID            int64  `json:"id"`
}

type UpsertPassRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CarModel  string `json:"carModel"`
	CarNumber string `json:"carNumber"`
	OfficeID  int64  `json:"officeId"`
}

type CreatePassResponse struct {
	ID int64 `json:"id"`
}

type OrdersResponse struct {
	Orders []Order `json:"orders"`
}

type OrdersNewResponse struct {
	Orders []OrderNew `json:"orders"`
}

type OrderNew map[string]interface{}
type Order map[string]interface{}

type OrdersListResponse struct {
	Next   int64   `json:"next"`
	Orders []Order `json:"orders"`
}

type OrdersIDsRequest struct {
	Orders []int64 `json:"orders"`
}

type OrdersStatusResponse struct {
	Orders []OrderStatus `json:"orders"`
}

type OrderStatus struct {
	ID             int64  `json:"id"`
	SupplierStatus string `json:"supplierStatus"`
	WBStatus       string `json:"wbStatus"`
}

type ReshipmentOrdersResponse struct {
	Orders []ReshipmentOrder `json:"orders"`
}

type ReshipmentOrder struct {
	SupplyID string `json:"supplyID"`
	OrderID  int64  `json:"orderID"`
}

type StickerQuery struct {
	Type   string `json:"-"`
	Width  int64  `json:"-"`
	Height int64  `json:"-"`
}

type OrdersStickersRequest struct {
	Orders []int64 `json:"orders"`
}

type OrdersStickersResponse struct {
	Stickers []OrderSticker `json:"stickers"`
}

type OrderSticker struct {
	OrderID  int64  `json:"orderId"`
	PartA    string `json:"partA"`
	PartB    string `json:"partB"`
	Barcode  string `json:"barcode"`
	File     string `json:"file"`
	Status   string `json:"status,omitempty"`
	ParcelID string `json:"parcelId,omitempty"`
}

type OrdersMetaRequest struct {
	Orders []int64 `json:"orders"`
}

type OrdersMetaResponse struct {
	Orders []OrderMeta `json:"orders"`
}

type OrderMeta struct {
	ID          int64         `json:"id"`
	MetaDetails []MetaDetails `json:"metaDetails"`
	Meta        interface{}   `json:"meta"`
}

type MetaDetails struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Decision string `json:"decision"`
}

type DeleteOrderMetaQuery struct {
	Key string `json:"-"`
}

type MetaSgtinRequest struct {
	Sgtins []string `json:"sgtins"`
}

type MetaUinRequest struct {
	Uin string `json:"uin"`
}

type MetaImeiRequest struct {
	Imei string `json:"imei"`
}

type MetaGtinRequest struct {
	Gtin string `json:"gtin"`
}

type MetaExpirationRequest struct {
	Expiration string `json:"expiration"`
}

type MetaCustomsDeclarationRequest struct {
	CustomsDeclaration string `json:"customsDeclaration"`
}

type OrdersStatusHistoryResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type OrdersClientInfoResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type CreateSupplyRequest struct {
	Name string `json:"name"`
}

type CreateSupplyResponse struct {
	ID string `json:"id"`
}

type SuppliesResponse struct {
	Next     int64    `json:"next"`
	Supplies []Supply `json:"supplies"`
}

type Supply struct {
	ID                  string `json:"id"`
	IsB2B               *bool  `json:"isB2b"`
	Done                bool   `json:"done"`
	CreatedAt           string `json:"createdAt"`
	ClosedAt            string `json:"closedAt"`
	ScanDt              string `json:"scanDt"`
	Name                string `json:"name"`
	CargoType           int    `json:"cargoType"`
	CrossBorderType     *int   `json:"crossBorderType"`
	DestinationOfficeID *int64 `json:"destinationOfficeId"`
}

type SupplyOrdersRequest struct {
	Orders []int64 `json:"orders"`
}

type SupplyOrderIDsResponse struct {
	OrderIDs []int64 `json:"orderIds"`
}

type SupplyBarcodeQuery struct {
	Type string `json:"-"`
}

type SupplyBarcodeResponse struct {
	Barcode string `json:"barcode"`
	File    string `json:"file"`
}

type SupplyTrbxResponse struct {
	Trbxes []SupplyTrbx `json:"trbxes"`
}

type SupplyTrbx struct {
	ID string `json:"id"`
}

type SupplyTrbxAddRequest struct {
	Amount int `json:"amount"`
}

type SupplyTrbxAddResponse struct {
	TrbxIDs []string `json:"trbxIds"`
}

type SupplyTrbxDeleteRequest struct {
	TrbxIDs []string `json:"trbxIds"`
}

type SupplyTrbxStickersRequest struct {
	TrbxIDs []string `json:"trbxIds"`
}

type SupplyTrbxStickersResponse struct {
	Stickers []TrbxSticker `json:"stickers"`
}

type TrbxSticker struct {
	Barcode string `json:"barcode"`
	File    string `json:"file"`
}

type DBWOrdersQuery struct {
	Limit    int64 `json:"-"`
	Next     int64 `json:"-"`
	DateFrom int64 `json:"-"`
	DateTo   int64 `json:"-"`
}

type DBWOrdersNewResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type DBWOrdersListResponse struct {
	Next   int64                    `json:"next"`
	Orders []map[string]interface{} `json:"orders"`
}

type DBWDeliveryDatesRequest struct {
	Orders []int64 `json:"orders"`
}

type DBWDeliveryDatesResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type DBWOrdersClientInfoResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type DBWOrdersCourierInfoResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type DBWOrderMetaResponse struct {
	Meta map[string]interface{} `json:"meta"`
}

type DBSOrdersNewResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type DBSOrdersListResponse struct {
	Next   int64                    `json:"next"`
	Orders []map[string]interface{} `json:"orders"`
}

type DBSOrderGroupsRequest struct {
	Groups []string `json:"groups"`
}

type DBSOrderGroupsResponse []map[string]interface{}

type DBSClientInfoResponse struct {
	Orders []map[string]interface{} `json:"orders"`
}

type DBSOrdersRequestV2 struct {
	OrdersIDs []int64 `json:"ordersIds"`
}

type DBSOrdersCodeRequest struct {
	Orders []DBSOrderCode `json:"orders"`
}

type DBSOrderCode struct {
	Code    string `json:"code"`
	OrderID int64  `json:"orderId"`
}

type DBSStatusSetResponses struct {
	RequestID string                   `json:"requestId"`
	Results   []DBSStatusSetResultItem `json:"results"`
}

type DBSStatusSetResultItem struct {
	OrderID int64                   `json:"orderId"`
	IsError bool                    `json:"isError"`
	Errors  []DBSBatchErrorResponse `json:"errors"`
	Data    map[string]interface{}  `json:"data"`
}

type DBSBatchErrorResponse struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type DBSOrdersMetaInfoResponse struct {
	Meta []map[string]interface{} `json:"meta"`
}

type DBSOrdersMetaDeleteRequest struct {
	Key      string  `json:"key"`
	OrderIDs []int64 `json:"orderIds"`
}

type DBSOrdersSGTINSetRequest struct {
	Orders []DBSOrderSGTIN `json:"orders"`
}

type DBSOrderSGTIN struct {
	OrderID int64    `json:"orderId"`
	SGTINs  []string `json:"sgtins"`
}

type DBSOrdersUINSetRequest struct {
	Orders []DBSOrderUIN `json:"orders"`
}

type DBSOrderUIN struct {
	OrderID int64  `json:"orderId"`
	UIN     string `json:"uin"`
}

type DBSOrdersIMEISetRequest struct {
	Orders []DBSOrderIMEI `json:"orders"`
}

type DBSOrderIMEI struct {
	OrderID int64  `json:"orderId"`
	IMEI    string `json:"imei"`
}

type DBSOrdersGTINSetRequest struct {
	Orders []DBSOrderGTIN `json:"orders"`
}

type DBSOrderGTIN struct {
	OrderID int64  `json:"orderId"`
	GTIN    string `json:"gtin"`
}

type DBSOrdersCustomsSetRequest struct {
	Orders []DBSOrderCustoms `json:"orders"`
}

type DBSOrderCustoms struct {
	OrderID            int64  `json:"orderId"`
	CustomsDeclaration string `json:"customsDeclaration"`
}

type ClickCollectOrdersRequestV2 struct {
	OrdersIDs []int64 `json:"ordersIds"`
}

type ClickCollectOrdersMetaDeleteRequest struct {
	Key      string  `json:"key"`
	OrdersID []int64 `json:"ordersIds"`
}

type ClickCollectCheckIdentityRequest struct {
	OrderCode string `json:"orderCode"`
	Passcode  string `json:"passcode"`
}

type ClickCollectCheckedIdentity struct {
	Ok bool `json:"ok"`
}

type ClickCollectOrderClientInfo struct {
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	OrderID   int64  `json:"orderID"`
	PhoneCode int64  `json:"phoneCode"`
}

type ClickCollectOrderClientInfoResponse struct {
	Orders []ClickCollectOrderClientInfo `json:"orders"`
}

type ClickCollectOrdersMetaResponse struct {
	Meta []ClickCollectOrderMetaV2 `json:"meta"`
}

type ClickCollectOrderMetaV2 struct {
	Error   string   `json:"error"`
	GTIN    *string  `json:"gtin"`
	IMEI    *string  `json:"imei"`
	OrderID int64    `json:"orderId"`
	SGTIN   []string `json:"sgtin"`
	UIN     *string  `json:"uin"`
}

type ClickCollectOrderStatusV2 struct {
	OrderID        int64                   `json:"orderId"`
	SupplierStatus string                  `json:"supplierStatus"`
	WBStatus       string                  `json:"wbStatus"`
	Errors         []DBSBatchErrorResponse `json:"errors"`
}

type ClickCollectOrderStatusesV2 struct {
	Orders []ClickCollectOrderStatusV2 `json:"orders"`
}

type ClickCollectOrderStatusResponse struct {
	Orders []OrderStatus `json:"orders"`
}

type ClickCollectOrdersListResponse struct {
	Next   int64                    `json:"next"`
	Orders []map[string]interface{} `json:"orders"`
}
