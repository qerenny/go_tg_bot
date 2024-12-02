package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ClientData struct {
	Id      int      `json:"id"`
	Setting Settings `json:"setting"`
}

type Settings struct {
	Clients []Client `json:"clients"`
}

type Client struct {
	Id         *string `json:"id"`
	Flow       string  `json:"flow"`
	Email      *string `json:"email"`
	LimitIp    int     `json:"limitIp"`
	TotalGb    int     `json:"totalGb"`
	ExpiryTime int     `json:"expiryTime"`
	Enable     bool    `json:"enable"`
	TgId       *string `json:"tgId"`
	SubId      *string `json:"subId"`
	Reset      int     `json:"reset"`
}

type ReceiptData struct {
	Receipt Receipt `json:"receipt"`
}

type Receipt struct {
	Items []Item `json:"items"`
}

type Item struct {
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Amount      Amount `json:"amount"`
	VatCode     int    `json:"vat_code"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

func JsonPostClientUpdater(jsonPath, tgId, uuid, email, subId string, date int) (string, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var data ClientData

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)

	}

	for i := range data.Setting.Clients {
		data.Setting.Clients[i].TgId = &tgId
		data.Setting.Clients[i].Id = &uuid
		data.Setting.Clients[i].Email = &email
		data.Setting.Clients[i].SubId = &subId
		data.Setting.Clients[i].ExpiryTime = date
	}

	updatedJson, err := json.Marshal(data.Setting)
	if err != nil {
		panic(err)
	}

	return string(updatedJson), nil

}

func ReceiptCreator(jsonPath, description string, amount int) (string, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var data ReceiptData

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}

	for i := range data.Receipt.Items {
		data.Receipt.Items[i].Description = description
		data.Receipt.Items[i].Amount.Value = fmt.Sprintf("%.2f", float64(amount)/100)
	}

	updatedJson, err := json.Marshal(data.Receipt)
	if err != nil {
		panic(err)
	}

	return string(updatedJson), nil
}
