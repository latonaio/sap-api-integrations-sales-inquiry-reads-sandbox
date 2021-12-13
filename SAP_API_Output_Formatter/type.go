package sap_api_output_formatter

type SalesInquiry struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	SalesInquiry  string `json:"sales_inquiry"`
	Deleted       bool   `json:"deleted"`
}    
    
type Header struct {
	SalesInquiry                   string      `json:"SalesInquiry"`
	SalesInquiryType               string      `json:"SalesInquiryType"`
	SalesOrganization              string      `json:"SalesOrganization"`
	DistributionChannel            string      `json:"DistributionChannel"`
	OrganizationDivision           string      `json:"OrganizationDivision"`
	SalesGroup                     string      `json:"SalesGroup"`
	SalesOffice                    string      `json:"SalesOffice"`
	SalesDistrict                  string      `json:"SalesDistrict"`
	SoldToParty                    string      `json:"SoldToParty"`
	CreationDate                   string      `json:"CreationDate"`
	LastChangeDate                 string      `json:"LastChangeDate"`
	PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
	CustomerPurchaseOrderType      string      `json:"CustomerPurchaseOrderType"`
	CustomerPurchaseOrderDate      string      `json:"CustomerPurchaseOrderDate"`
	SalesInquiryDate               string      `json:"SalesInquiryDate"`
	TotalNetAmount                 string      `json:"TotalNetAmount"`
	TransactionCurrency            string      `json:"TransactionCurrency"`
	SDDocumentReason               string      `json:"SDDocumentReason"`
	PricingDate                    string      `json:"PricingDate"`
	HeaderBillingBlockReason       string      `json:"HeaderBillingBlockReason"`
	BindingPeriodValidityStartDate string      `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   string      `json:"BindingPeriodValidityEndDate"`
	HdrOrderProbabilityInPercent   string      `json:"HdrOrderProbabilityInPercent"`
	ExpectedOrderNetAmount         string      `json:"ExpectedOrderNetAmount"`
	IncotermsClassification        string      `json:"IncotermsClassification"`
	CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
	PaymentMethod                  string      `json:"PaymentMethod"`
	OverallSDProcessStatus         string      `json:"OverallSDProcessStatus"`
	TotalCreditCheckStatus         string      `json:"TotalCreditCheckStatus"`
	OverallSDDocumentRejectionSts  string      `json:"OverallSDDocumentRejectionSts"`
	ToHeaderPartner                string      `json:"to_Partner"`
	ToItem                         string      `json:"to_Item"`	
}

type Item struct {
	SalesInquiry                  string `json:"SalesInquiry"`
	SalesInquiryItem              string `json:"SalesInquiryItem"`
	SalesInquiryItemCategory      string `json:"SalesInquiryItemCategory"`
	SalesInquiryItemText          string `json:"SalesInquiryItemText"`
	PurchaseOrderByCustomer       string `json:"PurchaseOrderByCustomer"`
	Material                      string `json:"Material"`
	MaterialByCustomer            string `json:"MaterialByCustomer"`
	RequestedQuantity             string `json:"RequestedQuantity"`
	RequestedQuantityUnit         string `json:"RequestedQuantityUnit"`
	ItemOrderProbabilityInPercent string `json:"ItemOrderProbabilityInPercent"`
	ItemGrossWeight               string `json:"ItemGrossWeight"`
	ItemNetWeight                 string `json:"ItemNetWeight"`
	ItemWeightUnit                string `json:"ItemWeightUnit"`
	ItemVolume                    string `json:"ItemVolume"`
	ItemVolumeUnit                string `json:"ItemVolumeUnit"`
	TransactionCurrency           string `json:"TransactionCurrency"`
	NetAmount                     string `json:"NetAmount"`
	MaterialGroup                 string `json:"MaterialGroup"`
	Batch                         string `json:"Batch"`
	IncotermsClassification       string `json:"IncotermsClassification"`
	CustomerPaymentTerms          string `json:"CustomerPaymentTerms"`
	SalesDocumentRjcnReason       string `json:"SalesDocumentRjcnReason"`
	WBSElement                    string `json:"WBSElement"`
	SDProcessStatus               string `json:"SDProcessStatus"`
	ToItemPricingElement          string `json:"to_PricingElement"`
}

type ToHeaderPartner struct {
	SalesInquiry    string `json:"SalesInquiry"`
	PartnerFunction string `json:"PartnerFunction"`
	Customer        string `json:"Customer"`
	Supplier        string `json:"Supplier"`
}

type ToItem struct {
	SalesInquiry                  string `json:"SalesInquiry"`
	SalesInquiryItem              string `json:"SalesInquiryItem"`
	SalesInquiryItemCategory      string `json:"SalesInquiryItemCategory"`
	SalesInquiryItemText          string `json:"SalesInquiryItemText"`
	PurchaseOrderByCustomer       string `json:"PurchaseOrderByCustomer"`
	Material                      string `json:"Material"`
	MaterialByCustomer            string `json:"MaterialByCustomer"`
	RequestedQuantity             string `json:"RequestedQuantity"`
	RequestedQuantityUnit         string `json:"RequestedQuantityUnit"`
	ItemOrderProbabilityInPercent string `json:"ItemOrderProbabilityInPercent"`
	ItemGrossWeight               string `json:"ItemGrossWeight"`
	ItemNetWeight                 string `json:"ItemNetWeight"`
	ItemWeightUnit                string `json:"ItemWeightUnit"`
	ItemVolume                    string `json:"ItemVolume"`
	ItemVolumeUnit                string `json:"ItemVolumeUnit"`
	TransactionCurrency           string `json:"TransactionCurrency"`
	NetAmount                     string `json:"NetAmount"`
	MaterialGroup                 string `json:"MaterialGroup"`
	Batch                         string `json:"Batch"`
	IncotermsClassification       string `json:"IncotermsClassification"`
	CustomerPaymentTerms          string `json:"CustomerPaymentTerms"`
	SalesDocumentRjcnReason       string `json:"SalesDocumentRjcnReason"`
	WBSElement                    string `json:"WBSElement"`
	SDProcessStatus               string `json:"SDProcessStatus"`
	ToItemPricingElement          string `json:"to_PricingElement"`
}

type ToItemPricingElement struct {
	SalesInquiry                  string `json:"SalesInquiry"`
	SalesInquiryItem              string `json:"SalesInquiryItem"`
	PricingProcedureStep          string `json:"PricingProcedureStep"`
	PricingProcedureCounter       string `json:"PricingProcedureCounter"`
	ConditionApplication          string `json:"ConditionApplication"`
	ConditionType                 string `json:"ConditionType"`
	PricingDateTime               string `json:"PricingDateTime"`
	ConditionCalculationType      string `json:"ConditionCalculationType"`
	ConditionBaseValue            string `json:"ConditionBaseValue"`
	ConditionRateValue            string `json:"ConditionRateValue"`
	ConditionCurrency             string `json:"ConditionCurrency"`
	ConditionQuantity             string `json:"ConditionQuantity"`
	ConditionQuantityUnit         string `json:"ConditionQuantityUnit"`
	ConditionCategory             string `json:"ConditionCategory"`
	PricingScaleType              string `json:"PricingScaleType"`
	ConditionRecord               string `json:"ConditionRecord"`
	ConditionSequentialNumber     string `json:"ConditionSequentialNumber"`
	TaxCode                       string `json:"TaxCode"`
	ConditionAmount               string `json:"ConditionAmount"`
	TransactionCurrency           string `json:"TransactionCurrency"`
	PricingScaleBasis             string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue      string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit       string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency   string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged    bool   `json:"ConditionIsManuallyChanged"`
}
