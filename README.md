# Go OpenAI
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/sashabaranov/go-openai)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-openai)](https://goreportcard.com/report/github.com/sashabaranov/go-openai)


This library provides Go clients for [OpenAI API](https://platform.openai.com/). We support:


Installation:
```
go get github.com/sashabaranov/go-openai
```


ChatGPT example usage:

```go
package main

import (
	"context"
	"fmt"
	 "github.com/sashabaranov/simpleGPT"
)

func main() {
	client := simpleGPT.NewClient("your token")
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



Other examples:

<details>
<summary>GPT-3 completion</summary>

```go
package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/simpleGPT"
)

func main() {
	client := simpleGPT.NewClient("your token")
    client.Ctx = context.Background()

	resp,err  := client.SendStreamMsg("How to be rich")
	if err != nil {
    		return
    }
    fmt.Println(resp)
	resp,err = client.SendStreamMsg("what can I do to archive that?")
    if err != nil {
    		return
    }
    fmt.Println(resp)
}
```
</details>
