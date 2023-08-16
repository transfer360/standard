package notice

// https://transfer360.dev/#/status/gethirerinformation

type NoticeTOL struct {
	CompanyName  string `json:"companyname,omitempty"`
	Name         string `json:"first-name,omitempty"`
	Surname      string `json:"last-name,omitempty"`
	AddressLine1 string `json:"addressline1,omitempty"`
	AddressLine2 string `json:"addressline2,omitempty"`
	AddressLine3 string `json:"addressline3,omitempty"`
	PostCode     string `json:"post_code,omitempty"`
	Country      string `json:"country,omitempty"`
	Primary      bool   `json:"primary"`
}

type NoticeUpdateFromLease struct {
	Status            int       `json:"status"`
	LeaseCompanyNote  string    `json:"lease_company_note"`
	ContactInfo       NoticeTOL `json:"contact_information"`
	ContractNumber    string    `json:"agreement_no,omitempty"`
	AgreementURL      string    `json:"agreement_url,omitempty"`
	ContractStartDate string    `json:"hire_from,omitempty"`
	ContractEndDate   string    `json:"hire_to,omitempty"`
}
