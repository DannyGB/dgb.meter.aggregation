package application

import (
	"dgb/meter.aggregation/internal/configuration"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const route = "/api"

const readingRoute = "/reading"
const oilRoute = "/oil"
const SKIP = "skip"
const TAKE = "take"
const SORT = "sort"
const FILTER = "filter"
const ACCESS_CLAIM = "access_as_user"

type Api struct {
	response   *Response
	elecApi    *ElecApi
	oilApi     *OilApi
	config     configuration.Configuration
	middleware *Middleware
}

func (api *Api) HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	subRoute := myRouter.PathPrefix(route).Subrouter()

	subRoute.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		api.response.Write(w, 200, "OK")
	}).Methods(http.MethodGet)

	subRoute.
		Path(api.getReadingRoute(readingRoute, "count")).
		Queries("filter", "{filter}").
		Methods(http.MethodGet, http.MethodOptions).
		HandlerFunc(api.middleware.Options(api.middleware.Authorize(api.elecApi.get, ACCESS_CLAIM)))

	subRoute.HandleFunc(api.getReadingRoute(readingRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.elecApi.get, ACCESS_CLAIM))).Methods(http.MethodGet, http.MethodOptions)
	subRoute.HandleFunc(api.getReadingRoute(readingRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.elecApi.update, ACCESS_CLAIM))).Methods(http.MethodPut, http.MethodOptions)
	subRoute.HandleFunc(api.getReadingRoute(readingRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.elecApi.delete, ACCESS_CLAIM))).Methods(http.MethodDelete, http.MethodOptions)
	subRoute.HandleFunc(api.getReadingRoute(readingRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.elecApi.create, ACCESS_CLAIM))).Methods(http.MethodPost, http.MethodOptions)

	subRoute.
		Path(api.getReadingRoute(readingRoute, "/")).
		Queries("skip", "{skip:[0-9]+}", "take", "{take:[0-9]+}", "sort", "{sort}", "filter", "{filter}").
		Methods(http.MethodGet, http.MethodOptions).
		HandlerFunc(api.middleware.Options(api.middleware.Authorize(api.elecApi.get, ACCESS_CLAIM)))

	subRoute.
		Path(api.getReadingRoute(readingRoute, "{year:[0-9]+}/total")).
		Methods(http.MethodGet, http.MethodOptions).
		HandlerFunc(api.middleware.Options(api.middleware.Authorize(api.elecApi.get, ACCESS_CLAIM)))

	// Oil

	subRoute.
		Path(api.getReadingRoute(oilRoute, "count")).
		Queries("filter", "{filter}").
		Methods(http.MethodGet, http.MethodOptions).
		HandlerFunc(api.middleware.Options(api.middleware.Authorize(api.oilApi.get, ACCESS_CLAIM)))

	subRoute.HandleFunc(api.getReadingRoute(oilRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.oilApi.get, ACCESS_CLAIM))).Methods(http.MethodGet, http.MethodOptions)
	subRoute.HandleFunc(api.getReadingRoute(oilRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.oilApi.update, ACCESS_CLAIM))).Methods(http.MethodPut, http.MethodOptions)
	subRoute.HandleFunc(api.getReadingRoute(oilRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.oilApi.delete, ACCESS_CLAIM))).Methods(http.MethodDelete, http.MethodOptions)
	subRoute.HandleFunc(api.getReadingRoute(oilRoute, "{id:[0-9a-zA\\-]+}"), api.middleware.Options(api.middleware.Authorize(api.oilApi.create, ACCESS_CLAIM))).Methods(http.MethodPost, http.MethodOptions)

	subRoute.
		Path(api.getReadingRoute(oilRoute, "/")).
		Queries("skip", "{skip:[0-9]+}", "take", "{take:[0-9]+}", "sort", "{sort}", "filter", "{filter}").
		Methods(http.MethodGet, http.MethodOptions).
		HandlerFunc(api.middleware.Options(api.middleware.Authorize(api.oilApi.get, ACCESS_CLAIM)))

	subRoute.
		Path(api.getReadingRoute(oilRoute, "{year:[0-9]+}/total")).
		Methods(http.MethodGet, http.MethodOptions).
		HandlerFunc(api.middleware.Options(api.middleware.Authorize(api.oilApi.get, ACCESS_CLAIM)))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", api.config.HTTP_PORT), subRoute))
}

func (api *Api) getReadingRoute(route string, subRoute string) string {
	return fmt.Sprintf("%s/%s", route, subRoute)
}

func NewApi(response *Response, elecApi *ElecApi, oilApi *OilApi, configuration configuration.Configuration, middleware *Middleware) *Api {
	return &Api{
		response,
		elecApi,
		oilApi,
		configuration,
		middleware,
	}
}
