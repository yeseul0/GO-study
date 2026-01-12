package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// UnMarshal (json->map)
func ex1() {
	doc := `
	{
		"name": "maria",
		"age": 10
	}
	`

	var data map[string]interface{}    //JSON 데이터를 저장할 공간을 맵으로 선언
	json.Unmarshal([]byte(doc), &data) //doc -> byte 슬라이스로 변환해서 넣음, &data 이 주소map에 doc 데이터가 채워질것!
	fmt.Println(data["name"], data["age"])
}

// MArshal (map->json)
func ex2() {
	data := make(map[string]interface{}) //빈 맵 공간 생성

	data["name"] = "maria"
	data["age"] = 10

	doc, _ := json.Marshal(data)                  //data 맵을 json 문서로 변환 ([]byte 형식임)
	doc2, _ := json.MarshalIndent(data, "", "  ") //2번째 인자는 prefix, 3번째 인자는 들여쓸 문자.
	fmt.Println(string(doc))                      //[]byte 형식이라서 string으로 변환!
	fmt.Println(string(doc2))
}

type Author struct {
	Name  string
	Email string
}
type Comment struct {
	Id      uint64
	Author  Author
	Content string
}

type Article struct {
	Id         uint64
	Title      string `json:"title"` //json 파일화 했을 때 키가 소문자로 시작하게 하려고 ㅎㅎ!!
	Author     Author
	Content    string
	Recommends []string
	Comments   []Comment
}

// 구조체 활용
func ex3() {
	doc := `
	[{
			"Id":1.
			"Title":"Hello, World!!",
			"Author": {
				"Name" : "Maria", 
				"Email" : "maria@example.com
			},
			"Content":"Hello~~"
			"Recommends" : [
				"Jogn",
				"Andrew"
			],
			"Comments" : [{
				"id":1",
				"Author": {
					"Name": "Andrew",
					"Email": "andrew@hello.com"
				},
				"Content":"Hello Maria"
			}]	
	}]
	`
	//json -> 구조체
	var data []Article //Json 문서의 데이터를 저장할 구조체 슬라이스
	json.Unmarshal([]byte(doc), &data)
	fmt.Println(data)
}

// 구조체 -> json Marshal, file 저장
func ex4() {

	data := make([]Article, 1) //Article 구조체 슬라이스(길이1)

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) //data(구조체 슬라이스) -> json

	fmt.Println(string(doc))
	// 참고!! 구조체 필드의 첫 글자를 소문자로 하면, json 문서에 해당 필드는 빠진다!
	err := os.WriteFile("./articles.json", doc, os.FileMode(0644)) //articles.json파일에 json 문서 저장
	if err != nil {
		fmt.Println(err)
		return
	}
}

// json file 읽어오기
func main() {
	b, err := os.ReadFile("./articles.json") //파일 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Article       //json 문서내용 담을 구조체 슬라이스
	json.Unmarshal(b, &data) //b 변환하여 data가 가리키는 슬라이스에 저장
	fmt.Println(data)
}
