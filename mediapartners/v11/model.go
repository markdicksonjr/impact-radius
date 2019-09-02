package v11

type CompanyInfoResponse struct {
	CompanyName                 string
	PrimaryPromotionalMethod    string
	PromotionalMethods          string
	Website                     string
	PrimaryPhoneNumber          string
	PrimaryPhoneNumberCountry   string
	SecondaryPhoneNumber        string
	SecondaryPhoneNumberCountry string
	MinimumContactRating        string
	Timezone                    string
	Currency                    string
	RegisteredForIndirectTax    string
	IndirectTaxNumber           string
	OrganizationType            string
	EinSsnForeignTaxId          string
	CorporateAddress            *Address `json:",omitempty"`
	BillingAddress              *Address `json:",omitempty"`
	FinanceContact              *Contact `json:",omitempty"`
	TechnicalContact            *Contact `json:",omitempty"`
	CommercialContact           *Contact `json:",omitempty"`
	Uri                         string
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PostalCode   string
	Country      string
}

type Contact struct {
	UserId                 string
	Name                   string
	Email                  string
	WorkPhoneNumber        string
	WorkPhoneNumberCountry string
	CellPhoneNumber        string
	CellPhoneNumberCountry string
}
