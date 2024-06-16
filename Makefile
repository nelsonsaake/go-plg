.PHONY: ncommit run get

VERSION = v0.0.46

ncommit:
	cls
	git add .
	git commit -m "ncommit"
	git push origin main 

run:
	cls
	cd xls && go run .

get:
	go get github.com/nelsonsaake/go@$(VERSION)
