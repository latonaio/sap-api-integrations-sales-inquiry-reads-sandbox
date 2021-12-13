package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-inquiry-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	SalesInquiry:                   data.SalesInquiry,
	SalesInquiryType:               data.SalesInquiryType,
	SalesOrganization:              data.SalesOrganization,
	DistributionChannel:            data.DistributionChannel,
	OrganizationDivision:           data.OrganizationDivision,
	SalesGroup:                     data.SalesGroup,
	SalesOffice:                    data.SalesOffice,
	SalesDistrict:                  data.SalesDistrict,
	SoldToParty:                    data.SoldToParty,
	CreationDate:                   data.CreationDate,
	LastChangeDate:                 data.LastChangeDate,
	PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
	CustomerPurchaseOrderType:      data.CustomerPurchaseOrderType,
	CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
	SalesInquiryDate:               data.SalesInquiryDate,
	TotalNetAmount:                 data.TotalNetAmount,
	TransactionCurrency:            data.TransactionCurrency,
	SDDocumentReason:               data.SDDocumentReason,
	PricingDate:                    data.PricingDate,
	HeaderBillingBlockReason:       data.HeaderBillingBlockReason,
	BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
	BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
	HdrOrderProbabilityInPercent:   data.HdrOrderProbabilityInPercent,
	ExpectedOrderNetAmount:         data.ExpectedOrderNetAmount,
	IncotermsClassification:        data.IncotermsClassification,
	CustomerPaymentTerms:           data.CustomerPaymentTerms,
	PaymentMethod:                  data.PaymentMethod,
	OverallSDProcessStatus:         data.OverallSDProcessStatus,
	TotalCreditCheckStatus:         data.TotalCreditCheckStatus,
	OverallSDDocumentRejectionSts:  data.OverallSDDocumentRejectionSts,
	ToHeaderPartner:                data.ToHeaderPartner.Deferred.URI,
	ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
	SalesInquiry:                  data.SalesInquiry,
	SalesInquiryItem:              data.SalesInquiryItem,
	SalesInquiryItemCategory:      data.SalesInquiryItemCategory,
	SalesInquiryItemText:          data.SalesInquiryItemText,
	PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
	Material:                      data.Material,
	MaterialByCustomer:            data.MaterialByCustomer,
	RequestedQuantity:             data.RequestedQuantity,
	RequestedQuantityUnit:         data.RequestedQuantityUnit,
	ItemOrderProbabilityInPercent: data.ItemOrderProbabilityInPercent,
	ItemGrossWeight:               data.ItemGrossWeight,
	ItemNetWeight:                 data.ItemNetWeight,
	ItemWeightUnit:                data.ItemWeightUnit,
	ItemVolume:                    data.ItemVolume,
	ItemVolumeUnit:                data.ItemVolumeUnit,
	TransactionCurrency:           data.TransactionCurrency,
	NetAmount:                     data.NetAmount,
	MaterialGroup:                 data.MaterialGroup,
	Batch:                         data.Batch,
	IncotermsClassification:       data.IncotermsClassification,
	CustomerPaymentTerms:          data.CustomerPaymentTerms,
	SalesDocumentRjcnReason:       data.SalesDocumentRjcnReason,
	WBSElement:                    data.WBSElement,
	SDProcessStatus:               data.SDProcessStatus,
	ToItemPricingElement:          data.ToItemPricingElement.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
	SalesInquiry:    data.SalesInquiry,
	PartnerFunction: data.PartnerFunction,
	Customer:        data.Customer,
	Supplier:        data.Supplier,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
	SalesInquiry:                  data.SalesInquiry,
	SalesInquiryItem:              data.SalesInquiryItem,
	SalesInquiryItemCategory:      data.SalesInquiryItemCategory,
	SalesInquiryItemText:          data.SalesInquiryItemText,
	PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
	Material:                      data.Material,
	MaterialByCustomer:            data.MaterialByCustomer,
	RequestedQuantity:             data.RequestedQuantity,
	RequestedQuantityUnit:         data.RequestedQuantityUnit,
	ItemOrderProbabilityInPercent: data.ItemOrderProbabilityInPercent,
	ItemGrossWeight:               data.ItemGrossWeight,
	ItemNetWeight:                 data.ItemNetWeight,
	ItemWeightUnit:                data.ItemWeightUnit,
	ItemVolume:                    data.ItemVolume,
	ItemVolumeUnit:                data.ItemVolumeUnit,
	TransactionCurrency:           data.TransactionCurrency,
	NetAmount:                     data.NetAmount,
	MaterialGroup:                 data.MaterialGroup,
	Batch:                         data.Batch,
	IncotermsClassification:       data.IncotermsClassification,
	CustomerPaymentTerms:          data.CustomerPaymentTerms,
	SalesDocumentRjcnReason:       data.SalesDocumentRjcnReason,
	WBSElement:                    data.WBSElement,
	SDProcessStatus:               data.SDProcessStatus,
	ToItemPricingElement:          data.ToItemPricingElement.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
	SalesInquiry:                  data.SalesInquiry,
	SalesInquiryItem:              data.SalesInquiryItem,
	PricingProcedureStep:          data.PricingProcedureStep,
	PricingProcedureCounter:       data.PricingProcedureCounter,
	ConditionApplication:          data.ConditionApplication,
	ConditionType:                 data.ConditionType,
	PricingDateTime:               data.PricingDateTime,
	ConditionCalculationType:      data.ConditionCalculationType,
	ConditionBaseValue:            data.ConditionBaseValue,
	ConditionRateValue:            data.ConditionRateValue,
	ConditionCurrency:             data.ConditionCurrency,
	ConditionQuantity:             data.ConditionQuantity,
	ConditionQuantityUnit:         data.ConditionQuantityUnit,
	ConditionCategory:             data.ConditionCategory,
	PricingScaleType:              data.PricingScaleType,
	ConditionRecord:               data.ConditionRecord,
	ConditionSequentialNumber:     data.ConditionSequentialNumber,
	TaxCode:                       data.TaxCode,
	ConditionAmount:               data.ConditionAmount,
	TransactionCurrency:           data.TransactionCurrency,
	PricingScaleBasis:             data.PricingScaleBasis,
	ConditionScaleBasisValue:      data.ConditionScaleBasisValue,
	ConditionScaleBasisUnit:       data.ConditionScaleBasisUnit,
	ConditionScaleBasisCurrency:   data.ConditionScaleBasisCurrency,
	ConditionIsManuallyChanged:    data.ConditionIsManuallyChanged,
		})
	}

	return toItemPricingElement, nil
}
