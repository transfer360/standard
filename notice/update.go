package notice

import "time"

// https://transfer360.dev/#/status/gethirerinformation

type NoticeTOL struct {
	CompanyName  string `json:"companyname,omitempty" firestore:"CompanyName"`
	Name         string `json:"first-name,omitempty" firestore:"Name"`
	Surname      string `json:"last-name,omitempty" firestore:"Surname"`
	AddressLine1 string `json:"addressline1,omitempty" firestore:"AddressLine1"`
	AddressLine2 string `json:"addressline2,omitempty" firestore:"AddressLine2"`
	AddressLine3 string `json:"addressline3,omitempty" firestore:"AddressLine3"`
	AddressLine4 string `json:"addressline4,omitempty" firestore:"AddressLine4"`
	PostCode     string `json:"post_code,omitempty" firestore:"PostCode"`
	Country      string `json:"country,omitempty" firestore:"Country"`
	Primary      bool   `json:"primary" firestore:"Primary"`
}

type NoticeUpdateFromLease struct {
	Status            int       `json:"status" firestore:"Status"`
	LeaseCompanyNote  string    `json:"lease_company_note" firestore:"LeaseCompanyNote"`
	ContactInfo       NoticeTOL `json:"contact_information" firestore:"ContactInfo"`
	ContractNumber    string    `json:"agreement_no,omitempty" firestore:"ContractNumber"`
	AgreementURL      string    `json:"agreement_url,omitempty" firestore:"AgreementURL"`
	ContractStartDate string    `json:"hire_from,omitempty" firestore:"ContractStartDate"`
	ContractEndDate   string    `json:"hire_to,omitempty" firestore:"ContractEndDate"`
}

type ParkingChargeNoticeSearchResult struct {
	FleetID      int                   `json:"fleet_id" firestore:"FleetID"`
	LeaseReturn  NoticeUpdateFromLease `json:"lease_return" firestore:"LeaseReturn"`
	DateOfUpdate time.Time             `json:"date_of_update" firestore:"DateOfUpdate"`
	Sref         string                `json:"sref" firestore:"Sref"`
	ClientID     string                `json:"client_id" firestore:"ClientID"`
}
