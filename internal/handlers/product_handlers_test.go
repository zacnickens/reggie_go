// productHandlers.go

package handlers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/square/go-square"
)

const squareAPIDomain = "https://connect.squareupservices.com"
const squareAPIToken string = "<your-square-api-token>"

func getSquareProducts() ([]*SquareProduct, error) {
    client := &square.Client{
        Token:   squareAPIToken,
        Domain:  squareAPIDomain,
        Version: "2020-09-16",
    }

    var response *square.ListCatalogObjectsResponse
    err := client.ListCatalogObjects(&response)
    if err != nil {
        return nil, err
    }

    squareProducts := make([]*SquareProduct, len(response.Objects))
    for i, obj := range response.Objects {
        p := &SquareProduct{}
        json.Unmarshal(obj.Data, p)
        squareProducts[i] = p
    }
    return squareProducts, nil
}

func getSquareProduct(id string) (*SquareProduct, error) {
    client := &square.Client{
        Token:   squareAPIToken,
        Domain:  squareAPIDomain,
        Version: "2020-09-16",
    }

    var response *square.RetrieveCatalogObjectResponse
    err := client.RetrieveCatalogObject(&response)
    if err != nil {
        return nil, err
    }

    p := &SquareProduct{}
    json.Unmarshal(response.Object.Data, p)
    return p, nil
}