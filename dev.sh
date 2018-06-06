# clean the binaries
go clean

# tests
go test -v -race -cover -timeout 30s $(go list ./... | grep -v /vendor/)

# build and run the application
go build 
./run.sh
