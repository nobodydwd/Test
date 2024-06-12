package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type RPCRequest struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      int             `json:"id"`
}

type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
	ID      int         `json:"id"`
}

var balances = map[string]int{
	"alex":  1000,
	"nasry": 500,
}

func transfer(params json.RawMessage) (interface{}, error) {
	var args []string
	if err := json.Unmarshal(params, &args); err != nil {
		return nil, err
	}

	from := args[0]
	to := args[1]
	amount, err := strconv.Atoi(args[2])
	if err != nil {
		return nil, err
	}

	if balances[from] < amount {
		return nil, fmt.Errorf("insufficient funds")
	}

	balances[from] -= amount
	balances[to] += amount

	return fmt.Sprintf("Transferred %d from %s to %s", amount, from, to), nil
}

func handleRPCRequest(w http.ResponseWriter, r *http.Request) {
	var req RPCRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result interface{}
	var err error

	switch req.Method {
	case "eth_sendTransaction":
		result, err = transfer(req.Params)
	default:
		err = fmt.Errorf("method not supported")
	}

	resp := RPCResponse{
		Jsonrpc: "2.0",
		Result:  result,
		ID:      req.ID,
	}

	if err != nil {
		resp.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func startRPCServer() {
	http.HandleFunc("/", handleRPCRequest)
	log.Println("Starting JSON-RPC server on :8546")
	log.Fatal(http.ListenAndServe(":8546", nil))
}
