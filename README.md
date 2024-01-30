# data-platform-api-production-order-pdf-generates-rmq-kube

data-platform-api-production-order-pdf-generates-rmq-kube は、周辺業務システム　を データ連携基盤 と統合することを目的に、API で製造指示書のPDF帳票データを出力するマイクロサービスです。  
https://xxx.xxx.io/api/API_FUNC_PRODUCTION_ORDER_PDF_SRV/generates/

## 動作環境

data-platform-api-production-order-pdf-generates-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
data-platform-api-production-order-pdf-generates-rmq-kube が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/API_FUNC_PRODUCTION_ORDER_PDF_SRV/generates/

## 本レポジトリ に 含まれる API名
data-platform-api-production-order-pdf-generates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_ProductionOrder（製造指示書 - 製造指示書データ）

## API への 値入力条件 の 初期値
data-platform-api-production-order-pdf-generates-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に出力したいPDF帳票の種別（＝APIの種別）を入力し、指定することができます。   

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ProductionOrder" が指定されています。    
  
```
	"api_schema": "DPFMProductionOrderPDFGenerates",
	"accepter": ["ProductionOrder"],
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncPDFCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	var response interface{}
	// SQL処理
	response = c.readSqlProcess(nil, &mtx, input, output, accepter, &errs, log)

	return response, nil
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は 製造指示書 の PDFデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ProductionOrder" ～ "IsMarkedForDeletion" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
XXX
```
