package quickbooks

// VendorObject the complete quickbooks vendor object type
type VendorObject struct {
	Vendor 	Vendor `json:"Vendor"`
	Time     string   `json:"time"`
}

// Vendor quickbooks vendor type
type Vendor struct {
	ID                      string   `json:"Id,omitempty"`
	Vendor1099				bool	 `json:"Vendor1099,omitempty"`
	BillAddr                *Address `json:"BillAddr,omitempty"`
	Domain                  string   `json:"domain,omitempty"`
	Sparse                  bool     `json:"sparse,omitempty"`
	SyncToken               string   `json:"SyncToken,omitempty"`
	GivenName               string   `json:"GivenName"`
	DisplayName             string   `json:"DisplayName"`
	PrintOnCheckName     	string   `json:"PrintOnCheckName,omitempty"`
	CompanyName             string   `json:"CompanyName,omitempty"`
	Active                  bool     `json:"Active,omitempty"`
	Balance                 float64  `json:"Balance,omitempty"`

	PrimaryPhone *struct {
		FreeFormNumber string `json:"FreeFormNumber"`
	} `json:"PrimaryPhone,omitempty"`
	PrimaryEmailAddr *struct {
		Address string `json:"Address"`
	} `json:"PrimaryEmailAddr,omitempty"`

	MetaData *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// VendorRef quickbooks vendor reference object
type VendorRef struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}