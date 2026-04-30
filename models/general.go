package models

type PingResponse struct {
	TS     string `json:"TS"`
	Status string `json:"Status"`
}

type NewsQuery struct {
	From   string `json:"-"`
	FromID uint64 `json:"-"`
}

type NewsResponse struct {
	Data []NewsItem `json:"data"`
}

type NewsItem struct {
	Content string     `json:"content"`
	Date    string     `json:"date"`
	Header  string     `json:"header"`
	ID      int64      `json:"id"`
	Types   []NewsType `json:"types"`
}

type NewsType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SellerInfo struct {
	Name      string `json:"name"`
	SID       string `json:"sid"`
	TIN       string `json:"tin"`
	TradeMark string `json:"tradeMark"`
}

type SupplierRating struct {
	FeedbackCount int64   `json:"feedbackCount"`
	Valuation     float64 `json:"valuation"`
}

type SubscriptionsJamInfo struct {
	State            string `json:"state"`
	ActivationSource string `json:"activationSource"`
	Level            string `json:"level"`
	Since            string `json:"since"`
	Till             string `json:"till"`
}

type ErrorResponse struct {
	Title     string  `json:"title"`
	Detail    string  `json:"detail"`
	RequestID string  `json:"requestId"`
	Origin    string  `json:"origin"`
	Status    float64 `json:"status"`
}

type ProblemResponse struct {
	Title      string  `json:"title"`
	Detail     string  `json:"detail"`
	Code       string  `json:"code"`
	RequestID  string  `json:"requestId"`
	Origin     string  `json:"origin"`
	Status     float64 `json:"status"`
	StatusText string  `json:"statusText"`
	Timestamp  string  `json:"timestamp"`
}
