package productdata

type PageInfo struct {
	Start           string `json:"@start"`
	End             string `json:"@end"`
	Uri             string `json:"@uri"`
	Total           string `json:"@total"`
	FirstPageUri    string `json:"@firstpageuri"`
	LastPageUri     string `json:"@lastpageuri"`
	NextPageUri     string `json:"@nextpageuri"`
	NumPages        string `json:"@numpages"`
	Page            string `json:"@page"`
	PageSize        string `json:"@pagesize"`
	PreviousPageUri string `json:"@previouspageuri"`
}

type CatalogsResponse struct {
	PageInfo
	Catalogs []CatalogInfo
}

type CatalogInfo struct {
	AdvertiserId    string
	AdvertiserName  string
	CampaignId      string
	CampaignName    string
	DateLastUpdated string
	Id              string
	ItemsUri        string
	Locations       []string
	Name            string
	NumberOfItems   string
	Uri             string
}

type CatalogItemsResponse struct {
	PageInfo
	Items []ItemInfo
}

type ItemInfo struct {
	AdditionalImageUrls      []string
	Adult                    bool
	AgeGroup                 string
	AgeRangeMax              string
	AgeRangeMin              string
	AgeRangeUnit             string
	Asin                     string
	Bullets                  []string // TODO: check type
	CampaignId               string
	CampaignName             string
	CatalogId                string
	CatalogItemId            string
	Category                 string
	Colors                   []string
	Condition                string
	Currency                 string
	CurrentPrice             string
	Description              string
	EstimatedShipDate        string
	ExpirationDate           string
	Gender                   string
	Gtin                     string
	GtinType                 string
	Id                       string
	ImageUrl                 string
	IsParent                 bool
	Labels                   []string // TODO: check type
	LaunchDate               string
	Manufacturer             string
	Material                 string
	MobileUrl                string
	Money1                   string
	Money2                   string
	Money3                   string
	Mpn                      string
	MultiPack                string
	Name                     string
	Numeric1                 string
	Numeric2                 string
	Numeric3                 string
	OriginalFormatCategory   string
	OriginalFormatCategoryId string
	OriginalPrice            string
	ParentName               string
	ParentSky                string
	Pattern                  string
	ProductBid               string
	ShippingHeight           string
	ShippingLabel            string
	ShippingLength           string
	ShippingLengthUnit       string
	ShippingRate             string
	ShippingWeight           string
	ShippingWeightUnit       string
	ShippingWidth            string
	Size                     string
	SizeUnit                 string
	StockAvailability        string
	Text1                    string
	Text2                    string
	Text3                    string
	Uri                      string
	Url                      string
	Weight                   string
	WeightUnit               string
}
