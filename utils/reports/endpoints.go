package reports

const (
	SupplierStocksEndpoint = "/api/v1/supplier/stocks"
	SupplierOrdersEndpoint = "/api/v1/supplier/orders"
	SupplierSalesEndpoint  = "/api/v1/supplier/sales"

	ExciseReportEndpoint = "/api/v1/analytics/excise-report"

	WarehouseRemainsCreateTaskEndpoint   = "/api/v1/warehouse_remains"
	WarehouseRemainsTaskStatusEndpoint   = "/api/v1/warehouse_remains/tasks/%s/status"
	WarehouseRemainsTaskDownloadEndpoint = "/api/v1/warehouse_remains/tasks/%s/download"

	MeasurementPenaltiesEndpoint  = "/api/analytics/v1/measurement-penalties"
	WarehouseMeasurementsEndpoint = "/api/analytics/v1/warehouse-measurements"
	DeductionsEndpoint            = "/api/analytics/v1/deductions"
	AntifraudDetailsEndpoint      = "/api/v1/analytics/antifraud-details"
	GoodsLabelingEndpoint         = "/api/v1/analytics/goods-labeling"

	AcceptanceReportCreateTaskEndpoint   = "/api/v1/acceptance_report"
	AcceptanceReportTaskStatusEndpoint   = "/api/v1/acceptance_report/tasks/%s/status"
	AcceptanceReportTaskDownloadEndpoint = "/api/v1/acceptance_report/tasks/%s/download"

	PaidStorageCreateTaskEndpoint   = "/api/v1/paid_storage"
	PaidStorageTaskStatusEndpoint   = "/api/v1/paid_storage/tasks/%s/status"
	PaidStorageTaskDownloadEndpoint = "/api/v1/paid_storage/tasks/%s/download"

	RegionSaleEndpoint                 = "/api/v1/analytics/region-sale"
	BrandShareBrandsEndpoint           = "/api/v1/analytics/brand-share/brands"
	BrandShareParentsEndpoint          = "/api/v1/analytics/brand-share/parent-subjects"
	BrandShareReportEndpoint           = "/api/v1/analytics/brand-share"
	BannedBlockedEndpoint              = "/api/v1/analytics/banned-products/blocked"
	BannedShadowedEndpoint             = "/api/v1/analytics/banned-products/shadowed"
	GoodsReturnEndpoint                = "/api/v1/analytics/goods-return"
	StocksReportWBWarehousesEndpoint   = "/api/analytics/v1/stocks-report/wb-warehouses"
	StocksReportGroupsEndpoint         = "/api/v2/stocks-report/products/groups"
	StocksReportProductsEndpoint       = "/api/v2/stocks-report/products/products"
	StocksReportSizesEndpoint          = "/api/v2/stocks-report/products/sizes"
	StocksReportOfficesEndpoint        = "/api/v2/stocks-report/offices"
	SalesFunnelProductsV3Endpoint      = "/api/analytics/v3/sales-funnel/products"
	SalesFunnelProductsHistoryEndpoint = "/api/analytics/v3/sales-funnel/products/history"
	SalesFunnelGroupedHistoryEndpoint  = "/api/analytics/v3/sales-funnel/grouped/history"
	NmReportDownloadsEndpoint          = "/api/v2/nm-report/downloads"
	NmReportDownloadsRetryEndpoint     = "/api/v2/nm-report/downloads/retry"
	NmReportDownloadFileEndpoint       = "/api/v2/nm-report/downloads/file/%s"
	SearchReportMainEndpoint           = "/api/v2/search-report/report"
	SearchReportTableGroupsEndpoint    = "/api/v2/search-report/table/groups"
	SearchReportTableDetailsEndpoint   = "/api/v2/search-report/table/details"
	SearchReportProductTextsEndpoint   = "/api/v2/search-report/product/search-texts"
	SearchReportProductOrdersEndpoint  = "/api/v2/search-report/product/orders"
)
