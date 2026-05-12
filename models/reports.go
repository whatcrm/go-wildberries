package models

type StatsDateFromQuery struct {
	DateFrom string
	Flag     *int
}

type StocksItem struct {
	LastChangeDate  string  `json:"lastChangeDate"`
	WarehouseName   string  `json:"warehouseName"`
	SupplierArticle string  `json:"supplierArticle"`
	NmID            int64   `json:"nmId"`
	Barcode         string  `json:"barcode"`
	Quantity        int     `json:"quantity"`
	InWayToClient   int     `json:"inWayToClient"`
	InWayFromClient int     `json:"inWayFromClient"`
	QuantityFull    int     `json:"quantityFull"`
	Category        string  `json:"category"`
	Subject         string  `json:"subject"`
	Brand           string  `json:"brand"`
	TechSize        string  `json:"techSize"`
	Price           float64 `json:"Price"`
	Discount        float64 `json:"Discount"`
	IsSupply        bool    `json:"isSupply"`
	IsRealization   bool    `json:"isRealization"`
	SCCode          string  `json:"SCCode"`
}

type OrdersItem struct {
	Date            string  `json:"date"`
	LastChangeDate  string  `json:"lastChangeDate"`
	WarehouseName   string  `json:"warehouseName"`
	WarehouseType   string  `json:"warehouseType"`
	CountryName     string  `json:"countryName"`
	OblastOkrugName string  `json:"oblastOkrugName"`
	RegionName      string  `json:"regionName"`
	SupplierArticle string  `json:"supplierArticle"`
	NmID            int64   `json:"nmId"`
	Barcode         string  `json:"barcode"`
	Category        string  `json:"category"`
	Subject         string  `json:"subject"`
	Brand           string  `json:"brand"`
	TechSize        string  `json:"techSize"`
	IncomeID        int64   `json:"incomeID"`
	IsSupply        bool    `json:"isSupply"`
	IsRealization   bool    `json:"isRealization"`
	TotalPrice      float64 `json:"totalPrice"`
	DiscountPercent int     `json:"discountPercent"`
	Spp             float64 `json:"spp"`
	FinishedPrice   float64 `json:"finishedPrice"`
	PriceWithDisc   float64 `json:"priceWithDisc"`
	IsCancel        bool    `json:"isCancel"`
	CancelDate      string  `json:"cancelDate"`
	Sticker         string  `json:"sticker"`
	GNumber         string  `json:"gNumber"`
	SRID            string  `json:"srid"`
}

type SalesItem struct {
	OrdersItem
	PaymentSaleAmount int     `json:"paymentSaleAmount"`
	ForPay            float64 `json:"forPay"`
	SaleID            string  `json:"saleID"`
}

type ExciseReportRequest struct {
	Countries []string `json:"countries,omitempty"`
}

type ExciseReportResponse struct {
	Response struct {
		Data []ExciseReportItem `json:"data"`
	} `json:"response"`
}

type ExciseReportItem struct {
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	CurrencyNameShort string  `json:"currency_name_short"`
	ExciseShort       string  `json:"excise_short"`
	Barcode           string  `json:"barcode"`
	NmID              int64   `json:"nm_id"`
	OperationTypeID   int     `json:"operation_type_id"`
	FiscalDocNumber   int64   `json:"fiscal_doc_number"`
	FiscalDt          string  `json:"fiscal_dt"`
	FiscalDriveNumber string  `json:"fiscal_drive_number"`
	Rid               int64   `json:"rid"`
	SRID              string  `json:"srid"`
}

type CreateTaskResponse struct {
	Data struct {
		TaskID string `json:"taskId"`
	} `json:"data"`
}

type GetTaskStatusResponse struct {
	Data struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"data"`
}

type WarehouseRemainsTaskQuery struct {
	Locale         string
	GroupByBrand   *bool
	GroupBySubject *bool
	GroupBySa      *bool
	GroupByNm      *bool
	GroupByBarcode *bool
	GroupBySize    *bool
	FilterPics     *int
	FilterVolume   *int
}

type WarehouseRemainsItem struct {
	Brand       string                      `json:"brand"`
	SubjectName string                      `json:"subjectName"`
	VendorCode  string                      `json:"vendorCode"`
	NmID        int64                       `json:"nmId"`
	Barcode     string                      `json:"barcode"`
	TechSize    string                      `json:"techSize"`
	Volume      float64                     `json:"volume"`
	Warehouses  []WarehouseRemainsWarehouse `json:"warehouses"`
}

type WarehouseRemainsWarehouse struct {
	WarehouseName string `json:"warehouseName"`
	Quantity      int    `json:"quantity"`
}

type RetentionsQuery struct {
	DateFrom string
	DateTo   string
	Limit    int
	Offset   int
}

type DeductionsQuery struct {
	DateFrom string
	DateTo   string
	Sort     string
	Order    string
	Limit    int
	Offset   int
}

type MeasurementPenaltiesResponse struct {
	Data struct {
		Reports []MeasurementPenaltyItem `json:"reports"`
		Total   int                      `json:"total"`
	} `json:"data"`
}

type MeasurementPenaltyItem struct {
	NmID           int64    `json:"nmId"`
	SubjectName    string   `json:"subjectName"`
	DimID          int64    `json:"dimId"`
	PrcOver        float64  `json:"prcOver"`
	Volume         float64  `json:"volume"`
	Width          int      `json:"width"`
	Length         int      `json:"length"`
	Height         int      `json:"height"`
	VolumeSup      float64  `json:"volumeSup"`
	WidthSup       int      `json:"widthSup"`
	LengthSup      int      `json:"lengthSup"`
	HeightSup      int      `json:"heightSup"`
	PhotoURLs      []string `json:"photoUrls"`
	DTBonus        string   `json:"dtBonus"`
	IsValid        bool     `json:"isValid"`
	IsValidDT      string   `json:"isValidDt"`
	ReversalAmount float64  `json:"reversalAmount"`
	PenaltyAmount  float64  `json:"penaltyAmount"`
}

type WarehouseMeasurementsResponse struct {
	Data struct {
		Reports []WarehouseMeasurementItem `json:"reports"`
		Total   int                        `json:"total"`
	} `json:"data"`
}

type WarehouseMeasurementItem struct {
	NmID        int64    `json:"nmId"`
	SubjectName string   `json:"subjectName"`
	DimID       int64    `json:"dimId"`
	Volume      float64  `json:"volume"`
	Width       int      `json:"width"`
	Length      int      `json:"length"`
	Height      int      `json:"height"`
	PhotoURLs   []string `json:"photoUrls"`
	DT          string   `json:"dt"`
}

type DeductionsResponse struct {
	Data struct {
		Reports []DeductionItem `json:"reports"`
		Total   int             `json:"total"`
	} `json:"data"`
}

type DeductionItem struct {
	DTBonus       string   `json:"dtBonus"`
	NmID          int64    `json:"nmId"`
	OldShkID      int64    `json:"oldShkId"`
	OldColor      string   `json:"oldColor"`
	OldSize       string   `json:"oldSize"`
	OldSKU        string   `json:"oldSku"`
	OldVendorCode string   `json:"oldVendorCode"`
	NewShkID      int64    `json:"newShkId"`
	NewColor      string   `json:"newColor"`
	NewSize       string   `json:"newSize"`
	NewSKU        string   `json:"newSku"`
	NewVendorCode string   `json:"newVendorCode"`
	BonusSumm     float64  `json:"bonusSumm"`
	BonusType     string   `json:"bonusType"`
	PhotoURLs     []string `json:"photoUrls"`
}

type AntifraudDetailsResponse struct {
	Details []struct {
		NmID     int64  `json:"nmID"`
		Sum      int64  `json:"sum"`
		Currency string `json:"currency"`
		DateFrom string `json:"dateFrom"`
		DateTo   string `json:"dateTo"`
	} `json:"details"`
}

type GoodsLabelingResponse struct {
	Report []GoodsLabelingItem `json:"report"`
}

type GoodsLabelingItem struct {
	Amount    float64  `json:"amount"`
	Date      string   `json:"date"`
	IncomeID  int64    `json:"incomeId"`
	NmID      int64    `json:"nmID"`
	PhotoURLs []string `json:"photoUrls"`
	ShkID     int64    `json:"shkID"`
	SKU       string   `json:"sku"`
}

type AcceptanceReportItem struct {
	Count         int     `json:"count"`
	GICreateDate  string  `json:"giCreateDate"`
	IncomeID      int64   `json:"incomeId"`
	NmID          int64   `json:"nmID"`
	ShkCreateDate string  `json:"shkCreateDate"`
	SubjectName   string  `json:"subjectName"`
	Total         float64 `json:"total"`
}

type PaidStorageItem struct {
	Date             string  `json:"date"`
	LogWarehouseCoef float64 `json:"logWarehouseCoef"`
	OfficeID         int64   `json:"officeId"`
	Warehouse        string  `json:"warehouse"`
	WarehouseCoef    float64 `json:"warehouseCoef"`
	GIID             int64   `json:"giId"`
	ChrtID           int64   `json:"chrtId"`
	Size             string  `json:"size"`
	Barcode          string  `json:"barcode"`
	Subject          string  `json:"subject"`
	Brand            string  `json:"brand"`
	VendorCode       string  `json:"vendorCode"`
	NmID             int64   `json:"nmId"`
	Volume           float64 `json:"volume"`
	CalcType         string  `json:"calcType"`
	WarehousePrice   float64 `json:"warehousePrice"`
	BarcodesCount    int     `json:"barcodesCount"`
	PalletPlaceCode  int     `json:"palletPlaceCode"`
	PalletCount      float64 `json:"palletCount"`
	OriginalDate     string  `json:"originalDate"`
	LoyaltyDiscount  float64 `json:"loyaltyDiscount"`
	TariffFixDate    string  `json:"tariffFixDate"`
	TariffLowerDate  string  `json:"tariffLowerDate"`
}

type RegionSaleResponse struct {
	Report []struct {
		CityName                 string  `json:"cityName"`
		CountryName              string  `json:"countryName"`
		FoName                   string  `json:"foName"`
		NmID                     int64   `json:"nmID"`
		RegionName               string  `json:"regionName"`
		Sa                       string  `json:"sa"`
		SaleInvoiceCostPrice     float64 `json:"saleInvoiceCostPrice"`
		SaleInvoiceCostPricePerc float64 `json:"saleInvoiceCostPricePerc"`
		SaleItemInvoiceQty       int     `json:"saleItemInvoiceQty"`
	} `json:"report"`
}

type BrandShareBrandsResponse struct {
	Data []string `json:"data"`
}

type ParentSubjectsResponse struct {
	Data []struct {
		ParentID   int64  `json:"parentId"`
		ParentName string `json:"parentName"`
	} `json:"data"`
}

type BrandShareResponse struct {
	Report []struct {
		ApplyDate    string  `json:"applyDate"`
		BrandRating  int     `json:"brandRating"`
		PricePercent float64 `json:"pricePercent"`
		QtyPercent   float64 `json:"qtyPercent"`
	} `json:"report"`
}

type BannedProductsResponse struct {
	Report []BannedProductItem `json:"report"`
}

type BannedProductItem struct {
	Brand      string   `json:"brand"`
	NmID       int64    `json:"nmId"`
	Title      string   `json:"title"`
	VendorCode string   `json:"vendorCode"`
	Reason     string   `json:"reason,omitempty"`
	NmRating   *float64 `json:"nmRating,omitempty"`
}

type GoodsReturnResponse struct {
	Report []GoodsReturnItem `json:"report"`
}

type GoodsReturnItem struct {
	Barcode          string  `json:"barcode"`
	Brand            string  `json:"brand"`
	CompletedDT      *string `json:"completedDt"`
	DstOfficeAddress string  `json:"dstOfficeAddress"`
	DstOfficeID      int64   `json:"dstOfficeId"`
	ExpiredDT        *string `json:"expiredDt"`
	IsStatusActive   int     `json:"isStatusActive"`
	NmID             int64   `json:"nmId"`
	OrderDT          string  `json:"orderDt"`
	OrderID          int64   `json:"orderId"`
	ReadyToReturnDT  *string `json:"readyToReturnDt"`
	Reason           string  `json:"reason"`
	ReturnType       string  `json:"returnType"`
	ShkID            int64   `json:"shkId"`
	SRID             string  `json:"srid"`
	Status           string  `json:"status"`
	StickerID        string  `json:"stickerId"`
	SubjectName      string  `json:"subjectName"`
	TechSize         string  `json:"techSize"`
}

type DateRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type TableOrderBy struct {
	Field string `json:"field"`
	Mode  string `json:"mode"`
}

type InventoryRequest struct {
	NmIDs   []int64 `json:"nmIds,omitempty"`
	ChrtIDs []int64 `json:"chrtIds,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	Offset  *int    `json:"offset,omitempty"`
}

type InventoryWbResponseEnvelope struct {
	Data InventoryWbResponse `json:"data"`
}

type InventoryWbResponse struct {
	Items []InventoryWbItem `json:"items"`
}

type InventoryWbItem struct {
	NmID            int64  `json:"nmId"`
	ChrtID          int64  `json:"chrtId"`
	WarehouseID     int64  `json:"warehouseId"`
	WarehouseName   string `json:"warehouseName"`
	RegionName      string `json:"regionName"`
	Quantity        uint64 `json:"quantity"`
	InWayToClient   uint64 `json:"inWayToClient"`
	InWayFromClient uint64 `json:"inWayFromClient"`
}

type CommonReportFilters struct {
	NmIDs               []int64      `json:"nmIDs,omitempty"`
	SubjectIDs          []int32      `json:"subjectIDs,omitempty"`
	BrandNames          []string     `json:"brandNames,omitempty"`
	TagIDs              []int64      `json:"tagIDs,omitempty"`
	CurrentPeriod       DateRange    `json:"currentPeriod"`
	StockType           string       `json:"stockType"`
	SkipDeletedNm       bool         `json:"skipDeletedNm"`
	AvailabilityFilters []string     `json:"availabilityFilters,omitempty"`
	OrderBy             TableOrderBy `json:"orderBy"`
}

type TableGroupRequestSt struct {
	CommonReportFilters
	Limit  *int `json:"limit,omitempty"`
	Offset int  `json:"offset"`
}

type TableGroupResponseEnvelope struct {
	Data TableGroupResponseSt `json:"data"`
}

type TableGroupResponseSt struct {
	Groups   []TableGroupItemSt `json:"groups"`
	Currency string             `json:"currency"`
}

type TableGroupItemSt struct {
	SubjectID   int64                `json:"subjectID"`
	SubjectName string               `json:"subjectName"`
	BrandName   string               `json:"brandName"`
	TagID       int64                `json:"tagID"`
	TagName     string               `json:"tagName"`
	Metrics     TableCommonMetrics   `json:"metrics"`
	Items       []TableProductItemSt `json:"items"`
}

type TableCommonMetrics struct {
	OrdersCount      uint64  `json:"ordersCount"`
	OrdersSum        uint64  `json:"ordersSum"`
	AvgOrders        float64 `json:"avgOrders"`
	BuyoutCount      uint64  `json:"buyoutCount"`
	BuyoutSum        uint64  `json:"buyoutSum"`
	BuyoutPercent    uint32  `json:"buyoutPercent"`
	StockCount       uint64  `json:"stockCount"`
	StockSum         uint64  `json:"stockSum"`
	ToClientCount    uint64  `json:"toClientCount"`
	FromClientCount  uint64  `json:"fromClientCount"`
	LostOrdersCount  float64 `json:"lostOrdersCount"`
	LostOrdersSum    float64 `json:"lostOrdersSum"`
	LostBuyoutsCount float64 `json:"lostBuyoutsCount"`
	LostBuyoutsSum   float64 `json:"lostBuyoutsSum"`
}

type TableProductItemSt struct {
	NmID        int64            `json:"nmID"`
	IsDeleted   bool             `json:"isDeleted"`
	SubjectName string           `json:"subjectName"`
	Name        string           `json:"name"`
	VendorCode  string           `json:"vendorCode"`
	BrandName   string           `json:"brandName"`
	MainPhoto   string           `json:"mainPhoto"`
	HasSizes    bool             `json:"hasSizes"`
	Metrics     ProductMetricsSt `json:"metrics"`
}

type ProductMetricsSt struct {
	TableCommonMetrics
	CurrentPrice struct {
		MinPrice uint64 `json:"minPrice"`
		MaxPrice uint64 `json:"maxPrice"`
	} `json:"currentPrice"`
	Availability string `json:"availability"`
}

type TableProductRequest struct {
	NmIDs               []int64      `json:"nmIDs,omitempty"`
	SubjectID           *int32       `json:"subjectID,omitempty"`
	BrandName           string       `json:"brandName,omitempty"`
	TagID               *int64       `json:"tagID,omitempty"`
	CurrentPeriod       DateRange    `json:"currentPeriod"`
	StockType           string       `json:"stockType"`
	SkipDeletedNm       bool         `json:"skipDeletedNm"`
	OrderBy             TableOrderBy `json:"orderBy"`
	AvailabilityFilters []string     `json:"availabilityFilters,omitempty"`
	Limit               *int         `json:"limit,omitempty"`
	Offset              int          `json:"offset"`
}

type TableProductResponseEnvelope struct {
	Data TableProductResponse `json:"data"`
}

type TableProductResponse struct {
	Items    []TableProductItemSt `json:"items"`
	Currency string               `json:"currency"`
}

type TableSizeRequest struct {
	NmID          int64        `json:"nmID"`
	CurrentPeriod DateRange    `json:"currentPeriod"`
	StockType     string       `json:"stockType"`
	OrderBy       TableOrderBy `json:"orderBy"`
	IncludeOffice bool         `json:"includeOffice"`
}

type TableSizeResponseEnvelope struct {
	Data TableSizeResponse `json:"data"`
}

type TableSizeResponse struct {
	Offices  []TableOfficeItem `json:"offices"`
	Sizes    []TableSizeItem   `json:"sizes"`
	Currency string            `json:"currency"`
}

type TableSizeItem struct {
	Name    string            `json:"name"`
	ChrtID  int64             `json:"chrtID"`
	Offices []TableOfficeItem `json:"offices"`
	Metrics ProductMetricsSt  `json:"metrics"`
}

type TableOfficeItem struct {
	RegionName string             `json:"regionName"`
	OfficeID   int64              `json:"officeID"`
	OfficeName string             `json:"officeName"`
	Metrics    TableCommonMetrics `json:"metrics"`
}

type TableShippingOfficeRequest struct {
	NmIDs         []int64   `json:"nmIDs,omitempty"`
	SubjectIDs    []int32   `json:"subjectIDs,omitempty"`
	BrandNames    []string  `json:"brandNames,omitempty"`
	TagIDs        []int64   `json:"tagIDs,omitempty"`
	CurrentPeriod DateRange `json:"currentPeriod"`
	StockType     string    `json:"stockType"`
	SkipDeletedNm bool      `json:"skipDeletedNm"`
}

type TableShippingOfficeResponseEnvelope struct {
	Data TableShippingOfficeResponse `json:"data"`
}

type TableShippingOfficeResponse struct {
	Regions  []TableShippingOfficeItem `json:"regions"`
	Currency string                    `json:"currency"`
}

type TableShippingOfficeItem struct {
	RegionName string                     `json:"regionName"`
	Metrics    TableShippingOfficeMetrics `json:"metrics"`
	Offices    []TableShippingOffice      `json:"offices"`
}

type TableShippingOffice struct {
	OfficeID   int64                      `json:"officeID"`
	OfficeName string                     `json:"officeName"`
	Metrics    TableShippingOfficeMetrics `json:"metrics"`
}

type TableShippingOfficeMetrics struct {
	StockCount      uint64 `json:"stockCount"`
	StockSum        uint64 `json:"stockSum"`
	ToClientCount   uint64 `json:"toClientCount"`
	FromClientCount uint64 `json:"fromClientCount"`
}
