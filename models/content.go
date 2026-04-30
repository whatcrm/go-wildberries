package models

type ContentBaseResponse struct {
	Error            bool        `json:"error"`
	ErrorText        string      `json:"errorText"`
	AdditionalErrors interface{} `json:"additionalErrors"`
}

type ObjectParent struct {
	Name      string `json:"name"`
	ID        int64  `json:"id"`
	IsVisible bool   `json:"isVisible"`
}

type ObjectParentResponse struct {
	Data []ObjectParent `json:"data"`
	ContentBaseResponse
}

type Subject struct {
	SubjectID   int64  `json:"subjectID"`
	ParentID    int64  `json:"parentID"`
	SubjectName string `json:"subjectName"`
	ParentName  string `json:"parentName"`
}

type SubjectsQuery struct {
	Locale   string `json:"-"`
	Name     string `json:"-"`
	Limit    int64  `json:"-"`
	Offset   int64  `json:"-"`
	ParentID int64  `json:"-"`
}

type SubjectsResponse struct {
	Data []Subject `json:"data"`
	ContentBaseResponse
}

type SubjectCharacteristic struct {
	CharcID     int64  `json:"charcID"`
	SubjectName string `json:"subjectName"`
	SubjectID   int64  `json:"subjectID"`
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	UnitName    string `json:"unitName"`
	MaxCount    int    `json:"maxCount"`
	Popular     bool   `json:"popular"`
	CharcType   int    `json:"charcType"`
	HasFilter   bool   `json:"hasFilter"`
}

type SubjectCharacteristicsResponse struct {
	Data []SubjectCharacteristic `json:"data"`
	ContentBaseResponse
}

type ColorValue struct {
	Name       string `json:"name"`
	ParentName string `json:"parentName"`
}

type ColorsResponse struct {
	Data []ColorValue `json:"data"`
	ContentBaseResponse
}

type StringListResponse struct {
	Data []string `json:"data"`
	ContentBaseResponse
}

type CountryValue struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

type CountriesResponse struct {
	Data []CountryValue `json:"data"`
	ContentBaseResponse
}

type TnvedQuery struct {
	SubjectID int64  `json:"-"`
	Search    string `json:"-"`
	Locale    string `json:"-"`
}

type TnvedItem struct {
	Tnved string `json:"tnved"`
	IsKiz bool   `json:"isKiz"`
}

type TnvedResponse struct {
	Data []TnvedItem `json:"data"`
	ContentBaseResponse
}

type BrandsQuery struct {
	SubjectID int64 `json:"-"`
	Next      int64 `json:"-"`
}

type Brand struct {
	ID      int64  `json:"id"`
	LogoURL string `json:"logoUrl"`
	Name    string `json:"name"`
}

type BrandsResponse struct {
	Brands []Brand `json:"brands"`
	Next   int64   `json:"next"`
	Total  int64   `json:"total"`
}

type Tag struct {
	ID    int64  `json:"id"`
	Color string `json:"color"`
	Name  string `json:"name"`
}

type TagsResponse struct {
	Data []Tag `json:"data"`
	ContentBaseResponse
}

type UpsertTagRequest struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

type LinkTagRequest struct {
	NmID   int64   `json:"nmID"`
	TagsID []int64 `json:"tagsIDs"`
}

type GenericDataResponse struct {
	Data interface{} `json:"data"`
	ContentBaseResponse
}

type CardsListRequest struct {
	Settings map[string]interface{} `json:"settings"`
}

type CardsListResponse struct {
	Cards  []map[string]interface{} `json:"cards"`
	Cursor map[string]interface{}   `json:"cursor"`
}

type PublicErrorsListRequest struct {
	Cursor map[string]interface{} `json:"cursor,omitempty"`
	Order  map[string]interface{} `json:"order,omitempty"`
}

type PublicErrorsListResponse struct {
	Data map[string]interface{} `json:"data"`
	ContentBaseResponse
}

type CardsUpdateItem map[string]interface{}

type MoveCardsRequest map[string]interface{}

type NmIDsRequest struct {
	NmIDs []int64 `json:"nmIDs"`
}

type CardsTrashRequest struct {
	Settings map[string]interface{} `json:"settings"`
}

type CardsLimitsResponse struct {
	Data struct {
		FreeLimits int64 `json:"freeLimits"`
		PaidLimits int64 `json:"paidLimits"`
	} `json:"data"`
	ContentBaseResponse
}

type GenerateBarcodesRequest struct {
	Count int `json:"count"`
}

type GenerateBarcodesResponse struct {
	Data []string `json:"data"`
	ContentBaseResponse
}

type UploadCardsRequest []map[string]interface{}

type UploadCardsAddRequest struct {
	ImtID      int64                    `json:"imtID"`
	CardsToAdd []map[string]interface{} `json:"cardsToAdd"`
}

type SaveMediaByLinksRequest struct {
	NmID int64    `json:"nmId"`
	Data []string `json:"data"`
}
