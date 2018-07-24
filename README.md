osuapi-go
=========

A Go implementation of the osu! api created by Wyatt Calandro  
[![GoDoc](https://godoc.org/github.com/wcalandro/osuapi-go?status.svg)](http://godoc.org/github.com/wcalandro/osuapi-go)
# Using the module

Install using `go get`
```shell
$ go get github.com/wcalandro/osuapi-go
```
In your code
```go
import "github.com/wcalandro/osuapi-go"
```

# Creating an instance of the API
```go
api := osuapi.NewAPI("Your API Key from https://osu.ppy.sh/p/api")
```

# Using `osuapi.M`
`osuapi.M` is a shortcut for `map[string]string`, and is used when passing options to the osu! api endpoint. The API key is automatically added to each query, however additional options such as beatmap IDs and usernames must be provided.
## Example
```go
// Adds &u=Arm1stice to the query of the API request
api.GetUser(osuapi.M{"u":"Arm1stice"})
```

# Caveats
This package doesn't implement the following API endpoints at this time:
- `/api/get_match`
- `/api/get_replay`  

If you would like to have these implemente in this package, feel free to submit a Pull Request
