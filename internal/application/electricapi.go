package application

import (
	"bytes"
	"dgb/meter.aggregation/internal/configuration"
	"encoding/json"
	"io"
	"net/http"
)

type ElecApi struct {
	response   *Response
	config     configuration.Configuration
	middleware *Middleware
}

func (api *ElecApi) upsert(w http.ResponseWriter, r *http.Request, method string) {
	defer r.Body.Close()

	var fields json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&fields); err != nil {
		api.response.BadRequest(ResponseParams{W: w})
		return
	}

	s := string(fields)

	req, _ := http.NewRequest(method, api.config.ELEC_BACKEND+r.RequestURI, bytes.NewBuffer([]byte(s)))
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

func (api *ElecApi) get(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, api.config.ELEC_BACKEND+r.RequestURI, nil)
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

func (api *ElecApi) update(w http.ResponseWriter, r *http.Request) {
	api.upsert(w, r, http.MethodPut)
}

func (api *ElecApi) create(w http.ResponseWriter, r *http.Request) {
	api.upsert(w, r, http.MethodPost)
}

func (api *ElecApi) delete(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest(http.MethodDelete, api.config.ELEC_BACKEND+r.RequestURI, nil)
	req.Header.Add("authorization", r.Header.Get("authorization"))
	response, err := http.DefaultClient.Do(req)

	if err != nil {
		api.response.ServerError(ResponseParams{W: w})
		return
	}

	api.response.AddHeaders(w)
	io.Copy(w, response.Body)
	response.Body.Close()
}

func NewElecApi(response *Response, configuration configuration.Configuration, middleware *Middleware) *ElecApi {
	return &ElecApi{
		response,
		configuration,
		middleware,
	}
}
