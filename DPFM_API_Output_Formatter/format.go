package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-production-order-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
)

func SetToPdfData(
	headerData *dpfm_api_input_reader.Header,
	inputItems []dpfm_api_input_reader.Items,
) *Header {
	pm := &Header{}

	var items []Items
	for _, dataItems := range inputItems {
		if headerData.ProductionOrder == dataItems.ProductionOrder {
			items = append(items, Items{
				OrderID:						dataItems.OrderID,
				OrderItem:						dataItems.OrderItem,
				Product:						dataItems.Product,
				ProductSpecification:			dataItems.ProductSpecification,
				SizeOrDimensionText:			dataItems.SizeOrDimensionText,
				OrderItemText:        			dataItems.OrderItemText,
				OrderQuantityInBaseUnit:		dataItems.OrderQuantityInBaseUnit,
				OrderQuantityInDeliveryUnit:	dataItems.OrderQuantityInDeliveryUnit,
				BaseUnit:						dataItems.BaseUnit,
				DeliveryUnit:        			dataItems.DeliveryUnit,
				RequestedDeliveryDate:        	dataItems.RequestedDeliveryDate,
				RequestedDeliveryTime:        	dataItems.RequestedDeliveryTime,
				NetAmount:						dataItems.NetAmount,
				TaxAmount:						dataItems.TaxAmount,
				GrossAmount:					dataItems.GrossAmount,
			})
		}
	}

	pm = &Header{
				OrderID:   					headerData.OrderID,
				OrderStatus:   				headerData.OrderStatus,
				OrderDate:   				headerData.OrderDate,
				OrderType:   				headerData.OrderType,
				Buyer: 						headerData.Buyer,
				BuyerName: 					headerData.BuyerName,
				Seller: 					headerData.Seller,
				SellerName:              	headerData.SellerName,
				RequestedDeliveryDate: 		headerData.RequestedDeliveryDate,
				RequestedDeliveryTime: 		headerData.RequestedDeliveryTime,
				TotalGrossAmount: 			headerData.TotalGrossAmount,
				Contract: 					headerData.Contract,
				ContractItem: 				headerData.ContractItem,
				Project: 					headerData.Project,
				WBSElement: 				headerData.WBSElement,
				Incoterms: 					headerData.Incoterms,
				PricingDate: 				headerData.PricingDate,
				PaymentTerms: 				headerData.PaymentTerms,
				PaymentTermsName: 			headerData.PaymentTermsName,
				PaymentMethod: 				headerData.PaymentMethod,
				TransactionCurrency: 		headerData.TransactionCurrency,
				Items:     					items,
	}

	return pm
}
