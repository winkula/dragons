go test -coverprofile $env:TEMP\cover.out .\... 
go tool cover -html $env:TEMP\cover.out
go tool cover -func $env:TEMP\cover.out
