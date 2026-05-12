package models

type AnalyticsPeriod struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type AnalyticsOrderBy struct {
	Field string `json:"field"`
	Mode  string `json:"mode"`
}

type MainRequest struct {
	CurrentPeriod          AnalyticsPeriod  `json:"currentPeriod"`
	PastPeriod             AnalyticsPeriod  `json:"pastPeriod,omitempty"`
	NmIDs                  []int64          `json:"nmIds,omitempty"`
	SubjectIDs             []int64          `json:"subjectIds,omitempty"`
	BrandNames             []string         `json:"brandNames,omitempty"`
	TagIDs                 []int64          `json:"tagIds,omitempty"`
	PositionCluster        string           `json:"positionCluster"`
	OrderBy                AnalyticsOrderBy `json:"orderBy"`
	IncludeSubstitutedSKUs *bool            `json:"includeSubstitutedSKUs,omitempty"`
	IncludeSearchTexts     *bool            `json:"includeSearchTexts,omitempty"`
	Limit                  int              `json:"limit"`
	Offset                 int              `json:"offset"`
}

type MainResponse struct {
	CommonInfo     AnalyticsCommonInfo     `json:"commonInfo"`
	PositionInfo   AnalyticsPositionInfo   `json:"positionInfo"`
	VisibilityInfo AnalyticsVisibilityInfo `json:"visibilityInfo"`
	Groups         []AnalyticsTableGroup   `json:"groups"`
	Currency       string                  `json:"currency"`
}

type TableGroupRequest struct {
	CurrentPeriod          AnalyticsPeriod  `json:"currentPeriod"`
	PastPeriod             AnalyticsPeriod  `json:"pastPeriod,omitempty"`
	NmIDs                  []int64          `json:"nmIds,omitempty"`
	SubjectIDs             []int64          `json:"subjectIds,omitempty"`
	BrandNames             []string         `json:"brandNames,omitempty"`
	TagIDs                 []int64          `json:"tagIds,omitempty"`
	OrderBy                AnalyticsOrderBy `json:"orderBy"`
	PositionCluster        string           `json:"positionCluster"`
	IncludeSubstitutedSKUs *bool            `json:"includeSubstitutedSKUs,omitempty"`
	IncludeSearchTexts     *bool            `json:"includeSearchTexts,omitempty"`
	Limit                  int              `json:"limit"`
	Offset                 int              `json:"offset"`
}

type TableGroupResponse struct {
	Groups   []AnalyticsTableGroup `json:"groups"`
	Currency string                `json:"currency"`
}

type TableDetailsRequest struct {
	CurrentPeriod          AnalyticsPeriod  `json:"currentPeriod"`
	PastPeriod             AnalyticsPeriod  `json:"pastPeriod,omitempty"`
	SubjectID              *int64           `json:"subjectId,omitempty"`
	BrandName              string           `json:"brandName,omitempty"`
	TagID                  *int64           `json:"tagId,omitempty"`
	NmIDs                  []int64          `json:"nmIds,omitempty"`
	OrderBy                AnalyticsOrderBy `json:"orderBy"`
	PositionCluster        string           `json:"positionCluster"`
	IncludeSubstitutedSKUs *bool            `json:"includeSubstitutedSKUs,omitempty"`
	IncludeSearchTexts     *bool            `json:"includeSearchTexts,omitempty"`
	Limit                  int              `json:"limit"`
	Offset                 int              `json:"offset"`
}

type TableDetailsResponse struct {
	Products []AnalyticsTableProduct `json:"products"`
	Currency string                  `json:"currency"`
}

type ProductSearchTextsRequest struct {
	CurrentPeriod          AnalyticsPeriod  `json:"currentPeriod"`
	PastPeriod             AnalyticsPeriod  `json:"pastPeriod,omitempty"`
	NmIDs                  []int64          `json:"nmIds"`
	TopOrderBy             string           `json:"topOrderBy"`
	IncludeSubstitutedSKUs *bool            `json:"includeSubstitutedSKUs,omitempty"`
	IncludeSearchTexts     *bool            `json:"includeSearchTexts,omitempty"`
	OrderBy                AnalyticsOrderBy `json:"orderBy"`
	Limit                  int              `json:"limit"`
}

type ProductSearchTextsResponse struct {
	Items    []AnalyticsSearchTextItem `json:"items"`
	Currency string                    `json:"currency"`
}

type ProductOrdersRequest struct {
	Period      AnalyticsPeriod `json:"period"`
	NmID        int64           `json:"nmId"`
	SearchTexts []string        `json:"searchTexts"`
}

type ProductOrdersResponse struct {
	Total []AnalyticsProductOrdersMetrics  `json:"total"`
	Items []AnalyticsProductOrdersTextItem `json:"items"`
}

type ProductsRequest struct {
	SelectedPeriod AnalyticsPeriod  `json:"selectedPeriod"`
	PastPeriod     AnalyticsPeriod  `json:"pastPeriod,omitempty"`
	NmIDs          []int64          `json:"nmIds,omitempty"`
	BrandNames     []string         `json:"brandNames,omitempty"`
	SubjectIDs     []int64          `json:"subjectIds,omitempty"`
	TagIDs         []int64          `json:"tagIds,omitempty"`
	SkipDeletedNm  *bool            `json:"skipDeletedNm,omitempty"`
	OrderBy        AnalyticsOrderBy `json:"orderBy,omitempty"`
	Limit          *int             `json:"limit,omitempty"`
	Offset         *int             `json:"offset,omitempty"`
}

type ProductHistoryRequest struct {
	SelectedPeriod   AnalyticsPeriod `json:"selectedPeriod"`
	NmIDs            []int64         `json:"nmIds"`
	SkipDeletedNm    *bool           `json:"skipDeletedNm,omitempty"`
	AggregationLevel string          `json:"aggregationLevel,omitempty"`
}

type GroupedHistoryRequest struct {
	SelectedPeriod   AnalyticsPeriod `json:"selectedPeriod"`
	BrandNames       []string        `json:"brandNames,omitempty"`
	SubjectIDs       []int64         `json:"subjectIds,omitempty"`
	TagIDs           []int64         `json:"tagIds,omitempty"`
	SkipDeletedNm    *bool           `json:"skipDeletedNm,omitempty"`
	AggregationLevel string          `json:"aggregationLevel,omitempty"`
}

type ProductsResponse struct {
	Products []AnalyticsProductWithStatistic `json:"products"`
	Currency string                          `json:"currency"`
}

type ProductHistoryResponse []AnalyticsProductHistoryItem
type GroupedHistoryResponse []AnalyticsProductHistoryItem

type NmReportRetryReportRequest struct {
	DownloadID string `json:"downloadId"`
}

type NmReportCreateReportResponse struct {
	Data string `json:"data"`
}

type NmReportGetReportsResponse struct {
	Data []NmReportGetReportsItem `json:"data"`
}

type NmReportGetReportsItem struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type NmReportRetryReportResponse struct {
	Data string `json:"data"`
}

type MainResponseEnvelope struct {
	Data MainResponse `json:"data"`
}

type TableGroupResponseEnvelopeV2 struct {
	Data TableGroupResponse `json:"data"`
}

type TableDetailsResponseEnvelope struct {
	Data TableDetailsResponse `json:"data"`
}

type ProductSearchTextsResponseEnvelope struct {
	Data ProductSearchTextsResponse `json:"data"`
}

type ProductOrdersResponseEnvelope struct {
	Data ProductOrdersResponse `json:"data"`
}

type ProductsResponseEnvelope struct {
	Data ProductsResponse `json:"data"`
}

type GroupedHistoryResponseEnvelope struct {
	Data GroupedHistoryResponse `json:"data"`
}

type SalesFunnelCSVRequest struct {
	ID             string      `json:"id"`
	ReportType     string      `json:"reportType"`
	UserReportName string      `json:"userReportName,omitempty"`
	Params         interface{} `json:"params"`
}

type AnalyticsMetric struct {
	Current  int `json:"current"`
	Dynamics int `json:"dynamics,omitempty"`
}

type AnalyticsCommonInfo struct {
	SupplierRating struct {
		Current  float64 `json:"current"`
		Dynamics float64 `json:"dynamics,omitempty"`
	} `json:"supplierRating"`
	AdvertisedProducts AnalyticsMetric `json:"advertisedProducts"`
	TotalProducts      int             `json:"totalProducts"`
}

type AnalyticsPositionChartItem struct {
	DT      string `json:"dt"`
	Average int    `json:"average"`
	Median  int    `json:"median"`
}

type AnalyticsPositionClusters struct {
	FirstHundred  AnalyticsMetric `json:"firstHundred"`
	SecondHundred AnalyticsMetric `json:"secondHundred"`
	Below         AnalyticsMetric `json:"below"`
}

type AnalyticsPositionInfo struct {
	Average    AnalyticsMetric              `json:"average"`
	Median     AnalyticsMetric              `json:"median"`
	ChartItems []AnalyticsPositionChartItem `json:"chartItems"`
	Clusters   AnalyticsPositionClusters    `json:"clusters"`
}

type AnalyticsVisibilityChartItem struct {
	DT         string `json:"dt"`
	Visibility int    `json:"visibility"`
	Open       int    `json:"open"`
}

type AnalyticsVisibilityInfo struct {
	Visibility AnalyticsMetric                `json:"visibility"`
	OpenCard   AnalyticsMetric                `json:"openCard"`
	ByDay      []AnalyticsVisibilityChartItem `json:"byDay"`
	ByWeek     []AnalyticsVisibilityChartItem `json:"byWeek"`
	ByMonth    []AnalyticsVisibilityChartItem `json:"byMonth"`
}

type AnalyticsPriceRange struct {
	MinPrice int `json:"minPrice"`
	MaxPrice int `json:"maxPrice"`
}

type AnalyticsMetricWithPercentile struct {
	Current    int `json:"current"`
	Dynamics   int `json:"dynamics,omitempty"`
	Percentile int `json:"percentile,omitempty"`
}

type AnalyticsProductMetrics struct {
	AvgPosition AnalyticsMetric `json:"avgPosition"`
	OpenCard    AnalyticsMetric `json:"openCard"`
	AddToCart   AnalyticsMetric `json:"addToCart"`
	OpenToCart  AnalyticsMetric `json:"openToCart"`
	Orders      AnalyticsMetric `json:"orders"`
	CartToOrder AnalyticsMetric `json:"cartToOrder"`
	Visibility  AnalyticsMetric `json:"visibility"`
}

type AnalyticsTableProduct struct {
	NmID             int64               `json:"nmId"`
	Name             string              `json:"name"`
	VendorCode       string              `json:"vendorCode"`
	SubjectName      string              `json:"subjectName"`
	BrandName        string              `json:"brandName"`
	MainPhoto        string              `json:"mainPhoto"`
	IsAdvertised     bool                `json:"isAdvertised"`
	IsSubstitutedSKU bool                `json:"isSubstitutedSKU,omitempty"`
	IsCardRated      bool                `json:"isCardRated"`
	Rating           float64             `json:"rating"`
	FeedbackRating   float64             `json:"feedbackRating"`
	Price            AnalyticsPriceRange `json:"price"`
	AnalyticsProductMetrics
}

type AnalyticsTableGroup struct {
	SubjectName string                  `json:"subjectName"`
	SubjectID   int64                   `json:"subjectId"`
	BrandName   string                  `json:"brandName"`
	TagName     string                  `json:"tagName"`
	TagID       int64                   `json:"tagId"`
	Metrics     AnalyticsProductMetrics `json:"metrics"`
	Items       []AnalyticsTableProduct `json:"items"`
}

type AnalyticsSearchTextItem struct {
	Text           string                        `json:"text"`
	NmID           int64                         `json:"nmId"`
	SubjectName    string                        `json:"subjectName"`
	BrandName      string                        `json:"brandName"`
	VendorCode     string                        `json:"vendorCode"`
	Name           string                        `json:"name"`
	IsCardRated    bool                          `json:"isCardRated"`
	Rating         float64                       `json:"rating"`
	FeedbackRating float64                       `json:"feedbackRating"`
	Price          AnalyticsPriceRange           `json:"price"`
	Frequency      AnalyticsMetric               `json:"frequency"`
	WeekFrequency  int                           `json:"weekFrequency"`
	MedianPosition AnalyticsMetric               `json:"medianPosition"`
	AvgPosition    AnalyticsMetric               `json:"avgPosition"`
	OpenCard       AnalyticsMetricWithPercentile `json:"openCard"`
	AddToCart      AnalyticsMetricWithPercentile `json:"addToCart"`
	OpenToCart     AnalyticsMetricWithPercentile `json:"openToCart"`
	Orders         AnalyticsMetricWithPercentile `json:"orders"`
	CartToOrder    AnalyticsMetricWithPercentile `json:"cartToOrder"`
	Visibility     AnalyticsMetric               `json:"visibility"`
}

type AnalyticsProductOrdersMetrics struct {
	DT          string `json:"dt"`
	AvgPosition int    `json:"avgPosition"`
	Orders      int    `json:"orders"`
}

type AnalyticsProductOrdersTextItem struct {
	Text      string                          `json:"text"`
	Frequency int                             `json:"frequency"`
	DateItems []AnalyticsProductOrdersMetrics `json:"dateItems"`
}

type AnalyticsTag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type AnalyticsProduct struct {
	NmID           int64          `json:"nmId"`
	Title          string         `json:"title"`
	VendorCode     string         `json:"vendorCode"`
	BrandName      string         `json:"brandName"`
	SubjectID      int64          `json:"subjectId"`
	SubjectName    string         `json:"subjectName"`
	Tags           []AnalyticsTag `json:"tags"`
	ProductRating  float64        `json:"productRating"`
	FeedbackRating float64        `json:"feedbackRating"`
	Stocks         struct {
		WB         int `json:"wb"`
		MP         int `json:"mp"`
		BalanceSum int `json:"balanceSum"`
	} `json:"stocks"`
}

type AnalyticsTimeToReady struct {
	Days  int `json:"days"`
	Hours int `json:"hours"`
	Mins  int `json:"mins"`
}

type AnalyticsConversions struct {
	AddToCartPercent   int `json:"addToCartPercent"`
	CartToOrderPercent int `json:"cartToOrderPercent"`
	BuyoutPercent      int `json:"buyoutPercent"`
}

type AnalyticsWbClubMetrics struct {
	OrderCount          int     `json:"orderCount"`
	OrderSum            int     `json:"orderSum"`
	BuyoutSum           int     `json:"buyoutSum"`
	BuyoutCount         int     `json:"buyoutCount"`
	CancelSum           int     `json:"cancelSum"`
	CancelCount         int     `json:"cancelCount"`
	AvgPrice            int     `json:"avgPrice"`
	BuyoutPercent       int     `json:"buyoutPercent"`
	AvgOrderCountPerDay float64 `json:"avgOrderCountPerDay"`
}

type AnalyticsStatistic struct {
	Period               AnalyticsPeriod        `json:"period"`
	OpenCount            int                    `json:"openCount"`
	CartCount            int                    `json:"cartCount"`
	OrderCount           int                    `json:"orderCount"`
	OrderSum             int                    `json:"orderSum"`
	BuyoutCount          int                    `json:"buyoutCount"`
	BuyoutSum            int                    `json:"buyoutSum"`
	CancelCount          int                    `json:"cancelCount"`
	CancelSum            int                    `json:"cancelSum"`
	AvgPrice             int                    `json:"avgPrice"`
	AvgOrdersCountPerDay float64                `json:"avgOrdersCountPerDay"`
	ShareOrderPercent    float64                `json:"shareOrderPercent"`
	AddToWishlist        int                    `json:"addToWishlist"`
	TimeToReady          AnalyticsTimeToReady   `json:"timeToReady"`
	LocalizationPercent  int                    `json:"localizationPercent"`
	WBClub               AnalyticsWbClubMetrics `json:"wbClub"`
	Conversions          AnalyticsConversions   `json:"conversions"`
}

type AnalyticsComparison struct {
	OpenCountDynamic            int                  `json:"openCountDynamic"`
	CartCountDynamic            int                  `json:"cartCountDynamic"`
	OrderCountDynamic           int                  `json:"orderCountDynamic"`
	OrderSumDynamic             int                  `json:"orderSumDynamic"`
	BuyoutCountDynamic          int                  `json:"buyoutCountDynamic"`
	BuyoutSumDynamic            int                  `json:"buyoutSumDynamic"`
	CancelCountDynamic          int                  `json:"cancelCountDynamic"`
	CancelSumDynamic            int                  `json:"cancelSumDynamic"`
	AvgOrdersCountPerDayDynamic int                  `json:"avgOrdersCountPerDayDynamic"`
	AvgPriceDynamic             int                  `json:"avgPriceDynamic"`
	ShareOrderPercentDynamic    int                  `json:"shareOrderPercentDynamic"`
	AddToWishlistDynamic        int                  `json:"addToWishlistDynamic"`
	TimeToReadyDynamic          AnalyticsTimeToReady `json:"timeToReadyDynamic"`
	LocalizationPercentDynamic  int                  `json:"localizationPercentDynamic"`
	Conversions                 AnalyticsConversions `json:"conversions"`
}

type AnalyticsProductWithStatistic struct {
	Product   AnalyticsProduct `json:"product"`
	Statistic struct {
		Selected   AnalyticsStatistic  `json:"selected"`
		Past       AnalyticsStatistic  `json:"past"`
		Comparison AnalyticsComparison `json:"comparison"`
	} `json:"statistic"`
}

type AnalyticsHistoryItem struct {
	Date                  string `json:"date"`
	OpenCount             int    `json:"openCount"`
	CartCount             int    `json:"cartCount"`
	OrderCount            int    `json:"orderCount"`
	OrderSum              int    `json:"orderSum"`
	BuyoutCount           int    `json:"buyoutCount"`
	BuyoutSum             int    `json:"buyoutSum"`
	BuyoutPercent         int    `json:"buyoutPercent"`
	AddToCartConversion   int    `json:"addToCartConversion"`
	CartToOrderConversion int    `json:"cartToOrderConversion"`
	AddToWishlistCount    int    `json:"addToWishlistCount"`
}

type AnalyticsHistoryProduct struct {
	NmID        int64  `json:"nmId"`
	Title       string `json:"title"`
	VendorCode  string `json:"vendorCode"`
	BrandName   string `json:"brandName"`
	SubjectID   int64  `json:"subjectId"`
	SubjectName string `json:"subjectName"`
}

type AnalyticsProductHistoryItem struct {
	Product  AnalyticsHistoryProduct `json:"product"`
	History  []AnalyticsHistoryItem  `json:"history"`
	Currency string                  `json:"currency"`
}
