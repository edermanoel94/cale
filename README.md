Rest GO - Helpful library  for Rest API
================================

A package that provide many helpful methods for working with rest api.

And if u are tired of written, this:

```go

func SomeHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Add("Content-Type", "application/json")

    product := &product{"Smart TV", 50.00}
    
    bytes, err := json.Marshal(product)
    
    if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
        message := fmt.Sprintf("{\"message\": \"%s\"}", err.Error())
    	w.Write([]byte(message))
	    return
    }
    
    w.WriteHeader(http.StatusOk)
    w.Write(bytes)
}
```

Get started:

  * Install rest GO with [one line of code](#installation), or [update it with another](#staying-up-to-date)


[`rest`](http://godoc.org/github.com/edermanoel94/rest-go "API documentation") package
-------------------------------------------------------------------------------------------

The `rest` package provides some helpful methods that allow you to write better rest api in GO.

  * Allows for very readable code

See it in action:

```go
package yours

import (
    "encoding/json"
    "github.com/edermanoel94/rest-go"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

func SomeHandler(w http.ResponseWriter, r *http.Request) {
	
    product := &product{"Smart TV", 50.00}

    bytes, err := json.Marshal(product)
    
    if err != nil {
        // JsonWithError let u write a custom error message on json formatted
        rest.JsonWithError(w, err, http.StatusInternalServerError)
        return
    }

    // Json use if u want customize your json.Marshal in someway
    rest.Json(w, bytes, http.StatusOK)
}
```

Simplified

```go
package yours

import (
    "encoding/json"
    "github.com/edermanoel94/rest-go"
    "net/http"
)

type product struct {
    Name  string `json:"name"`
    Price float32 `json:"price"`
}

func SomeHandler(w http.ResponseWriter, r *http.Request) {
	
    product := &product{"Smart TV", 50.00}

    // JsonMarshalled marshall the struct for u and respond json.
    rest.JsonMarshalled(w, product, http.StatusOK)
}
```

Installation
============

To install, use `go get`:

    go get github.com/edermanoel94/rest-go


Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

------

License
=======

This project is licensed under the terms of the MIT license.
