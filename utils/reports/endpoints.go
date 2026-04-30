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

	RegionSaleEndpoint        = "/api/v1/analytics/region-sale"
	BrandShareBrandsEndpoint  = "/api/v1/analytics/brand-share/brands"
	BrandShareParentsEndpoint = "/api/v1/analytics/brand-share/parent-subjects"
	BrandShareReportEndpoint  = "/api/v1/analytics/brand-share"
	BannedBlockedEndpoint     = "/api/v1/analytics/banned-products/blocked"
	BannedShadowedEndpoint    = "/api/v1/analytics/banned-products/shadowed"
	GoodsReturnEndpoint       = "/api/v1/analytics/goods-return"
)
