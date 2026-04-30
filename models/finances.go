package models

// Finance: Balance
type BalanceResponse struct {
	Currency    string  `json:"currency"`
	Current     float64 `json:"current"`
	ForWithdraw float64 `json:"for_withdraw"`
}

// Finance: Sales reports
type SalesReportListReq struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`

	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Period string `json:"period,omitempty"`
}

type SalesReportListRes struct {
	ReportID   int64  `json:"reportId"`
	SellerName string `json:"sellerFinanceName"`
	DateFrom   string `json:"dateFrom"`
	DateTo     string `json:"dateTo"`
	CreateDate string `json:"createDate"`
	Currency   string `json:"currency"`
	ReportType int    `json:"reportType"`

	RetailAmountSum             string  `json:"retailAmountSum"`
	ForPaySum                   string  `json:"forPaySum"`
	AvgSalePercent              float64 `json:"avgSalePercent"`
	DeliveryServiceSum          string  `json:"deliveryServiceSum"`
	PaidStorageSum              string  `json:"paidStorageSum"`
	PaidAcceptanceSum           string  `json:"paidAcceptanceSum"`
	DeductionSum                string  `json:"deductionSum"`
	PenaltySum                  string  `json:"penaltySum"`
	AdditionalPaymentSum        string  `json:"additionalPaymentSum"`
	CashbackAmountSum           string  `json:"cashbackAmountSum"`
	CashbackDiscountSum         string  `json:"cashbackDiscountSum"`
	CashbackCommissionChangeSum string  `json:"cashbackCommissionChangeSum"`

	PaymentSchedule string `json:"paymentSchedule"`
	BankPaymentSum  string `json:"bankPaymentSum"`
}

type FinancialReportsDetailedReportIdReq struct {
	Limit  int      `json:"limit,omitempty"`
	RrdID  int      `json:"rrdId,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

type SalesReportsDetailedReq struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`

	Limit  int      `json:"limit,omitempty"`
	RrdID  int      `json:"rrdId,omitempty"`
	Period string   `json:"period,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

// Finance: Detailed sales report item.
// Note: Wildberries may extend the payload; unknown fields will be ignored by the decoder.
type SalesReportsDetailedRes struct {
	ReportID   int64  `json:"reportId"`
	DateFrom   string `json:"dateFrom"`
	DateTo     string `json:"dateTo"`
	CreateDate string `json:"createDate"`
	Currency   string `json:"currency"`
	ReportType int    `json:"reportType"`

	RrdID int64 `json:"rrdId"`
	GiID  int64 `json:"giId"`

	DlvPrc float64 `json:"dlvPrc"`

	FixTariffDateFrom string `json:"fixTariffDateFrom"`
	FixTariffDateTo   string `json:"fixTariffDateTo"`

	SubjectName string `json:"subjectName"`
	NmID        int64  `json:"nmId"`
	BrandName   string `json:"brandName"`
	VendorCode  string `json:"vendorCode"`
	Title       string `json:"title"`
	TechSize    string `json:"techSize"`
	SKU         string `json:"sku"`
	DocTypeName string `json:"docTypeName"`

	Quantity          int64   `json:"quantity"`
	RetailPrice       string  `json:"retailPrice"`
	RetailAmount      string  `json:"retailAmount"`
	SalePercent       int     `json:"salePercent"`
	CommissionPercent float64 `json:"commissionPercent"`

	OfficeName     string `json:"officeName"`
	SellerOperName string `json:"sellerOperName"`
	OrderDt        string `json:"orderDt"`
	SaleDt         string `json:"saleDt"`
	RrDate         string `json:"rrDate"`
	ShkID          int64  `json:"shkId"`

	RetailPriceWithDisc      string  `json:"retailPriceWithDisc"`
	DeliveryAmount           int64   `json:"deliveryAmount"`
	ReturnAmount             int64   `json:"returnAmount"`
	DeliveryService          string  `json:"deliveryService"`
	GiBoxTypeName            string  `json:"giBoxTypeName"`
	ProductDiscountForReport float64 `json:"productDiscountForReport"`
	SellerPromo              string  `json:"sellerPromo"`
	Spp                      float64 `json:"spp"`
	KvwBase                  float64 `json:"kvwBase"`
	Kvw                      float64 `json:"kvw"`
	SupRatingUp              float64 `json:"supRatingUp"`
	IsKgvpV2                 float64 `json:"isKgvpV2"`

	PpvzSalesCommission string `json:"ppvzSalesCommission"`
	ForPay              string `json:"forPay"`
	PpvzReward          string `json:"ppvzReward"`

	AcquiringFee      string  `json:"acquiringFee"`
	AcquiringPercent  float64 `json:"acquiringPercent"`
	PaymentProcessing string  `json:"paymentProcessing"`
	AcquiringBank     string  `json:"acquiringBank"`

	Vw    string `json:"vw"`
	VwNds string `json:"vwNds"`

	PpvzOfficeName   string `json:"ppvzOfficeName"`
	PpvzOfficeID     int64  `json:"ppvzOfficeId"`
	PpvzSupplierName string `json:"ppvzSupplierName"`
	PpvzSupplierInn  string `json:"ppvzSupplierInn"`

	DeclarationNumber string `json:"declarationNumber"`
	StickerID         string `json:"stickerId"`
	Country           string `json:"country"`
	SrvDbs            bool   `json:"srvDbs"`

	Penalty            string `json:"penalty"`
	AdditionalPayment  string `json:"additionalPayment"`
	RebillLogisticCost string `json:"rebillLogisticCost"`
	PaidStorage        string `json:"paidStorage"`
	Deduction          string `json:"deduction"`
	PaidAcceptance     string `json:"paidAcceptance"`

	OrderID int64 `json:"orderId"`
	IsB2b   bool  `json:"isB2b"`

	TrbxID string `json:"trbxId"`

	InstallmentCofinancingAmount string  `json:"installmentCofinancingAmount"`
	WibesDiscountPercent         float64 `json:"wibesDiscountPercent"`

	CashbackAmount           string `json:"cashbackAmount"`
	CashbackDiscount         string `json:"cashbackDiscount"`
	CashbackCommissionChange string `json:"cashbackCommissionChange"`
	PaymentSchedule          string `json:"paymentSchedule"`
	DeliveryMethod           string `json:"deliveryMethod"`

	SellerPromoID       int64   `json:"sellerPromoId"`
	SellerPromoDiscount float64 `json:"sellerPromoDiscount"`

	LoyaltyID       int64   `json:"loyaltyId"`
	LoyaltyDiscount float64 `json:"loyaltyDiscount"`

	UuidPromocode                  string  `json:"uuidPromocode"`
	SalePricePromocodeDiscountPrc  float64 `json:"salePricePromocodeDiscountPrc"`
	ArticleSubstitution            string  `json:"articleSubstitution"`
	SalePriceAffiliatedDiscountPrc float64 `json:"salePriceAffiliatedDiscountPrc"`
	SalePriceWholesaleDiscountPrc  float64 `json:"salePriceWholesaleDiscountPrc"`

	OrderUid string `json:"orderUid"`
	SRID     string `json:"srid"`

	// Optional fields shown in the OpenAPI excerpt.
	AgencyVat     *float64 `json:"agencyVat,omitempty"`
	Kiz           *string  `json:"kiz,omitempty"`
	BonusTypeName *string  `json:"bonusTypeName,omitempty"`
}

// Finance: Acquiring reports
type AcquiringReportListReq struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`

	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type AcquiringReportListRes struct {
	ReportID          int64  `json:"reportId"`
	SellerFinanceName string `json:"sellerFinanceName"`
	DateFrom          string `json:"dateFrom"`
	DateTo            string `json:"dateTo"`
	CreateDate        string `json:"createDate"`
	Currency          string `json:"currency"`

	AcquiringFeeSum    string `json:"acquiringFeeSum"`
	AcquiringFeeVatSum string `json:"acquiringFeeVatSum"`
}

type AcquiringReportsDetailedReq struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`

	Limit  int      `json:"limit,omitempty"`
	RrdID  int      `json:"rrdId,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

type AcquiringReportsDetailedRes struct {
	ReportID     int64  `json:"reportId"`
	Currency     string `json:"currency"`
	RrdID        int64  `json:"rrdId"`
	NmID         int64  `json:"nmId"`
	DocTypeName  string `json:"docTypeName"`
	RetailAmount string `json:"retailAmount"`

	SaleDate string `json:"saleDate"`
	AcqDate  string `json:"acqDate"`
	ShkID    int64  `json:"shkId"`

	AcquiringFee              string `json:"acquiringFee"`
	AcquiringFeeVat           string `json:"acquiringFeeVat"`
	AcquiringBank             string `json:"acquiringBank"`
	Tin                       string `json:"tin"`
	TaxRegistrationReasonCode string `json:"taxRegistrationReasonCode"`
	InvoiceNumber             string `json:"invoiceNumber"`
	InvoiceDate               string `json:"invoiceDate"`
	SRID                      string `json:"srid"`
}

// Deprecated v5 report detail by period item.
type DetailReportItem struct {
	RealizationReportID int64  `json:"realizationreport_id"`
	DateFrom            string `json:"date_from"`
	DateTo              string `json:"date_to"`
	CreateDT            string `json:"create_dt"`
	CurrencyName        string `json:"currency_name"`

	// Optional in schema.
	SupplierContractCode interface{} `json:"suppliercontract_code"`

	RrdID             int64   `json:"rrd_id"`
	GiID              int64   `json:"gi_id"`
	DlvPrc            float64 `json:"dlv_prc"`
	FixTariffDateFrom string  `json:"fix_tariff_date_from"`
	FixTariffDateTo   string  `json:"fix_tariff_date_to"`

	SubjectName string `json:"subject_name"`
	NmID        int64  `json:"nm_id"`
	BrandName   string `json:"brand_name"`
	SaName      string `json:"sa_name"`
	TsName      string `json:"ts_name"`

	Barcode     string `json:"barcode"`
	DocTypeName string `json:"doc_type_name"`
	Quantity    int64  `json:"quantity"`

	RetailPrice  float64 `json:"retail_price"`
	RetailAmount float64 `json:"retail_amount"`

	SalePercent       int     `json:"sale_percent"`
	CommissionPercent float64 `json:"commission_percent"`

	OfficeName       string `json:"office_name"`
	SupplierOperName string `json:"supplier_oper_name"`
	OrderDt          string `json:"order_dt"`
	SaleDt           string `json:"sale_dt"`
	RrDt             string `json:"rr_dt"`
	ShkID            int64  `json:"shk_id"`

	RetailPriceWithdiscRub   float64 `json:"retail_price_withdisc_rub"`
	DeliveryAmount           int64   `json:"delivery_amount"`
	ReturnAmount             int64   `json:"return_amount"`
	DeliveryRub              float64 `json:"delivery_rub"`
	GiBoxTypeName            string  `json:"gi_box_type_name"`
	ProductDiscountForReport float64 `json:"product_discount_for_report"`

	// Keep the rest optional to avoid strictness mismatch.
	Raw map[string]interface{} `json:"-"`
}

type SupplierReportDetailByPeriodQuery struct {
	DateFrom string `json:"dateFrom,omitempty"`
	DateTo   string `json:"dateTo,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	RrdID    int    `json:"rrdid,omitempty"`
	Period   string `json:"period,omitempty"` // weekly|daily
}

// Finance: Documents

type DocumentsCategoriesQuery struct {
	Locale string
}

type GetCategoriesResponse struct {
	Data GetCategoriesData `json:"data"`
}

type GetCategoriesData struct {
	Categories []GetCategoryItem `json:"categories"`
}

type GetCategoryItem struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type DocumentsListQuery struct {
	Locale    string
	BeginTime string `json:"beginTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`

	Sort  string `json:"sort,omitempty"`  // date|category
	Order string `json:"order,omitempty"` // desc|asc

	Category    string `json:"category,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`

	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type GetListResponse struct {
	Data GetListData `json:"data"`
}

type GetListData struct {
	Documents []DocumentListItem `json:"documents"`
}

type DocumentListItem struct {
	ServiceName  string   `json:"serviceName"`
	Name         string   `json:"name"`
	Category     string   `json:"category"`
	Extensions   []string `json:"extensions"`
	CreationTime string   `json:"creationTime"`
	Viewed       bool     `json:"viewed"`
}

type GetDocResponse struct {
	Data GetDocData `json:"data"`
}

type GetDocData struct {
	FileName  string `json:"fileName"`
	Extension string `json:"extension"`
	Document  string `json:"document"`
}

type GetDocsResponse struct {
	Data GetDocsData `json:"data"`
}

type GetDocsData struct {
	FileName  string `json:"fileName"`
	Extension string `json:"extension"`
	Document  string `json:"document"`
}

type RequestDownload struct {
	Params []RequestDownloadParam `json:"params"`
}

type RequestDownloadParam struct {
	Extension   string `json:"extension"`
	ServiceName string `json:"serviceName"`
}
