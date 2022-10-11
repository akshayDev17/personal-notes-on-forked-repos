package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	urlpkg "net/url"
	"reflect"
)

const APIBase string = "https://api.sandbox.paypal.com"

type parent struct {
	Value int64
}

type child struct {
	parent
	Details string
}

func (p *parent) SomeMethod() {
	fmt.Println("Hi, details are:", p)
}

func main() {

	// testing promotion.
	var sampleParent parent
	sampleParent = parent{
		Value: 37,
	}
	var sampleChild child
	sampleChild = child{
		parent:  sampleParent,
		Details: "Cristiano Ronaldo",
	}

	sampleChild.SomeMethod()

	// type assertion
	var greeting interface{} = "hello world"
	greetingStr, assertion_err := greeting.(string)
	fmt.Println("Assertion err =", assertion_err)
	fmt.Println("Greeting \"", greetingStr, "\"of type", reflect.TypeOf(greetingStr))

	// // Create a client instance
	// c, err := paypal.NewClient("clientID", "secretID", APIBase)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// _, err = c.GetAccessToken(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	type someStruct struct {
		Key   int64  `json:"key"`
		Value string `json:"value"`
	}

	var x = someStruct{
		Key:   1,
		Value: "hey",
	}
	data, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
	}
	var body io.Reader
	body = bytes.NewBuffer(data)
	fmt.Println(body)
	url := fmt.Sprintf("%s%s", APIBase, "/v2/checkout/orders")
	_, err = urlpkg.Parse(url)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(u) // https://api.sandbox.paypal.com/v2/checkout/orders
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = io.NopCloser(body)
		fmt.Println("error occurred")
		fmt.Println(ok)
	} else {
		fmt.Println(rc)
	}
}
