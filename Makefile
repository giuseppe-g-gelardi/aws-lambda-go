.PHONY: all clean

build: 
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lambdas/Get/bootstrap lambdas/Get/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lambdas/Goodbye/bootstrap lambdas/Goodbye/main.go 
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lambdas/Hello/bootstrap lambdas/Hello/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o lambdas/Numsdouble/bootstrap lambdas/Numsdouble/main.go

clean:
	@find lambdas -name "bootstrap" -type f -exec rm -f {} +

deploy: clean build
	cdk deploy --all
