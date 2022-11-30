# convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm-payment-requisition

convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm-payment-requisition は、周辺業務システム　を データ連携基盤 と統合することを目的に データ連携基盤 支払依頼データを 三菱UFJ銀行 支払依頼データ に変換するマイクロサービスです。  
https://xxx.xxx.io/api/FUNCTION_PAYMENT_REQUISITION_SRV/creates/

## 動作環境

convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm-payment-requisition の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm-payment-requisition が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/FUNCTION_PAYMENT_REQUISITION_SRV/creates/

## 本レポジトリ に 含まれる API名
convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm-payment-requisition には、次の API をコールするためのリソースが含まれています。  

* A_Header（支払依頼 - ヘッダデータ）
* A_Item（支払依頼 - 明細データ）

## API への 値入力条件 の 初期値
convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm-payment-requisition において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "DPFMPaymentRequisitionCreates",
	"accepter": ["Header"],
	"payment_requisition_id": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMPaymentRequisitionCreates",
	"accepter": ["All"],
	"payment_requisition_id": null,
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncPRMitsubishiCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,

	log *logger.Logger,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	convertFin := make(chan error)

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, convertFin, log, &errs, input, output)
		case "Item":
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}

	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-convertFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.Errorf("time out"))
		return errs
	}

	return nil
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は データ連携基盤 支払依頼 データが 三菱UFJ銀行 支払依頼 データ に変換された結果の JSON の例です。  
以下の項目のうち、"Payer" ～ "totalAmount" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。  

```
{
  "connection_key": "request",
  "result": true,
  "redis_key": "abcdefg",
  "filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
  "api_status_code": 200,
  "runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
  "business_partner": 201,
  "service_label": "PAYMENT_REQUISITION",
  "Message": {
      "Header": {
          "Payer": 101,
          "PayerPaymentDate": "2022-11-30",
          "PayerPaymentRequisitionID": 1,
          "transactionCode": null,
          "paymentApplicantCode": "1234567890",
          "paymentApplicantName": "ヤマダショウテン",
          "paymentDate": "2022-11-30",
          "payingBankNo": "0001",
          "payingBankName": "みずほ銀行",
          "payingBranchNo": "001",
          "payingBranchName": "東京営業部",
          "payingDepositType": "01",
          "payingAccountNo": "1234567",
          "totalNumber": null,
          "totalAmount": null
      }
  },
  "api_schema": "DPFMPaymentRequisitionCreates",
  "accepter": [
      "Header"
  ],
  "order_id": null,
  "deleted": false
}
```
