package models

type FBWGood struct {
	Quantity int    `json:"quantity"`
	Barcode  string `json:"barcode"`
}

// OpenAPI component alias: models.Good
type Good = FBWGood

type FBWError struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// OpenAPI component alias: models.ErrorModel
// Used for 400 responses from supplies API.
type ErrorModel struct {
	Status    int    `json:"status"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	RequestID string `json:"requestId"`
	Origin    string `json:"origin"`
}

type FBWWarehouseOption struct {
	WarehouseID   int64 `json:"warehouseID"`
	CanBox        bool  `json:"canBox"`
	CanMonopallet bool  `json:"canMonopallet"`
	CanSupersafe  bool  `json:"canSupersafe"`
	IsBoxOnPallet bool  `json:"isBoxOnPallet"`
}

type FBWOptionsItem struct {
	Barcode    string               `json:"barcode"`
	Error      *FBWError            `json:"error,omitempty"`
	IsError    bool                 `json:"isError,omitempty"`
	Warehouses []FBWWarehouseOption `json:"warehouses"`
}

type FBWAcceptanceOptionsResponse struct {
	Result    []FBWOptionsItem `json:"result"`
	RequestID string           `json:"requestId"`
}

// OpenAPI component alias: models.OptionsResultModel
type OptionsResultModel = FBWAcceptanceOptionsResponse

type FBWWarehouse struct {
	ID              int64  `json:"ID"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	WorkTime        string `json:"workTime"`
	IsActive        bool   `json:"isActive"`
	IsTransitActive bool   `json:"isTransitActive"`
}

// OpenAPI component alias: models.WarehousesResultItems
type WarehousesResultItems = FBWWarehouse

type FBWVolumeTariff struct {
	From  int     `json:"from"`
	To    int     `json:"to"`
	Value float64 `json:"value"`
}

// OpenAPI component alias: models.VolumeTariff
type VolumeTariff = FBWVolumeTariff

type FBWTransitTariff struct {
	TransitWarehouseName     string            `json:"transitWarehouseName"`
	DestinationWarehouseName string            `json:"destinationWarehouseName"`
	ActiveFrom               string            `json:"activeFrom"`
	BoxTariff                []FBWVolumeTariff `json:"boxTariff"`
	PalletTariff             int               `json:"palletTariff"`
}

// OpenAPI component alias: models.TransitTariff
type TransitTariff = FBWTransitTariff

type FBWSuppliesQuery struct {
	Limit  int64 `json:"-"`
	Offset int64 `json:"-"`
}

type FBWDateFilterType string

const (
	FBWDateFilterFactDate   FBWDateFilterType = "factDate"
	FBWDateFilterCreateDate FBWDateFilterType = "createDate"
	FBWDateFilterSupplyDate FBWDateFilterType = "supplyDate"
	FBWDateFilterUpdateDate FBWDateFilterType = "updatedDate"
)

type FBWDateFilter struct {
	From string            `json:"from,omitempty"`
	Till string            `json:"till,omitempty"`
	Type FBWDateFilterType `json:"type"`
}

// OpenAPI component alias: models.DateFilterRequest
type DateFilterRequest = FBWDateFilter

type FBWStatusID int

const (
	FBWSupplyStatusNotPlanned     FBWStatusID = 1
	FBWSupplyStatusPlanned        FBWStatusID = 2
	FBWSupplyStatusAllowUnload    FBWStatusID = 3
	FBWSupplyStatusAccepting      FBWStatusID = 4
	FBWSupplyStatusAccepted       FBWStatusID = 5
	FBWSupplyStatusUnloadedOnGate FBWStatusID = 6
)

type FBWSuppliesFiltersRequest struct {
	Dates     []FBWDateFilter `json:"dates,omitempty"`
	StatusIDs []FBWStatusID   `json:"statusIDs,omitempty"`
}

// OpenAPI component alias: models.SuppliesFiltersRequest
type SuppliesFiltersRequest = FBWSuppliesFiltersRequest

// OpenAPI component alias: models.HandySupplyStatus
type HandySupplyStatus = FBWStatusID

type FBWSupply struct {
	Phone         string  `json:"phone"`
	SupplyID      *int64  `json:"supplyID"`
	PreorderID    int64   `json:"preorderID"`
	CreateDate    string  `json:"createDate"`
	SupplyDate    *string `json:"supplyDate"`
	FactDate      *string `json:"factDate"`
	UpdatedDate   *string `json:"updatedDate"`
	StatusID      int     `json:"statusID"`
	BoxTypeID     int     `json:"boxTypeID"`
	IsBoxOnPallet *bool   `json:"isBoxOnPallet,omitempty"`
}

// OpenAPI component: models.Supply
// NOTE: we can't alias it to `Supply` because `models.Supply` already exists
// for marketplace "supplies" endpoints in this SDK.
type OrdersFBWSupply = FBWSupply

type FBWSupplyDetails struct {
	Phone                     string   `json:"phone"`
	StatusID                  int      `json:"statusID"`
	VirtualTypeID             *int     `json:"virtualTypeID,omitempty"`
	BoxTypeID                 int      `json:"boxTypeID"`
	CreateDate                string   `json:"createDate"`
	SupplyDate                *string  `json:"supplyDate"`
	FactDate                  *string  `json:"factDate"`
	UpdatedDate               *string  `json:"updatedDate"`
	WarehouseID               int64    `json:"warehouseID"`
	WarehouseName             string   `json:"warehouseName"`
	ActualWarehouseID         *int64   `json:"actualWarehouseID"`
	ActualWarehouseName       string   `json:"actualWarehouseName"`
	TransitWarehouseID        *int64   `json:"transitWarehouseID"`
	TransitWarehouseName      string   `json:"transitWarehouseName"`
	AcceptanceCost            *float64 `json:"acceptanceCost"`
	PaidAcceptanceCoefficient *float64 `json:"paidAcceptanceCoefficient"`
	RejectReason              *string  `json:"rejectReason"`
	SupplierAssignName        string   `json:"supplierAssignName"`
	StorageCoef               *string  `json:"storageCoef"`
	DeliveryCoef              *string  `json:"deliveryCoef"`
	Quantity                  int      `json:"quantity"`
	ReadyForSaleQuantity      int      `json:"readyForSaleQuantity"`
	AcceptedQuantity          int      `json:"acceptedQuantity"`
	UnloadingQuantity         int      `json:"unloadingQuantity"`
	DepersonalizedQuantity    *int     `json:"depersonalizedQuantity"`
	IsBoxOnPallet             *bool    `json:"isBoxOnPallet,omitempty"`
}

// OpenAPI component alias: models.SupplyDetails
type SupplyDetails = FBWSupplyDetails

type FBWSupplyDetailsQuery struct {
	IsPreorderID bool `json:"-"`
}

type FBWSupplyGoodsQuery struct {
	Limit        int64 `json:"-"`
	Offset       int64 `json:"-"`
	IsPreorderID bool  `json:"-"`
}

type FBWGoodInSupply struct {
	Barcode              string  `json:"barcode"`
	VendorCode           string  `json:"vendorCode"`
	NMID                 int64   `json:"nmID"`
	NeedKiz              bool    `json:"needKiz"`
	Tnved                *string `json:"tnved"`
	TechSize             string  `json:"techSize"`
	Color                *string `json:"color"`
	SupplierBoxAmount    *int    `json:"supplierBoxAmount"`
	Quantity             int     `json:"quantity"`
	ReadyForSaleQuantity *int    `json:"readyForSaleQuantity"`
	AcceptedQuantity     *int    `json:"acceptedQuantity"`
	UnloadingQuantity    *int    `json:"unloadingQuantity"`
}

// OpenAPI component alias: models.GoodInSupply
type GoodInSupply = FBWGoodInSupply

type FBWGoodInBox struct {
	Barcode  string `json:"barcode"`
	Quantity int    `json:"quantity"`
}

// OpenAPI component alias: models.GoodInBox
type GoodInBox = FBWGoodInBox

type FBWBox struct {
	PackageCode string         `json:"packageCode"`
	Quantity    int            `json:"quantity"`
	Barcodes    []FBWGoodInBox `json:"barcodes"`
}

// OpenAPI component alias: models.Box
type Box = FBWBox
