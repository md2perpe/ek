package req

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2015 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	. "gopkg.in/check.v1"
	"net"
	"net/http"
	"testing"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	_URL_GET          = "/get"
	_URL_POST         = "/post"
	_URL_PUT          = "/put"
	_URL_HEAD         = "/head"
	_URL_PATCH        = "/patch"
	_URL_DELETE       = "/delete"
	_URL_QUERY        = "/query"
	_URL_HEADERS      = "/headers"
	_URL_CONTENT_TYPE = "/content-type"
	_URL_ACCEPT       = "/accept"
	_URL_USER_AGENT   = "/user-agent"
	_URL_BASIC_AUTH   = "/basic-auth"
	_URL_STRING_RESP  = "/string-response"
	_URL_JSON_RESP    = "/json-response"
)

const (
	_TEST_USER_AGENT      = "REQ TEST USER AGENT"
	_TEST_CONTENT_TYPE    = "application/json"
	_TEST_ACCEPT          = "application/vnd.example.api+json;version=2"
	_TEST_BASIC_AUTH_USER = "admin"
	_TEST_BASIC_AUTH_PASS = "password"
	_TEST_STRING_RESP     = "Test String Response"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type ReqSuite struct {
	url      string
	listener net.Listener
}

type TestStruct struct {
	String  string `json:"string"`
	Integer int    `json:"integer"`
	Boolean bool   `json:"boolean"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ReqSuite{url: "http://127.0.0.1:30000"})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ReqSuite) SetUpSuite(c *C) {
	go runHTTPServer(s, c)
}

func (s *ReqSuite) TearDownSuite(c *C) {
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *ReqSuite) TestMethodGet(c *C) {
	getResp, err := Request{URL: s.url + _URL_GET, Method: GET}.Do()

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)

	getResp, err = Request{URL: s.url + _URL_GET}.Do()

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)

	getResp, err = Request{URL: s.url + _URL_GET}.Get()

	c.Assert(err, IsNil)
	c.Assert(getResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodPost(c *C) {
	postResp, err := Request{URL: s.url + _URL_POST, Method: POST}.Do()

	c.Assert(err, IsNil)
	c.Assert(postResp.StatusCode, Equals, 200)

	postResp, err = Request{URL: s.url + _URL_POST}.Post()

	c.Assert(err, IsNil)
	c.Assert(postResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodPut(c *C) {
	putResp, err := Request{URL: s.url + _URL_PUT, Method: PUT}.Do()

	c.Assert(err, IsNil)
	c.Assert(putResp.StatusCode, Equals, 200)

	putResp, err = Request{URL: s.url + _URL_PUT}.Put()

	c.Assert(err, IsNil)
	c.Assert(putResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodHead(c *C) {
	headResp, err := Request{URL: s.url + _URL_HEAD, Method: HEAD}.Do()

	c.Assert(err, IsNil)
	c.Assert(headResp.StatusCode, Equals, 200)

	headResp, err = Request{URL: s.url + _URL_HEAD}.Head()

	c.Assert(err, IsNil)
	c.Assert(headResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodPatch(c *C) {
	patchResp, err := Request{URL: s.url + _URL_PATCH, Method: PATCH}.Do()

	c.Assert(err, IsNil)
	c.Assert(patchResp.StatusCode, Equals, 200)

	patchResp, err = Request{URL: s.url + _URL_PATCH}.Patch()

	c.Assert(err, IsNil)
	c.Assert(patchResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestMethodDelete(c *C) {
	deleteResp, err := Request{URL: s.url + _URL_DELETE, Method: DELETE}.Do()

	c.Assert(err, IsNil)
	c.Assert(deleteResp.StatusCode, Equals, 200)

	deleteResp, err = Request{URL: s.url + _URL_DELETE}.Delete()

	c.Assert(err, IsNil)
	c.Assert(deleteResp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestQuery(c *C) {
	resp, err := Request{
		URL: s.url + _URL_QUERY,
		Query: Query{
			"user": "john",
			"id":   "912",
			"root": "",
		},
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestHeaders(c *C) {
	resp, err := Request{
		URL: s.url + _URL_HEADERS,
		Headers: Headers{
			"Header1": "Value1",
			"Header2": "Value2",
		},
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestContentType(c *C) {
	resp, err := Request{
		URL:         s.url + _URL_CONTENT_TYPE,
		ContentType: _TEST_CONTENT_TYPE,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestAccept(c *C) {
	resp, err := Request{
		URL:    s.url + _URL_ACCEPT,
		Accept: _TEST_ACCEPT,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestUserAgent(c *C) {
	resp, err := Request{
		URL:       s.url + _URL_USER_AGENT,
		UserAgent: _TEST_USER_AGENT,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	UserAgent = _TEST_USER_AGENT

	resp, err = Request{
		URL: s.url + _URL_USER_AGENT,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestBasicAuth(c *C) {
	resp, err := Request{
		URL:               s.url + _URL_BASIC_AUTH,
		BasicAuthUsername: _TEST_BASIC_AUTH_USER,
		BasicAuthPassword: _TEST_BASIC_AUTH_PASS,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *ReqSuite) TestStringResp(c *C) {
	resp, err := Request{
		URL: s.url + _URL_STRING_RESP,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)
	c.Assert(resp.String(), Equals, _TEST_STRING_RESP)
}

func (s *ReqSuite) TestJSONResp(c *C) {
	resp, err := Request{
		URL: s.url + _URL_JSON_RESP,
	}.Do()

	c.Assert(err, IsNil)
	c.Assert(resp.StatusCode, Equals, 200)

	testStruct := &TestStruct{}

	err = resp.JSON(testStruct)

	c.Assert(err, IsNil)
	c.Assert(testStruct.String, Equals, "test")
	c.Assert(testStruct.Integer, Equals, 912)
	c.Assert(testStruct.Boolean, Equals, true)
}

func (s *ReqSuite) TestErrors(c *C) {
	resp, err := Request{}.Do()

	c.Assert(resp, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Can't create request struct (URL property can't be empty and must be set)")

	resp, err = Request{URL: "ABCD"}.Do()

	c.Assert(resp, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Can't send request (Get ABCD: unsupported protocol scheme \"\")")

	resp, err = Request{URL: "http://127.0.0.1:60000"}.Do()

	c.Assert(resp, IsNil)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Can't send request (Get http://127.0.0.1:60000: dial tcp 127.0.0.1:60000: getsockopt: connection refused)")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func runHTTPServer(s *ReqSuite, c *C) {
	server := &http.Server{
		Handler:        http.NewServeMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	listener, err := net.Listen("tcp", ":30000")

	if err != nil {
		c.Fatal(err.Error())
	}

	s.listener = listener

	server.Handler.(*http.ServeMux).HandleFunc(_URL_GET, getRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_POST, postRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_PUT, putRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_HEAD, headRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_PATCH, patchRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_DELETE, deleteRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_QUERY, queryRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_HEADERS, headersRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_CONTENT_TYPE, contentTypeRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_ACCEPT, acceptRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_USER_AGENT, uaRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_BASIC_AUTH, basicAuthRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_STRING_RESP, stringRespRequestHandler)
	server.Handler.(*http.ServeMux).HandleFunc(_URL_JSON_RESP, jsonRespRequestHandler)

	err = server.Serve(listener)

	if err != nil {
		c.Fatal(err.Error())
	}
}

func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != GET {
		w.WriteHeader(801)
		return
	}

	w.WriteHeader(200)
}

func postRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		w.WriteHeader(802)
		return
	}

	w.WriteHeader(200)
}

func putRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != PUT {
		w.WriteHeader(803)
		return
	}

	w.WriteHeader(200)
}
func headRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != HEAD {
		w.WriteHeader(804)
		return
	}

	w.WriteHeader(200)
}

func patchRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != PATCH {
		w.WriteHeader(805)
		return
	}

	w.WriteHeader(200)
}

func deleteRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != DELETE {
		w.WriteHeader(806)
		return
	}

	w.WriteHeader(200)
}

func queryRequestHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if query.Get("user") != "john" {
		w.WriteHeader(900)
		return
	}

	if query.Get("id") != "912" {
		w.WriteHeader(901)
		return
	}

	_, root := query["root"]

	if !root {
		w.WriteHeader(902)
		return
	}

	w.WriteHeader(200)
}

func headersRequestHandler(w http.ResponseWriter, r *http.Request) {
	header1 := r.Header["Header1"]

	if len(header1) != 1 {
		w.WriteHeader(910)
		return
	}

	if header1[0] != "Value1" {
		w.WriteHeader(911)
		return
	}

	header2 := r.Header["Header2"]

	if len(header2) != 1 {
		w.WriteHeader(912)
		return
	}

	if header2[0] != "Value2" {
		w.WriteHeader(913)
		return
	}

	w.WriteHeader(200)
}

func contentTypeRequestHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header["Content-Type"]

	if len(header) != 1 {
		w.WriteHeader(920)
		return
	}

	if header[0] != _TEST_CONTENT_TYPE {
		w.WriteHeader(921)
		return
	}

	w.WriteHeader(200)
}

func acceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header["Accept"]

	if len(header) != 1 {
		w.WriteHeader(930)
		return
	}

	if header[0] != _TEST_ACCEPT {
		w.WriteHeader(931)
		return
	}

	w.WriteHeader(200)
}

func uaRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() != _TEST_USER_AGENT {
		w.WriteHeader(940)
		return
	}

	w.WriteHeader(200)
}

func basicAuthRequestHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(950)
		return
	}

	if user != _TEST_BASIC_AUTH_USER {
		w.WriteHeader(951)
		return
	}

	if pass != _TEST_BASIC_AUTH_PASS {
		w.WriteHeader(952)
		return
	}

	w.WriteHeader(200)
}

func stringRespRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(_TEST_STRING_RESP))
}

func jsonRespRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
  "string": "test",
  "integer": 912,
  "boolean": true }`,
	))
}
