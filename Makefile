.PHONY: ncommit run get

VERSION = v0.0.41

ncommit:
	cls
	git add .
	git commit -m "ncommit"
	git push origin main 

run:
	cls
	go run .

get:
	go get github.com/nelsonsaake/go@$(VERSION)
