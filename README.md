# goTunes
## Go package for interacting with Apple's iTunes Store API

### Install
```
go get https://github.com/metaprogrammed/gotunes.git
``` 


### Search iTunes
```
search_request := ItunesSearchRequest{Term: "up", Media: "movie", Entity: "movie"}
var search_results ItunesResponse = ItunesSearch(request) 

```

### iTunes ItemLookup
```
find_request := ItunesFindRequest{ItunesId: 284910350}
var find_results ItunesResponse = ItunesFind(request)

```

[iTunes Store API Reference](https://www.apple.com/itunes/affiliates/resources/documentation/itunes-store-web-service-search-api.html).