package ripple

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/rubblelabs/ripple/crypto"
	"github.com/rubblelabs/ripple/data"
)

func executeCommand(req []byte) (res []byte) {

	c, _, e := websocket.DefaultDialer.Dial("wss://s1.ripple.com", nil)
	if e != nil {
		log.Fatal(e)
	}
	defer c.Close()

	e = c.WriteMessage(websocket.TextMessage, req)
	if e != nil {
		log.Fatal(e)
	}

	_, res, e = c.ReadMessage()
	if e != nil {
		log.Fatal(e)
	}

	e = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if e != nil {
		log.Fatal(e)
	}

	return
}

func formatJSON(b []byte) string {
	var buffer bytes.Buffer
	e := json.Indent(&buffer, b, "", "  ")
	if e != nil {
		log.Fatal(e)
	}
	return string(buffer.Bytes())
}

// Handle executes a ripple command.
func Handle(w http.ResponseWriter, r *http.Request, getCommand func(r *http.Request) interface{}) interface{} {

	var req string
	var res string

	if r.Method == "POST" {

		r.ParseForm()
		command := getCommand(r)

		bytes, e := json.Marshal(command)
		if e != nil {
			log.Fatal(e)
		}

		req = formatJSON(bytes)
		res = formatJSON(executeCommand(bytes))
	}

	return struct {
		Req string
		Res string
	}{
		req,
		res,
	}
}

// GetCommandPing creates a "ping" command.
func GetCommandPing(r *http.Request) interface{} {
	return struct {
		ID      string `json:"id"`
		Command string `json:"command"`
	}{
		ID:      r.Form.Get("id"),
		Command: r.Form.Get("command"),
	}
}

// GetCommandAccountInfo creates an "account_info" command.
func GetCommandAccountInfo(r *http.Request) interface{} {
	return struct {
		ID      string `json:"id"`
		Command string `json:"command"`
		Ledger  string `json:"ledger_index"`
		Account string `json:"account"`
	}{
		ID:      r.Form.Get("id"),
		Command: r.Form.Get("command"),
		Ledger:  r.Form.Get("ledger"),
		Account: r.Form.Get("account"),
	}
}

// GetCommandAccountLines creates an "account_lines" command.
func GetCommandAccountLines(r *http.Request) interface{} {
	return struct {
		ID      string `json:"id"`
		Command string `json:"command"`
		Ledger  string `json:"ledger_index"`
		Account string `json:"account"`
	}{
		ID:      r.Form.Get("id"),
		Command: r.Form.Get("command"),
		Ledger:  r.Form.Get("ledger"),
		Account: r.Form.Get("account"),
	}
}

// GetCommandSubmitTrustSet creates a "submit" command for a TrustSet transaction.
func GetCommandSubmitTrustSet(r *http.Request) interface{} {

	transactionType := getFormValue(r, "transactionType").toUint16()
	account := getFormValue(r, "account").toString()
	secretKey := getFormValue(r, "secretKey").toString()
	fee := getFormValue(r, "fee").toString()
	sequence := getFormValue(r, "sequence").toUint32()
	lastLedgerSequence := getFormValue(r, "lastLedgerSequence").toUint32()

	flag := uint32(0)
	flag |= normalizeFlag(getFormValue(r, "flagFullyCanonicalSig")).toUint32()
	flag |= normalizeFlag(getFormValue(r, "flagSetAuth")).toUint32()
	flag |= normalizeFlag(getFormValue(r, "flagSetNoRipple")).toUint32()
	flag |= normalizeFlag(getFormValue(r, "flagClearNoRipple")).toUint32()
	flag |= normalizeFlag(getFormValue(r, "flagSetFreeze")).toUint32()
	flag |= normalizeFlag(getFormValue(r, "flagClearFreeze")).toUint32()

	value := getFormValue(r, "limitAmountValue").toString()
	currency := getFormValue(r, "limitAmountCurrency").toString()
	issuer := getFormValue(r, "limitAmountIssuer").toString()
	limitAmount := value + "/" + currency + "/" + issuer

	txTransactionType := data.TransactionType(transactionType)

	txAccount, e := data.NewAccountFromAddress(account)
	if e != nil {
		log.Fatal(e)
	}

	txFee, e := data.NewValue(fee, true)
	if e != nil {
		log.Fatal(e)
	}

	txSequence := sequence
	txLastLedgerSequence := lastLedgerSequence
	txTransactionFlag := data.TransactionFlag(flag)

	txLimitAmount, e := data.NewAmount(limitAmount)
	if e != nil {
		log.Fatal(e)
	}

	txBase := data.TxBase{
		TransactionType:    data.TransactionType(txTransactionType),
		Account:            *txAccount,
		Fee:                *txFee,
		Sequence:           txSequence,
		LastLedgerSequence: &txLastLedgerSequence,
		Flags:              &txTransactionFlag,
	}

	txTrustSet := data.TrustSet{
		TxBase:      txBase,
		LimitAmount: *txLimitAmount,
	}

	seed, e := crypto.NewRippleHash(secretKey)
	if e != nil {
		log.Fatal(e)
	}

	key, e := crypto.NewECDSAKey(seed.Payload())
	if e != nil {
		log.Fatal(e)
	}

	keySequence := uint32(0)
	data.Sign(&txTrustSet, key, &keySequence)

	_, raw, _ := data.Raw(&txTrustSet)
	txBlob := fmt.Sprintf("%X", raw)

	return struct {
		ID      string `json:"id"`
		Command string `json:"command"`
		TxBlob  string `json:"tx_blob"`
	}{
		ID:      r.Form.Get("id"),
		Command: r.Form.Get("command"),
		TxBlob:  txBlob,
	}
}

type formValue string

func getFormValue(r *http.Request, key string) formValue {
	s := r.Form.Get(key)
	return formValue(s)
}

func normalizeFlag(v formValue) formValue {
	s := string(v)
	if s == "" {
		return formValue("0x00000000")
	}
	return v
}

func (v formValue) toString() string {
	return string(v)
}

func (v formValue) toUint16() uint16 {
	s := string(v)
	value, e := strconv.ParseUint(s, 0, 16)
	if e != nil {
		log.Fatal(e)
	}
	return uint16(value)
}

func (v formValue) toUint32() uint32 {
	s := string(v)
	value, e := strconv.ParseUint(s, 0, 32)
	if e != nil {
		log.Fatal(e)
	}
	return uint32(value)
}
