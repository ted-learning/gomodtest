package main

import (
	json "encoding/json"
	"fmt"
)

type Order struct {
	ID         string      `json:"id"`
	TotalPrice float64     `json:"total_price"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
}

func parseNLP() {
	jsonStr := `{
					"data": [
						{
							"synonym":"",
							"weight":"0.6",
							"word": "真丝",
							"tag":"材质"
						},
						{
							"synonym":"",
							"weight":"0.8",
							"word": "韩都衣舍",
							"tag":"品牌"
						},
						{
							"synonym":"连身裙;联衣裙",
							"weight":"1.0",
							"word": "连衣裙",
							"tag":"品类"
						}
					]
				}`

	bytes := []byte(jsonStr)
	s := string(bytes)
	fmt.Println(s == jsonStr)

	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Tag     string `json:"tag"`
			Weight  string `json:"weight"`
			Word    string `json:"word"`
		} `json:"data"`
	}{}

	err := json.Unmarshal(bytes, &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m.Data)
	fmt.Println(m.Data[1].Tag)
}
func main() {
	order := Order{
		ID:         "1",
		TotalPrice: 100,
		Items: []OrderItem{
			{
				ID:    "o1",
				Price: 50,
				Name:  "name1",
			},
			{
				ID:          "o2",
				Price:       50,
				Name:        "name2",
				Description: "des...",
			},
		},
	}

	marshal, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", marshal)

	jsonstr := "{\"id\":\"1\",\"total_price\":100,\"items\":[{\"id\":\"o1\",\"name\":\"name1\",\"price\":50},{\"id\":\"o2\",\"name\":\"name2\",\"price\":50,\"description\":\"des...\"}]}"
	unmarshal := Order{}
	err = json.Unmarshal([]byte(jsonstr), &unmarshal)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", unmarshal)

	fmt.Println("------------------")
	parseNLP()
}
