package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			ToHeaderPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItem                         struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
