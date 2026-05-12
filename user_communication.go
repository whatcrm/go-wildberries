package wildberries

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"github.com/whatcrm/go-wildberries/models"
	"github.com/whatcrm/go-wildberries/utils/feedbacks"
)

func (c *Client) CreateFeedbackAnswer(ctx context.Context, request models.FeedbackAnswerRequest) error {
	return c.sendFeedbackAnswer(ctx, http.MethodPost, request)
}

func (c *Client) UpdateFeedbackAnswer(ctx context.Context, request models.FeedbackAnswerRequest) error {
	return c.sendFeedbackAnswer(ctx, http.MethodPatch, request)
}

func (c *Client) RequestFeedbackOrderReturn(ctx context.Context, request models.FeedbackOrderReturnRequest) (*models.FeedbackResultResponse, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.FeedbacksBaseURL+feedbacks.FeedbackOrderReturnEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	var out models.FeedbackResultResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbackByID(ctx context.Context, id string) (*models.FeedbackByIDResponse, error) {
	reqURL, err := url.Parse(c.FeedbacksBaseURL + feedbacks.FeedbackByIDEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("id", id)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.FeedbackByIDResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetNewFeedbacksQuestions(ctx context.Context) (*models.FeedbackQuestionsStatusResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.FeedbacksBaseURL+feedbacks.NewFeedbacksQuestionsEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var out models.FeedbackQuestionsStatusResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetQuestionsCountUnanswered(ctx context.Context) (*models.UnansweredCountResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.FeedbacksBaseURL+feedbacks.QuestionsCountUnansweredEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var out models.UnansweredCountResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetQuestionsCount(ctx context.Context, query models.TimestampCountQuery) (*models.CountValueResponse, error) {
	if err := validateTimestampCountQuery(query); err != nil {
		return nil, err
	}
	reqURL, err := c.buildCountURL(feedbacks.QuestionsCountEndpoint, query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.CountValueResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetQuestions(ctx context.Context, query models.QuestionsListQuery) (*models.QuestionsListResponse, error) {
	if err := validateQuestionsListQuery(query); err != nil {
		return nil, err
	}
	reqURL, err := url.Parse(c.FeedbacksBaseURL + feedbacks.QuestionsEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("isAnswered", strconv.FormatBool(query.IsAnswered))
	q.Set("take", strconv.Itoa(query.Take))
	q.Set("skip", strconv.Itoa(query.Skip))
	if query.NmID != nil {
		q.Set("nmId", strconv.FormatInt(*query.NmID, 10))
	}
	if query.Order != "" {
		q.Set("order", string(query.Order))
	}
	if query.DateFrom != nil {
		q.Set("dateFrom", strconv.FormatInt(*query.DateFrom, 10))
	}
	if query.DateTo != nil {
		q.Set("dateTo", strconv.FormatInt(*query.DateTo, 10))
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.QuestionsListResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) PatchQuestion(ctx context.Context, request models.QuestionPatchRequest) (*models.FeedbackResultResponse, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.FeedbacksBaseURL+feedbacks.QuestionsEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	var out models.FeedbackResultResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetQuestionByID(ctx context.Context, id string) (*models.QuestionByIDResponse, error) {
	reqURL, err := url.Parse(c.FeedbacksBaseURL + feedbacks.QuestionByIDEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("id", id)
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.QuestionByIDResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbacksCountUnanswered(ctx context.Context) (*models.UnansweredCountResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.FeedbacksBaseURL+feedbacks.FeedbacksCountUnansweredEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var out models.UnansweredCountResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbacksCount(ctx context.Context, query models.TimestampCountQuery) (*models.CountValueResponse, error) {
	if err := validateTimestampCountQuery(query); err != nil {
		return nil, err
	}
	reqURL, err := c.buildCountURL(feedbacks.FeedbacksCountEndpoint, query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.CountValueResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbacks(ctx context.Context, query models.FeedbacksListQuery) (*models.FeedbacksListResponse, error) {
	if err := validateFeedbacksListQuery(query); err != nil {
		return nil, err
	}
	reqURL, err := url.Parse(c.FeedbacksBaseURL + feedbacks.FeedbacksEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("isAnswered", strconv.FormatBool(query.IsAnswered))
	q.Set("take", strconv.Itoa(query.Take))
	q.Set("skip", strconv.Itoa(query.Skip))
	if query.NmID != nil {
		q.Set("nmId", strconv.FormatInt(*query.NmID, 10))
	}
	if query.Order != "" {
		q.Set("order", string(query.Order))
	}
	if query.DateFrom != nil {
		q.Set("dateFrom", strconv.FormatInt(*query.DateFrom, 10))
	}
	if query.DateTo != nil {
		q.Set("dateTo", strconv.FormatInt(*query.DateTo, 10))
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.FeedbacksListResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbackArchive(ctx context.Context, query models.FeedbackArchiveQuery) (*models.FeedbackArchiveResponse, error) {
	reqURL, err := url.Parse(c.FeedbacksBaseURL + feedbacks.FeedbackArchiveEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	if query.NmID != nil {
		q.Set("nmId", strconv.FormatInt(*query.NmID, 10))
	}
	q.Set("take", strconv.Itoa(query.Take))
	q.Set("skip", strconv.Itoa(query.Skip))
	if query.Order != "" {
		q.Set("order", query.Order)
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.FeedbackArchiveResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbackPins(ctx context.Context, query models.PinsQuery) (*models.PinsResponse, error) {
	reqURL, err := c.buildPinsURL(feedbacks.FeedbackPinsEndpoint, query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.PinsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) PinFeedbacks(ctx context.Context, items []models.PinReviewItem) (*models.PinReviewsResponse, error) {
	jsonBody, err := json.Marshal(items)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.FeedbacksBaseURL+feedbacks.FeedbackPinsEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	var out models.PinReviewsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) UnpinFeedbacks(ctx context.Context, pinIDs []int64) (*models.UnpinReviewsResponse, error) {
	jsonBody, err := json.Marshal(pinIDs)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.FeedbacksBaseURL+feedbacks.FeedbackPinsEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	var out models.UnpinReviewsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbackPinsCount(ctx context.Context, query models.PinsQuery) (*models.PinsCountResponse, error) {
	reqURL, err := c.buildPinsURL(feedbacks.FeedbackPinsCountEndpoint, query)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.PinsCountResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetFeedbackPinsLimits(ctx context.Context) (*models.PinsLimitsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.FeedbacksBaseURL+feedbacks.FeedbackPinsLimitsEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var out models.PinsLimitsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSellerChats(ctx context.Context) (*models.ChatsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BuyerChatBaseURL+feedbacks.SellerChatsEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var out models.ChatsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) GetSellerEvents(ctx context.Context, next *int64) (*models.EventsResponse, error) {
	reqURL, err := url.Parse(c.BuyerChatBaseURL + feedbacks.SellerEventsEndpoint)
	if err != nil {
		return nil, err
	}
	if next != nil {
		q := reqURL.Query()
		q.Set("next", strconv.FormatInt(*next, 10))
		reqURL.RawQuery = q.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.EventsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) SendSellerMessage(ctx context.Context, request models.SendMessageRequest) (*models.MessageResponse, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	if err := writer.WriteField("replySign", request.ReplySign); err != nil {
		return nil, err
	}
	if request.Message != "" {
		if err := writer.WriteField("message", request.Message); err != nil {
			return nil, err
		}
	}
	for _, file := range request.Files {
		part, err := writer.CreateFormFile("file", file.Name)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(part, file.Reader); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BuyerChatBaseURL+feedbacks.SellerMessageEndpoint, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	var out models.MessageResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DownloadSellerFile(ctx context.Context, id string, out *bytes.Buffer) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BuyerChatBaseURL+fmt.Sprintf(feedbacks.SellerDownloadFileEndpoint, id), nil)
	if err != nil {
		return err
	}
	return c.Send(req, out)
}

func (c *Client) GetClaims(ctx context.Context, query models.ClaimQuery) (*models.ClaimsResponse, error) {
	reqURL, err := url.Parse(c.ReturnsBaseURL + feedbacks.ReturnsClaimsEndpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	q.Set("is_archive", strconv.FormatBool(query.IsArchive))
	if query.ID != "" {
		q.Set("id", query.ID)
	}
	if query.Limit != nil {
		q.Set("limit", strconv.Itoa(*query.Limit))
	}
	if query.Offset != nil {
		q.Set("offset", strconv.Itoa(*query.Offset))
	}
	if query.NmID != nil {
		q.Set("nm_id", strconv.FormatInt(*query.NmID, 10))
	}
	reqURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var out models.ClaimsResponse
	if err = c.Send(req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) PatchClaim(ctx context.Context, request models.PatchClaimRequest) error {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, c.ReturnsBaseURL+feedbacks.ReturnsClaimPatchEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) sendFeedbackAnswer(ctx context.Context, method string, request models.FeedbackAnswerRequest) error {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, method, c.FeedbacksBaseURL+feedbacks.FeedbackAnswerEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	return c.Send(req, nil)
}

func (c *Client) buildPinsURL(endpoint string, query models.PinsQuery) (*url.URL, error) {
	reqURL, err := url.Parse(c.FeedbacksBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	if query.State != "" {
		q.Set("state", query.State)
	}
	if query.PinOn != "" {
		q.Set("pinOn", query.PinOn)
	}
	if query.ImtID != nil {
		q.Set("imtId", strconv.FormatInt(*query.ImtID, 10))
	}
	if query.NmID != nil {
		q.Set("nmId", strconv.FormatInt(*query.NmID, 10))
	}
	if query.FeedbackID != nil {
		q.Set("feedbackId", strconv.FormatInt(*query.FeedbackID, 10))
	}
	if query.DateFrom != "" {
		q.Set("dateFrom", query.DateFrom)
	}
	if query.DateTo != "" {
		q.Set("dateTo", query.DateTo)
	}
	if query.Next != nil {
		q.Set("next", strconv.FormatInt(*query.Next, 10))
	}
	if query.Limit != nil {
		q.Set("limit", strconv.Itoa(*query.Limit))
	}
	reqURL.RawQuery = q.Encode()
	return reqURL, nil
}

func (c *Client) buildCountURL(endpoint string, query models.TimestampCountQuery) (*url.URL, error) {
	reqURL, err := url.Parse(c.FeedbacksBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	q := reqURL.Query()
	if query.DateFrom != nil {
		q.Set("dateFrom", strconv.FormatInt(*query.DateFrom, 10))
	}
	if query.DateTo != nil {
		q.Set("dateTo", strconv.FormatInt(*query.DateTo, 10))
	}
	if query.IsAnswered != nil {
		q.Set("isAnswered", strconv.FormatBool(*query.IsAnswered))
	}
	reqURL.RawQuery = q.Encode()
	return reqURL, nil
}

func validateTimestampCountQuery(query models.TimestampCountQuery) error {
	if query.DateFrom != nil && *query.DateFrom < 0 {
		return errors.New("dateFrom must be >= 0")
	}
	if query.DateTo != nil && *query.DateTo < 0 {
		return errors.New("dateTo must be >= 0")
	}
	if query.DateFrom != nil && query.DateTo != nil && *query.DateFrom > *query.DateTo {
		return errors.New("dateFrom must be <= dateTo")
	}
	return nil
}

func validateQuestionsListQuery(query models.QuestionsListQuery) error {
	if query.Take <= 0 {
		return errors.New("take must be > 0")
	}
	if query.Skip < 0 {
		return errors.New("skip must be >= 0")
	}
	if query.Take+query.Skip > 10000 {
		return errors.New("take + skip must be <= 10000 for questions")
	}
	if query.Order != "" && query.Order != models.SortDateAsc && query.Order != models.SortDateDesc {
		return errors.New("order must be dateAsc or dateDesc")
	}
	if query.DateFrom != nil && *query.DateFrom < 0 {
		return errors.New("dateFrom must be >= 0")
	}
	if query.DateTo != nil && *query.DateTo < 0 {
		return errors.New("dateTo must be >= 0")
	}
	if query.DateFrom != nil && query.DateTo != nil && *query.DateFrom > *query.DateTo {
		return errors.New("dateFrom must be <= dateTo")
	}
	return nil
}

func validateFeedbacksListQuery(query models.FeedbacksListQuery) error {
	if query.Take <= 0 {
		return errors.New("take must be > 0")
	}
	if query.Take > 5000 {
		return errors.New("take must be <= 5000 for feedbacks")
	}
	if query.Skip < 0 {
		return errors.New("skip must be >= 0")
	}
	if query.Skip > 199990 {
		return errors.New("skip must be <= 199990 for feedbacks")
	}
	if query.Order != "" && query.Order != models.SortDateAsc && query.Order != models.SortDateDesc {
		return errors.New("order must be dateAsc or dateDesc")
	}
	if query.DateFrom != nil && *query.DateFrom < 0 {
		return errors.New("dateFrom must be >= 0")
	}
	if query.DateTo != nil && *query.DateTo < 0 {
		return errors.New("dateTo must be >= 0")
	}
	if query.DateFrom != nil && query.DateTo != nil && *query.DateFrom > *query.DateTo {
		return errors.New("dateFrom must be <= dateTo")
	}
	return nil
}
