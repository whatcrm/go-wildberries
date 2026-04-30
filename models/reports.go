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
		Reports []map[string]interface{} `json:"reports"`
		Total   int                      `json:"total"`
	} `json:"data"`
}

type WarehouseMeasurementsResponse = MeasurementPenaltiesResponse
type DeductionsResponse = MeasurementPenaltiesResponse

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
	Report []map[string]interface{} `json:"report"`
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
	Report []map[string]interface{} `json:"report"`
}

type GoodsReturnResponse struct {
	Report []map[string]interface{} `json:"report"`
}
