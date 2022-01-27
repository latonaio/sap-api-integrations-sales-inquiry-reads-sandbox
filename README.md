# sap-api-integrations-sales-inquiry-reads 
sap-api-integrations-sales-inquiry-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で販売引合データ を取得するマイクロサービスです。    
sap-api-integrations-sales-inquiry-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-sales-inquiry-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_SALES_INQUIRY_SRV_0001/overview

## 動作環境  
sap-api-integrations-sales-inquiry-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-sales-inquiry-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。 

## 本レポジトリ が 対応する API サービス
sap-api-integrations-sales-inquiry-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_SALES_INQUIRY_SRV_0001/overview
* APIサービス名(=baseURL): API_SALES_INQUIRY_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-sales-inquiry-reads には、次の API をコールするためのリソースが含まれています。  

* A_SalesInquiry（販売引合 - ヘッダ）※販売引合の詳細データを取得するために、ToHeaderPartner、ToItem、と合わせて利用されます。
* ToItem（販売引合 - 明細）※販売引合明細の詳細データを取得するために、ToItemPricingElementと合わせて利用されます。
* ToHeaderPartner（販売引合 - ヘッダ取引先）
* ToItemPricingElement（販売引合 - 明細価格条件）
* A_SalesInquiryItem（販売引合明細）※販売引合明細の詳細データを取得するために、ToItemPricingElementと合わせて利用されます。
* ToItemPricingElement（販売引合明細 - 明細価格条件）

## API への 値入力条件 の 初期値
sap-api-integrations-sales-inquiry-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.SalesInquiry.SalesInquiry（販売引合番号）
* inoutSDC.SalesInquiry.SalesInquiryItem.SalesInquiryItem（販売引合明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、 "Header", "Item" が指定されています。    
  
```
"api_schema": "sap.s4.beh.salesinquiry.v1.SalesInquiry.Created.v1",
"accepter": ["Header"],
"sales_inquiry": "10000000",
"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
"api_schema": "sap.s4.beh.salesinquiry.v1.SalesInquiry.Created.v1",
"accepter": ["All"],
"sales_inquiry": "10000000",
"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetSalesInquiry(salesInquiry, salesInquiryItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(salesInquiry)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(salesInquiry, salesInquiryItem)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 販売引合 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"SalesInquiry" ～ "ToHeaderPartner" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-sales-inquiry-reads/SAP_API_Caller/caller.go#L59",
	"function": "sap-api-integrations-sales-inquiry-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"SalesInquiry": "10000000",
			"SalesInquiryType": "IN",
			"SalesOrganization": "1710",
			"DistributionChannel": "10",
			"OrganizationDivision": "00",
			"SalesGroup": "",
			"SalesOffice": "",
			"SalesDistrict": "",
			"SoldToParty": "17100003",
			"CreationDate": "2017-02-28T09:00:00+09:00",
			"LastChangeDate": "2017-02-28T09:00:00+09:00",
			"PurchaseOrderByCustomer": "Test Inquiry",
			"CustomerPurchaseOrderType": "",
			"CustomerPurchaseOrderDate": "",
			"SalesInquiryDate": "2017-02-28T09:00:00+09:00",
			"TotalNetAmount": "0.00",
			"TransactionCurrency": "USD",
			"SDDocumentReason": "",
			"PricingDate": "2017-02-28T09:00:00+09:00",
			"HeaderBillingBlockReason": "",
			"BindingPeriodValidityStartDate": "",
			"BindingPeriodValidityEndDate": "",
			"HdrOrderProbabilityInPercent": "0",
			"ExpectedOrderNetAmount": "0.00",
			"IncotermsClassification": "EXW",
			"CustomerPaymentTerms": "0004",
			"PaymentMethod": "",
			"OverallSDProcessStatus": "C",
			"TotalCreditCheckStatus": "",
			"OverallSDDocumentRejectionSts": "C",
			"to_Partner": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SALES_INQUIRY_SRV/A_SalesInquiry('10000000')/to_Partner",
			"to_Item": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_SALES_INQUIRY_SRV/A_SalesInquiry('10000000')/to_Item"
		}
	],
	"time": "2022-01-27T22:15:35+09:00"
}

```
