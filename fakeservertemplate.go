package fakeexample

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

type Tx struct {
	Request struct {
		FullURI  string
		Body     []byte
		HostName string
		ApiToken string
	}
	Response struct {
		Body   string
		Status int
	}
}

const (
	ROUTE1 = iota
	ROUTE2
	ROUTE_THAT_RETURNS_FILE
	MAX_ROUTES
)

type RouteCall struct {
	URI       string
	CallCount int
	Tx        []Tx
}

var (
	Api      [MAX_ROUTES]RouteCall
	ApiRoute RouteCall

	testServer *httptest.Server
)

func FakeHandler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.RequestURI, ApiRoute.URI) {

		ApiRoute.Tx[0].Request.HostName = "http://" + r.Host
		ApiRoute.Tx[0].Request.ApiToken = getApiKey(r.RequestURI)
		ApiRoute.Tx[0].Request.Body, _ = ioutil.ReadAll(r.Body)
		w.WriteHeader(ApiRoute.Tx[0].Response.Status)
		fmt.Fprintf(w, ApiRoute.Tx[0].Response.Body)
		ApiRoute.CallCount++

		return
	}

	for i := range Api {
		if strings.Contains(r.RequestURI, Api[i].URI) {
			if len(Api[i].Tx) == Api[i].CallCount {
				AddApiTransaction(i)
			}

			Api[i].Tx[Api[i].CallCount].Request.Body, _ = ioutil.ReadAll(r.Body)
			Api[i].Tx[Api[i].CallCount].Request.FullURI = r.RequestURI
			Api[i].Tx[Api[i].CallCount].Request.HostName = "http://" + r.Host
			Api[i].Tx[Api[i].CallCount].Request.ApiToken = getApiKey(r.RequestURI)

			w.WriteHeader(Api[i].Tx[Api[i].CallCount].Response.Status)
			fmt.Fprintf(w, Api[i].Tx[Api[i].CallCount].Response.Body)
			Api[i].CallCount++

			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func InitServer() {

	testServer = httptest.NewServer(http.HandlerFunc(FakeHandler))

	os.Setenv("REAL_API_ENDPOINT", testServer.URL)
	os.Setenv("REAL_API_TOKEN", "TEST_TOKEN")

	Api = [MAX_ROUTES]RouteCall{
		ROUTE1:                  {URI: "/route1"},
		ROUTE2:                  {URI: "/route2"},
		ROUTE_THAT_RETURNS_FILE: {URI: "/routeThatReturnsAFile"},
	}

	ApiRoute = RouteCall{URI: "/"}
	ApiRoute.Tx = make([]Tx, 1, 1)
	ApiRoute.Tx[0].Response.Status = http.StatusOK
}

func CloseServer() {
	testServer.Close()
}

func AddApiTransaction(routeId int) int {

	Api[routeId].Tx = append(Api[routeId].Tx, Tx{})

	idx := len(Api[routeId].Tx) - 1
	Api[routeId].Tx[idx].Response.Status = http.StatusOK
	if routeId == ROUTE_THAT_RETURNS_FILE {
		file, _ := ioutil.ReadFile("fake.xml")
		Api[routeId].Tx[idx].Response.Body = string(file)
	}

	return idx
}

func getApiKey(uri string) string {
	s := strings.Split(uri, "&")
	if len(s) > 0 {
		for j := 0; j < len(s); j++ {
			if strings.Index(s[j], "key=") > -1 {
				return strings.Trim(s[j], "key=")
			}
		}
	}
	return ""
}
