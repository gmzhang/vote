all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o vote.linux server.go

clean:
	rm -f vote.*
