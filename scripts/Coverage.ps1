& go test -coverprofile c.out .\... 
& go tool cover -html c.out
& go tool cover -func c.out
