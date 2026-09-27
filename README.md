<h1 align="center">JEV Go</h1>

<p align="center">Go client for the JEV API.</p>

<br>

## Installation

```bash
go get github.com/alexdenkk/jev-go
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/alexdenkk/jev-go/jev"
)

func main() {
	api := jev.NewJevAPI(
		"<BASE_URL>/v1/systemone",
		"jev-latest",
		"<API_KEY>",
	)

	response, err := api.SystemOne(jev.Request{
		State: "The customer says: 'I was charged twice for my monthly subscription and I need one of the duplicate charges refunded.'",

		Questions: map[string]jev.Question{
			"urgent": jev.Noul(
				"Is this request valid?",
			),

			"department": jev.Choice(
				"Which team should handle this customer request?",
				map[string]string{
					"Billing":   "Billing team",
					"Sales":     "Sales team",
					"Technical": "Technical team",
					"Other":     "Any other team",
				},
			),

			"priority": jev.Score(
				"Urgency of this request?",
				[]string{
					"No urgency",
					"Low",
					"Moderate",
					"High",
					"Critical",
				},
			),
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
```

## Creating the client

Create a client with `NewJevAPI`:

```go
api := jev.NewJevAPI(
	"<BASE_URL>/v1/systemone",
	"jev-latest",
	"<API_KEY>",
)
```

The constructor accepts:

| Argument | Description                     |
| -------- | ------------------------------- |
| `url`    | JEV API endpoint                |
| `model`  | Model to use for the request    |
| `apiKey` | API key used for authentication |

The SDK uses the API key as a Bearer token and sends requests as JSON.

## Questions

### Noul

Use `Noul` for a boolean-style question.

```go
"valid": jev.Noul(
	"Is this request valid?",
)
```

### Choice

Use `Choice` when the answer should be selected from a predefined set of options.

```go
"department": jev.Choice(
	"Which team should handle this customer request?",
	map[string]string{
		"Billing":   "Billing team",
		"Sales":     "Sales team",
		"Technical": "Technical team",
		"Other":     "Any other team",
	},
)
```

### Score

Use `Score` when the answer should be selected from an ordered list of levels.

```go
"priority": jev.Score(
	"Urgency of this request?",
	[]string{
		"No urgency",
		"Low",
		"Moderate",
		"High",
		"Critical",
	},
)
```

## Error handling

`SystemOne` returns both the decoded response and an error:

```go
response, err := api.SystemOne(request)
if err != nil {
	return err
}
```

The SDK returns an error when:

* the request cannot be serialized;
* the HTTP request cannot be created or sent;
* the response body cannot be read;
* the API returns a non-2xx status code;
* the response cannot be decoded.

## API

### `NewJevAPI`

```go
func NewJevAPI(url, model, apiKey string) *JevAPI
```

Creates a JEV API client.

### `SystemOne`

```go
func (jev *JevAPI) SystemOne(request Request) (*Response, error)
```

Sends a request to the configured SystemOne endpoint and returns the structured JEV response.

The configured model is applied to the request automatically.

## Authentication

The SDK authenticates requests using the configured API key:

```http
Authorization: Bearer <API_KEY>
```
