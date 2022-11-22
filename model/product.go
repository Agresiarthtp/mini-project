package model

import (
	"errors"
	"math/rand"
	"strconv"
)

// struct product
type Product struct {
	Id    int
	Name  string  `required:"true"`
	Price float64 `required:"price"`
}


// Data yg disediakan
var Products []Product = []Product{
	{
		Id:    1,
		Name:  "Kaos Phincon",
		Price: 150000,
	},
	{
		Id:    2,
		Name:  "Lanyard Phincon",
		Price: 20000,
	},
	{
		Id:    3,
		Name:  "Tumbler Phincon",
		Price: 80000,
	},
}

type TransactionDetail struct {
	Id       int
	Item     string
	Price    float64
	Quantity int
	Total    float64
	Transaction
}

type Transaction struct {
	Id int
	TransactionNumber string
	Name string
	Quantity int
	Total float64
}

var Transactions []Transaction 
var TransactionDetails []TransactionDetail


// Get Id dari data yg diinput
func GetLastId() int{
	if Transactions == nil {
		return 0
	} else {
		var tempId int
		for _, v := range Transactions {
			if tempId < v.Id {
				tempId = v.Id
			}
		}
		return tempId 
	}
}

func GetLastIdDetail() int{
	if TransactionDetails == nil {
		return 0
	} else {
		var tempID int
		for _, v := range TransactionDetails {
			if tempID < v.Id {
				tempID = v.Id
			}
		}
		return tempID
	}
}

// Search byId menggunakan GoRoutine
// func (pro *TransactionDetail) SearchbyName() {
// 	// membuat channel dari tipe data yang ada distruct

// }

//func (pro *TransactionDetail)  Add(datas []map[string]interface{}){
func (pro *TransactionDetail)  Add(name string, quan int){
	pro.Id = GetLastId() + 1
	pro.Item = name
	pro.Price = GetPrice(name)
	pro.Quantity = quan
	pro.Total = pro.Price * float64(quan) 

	TransactionDetails = append(TransactionDetails, *pro)
}

func GetPrice(nama string) float64 {
	for _, v := range Products {
		if v.Name == nama {
			return v.Price
		}
	}
	return 0
}

func GenerateTrxNumber() string {
	var tempTrxNumber string = "TRXPHC-"
	var tempTrxRandom int
	min := 100
	max := 800
	// Intn , tidak ada angka minus/negatif
	tempTrxRandom = rand.Intn(max-min+10) + 1

	return tempTrxNumber + strconv.Itoa(tempTrxRandom)
}

func (pro *Transaction) AddTx(trxD []TransactionDetail, name string){
	pro.Id = GetLastId() + 1
	pro.TransactionNumber = GenerateTrxNumber()
	pro.Name = name
	for _, v := range trxD {
		pro.Quantity += v.Quantity
		pro.Total += v.Total
	}
}

func SearchbyTrxNumber(pro *string) ([]TransactionDetail, error) {
	var tempAllTD []TransactionDetail
	for _, v := range TransactionDetails {
		if v.TransactionNumber == *pro {
			tempAllTD = append(tempAllTD, v)
			return tempAllTD, nil
		}
	}
	return tempAllTD, errors.New("data tidak ada")
}
	
