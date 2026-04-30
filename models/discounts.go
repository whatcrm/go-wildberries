package models

type ResponseError struct {
	Data      interface{} `json:"data"`
	Error     bool        `json:"error"`
	ErrorText string      `json:"errorText"`
}

type BufferTaskResponse struct {
	Data      *SupplierTaskMetadataBuffer `json:"data"`
	Error     bool                        `json:"error"`
	ErrorText string                      `json:"errorText"`
}

type SupplierTaskMetadataBuffer struct {
	UploadID         int64  `json:"uploadID"`
	Status           int    `json:"status"`
	UploadDate       string `json:"uploadDate"`
	ActivationDate   string `json:"activationDate"`
	OverAllGoodsNumb int    `json:"overAllGoodsNumber"`
	SuccessGoodsNumb int    `json:"successGoodsNumber"`
}

type BufferGoodsTaskResponse struct {
	Data struct {
		UploadID    int64               `json:"uploadID"`
		BufferGoods []GoodBufferHistory `json:"bufferGoods"`
	} `json:"data"`
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type GoodBufferHistory struct {
	NmID            int64  `json:"nmID"`
	VendorCode      string `json:"vendorCode"`
	SizeID          int64  `json:"sizeID"`
	TechSizeName    string `json:"techSizeName"`
	Price           int    `json:"price"`
	CurrencyIsoCode string `json:"currencyIsoCode4217"`
	Discount        int    `json:"discount"`
	ClubDiscount    int    `json:"clubDiscount"`
	Status          int    `json:"status"`
	ErrorText       string `json:"errorText"`
}

type GoodsFilterQuery struct {
	Limit      int64 `json:"-"`
	Offset     int64 `json:"-"`
	FilterNmID int64 `json:"-"`
}

type GoodsFilterRequest struct {
	NmList []int64 `json:"nmList"`
}

type GoodsListResponse struct {
	Data struct {
		ListGoods []GoodsListItem `json:"listGoods"`
	} `json:"data"`
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type GoodsListItem struct {
	NmID              int64           `json:"nmID"`
	VendorCode        string          `json:"vendorCode"`
	Sizes             []GoodsListSize `json:"sizes"`
	CurrencyIsoCode   string          `json:"currencyIsoCode4217"`
	Discount          int             `json:"discount"`
	ClubDiscount      int             `json:"clubDiscount"`
	EditableSizePrice bool            `json:"editableSizePrice"`
	IsBadTurnover     bool            `json:"isBadTurnover"`
}

type GoodsListSize struct {
	SizeID              int64   `json:"sizeID"`
	Price               int     `json:"price"`
	DiscountedPrice     float64 `json:"discountedPrice"`
	ClubDiscountedPrice float64 `json:"clubDiscountedPrice"`
	TechSizeName        string  `json:"techSizeName"`
}

type SizeGoodsQuery struct {
	Limit  int64 `json:"-"`
	Offset int64 `json:"-"`
	NmID   int64 `json:"-"`
}

type SizeGoodsResponse struct {
	Data struct {
		ListGoods []SizeGood `json:"listGoods"`
	} `json:"data"`
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type SizeGood struct {
	NmID                int64   `json:"nmID"`
	SizeID              int64   `json:"sizeID"`
	VendorCode          string  `json:"vendorCode"`
	Price               int     `json:"price"`
	CurrencyIsoCode     string  `json:"currencyIsoCode4217"`
	DiscountedPrice     float64 `json:"discountedPrice"`
	ClubDiscountedPrice float64 `json:"clubDiscountedPrice"`
	Discount            int     `json:"discount"`
	ClubDiscount        int     `json:"clubDiscount"`
	TechSizeName        string  `json:"techSizeName"`
	EditableSizePrice   bool    `json:"editableSizePrice"`
	IsBadTurnover       bool    `json:"isBadTurnover"`
}

type QuarantineGoodsQuery struct {
	Limit  int64 `json:"-"`
	Offset int64 `json:"-"`
}

type QuarantineGoodsResponse struct {
	Data struct {
		QuarantineGoods []QuarantineGood `json:"quarantineGoods"`
	} `json:"data"`
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type QuarantineGood struct {
	NmID            int64   `json:"nmID"`
	SizeID          *int64  `json:"sizeID"`
	TechSizeName    string  `json:"techSizeName"`
	CurrencyIsoCode string  `json:"currencyIsoCode4217"`
	NewPrice        float64 `json:"newPrice"`
	OldPrice        float64 `json:"oldPrice"`
	NewDiscount     int     `json:"newDiscount"`
	OldDiscount     int     `json:"oldDiscount"`
	PriceDiff       float64 `json:"priceDiff"`
}

type TaskCreatedResponse struct {
	Data struct {
		ID            int64 `json:"id"`
		AlreadyExists bool  `json:"alreadyExists"`
	} `json:"data"`
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type GoodsUploadRequest struct {
	Data []GoodUpload `json:"data"`
}

type GoodUpload struct {
	NmID     int64 `json:"nmID"`
	Price    int   `json:"price,omitempty"`
	Discount int   `json:"discount,omitempty"`
}

type SizeUploadRequest struct {
	Data []SizeUpload `json:"data"`
}

type SizeUpload struct {
	NmID   int64 `json:"nmID"`
	SizeID int64 `json:"sizeID"`
	Price  int   `json:"price"`
}

type ClubDiscountUploadRequest struct {
	Data []ClubDiscountUpload `json:"data"`
}

type ClubDiscountUpload struct {
	NmID         int64 `json:"nmID"`
	ClubDiscount int   `json:"clubDiscount"`
}

type TaskHistoryResponse struct {
	Data      *SupplierTaskMetadata `json:"data"`
	Error     bool                  `json:"error"`
	ErrorText string                `json:"errorText"`
}

type SupplierTaskMetadata struct {
	UploadID         int64  `json:"uploadID"`
	Status           int    `json:"status"`
	UploadDate       string `json:"uploadDate"`
	ActivationDate   string `json:"activationDate"`
	OverAllGoodsNumb int    `json:"overAllGoodsNumber"`
	SuccessGoodsNumb int    `json:"successGoodsNumber"`
}

type GoodsTaskHistoryResponse struct {
	Data struct {
		UploadID     int64         `json:"uploadID"`
		HistoryGoods []GoodHistory `json:"historyGoods"`
	} `json:"data"`
	Error     bool   `json:"error"`
	ErrorText string `json:"errorText"`
}

type GoodHistory struct {
	NmID            int64  `json:"nmID"`
	VendorCode      string `json:"vendorCode"`
	SizeID          int64  `json:"sizeID"`
	TechSizeName    string `json:"techSizeName"`
	Price           int    `json:"price"`
	CurrencyIsoCode string `json:"currencyIsoCode4217"`
	Discount        int    `json:"discount"`
	ClubDiscount    int    `json:"clubDiscount"`
	Status          int    `json:"status"`
	ErrorText       string `json:"errorText"`
}
