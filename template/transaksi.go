package template

import (
	"exam-go-routine/helper"
	"exam-go-routine/model"
	"fmt"
)

func ViewTransaction() {
	var choice string
	helper.ClearScreen()
	fmt.Println("============================================================")
	fmt.Println("ID\tOrder Number")
	fmt.Println("============================================================")
	if len(model.Transactions) == 0 {
		fmt.Println("Data Kosong")
	} else {
		for _, v := range model.Transactions {
			fmt.Printf("%v\t%v\n", v.Id, v.TransactionNumber)
		}
		fmt.Println("\nMasukkan transaction number yang ingin dipilih:")
		fmt.Scanln(&choice)
		ch, err := model.SearchbyTrxNumber(&choice)
		if err == nil{
			fmt.Println("==================================================================")
			fmt.Println("No.Transaksi  : ", ch[0].Transaction.TransactionNumber)
			fmt.Println("Nama Pembeli  : ", ch[0].Transaction.Name)
			for _, v := range ch {
			fmt.Println("Id Product    : ", v.Transaction.Id)
			fmt.Println("Nama Product  : ", v.Item)
			fmt.Println("Harga Product : ", v.Price)
			fmt.Println("Quantity      : ", v.Quantity)
			fmt.Println("Total Harga   : ", v.Total)
			}
			fmt.Println("==================================================================")
			// for _, v := range ch{
			// 	fmt.Printf("%v\t%v\t%.2f\t%v\t%.2f\n",v.Id,v.Item,v.Price,v.Quantity,v.Total)
			// }
				
		}
	}

	helper.BackHandler()
	Menu()
}
