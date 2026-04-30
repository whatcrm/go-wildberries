package wildberries

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"github.com/whatcrm/go-wildberries/models"
	"github.com/whatcrm/go-wildberries/utils/content"
)

func (c *Client) GetParentCategories(ctx context.Context, locale string) (*models.ObjectParentResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + content.ObjectParentAllEndpoint)
	if err != nil {
		return nil, err
	}
	if locale != "" {
		params := requestURL.Query()
		params.Set("locale", locale)
		requestURL.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.ObjectParentResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetSubjects(ctx context.Context, query models.SubjectsQuery) (*models.SubjectsResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + content.ObjectAllEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	if query.Locale != "" {
		params.Set("locale", query.Locale)
	}
	if query.Name != "" {
		params.Set("name", query.Name)
	}
	if query.Limit > 0 {
		params.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	if query.Offset > 0 {
		params.Set("offset", strconv.FormatInt(query.Offset, 10))
	}
	if query.ParentID > 0 {
		params.Set("parentID", strconv.FormatInt(query.ParentID, 10))
	}
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.SubjectsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetSubjectCharacteristics(ctx context.Context, subjectID int64, locale string) (*models.SubjectCharacteristicsResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + fmt.Sprintf(content.ObjectCharcsEndpoint, subjectID))
	if err != nil {
		return nil, err
	}
	if locale != "" {
		params := requestURL.Query()
		params.Set("locale", locale)
		requestURL.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.SubjectCharacteristicsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetColors(ctx context.Context, locale string) (*models.ColorsResponse, error) {
	return c.getColorsLike(ctx, content.ColorsEndpoint, locale)
}

func (c *Client) GetKinds(ctx context.Context, locale string) (*models.StringListResponse, error) {
	return c.getStringList(ctx, content.KindsEndpoint, locale)
}

func (c *Client) GetCountries(ctx context.Context, locale string) (*models.CountriesResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + content.CountriesEndpoint)
	if err != nil {
		return nil, err
	}
	if locale != "" {
		params := requestURL.Query()
		params.Set("locale", locale)
		requestURL.RawQuery = params.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.CountriesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetSeasons(ctx context.Context, locale string) (*models.StringListResponse, error) {
	return c.getStringList(ctx, content.SeasonsEndpoint, locale)
}

func (c *Client) GetVATRates(ctx context.Context, locale string) (*models.StringListResponse, error) {
	return c.getStringList(ctx, content.VATEndpoint, locale)
}

func (c *Client) GetTnved(ctx context.Context, query models.TnvedQuery) (*models.TnvedResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + content.TNVEDEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("subjectID", strconv.FormatInt(query.SubjectID, 10))
	if query.Search != "" {
		params.Set("search", query.Search)
	}
	if query.Locale != "" {
		params.Set("locale", query.Locale)
	}
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.TnvedResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetBrands(ctx context.Context, query models.BrandsQuery) (*models.BrandsResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + content.BrandsEndpoint)
	if err != nil {
		return nil, err
	}
	params := requestURL.Query()
	params.Set("subjectId", strconv.FormatInt(query.SubjectID, 10))
	if query.Next > 0 {
		params.Set("next", strconv.FormatInt(query.Next, 10))
	}
	requestURL.RawQuery = params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.BrandsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetTags(ctx context.Context) (*models.TagsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.ContentBaseURL+content.TagsEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response models.TagsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CreateTag(ctx context.Context, request models.UpsertTagRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.TagEndpoint, request)
}

func (c *Client) UpdateTag(ctx context.Context, tagID int64, request models.UpsertTagRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPatch, c.ContentBaseURL+fmt.Sprintf(content.TagByIDEndpoint, tagID), request)
}

func (c *Client) DeleteTag(ctx context.Context, tagID int64) (*models.GenericDataResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.ContentBaseURL+fmt.Sprintf(content.TagByIDEndpoint, tagID), nil)
	if err != nil {
		return nil, err
	}
	var response models.GenericDataResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) LinkTagsToCard(ctx context.Context, request models.LinkTagRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.TagNomenclatureEndpoint, request)
}

func (c *Client) GetCardsList(ctx context.Context, locale string, request models.CardsListRequest) (*models.CardsListResponse, error) {
	return c.getCardsWithLocale(ctx, locale, content.GetCardsListEndpoint, request)
}

func (c *Client) GetCardsErrors(ctx context.Context, locale string, request models.PublicErrorsListRequest) (*models.PublicErrorsListResponse, error) {
	response, err := c.postWithLocale(ctx, locale, content.CardsErrorsListEndpoint, request, &models.PublicErrorsListResponse{})
	if err != nil {
		return nil, err
	}
	return response.(*models.PublicErrorsListResponse), nil
}

func (c *Client) UpdateCards(ctx context.Context, request []models.CardsUpdateItem) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.CardsUpdateEndpoint, request)
}

func (c *Client) MoveCards(ctx context.Context, request models.MoveCardsRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.CardsMoveNMEndpoint, request)
}

func (c *Client) DeleteCardsToTrash(ctx context.Context, request models.NmIDsRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.CardsDeleteTrashEndpoint, request)
}

func (c *Client) RecoverCards(ctx context.Context, request models.NmIDsRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.CardsRecoverEndpoint, request)
}

func (c *Client) GetCardsTrash(ctx context.Context, locale string, request models.CardsTrashRequest) (*models.CardsListResponse, error) {
	return c.getCardsWithLocale(ctx, locale, content.GetCardsTrashEndpoint, request)
}

func (c *Client) GetCardsLimits(ctx context.Context) (*models.CardsLimitsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.ContentBaseURL+content.CardsLimitsEndpoint, nil)
	if err != nil {
		return nil, err
	}
	var response models.CardsLimitsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GenerateBarcodes(ctx context.Context, request models.GenerateBarcodesRequest) (*models.GenerateBarcodesResponse, error) {
	return c.postContentNoLocale(ctx, content.BarcodesEndpoint, request, &models.GenerateBarcodesResponse{})
}

func (c *Client) UploadCards(ctx context.Context, request models.UploadCardsRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.CardsUploadEndpoint, request)
}

func (c *Client) UploadCardsAdd(ctx context.Context, request models.UploadCardsAddRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.CardsUploadAddEndpoint, request)
}

func (c *Client) UploadMediaByLinks(ctx context.Context, request models.SaveMediaByLinksRequest) (*models.GenericDataResponse, error) {
	return c.sendContentJSON(ctx, http.MethodPost, c.ContentBaseURL+content.MediaSaveEndpoint, request)
}

func (c *Client) UploadMediaFile(ctx context.Context, nmID int64, photoNumber int, filename string, body io.Reader) (*models.GenericDataResponse, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, body); err != nil {
		return nil, err
	}
	if err = writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ContentBaseURL+content.MediaFileEndpoint, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Nm-Id", strconv.FormatInt(nmID, 10))
	req.Header.Set("X-Photo-Number", strconv.Itoa(photoNumber))

	var response models.GenericDataResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) getColorsLike(ctx context.Context, endpoint, locale string) (*models.ColorsResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	if locale != "" {
		params := requestURL.Query()
		params.Set("locale", locale)
		requestURL.RawQuery = params.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.ColorsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) getStringList(ctx context.Context, endpoint, locale string) (*models.StringListResponse, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	if locale != "" {
		params := requestURL.Query()
		params.Set("locale", locale)
		requestURL.RawQuery = params.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	var response models.StringListResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) getCardsWithLocale(ctx context.Context, locale, endpoint string, payload interface{}) (*models.CardsListResponse, error) {
	response, err := c.postWithLocale(ctx, locale, endpoint, payload, &models.CardsListResponse{})
	if err != nil {
		return nil, err
	}
	return response.(*models.CardsListResponse), nil
}

func (c *Client) sendContentJSON(ctx context.Context, method, endpoint string, payload interface{}) (*models.GenericDataResponse, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	var response models.GenericDataResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) postContentNoLocale(ctx context.Context, endpoint string, payload interface{}, out interface{}) (*models.GenerateBarcodesResponse, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.ContentBaseURL+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	if err = c.Send(req, out); err != nil {
		return nil, err
	}
	return out.(*models.GenerateBarcodesResponse), nil
}

func (c *Client) postWithLocale(ctx context.Context, locale, endpoint string, payload interface{}, out interface{}) (interface{}, error) {
	requestURL, err := url.Parse(c.ContentBaseURL + endpoint)
	if err != nil {
		return nil, err
	}
	if locale != "" {
		params := requestURL.Query()
		params.Set("locale", locale)
		requestURL.RawQuery = params.Encode()
	}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	if err = c.Send(req, out); err != nil {
		return nil, err
	}
	return out, nil
}
