.PHONY: ncommit run get

VERSION = v0.0.47

ncommit:
	cls
	git add .
	git commit -m "ncommit"
	git push origin main 

run:
	cls
	cd tw-spacing && go run .
	cls
	git add .
	git commit -m "ncommit"
	git push origin main 

get:
	go get github.com/nelsonsaake/go@$(VERSION)
