package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    var data struct {
        Items []struct {
            Name              string
            Count             int
            Is_required       bool
            Is_moderator_only bool
            Has_synonyms      bool
        }
    }

    r, _ := http.Get("https://api.stackexchange.com/2.2/tags?page=1&pagesize=100&order=desc&sort=popular&site=stackoverflow")
    defer r.Body.Close()

    dec := json.NewDecoder(r.Body)
    dec.Decode(&data)

    for _, item := range data.Items {
        fmt.Printf("%s = %d\n", item.Name, item.Count)
    }

}