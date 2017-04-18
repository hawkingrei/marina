package bitbucket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hawkingrei/marina/conf"
	"io/ioutil"
	"net/http"
)

type HTTPServer struct {
	config *conf.Config
	server *http.Server
}

func NewHTTPServer(c *conf.Config) *HTTPServer {
	s := &HTTPServer{
		config: c,
	}

	router := mux.NewRouter()
	router.HandleFunc("/bitbucket", s.recWebhook).Methods("POST")

	s.server = &http.Server{Addr: c.HttpAddr, Handler: router}

	return s
}

func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) recWebhook(w http.ResponseWriter, req *http.Request) {
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	var eventPayload Payload
	json.Unmarshal(payload, &eventPayload)
	fmt.Println(eventPayload.Changesets.Values[0].ToCommit.Author.Name)
	fmt.Println(eventPayload.Changesets.Values[0].ToCommit.Author.EmailAddress)
	fmt.Println(eventPayload.Changesets.Values[0].ToCommit.Message)
	fmt.Println(eventPayload.Changesets.Values[0].ToCommit.Id)
}
