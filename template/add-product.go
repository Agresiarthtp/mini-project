package template

import (
	"bufio"
	"exam-go-routine/helper"
	"exam-go-routine/model"
	"fmt"
	"os"
	"reflect"
)

var tempAllTD []model.TransactionDetail

func AddOrder() {
	helper.ClearScreen()
	fmt.Println("----------------------------------------------")
	for _, v := range model.Products {
		fmt.Println(v.Id, v.Name, v.Price)
	}
	fmt.Println("----------------------------------------------")
	name := InputName()
	qty := InputQuantity()

	// proses product yg diinput
	var TransDetail model.TransactionDetail
	var Transaction model.Transaction

	// tempAllTD append dengan trandetail dari method Add
	TransDetail.Add(name, qty)
	tempAllTD = append(tempAllTD, TransDetail)

	var confirm string
	fmt.Println("Ingin menambah pembelian? yes/no")
	fmt.Scanln(&confirm)
	switch confirm {
	case "yes":
		helper.ClearScreen()
		AddOrder()
	case "no":
	}

	var nama string
	fmt.Print("Masukkan nama anda: ")
	fmt.Scanln(&nama)
	// tempAllTD append dengan transaction dari method AddTx
	// dari hasil append tempAllTD TransDetil dan tempAllTD Transaction,
	// akan disatukan kedalam tempAllTD
	Transaction.AddTx(tempAllTD, nama)

	for i := range tempAllTD {
		// mencari index dari tempAllTd untuk transaction
		tempAllTD[i].Transaction = Transaction
	}

	// kembali ke transaction details, untuk menambahkan produk baru
	model.TransactionDetails = append(model.TransactionDetails, tempAllTD...)
	fmt.Println("Ini hasil transaction detail ", model.TransactionDetails)
	model.Transactions = append(model.Transactions, Transaction)
	fmt.Println("INi hasil transaction", Transaction)

	tempAllTD = []model.TransactionDetail{}
	fmt.Println("Data berhasil diinput")
	helper.BackHandler()
	Menu()

}

func InputName() string {
	fmt.Print("Input nama produk yang ingin dibeli : ")
	// saat input nama produk bisa menggunakan spasi
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	if !ValidateName(&name) {
		fmt.Println("Nama produk tidak boleh kosong")
		InputName()
	}
	return name
}

func ValidateName(name *string) bool {
	var pro model.Product
	typeOf := reflect.TypeOf(pro)
	if typeOf.Field(1).Tag.Get("required") == "true" {
		if *name == "" {
			return false
		}
	}
	return true
}

func InputQuantity() int {
	var quantity int
	fmt.Print("Input jumlah barang yang ingin dibeli : ")
	fmt.Scanln(&quantity)
	return quantity
}
