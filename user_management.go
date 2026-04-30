package wildberries

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/whatcrm/go-wildberries/models"
	"github.com/whatcrm/go-wildberries/utils/usermanagement"
)

func (c *Client) CreateInvite(ctx context.Context, invite models.CreateInviteRequest) (*models.CreateInviteResponse, error) {
	requestURL := c.UserManagementBaseURL + usermanagement.InviteEndpoint

	jsonBody, err := json.Marshal(invite)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.CreateInviteResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetUsers(ctx context.Context, query models.GetUsersQuery) (*models.GetUsersResponse, error) {
	requestURL, err := url.Parse(c.UserManagementBaseURL + usermanagement.UsersEndpoint)
	if err != nil {
		return nil, err
	}

	params := requestURL.Query()
	if query.Limit > 0 {
		params.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	if query.Offset > 0 {
		params.Set("offset", strconv.FormatInt(query.Offset, 10))
	}
	if query.IsInviteOnly != nil {
		params.Set("isInviteOnly", strconv.FormatBool(*query.IsInviteOnly))
	}
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	var response models.GetUsersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateUsersAccess(ctx context.Context, request models.UpdateUserAccessRequest) error {
	requestURL := c.UserManagementBaseURL + usermanagement.UsersAccessEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}

func (c *Client) DeleteUser(ctx context.Context, deletedUserID int64) error {
	requestURL, err := url.Parse(c.UserManagementBaseURL + usermanagement.UserEndpoint)
	if err != nil {
		return err
	}

	params := requestURL.Query()
	params.Set("deletedUserID", strconv.FormatInt(deletedUserID, 10))
	requestURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL.String(), nil)
	if err != nil {
		return err
	}

	return c.Send(req, nil)
}
