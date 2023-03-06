# Go OpenAI
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/stone-stones/simpleGPT)
[![Go Report Card](https://goreportcard.com/badge/github.com/stone-stones/simpleGPT)](https://goreportcard.com/report/github.com/stone-stones/simpleGPT)


This library provides Go clients for [OpenAI API](https://platform.openai.com/docs/guides/chat/introduction).


Installation:
```
go get github.com/stone-stones/simpleGPT
```

Run Test:
```
go test -gcflags all=-l
```


ChatGPT example usage:

```go
package main 
import (
	"fmt"
	simpleGPT "github.com/stone-stones/simpleGPT"
)

func main() {
	client, err := simpleGPT.NewClient("your_token")
	if err != nil {
		return
	}
	res, err := client.SendMsg("Who is the American president")

	if err != nil {
		return
	}
	fmt.Println(res)

	//you can go on ask more question with pre questions
	res, err = client.SendMsg("tell me about him")
	if err != nil {
		return
	}
	fmt.Println(res)

}

```




<details>
<summary>chatGPT stream message</summary>

```go
package main

import (
	"fmt"
	simpleGPT "github.com/stone-stones/simpleGPT"
)

func main() {
	client, err := simpleGPT.NewClient("your_token")
	if err != nil {
		return
	}
	res, err := client.SendStreamMsg("Who is the American president")

	if err != nil {
		return
	}
	fmt.Println(res)

	//you can go on ask more question with pre questions
	res, err = client.SendStreamMsg("tell me about him")
	if err != nil {
		return
	}
	fmt.Println(res)

}
```
</details>


<details>
<summary>chatGPT with config</summary>

```go
package main

import (
	"context"
	"fmt"
	simpleGPT "github.com/stone-stones/simpleGPT"
)

func main() {
	err := simpleGPT.LoadConfigFromFile("path_to_your_config", simpleGPT.GetLogger())
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := simpleGPT.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	client.Ctx = context.Background()

	resp, err := client.SendMsg("How to be rich")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	resp, err = client.SendMsg("what can I do to archive that?")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
```
</details>


