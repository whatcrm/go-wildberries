package wildberries

import (
	"errors"

	wbutils "github.com/whatcrm/go-wildberries/utils"
)

type APIError = wbutils.APIError

func AsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	ok := errors.As(err, &apiErr)
	return apiErr, ok
}
