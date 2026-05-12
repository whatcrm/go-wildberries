package models

import "io"

type SortOrder string

const (
	SortDateAsc  SortOrder = "dateAsc"
	SortDateDesc SortOrder = "dateDesc"
)

type FeedbackAnswerRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type FeedbackOrderReturnRequest struct {
	FeedbackID string `json:"feedbackId"`
}

type FeedbackResultResponse struct {
	Data             interface{} `json:"data"`
	Error            bool        `json:"error"`
	ErrorText        string      `json:"errorText"`
	AdditionalErrors []string    `json:"additionalErrors"`
	RequestID        string      `json:"requestId,omitempty"`
}

type FeedbackQuestionsStatusResponse struct {
	Data struct {
		HasNewQuestions bool `json:"hasNewQuestions"`
		HasNewFeedbacks bool `json:"hasNewFeedbacks"`
	} `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}

type UnansweredCountResponse struct {
	Data struct {
		CountUnanswered      int `json:"countUnanswered"`
		CountUnansweredToday int `json:"countUnansweredToday"`
	} `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}

type TimestampCountQuery struct {
	DateFrom   *int64
	DateTo     *int64
	IsAnswered *bool
}

type CountValueResponse struct {
	Data             int      `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}

type QuestionsListQuery struct {
	IsAnswered bool
	NmID       *int64
	Take       int
	Skip       int
	Order      SortOrder
	DateFrom   *int64
	DateTo     *int64
}

type QuestionByIDResponse struct {
	Data             QuestionItem `json:"data"`
	Error            bool         `json:"error"`
	ErrorText        string       `json:"errorText"`
	AdditionalErrors []string     `json:"additionalErrors"`
}

type QuestionsListResponse struct {
	Data struct {
		CountUnanswered int            `json:"countUnanswered"`
		CountArchive    int            `json:"countArchive"`
		Questions       []QuestionItem `json:"questions"`
	} `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}

type QuestionPatchRequest struct {
	ID        string          `json:"id"`
	WasViewed *bool           `json:"wasViewed,omitempty"`
	Answer    *QuestionAnswer `json:"answer,omitempty"`
	State     string          `json:"state,omitempty"`
}

type QuestionAnswer struct {
	Text       string `json:"text"`
	Editable   bool   `json:"editable,omitempty"`
	CreateDate string `json:"createDate,omitempty"`
}

type QuestionItem struct {
	ID             string          `json:"id"`
	Text           string          `json:"text"`
	CreatedDate    string          `json:"createdDate"`
	State          string          `json:"state"`
	Answer         *QuestionAnswer `json:"answer"`
	ProductDetails FeedbackProduct `json:"productDetails"`
	WasViewed      bool            `json:"wasViewed"`
	IsWarned       bool            `json:"isWarned"`
}

type FeedbacksListQuery struct {
	IsAnswered bool
	NmID       *int64
	Take       int
	Skip       int
	Order      SortOrder
	DateFrom   *int64
	DateTo     *int64
}

type FeedbacksListResponse struct {
	Data struct {
		CountUnanswered int            `json:"countUnanswered"`
		CountArchive    int            `json:"countArchive"`
		Feedbacks       []FeedbackItem `json:"feedbacks"`
	} `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}

type FeedbackByIDResponse struct {
	Data             FeedbackItem `json:"data"`
	Error            bool         `json:"error"`
	ErrorText        string       `json:"errorText"`
	AdditionalErrors []string     `json:"additionalErrors"`
}

type FeedbackArchiveQuery struct {
	NmID  *int64
	Take  int
	Skip  int
	Order string
}

type FeedbackArchiveResponse struct {
	Data struct {
		Feedbacks []FeedbackItem `json:"feedbacks"`
	} `json:"data"`
	Error            bool     `json:"error"`
	ErrorText        string   `json:"errorText"`
	AdditionalErrors []string `json:"additionalErrors"`
}

type FeedbackItem struct {
	ID                        string              `json:"id"`
	Text                      string              `json:"text"`
	Pros                      string              `json:"pros"`
	Cons                      string              `json:"cons"`
	MatchingSize              string              `json:"matchingSize"`
	ProductValuation          int                 `json:"productValuation"`
	CreatedDate               string              `json:"createdDate"`
	Answer                    *FeedbackAnswer     `json:"answer"`
	State                     string              `json:"state"`
	ProductDetails            FeedbackProduct     `json:"productDetails"`
	PhotoLinks                []FeedbackPhotoLink `json:"photoLinks"`
	UserName                  string              `json:"userName"`
	OrderStatus               string              `json:"orderStatus"`
	Video                     *FeedbackVideo      `json:"video"`
	WasViewed                 bool                `json:"wasViewed"`
	IsAbleSupplierFeedbackVal bool                `json:"isAbleSupplierFeedbackValuation"`
	SupplierFeedbackValuation int                 `json:"supplierFeedbackValuation"`
	IsAbleSupplierProductVal  bool                `json:"isAbleSupplierProductValuation"`
	SupplierProductValuation  int                 `json:"supplierProductValuation"`
	IsAbleReturnProductOrders bool                `json:"isAbleReturnProductOrders"`
	ReturnProductOrdersDate   *string             `json:"returnProductOrdersDate"`
	Bables                    []string            `json:"bables"`
	LastOrderShkID            int64               `json:"lastOrderShkId"`
	LastOrderCreatedAt        string              `json:"lastOrderCreatedAt"`
	Color                     string              `json:"color"`
	SubjectID                 int64               `json:"subjectId"`
	SubjectName               string              `json:"subjectName"`
	ParentFeedbackID          *string             `json:"parentFeedbackId"`
	ChildFeedbackID           *string             `json:"childFeedbackId"`
}

type FeedbackAnswer struct {
	Text     string `json:"text"`
	State    string `json:"state"`
	Editable bool   `json:"editable"`
}

type FeedbackProduct struct {
	NmID            int64   `json:"nmId"`
	ImtID           int64   `json:"imtId"`
	ProductName     string  `json:"productName"`
	SupplierArticle *string `json:"supplierArticle"`
	SupplierName    *string `json:"supplierName"`
	BrandName       *string `json:"brandName"`
	Size            string  `json:"size"`
}

type FeedbackPhotoLink struct {
	FullSize string `json:"fullSize"`
	MiniSize string `json:"miniSize"`
}

type FeedbackVideo struct {
	PreviewImage string `json:"previewImage"`
	Link         string `json:"link"`
	DurationSec  int    `json:"durationSec"`
}

type PinsQuery struct {
	State      string
	PinOn      string
	ImtID      *int64
	NmID       *int64
	FeedbackID *int64
	DateFrom   string
	DateTo     string
	Next       *int64
	Limit      *int
}

type PinsResponse struct {
	Data []PinnedReviewItem `json:"data"`
	Next *int64             `json:"next,omitempty"`
}

type PinnedReviewItem struct {
	ChangeStateAt string `json:"changeStateAt"`
	ImtID         int64  `json:"imtId"`
	NmID          int64  `json:"nmId"`
	PinID         int64  `json:"pinId"`
	PinMethod     string `json:"pinMethod"`
	PinOn         string `json:"pinOn"`
	FeedbackID    string `json:"feedbackId"`
	State         string `json:"state"`
	UnpinnedCause string `json:"unpinnedCause,omitempty"`
}

type PinReviewItem struct {
	PinMethod  string `json:"pinMethod"`
	PinOn      string `json:"pinOn"`
	FeedbackID string `json:"feedbackId"`
}

type PinReviewsResponse struct {
	Data []PinReviewResult `json:"data"`
}

type PinReviewResult struct {
	FeedbackID string             `json:"feedbackId"`
	PinID      *int64             `json:"pinId,omitempty"`
	PinMethod  string             `json:"pinMethod"`
	PinOn      string             `json:"pinOn"`
	IsErrors   bool               `json:"isErrors"`
	Errors     []OpenAPIResultErr `json:"errors,omitempty"`
}

type OpenAPIResultErr struct {
	Origin    string `json:"origin"`
	Detail    string `json:"detail"`
	RequestID string `json:"requestId"`
	Status    string `json:"status"`
	Title     string `json:"title"`
}

type UnpinReviewsResponse struct {
	Data []int64 `json:"data"`
}

type PinsCountResponse struct {
	Data int `json:"data"`
}

type PinsLimitsResponse struct {
	Data struct {
		Subscription *SellerLimit `json:"subscription"`
		Tariff       *SellerLimit `json:"tariff"`
	} `json:"data"`
}

type SellerLimit struct {
	PerUnitLimit int  `json:"perUnitLimit"`
	Remaining    int  `json:"remaining"`
	TotalLimit   int  `json:"totalLimit"`
	Unlimited    bool `json:"unlimited"`
	Used         int  `json:"used"`
}

type ChatsResponse struct {
	Result []ChatItem `json:"result"`
	Errors []string   `json:"errors"`
}

type ChatItem struct {
	ChatID      string      `json:"chatID"`
	ReplySign   string      `json:"replySign"`
	ClientName  string      `json:"clientName"`
	GoodCard    GoodCard    `json:"goodCard"`
	LastMessage LastMessage `json:"lastMessage"`
}

type GoodCard struct {
	Date          string `json:"date"`
	NmID          int64  `json:"nmID"`
	Price         int64  `json:"price"`
	PriceCurrency string `json:"priceCurrency"`
	RID           string `json:"rid"`
	Size          string `json:"size"`
}

type LastMessage struct {
	Text         string `json:"text"`
	AddTimestamp int64  `json:"addTimestamp"`
}

type EventsResponse struct {
	Result EventsResult `json:"result"`
	Errors []string     `json:"errors"`
}

type EventsResult struct {
	Next            int64       `json:"next"`
	NewestEventTime string      `json:"newestEventTime"`
	OldestEventTime string      `json:"oldestEventTime"`
	TotalEvents     int         `json:"totalEvents"`
	Events          []ChatEvent `json:"events"`
}

type ChatEvent struct {
	ChatID       string        `json:"chatID"`
	EventID      string        `json:"eventID"`
	EventType    string        `json:"eventType"`
	IsNewChat    bool          `json:"isNewChat"`
	Message      *EventMessage `json:"message"`
	Source       string        `json:"source"`
	AddTimestamp int64         `json:"addTimestamp"`
	AddTime      string        `json:"addTime"`
	ReplySign    string        `json:"replySign,omitempty"`
	Sender       string        `json:"sender"`
	ClientName   string        `json:"clientName,omitempty"`
}

type EventMessage struct {
	Attachments *EventAttachments `json:"attachments"`
	Text        string            `json:"text"`
}

type EventAttachments struct {
	GoodCard *GoodCard   `json:"goodCard"`
	Files    []ChatFile  `json:"files"`
	Images   []ChatImage `json:"images"`
}

type ChatFile struct {
	ContentType string `json:"contentType"`
	Date        string `json:"date"`
	DownloadID  string `json:"downloadID"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Size        int64  `json:"size"`
}

type ChatImage struct {
	Date       string `json:"date"`
	DownloadID string `json:"downloadID"`
	URL        string `json:"url"`
}

type SendMessageRequest struct {
	ReplySign string
	Message   string
	Files     []SendMessageFile
}

type SendMessageFile struct {
	Name   string
	Reader io.Reader
}

type MessageResponse struct {
	Errors []string `json:"errors"`
	Result struct {
		AddTime int64  `json:"addTime"`
		ChatID  string `json:"chatID"`
		Sign    string `json:"sign"`
	} `json:"result"`
}

type ClaimQuery struct {
	IsArchive bool
	ID        string
	Limit     *int
	Offset    *int
	NmID      *int64
}

type ClaimsResponse struct {
	Claims []ClaimItem `json:"claims"`
	Total  int         `json:"total"`
}

type ClaimItem struct {
	ID           string   `json:"id"`
	ClaimType    int      `json:"claim_type"`
	Status       int      `json:"status"`
	StatusEx     int      `json:"status_ex"`
	NmID         int64    `json:"nm_id"`
	UserComment  string   `json:"user_comment"`
	WBComment    *string  `json:"wb_comment"`
	DT           string   `json:"dt"`
	ImtName      *string  `json:"imt_name"`
	OrderDT      string   `json:"order_dt"`
	DTUpdate     string   `json:"dt_update"`
	Photos       []string `json:"photos"`
	VideoPaths   []string `json:"video_paths"`
	Actions      []string `json:"actions"`
	Price        float64  `json:"price"`
	CurrencyCode string   `json:"currency_code"`
	SRID         string   `json:"srid"`
	OriginIDInfo *string  `json:"origin_id_info"`
	DeliveryDT   string   `json:"delivery_dt"`
}

type PatchClaimRequest struct {
	ID      string `json:"id"`
	Action  string `json:"action"`
	Comment string `json:"comment,omitempty"`
}
