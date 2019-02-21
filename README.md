# London Public Transport Monitor

This is a test app to play around with golang.

It consumes https://api.tfl.gov.uk/ API. 

## Prerequisites
1. [Golang installed](https://golang.org/doc/install)
2. [Dep](https://github.com/golang/dep) (Golang dependency management tool) installed
3. Create `.env` file
```
API_KEY=<your api key for tfl api>
APP_ID=<your app id for tfl api>
```



## Development
1. Install dependencies
```
make install
```
2. Build the app
```
make build
```
3. Run the app
```
make start
```

### or using Docker
```
docker build -t go-tubes .
docker run go-tubes
```