package application

import (
	"bytes"
	"dgb/meter.aggregation/internal/configuration"
	"encoding/json"
	"io"
	"net/http"
)

type OilApi struct {
	response   *Response
	config     configuration.Configuration
	middleware *Middleware
}

func (api *OilApi) upsert(w http.ResponseWriter, r *http.Request, method string) {
	defer r.Body.Close()

	var fields json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&fields); err != nil {
		api.response.BadRequest(ResponseParams{W: w})
		return
	}

	s := string(fields)

	req, _ := http.NewRequest(method, api.config.OIL_BACKEND+r.RequestURI, bytes.NewBuffer([]byte(s)))
	req.Header.Add("authorization", r.Header.Get("authorization"))
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		api.response.ServerError(ResponseParams{W: w})
		return
	}

	api.response.AddHeaders(w)
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func (api *OilApi) get(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, api.config.OIL_BACKEND+r.RequestURI, nil)
	req.Header.Add("authorization", r.Header.Get("authorization"))
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		api.response.ServerError(ResponseParams{W: w})
		return
	}

	api.response.AddHeaders(w)
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func (api *OilApi) update(w http.ResponseWriter, r *http.Request) {
	api.upsert(w, r, http.MethodPut)
}

func (api *OilApi) create(w http.ResponseWriter, r *http.Request) {
	api.upsert(w, r, http.MethodPost)
}

func (api *OilApi) delete(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodDelete, api.config.OIL_BACKEND+r.RequestURI, nil)
	req.Header.Add("authorization", r.Header.Get("authorization"))
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		api.response.ServerError(ResponseParams{W: w})
		return
	}

	api.response.AddHeaders(w)
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func NewOilApi(response *Response, configuration configuration.Configuration, middleware *Middleware) *OilApi {
	return &OilApi{
		response,
		configuration,
		middleware,
	}
}
