package models

type TariffsDateQuery struct {
	Date string `json:"-"`
}

type CommissionResponse struct {
	Report []CommissionItem `json:"report"`
}

type CommissionItem struct {
	KgvpBooking       *float64 `json:"kgvpBooking,omitempty"`
	KgvpMarketplace   *float64 `json:"kgvpMarketplace,omitempty"`
	KgvpPickup        *float64 `json:"kgvpPickup,omitempty"`
	KgvpSupplier      *float64 `json:"kgvpSupplier,omitempty"`
	KgvpSupplierExpr  *float64 `json:"kgvpSupplierExpress,omitempty"`
	PaidStorageKgvp   *float64 `json:"paidStorageKgvp,omitempty"`
	KgvpChina         *float64 `json:"kgvpChina,omitempty"`
	KgvpTurkey        *float64 `json:"kgvpTurkey,omitempty"`
	KgvpMarketplaceUz *float64 `json:"kgvpMarketplaceUz,omitempty"`
	KgvpPaidStorageUz *float64 `json:"kgvpPaidStorageUz,omitempty"`
	KgvpSupplierUz    *float64 `json:"kgvpSupplierUz,omitempty"`
	KgvpUAE           *float64 `json:"kgvpUAE,omitempty"`
	ParentID          int64    `json:"parentID"`
	ParentName        string   `json:"parentName"`
	SubjectID         int64    `json:"subjectID"`
	SubjectName       string   `json:"subjectName"`
}

type TariffsBoxResponse struct {
	Response TariffsBoxEnvelope `json:"response"`
}

type TariffsBoxEnvelope struct {
	Data WarehousesBoxRates `json:"data"`
}

type WarehousesBoxRates struct {
	DTNextBox     string              `json:"dtNextBox"`
	DTTillMax     string              `json:"dtTillMax"`
	WarehouseList []WarehouseBoxRates `json:"warehouseList"`
}

type WarehouseBoxRates struct {
	BoxDeliveryBase             string `json:"boxDeliveryBase"`
	BoxDeliveryCoefExpr         string `json:"boxDeliveryCoefExpr"`
	BoxDeliveryLiter            string `json:"boxDeliveryLiter"`
	BoxDeliveryMarketplaceBase  string `json:"boxDeliveryMarketplaceBase"`
	BoxDeliveryMarketplaceCoef  string `json:"boxDeliveryMarketplaceCoefExpr"`
	BoxDeliveryMarketplaceLiter string `json:"boxDeliveryMarketplaceLiter"`
	BoxStorageBase              string `json:"boxStorageBase"`
	BoxStorageCoefExpr          string `json:"boxStorageCoefExpr"`
	BoxStorageLiter             string `json:"boxStorageLiter"`
	GeoName                     string `json:"geoName"`
	WarehouseName               string `json:"warehouseName"`
}

type TariffsPalletResponse struct {
	Response TariffsPalletEnvelope `json:"response"`
}

type TariffsPalletEnvelope struct {
	Data WarehousesPalletRates `json:"data"`
}

type WarehousesPalletRates struct {
	DTNextPallet  string                 `json:"dtNextPallet"`
	DTTillMax     string                 `json:"dtTillMax"`
	WarehouseList []WarehousePalletRates `json:"warehouseList"`
}

type WarehousePalletRates struct {
	PalletDeliveryExpr      string `json:"palletDeliveryExpr"`
	PalletDeliveryValueBase string `json:"palletDeliveryValueBase"`
	PalletDeliveryValueLit  string `json:"palletDeliveryValueLiter"`
	PalletStorageExpr       string `json:"palletStorageExpr"`
	PalletStorageValueExpr  string `json:"palletStorageValueExpr"`
	WarehouseName           string `json:"warehouseName"`
}

type ReturnTariffsResponse struct {
	Response ReturnTariffsEnvelope `json:"response"`
}

type ReturnTariffsEnvelope struct {
	Data WarehousesReturnRates `json:"data"`
}

type WarehousesReturnRates struct {
	DTNextDeliveryDumpKgt string                 `json:"dtNextDeliveryDumpKgt"`
	DTNextDeliveryDumpSrg string                 `json:"dtNextDeliveryDumpSrg"`
	DTNextDeliveryDumpSup string                 `json:"dtNextDeliveryDumpSup"`
	WarehouseList         []WarehouseReturnRates `json:"warehouseList"`
}

type WarehouseReturnRates struct {
	DeliveryDumpKgtOfficeBase  string `json:"deliveryDumpKgtOfficeBase"`
	DeliveryDumpKgtOfficeLiter string `json:"deliveryDumpKgtOfficeLiter"`
	DeliveryDumpKgtReturnExpr  string `json:"deliveryDumpKgtReturnExpr"`
	DeliveryDumpSrgOfficeExpr  string `json:"deliveryDumpSrgOfficeExpr"`
	DeliveryDumpSrgReturnExpr  string `json:"deliveryDumpSrgReturnExpr"`
	DeliveryDumpSupCourierBase string `json:"deliveryDumpSupCourierBase"`
	DeliveryDumpSupCourierLit  string `json:"deliveryDumpSupCourierLiter"`
	DeliveryDumpSupOfficeBase  string `json:"deliveryDumpSupOfficeBase"`
	DeliveryDumpSupOfficeLiter string `json:"deliveryDumpSupOfficeLiter"`
	DeliveryDumpSupReturnExpr  string `json:"deliveryDumpSupReturnExpr"`
	WarehouseName              string `json:"warehouseName"`
}

type AcceptanceCoefficient struct {
	Date                    string  `json:"date"`
	Coefficient             float64 `json:"coefficient"`
	WarehouseID             int64   `json:"warehouseID"`
	WarehouseName           string  `json:"warehouseName"`
	AllowUnload             bool    `json:"allowUnload"`
	BoxTypeID               *int    `json:"boxTypeID,omitempty"`
	StorageCoef             *string `json:"storageCoef"`
	DeliveryCoef            *string `json:"deliveryCoef"`
	DeliveryBaseLiter       *string `json:"deliveryBaseLiter"`
	DeliveryAdditionalLiter *string `json:"deliveryAdditionalLiter"`
	StorageBaseLiter        *string `json:"storageBaseLiter"`
	StorageAdditionalLiter  *string `json:"storageAdditionalLiter"`
	IsSortingCenter         bool    `json:"isSortingCenter"`
}
