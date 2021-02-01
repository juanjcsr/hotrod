package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Event struct {
	Text string `json:"text"`
}

func HandleRequest(ctx context.Context, e Event) (string, error) {
	return fmt.Sprintf("Hello %s", e.Text), nil
}

func HandleApiGateway(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	p := e.QueryStringParameters["text"]
	t := fmt.Sprintf("Hello %s", p)
	re := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	re.Body = t
	return re, nil
}

func main() {
	fmt.Println("HOLA")
	//route.Hello()
	//lambda.Start(HandleApiGateway)
	fmt.Println(A11ysor("Wizeline"))
	fmt.Println(A11ysor("Wizeline") == "W6e")
	fmt.Println(A11ysor("Accessibility") == "A11y")
	fmt.Println(A11ysor("are") == "are")

	fmt.Println(A11ysorSentences("We are Wizeline"))
	fmt.Println(A11ysorSentences("We become what we want to be") == "We b4e w2t we w2t to be")
	fmt.Println(A11ysorSentences("We become what we want to be"))
}

func A11ysor(s string) string {
	// Access -> A4s
	// Wizeline -> W6e

	splited := strings.Split(s, "")
	l := len(splited)
	if l > 3 {
		n := l - 2
		return fmt.Sprintf("%s%d%s", splited[0], n, splited[l-1])
	}
	return s
}

func A11ysorSentences(s string) string {
	splitted := strings.Split(s, " ") // ["We", "are", "Wizeline"]
	sbuilder := strings.Builder{}
	dict := make(map[string]bool)
	for i, v := range splitted {
		converted := A11ysor(v)
		_, ok := dict[converted]
		if ok {
			sbuilder.WriteString(v)
		} else {
			sbuilder.WriteString(converted)
			dict[converted] = true
		}

		if i < len(splitted)-1 {
			sbuilder.WriteString(" ")
		}

	}
	return sbuilder.String()
}
