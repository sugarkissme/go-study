package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

//请求体结构
type requestBody struct {
	Key string `json:"key"`
	Info string `json:"info"`
	UserId string `json:"userid"`
}
type responseBody struct {
	Code int `json:"code"`
	Text string `json:"text"`
	List []string `json:"list"`
	Url string `json:"url"`
}
func process(inputChan <-chan string,userid string) {
	for {
		input := <-inputChan
		if input == "EOF" {
			break
		}
		//构建请求体
		reqData := &requestBody{
			Key:    "40d56dcf5e1d4edc8a891eb824a11437",
			Info:   input,
			UserId: userid,
		}
		byteData, _ := json.Marshal(&reqData)
		req, err := http.NewRequest("POST", "http://www.tuling123.com/openapi/api", bytes.NewReader(byteData))
		req.Header.Set("Content-type", "application/json;charset=UTF-8")
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("net work error!")
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			var respData responseBody
			json.Unmarshal(body, &respData)
			fmt.Println("AI:" + respData.Text)
		}
		resp.Body.Close()
	}
}

func main(){
	var input string
	fmt.Println("Enter 'EOF' to shut down: ")
	//创建通道
	channel :=make(chan string)
	//main 结束时关闭通道
	defer close(channel)
	go process(channel,string(rand.Int63()))
	for  {
		//从命令行中读取输入
		fmt.Scanf("%s",&input)
		channel <-input
		//结束程序
		if input =="EOF"{
			fmt.Println("Bye!")
			break
		}
	}

}
