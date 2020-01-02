package quickbooks

import (
	"encoding/json"
	"fmt"
)

// PaymentObject the complete quickbooks payment object type
type PaymentObject struct {
	Payment Payment `json:"Payment"`
	Time    string  `json:"time"`
}

// Payment quickbooks payment type
type Payment struct {
	ID                  string        `json:"Id,omitempty"`
	CustomerRef         *CustomerRef  `json:"CustomerRef,omitempty"`
	DepositToAccountRef *AccountRef   `json:"DepositToAccountRef,omitempty"`
	TotalAmt            float64       `json:"TotalAmt"`
	UnappliedAmt        float64       `json:"UnappliedAmt,omitempty"`
	ProcessPayment      bool          `json:"ProcessPayment,omitempty"`
	Domain              string        `json:"domain,omitempty"`
	Sparse              bool          `json:"sparse,omitempty"`
	SyncToken           string        `json:"SyncToken,omitempty"`
	TxnDate             string        `json:"TxnDate,omitempty"`
	Line                []PaymentLine `json:"Line"`
	MetaData            *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// PaymentLine quickbooks payment line object
type PaymentLine struct {
	Amount    float64     `json:"Amount"`
	LinkedTxn []LinkedTxn `json:"LinkedTxn"`
	LineEx    *LineEx     `json:"LineEx,omitempty"`
}

// LineEx quickbooks payment LineEx object
type LineEx struct {
	Any []struct {
		Name         string `json:"name"`
		DeclaredType string `json:"declaredType"`
		Scope        string `json:"scope"`
		Value        struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"value"`
		Nil             bool `json:"nil"`
		GlobalScope     bool `json:"globalScope"`
		TypeSubstituted bool `json:"typeSubstituted"`
	} `json:"any"`
}

// GetPayment returns payment from PaymentID
func (q *Quickbooks) GetPayment(PaymentID string) (*PaymentObject, error) {
	endpoint := fmt.Sprintf("/company/%s/payment/%s", q.RealmID, PaymentID)

	res, err := q.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	paymentObject := PaymentObject{}
	err = json.NewDecoder(res.Body).Decode(&paymentObject)
	if err != nil {
		return nil, err
	}

	return &paymentObject, nil
}
