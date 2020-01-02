package quickbooks

import (
	"encoding/json"
	"fmt"
)

// BillPaymentObject the complete quickbooks billpayment object type
type BillPaymentObject struct {
	BillPayment BillPayment `json:"BillPayment"`
	Time    string  `json:"time"`
}

// BillPayment quickbooks payment type
type BillPayment struct {
	ID                  string        `json:"Id,omitempty"`
	VendorRef         	*VendorRef    `json:"VendorRef,omitempty"`
	DocNumber 			string   	  `json:"DocNumber,omitempty"`
	TotalAmt            float64       `json:"TotalAmt"`
	ProcessBillPayment  bool          `json:"ProcessBillPayment,omitempty"`
	Domain              string        `json:"domain,omitempty"`
	PrivateNote         string        `json:"PrivateNote,omitempty"`
	PayType         	string        `json:"PayType,omitempty"`
	Sparse              bool          `json:"sparse,omitempty"`
	SyncToken           string        `json:"SyncToken,omitempty"`
	TxnDate             string        `json:"TxnDate,omitempty"`
	Line                []BillPaymentLine `json:"Line"`
	MetaData            *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// BillPaymentLine quickbooks payment line object
type BillPaymentLine struct {
	Amount    float64     `json:"Amount"`
	LinkedTxn []LinkedTxn `json:"LinkedTxn"`
}

// GetBillPayment returns payment from BillPaymentID
func (q *Quickbooks) GetBillPayment(BillPaymentID string) (*BillPaymentObject, error) {
	endpoint := fmt.Sprintf("/company/%s/billpayment/%s", q.RealmID, BillPaymentID)

	res, err := q.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	billPaymentObject := BillPaymentObject{}
	err = json.NewDecoder(res.Body).Decode(&billPaymentObject)
	if err != nil {
		return nil, err
	}

	return &billPaymentObject, nil
}