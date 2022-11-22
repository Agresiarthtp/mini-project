package controller

import "exam-go-routine/interfaces"

func InserProductCheck(product interfaces.ProductInterface, name string, quantity int) {
	var datas = map[string]any{
		"name": name,
		"quantity" : quantity,
	}
	var datasSlice []map[string]any
	datasSlice = append(datasSlice, datas)
	
	product.Add(datasSlice)
}
