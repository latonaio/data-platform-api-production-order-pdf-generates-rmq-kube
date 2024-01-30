package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
	ProductionOrderPdf	string      `json:"production_order_pdf"`
}

type Message struct {
	Header     []Header   `json:"Header"`
}

type Header struct {
	OrderID               int     `json:"OrderID"`
	OrderStatus           string  `json:"OrderStatus"`
	OrderDate             string  `json:"OrderDate"`
	OrderType             string  `json:"OrderType"`
	Buyer                 int     `json:"Buyer"`
	BuyerName             string  `json:"BuyerName"`
	Seller                int     `json:"Seller"`
	SellerName            string  `json:"SellerName"`
	RequestedDeliveryDate string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime string  `json:"RequestedDeliveryTime"`
	TotalGrossAmount      float32 `json:"TotalGrossAmount"`
	Contract              *int    `json:"Contract"`
	ContractItem          *int    `json:"ContractItem"`
	Project               *int    `json:"Project"`
	WBSElement            *int    `json:"WBSElement"`
	Incoterms			  *string `json:"Incoterms"`
	PricingDate			  string  `json:"PricingDate"`
	PaymentTerms          string  `json:"PaymentTerms"`
	PaymentTermsName      string  `json:"PaymentTermsName"`
	PaymentMethod         string  `json:"PaymentMethod"`
	TransactionCurrency   string  `json:"TransactionCurrency"`
	Items				  []Items `json:"Items"`
}

type Items struct {
	OrderID               		int     `json:"OrderID"`
	OrderItem                   int     `json:"OrderItem"`
//	OrderStatus                 string  `json:"OrderStatus"`
//	OrderItemCategory           string  `json:"OrderItemCategory"`
	Product                     string  `json:"Product"`
	ProductSpecification		*string `json:"ProductSpecification"`
	SizeOrDimensionText			*string `json:"SizeOrDimensionText"`
	OrderItemText               string  `json:"OrderItemText"`
//	OrderItemTextByBuyer        string  `json:"OrderItemTextByBuyer"`
//	OrderItemTextBySeller       string  `json:"OrderItemTextBySeller"`
	OrderQuantityInBaseUnit     float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
	BaseUnit                    string  `json:"BaseUnit"`
	DeliveryUnit                string  `json:"DeliveryUnit"`
	RequestedDeliveryDate       string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       string  `json:"RequestedDeliveryTime"`	
	NetAmount                   float32 `json:"NetAmount"`
	TaxAmount                   float32 `json:"TaxAmount"`
	GrossAmount                 float32 `json:"GrossAmount"`
}
