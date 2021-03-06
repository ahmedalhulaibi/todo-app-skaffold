package httprouter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"

	"github.com/shaxbee/todo-app-skaffold/pkg/httperror"
)

// JSONRequest expects application/json content-type and attempts
// to unmarshal request body into dst.
// 415 Unsupported Media Type is returned if invalid content-type was provided
// 400 Bad Request is returned if request body failed to unmarshal
func JSONRequest(req *http.Request, dst interface{}) error {
	mt, _, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if err != nil || mt != "application/json" {
		return httperror.New(http.StatusUnsupportedMediaType)
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body")
	}

	if err := json.Unmarshal(data, dst); err != nil {
		return httperror.New(
			http.StatusBadRequest,
			httperror.Message("Failed to unmarshal request body"),
			httperror.Cause(err),
		)
	}

	return nil
}

// JSONResponse sets content-type to application/json and marshals src
// as json to response body
// 500 Internal Server Error is returned if src could not be marshaled
func JSONResponse(w http.ResponseWriter, status int, src interface{}) error {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
	}

	data, err := json.Marshal(src)
	if err != nil {
		return httperror.New(
			http.StatusInternalServerError,
			httperror.Message("Failed to marshal response body"),
			httperror.Cause(err),
		)
	}

	w.WriteHeader(status)
	_, _ = w.Write(data)

	return nil
}
