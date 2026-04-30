package finance

const (
	BalanceEndpoint = "/api/v1/account/balance"

	SalesReportsListEndpoint         = "/api/finance/v1/sales-reports/list"
	SalesReportsDetailedEndpoint     = "/api/finance/v1/sales-reports/detailed"
	SalesReportsDetailedByIDEndpoint = "/api/finance/v1/sales-reports/detailed/%d"

	AcquiringListEndpoint         = "/api/finance/v1/acquiring/list"
	AcquiringDetailedEndpoint     = "/api/finance/v1/acquiring/detailed"
	AcquiringDetailedByIDEndpoint = "/api/finance/v1/acquiring/detailed/%d"

	SupplierReportDetailByPeriodEndpoint = "/api/v5/supplier/reportDetailByPeriod"
)
