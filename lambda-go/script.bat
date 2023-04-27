go get -v all
SET GOOS=linux
SET GOARCH=amd64
go build -o build/main cmd/main.go