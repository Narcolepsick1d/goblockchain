package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

type WalletServer struct {
	port    uint16
	gateway string
}

const tempDir = "wallet_server/templates"

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port: port, gateway: gateway}
}
func (ws *WalletServer) Port() uint16 {
	return ws.port
}
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}
func (ws *WalletServer) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("error: invalid http method ")
	}
}
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.port)), nil))
}
