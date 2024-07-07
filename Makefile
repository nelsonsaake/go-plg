.PHONY: ncommit run get dep run-dep

VERSION = v0.0.51

ncommit:
	cls
	git add .
	git commit -m "ncommit"
	git push origin main 

run:	
	cls
	cd font-family && go run .

dep:
	cls
	git add .
	git commit -m "ncommit"
	git push origin main 

run-dep:
	make run
	make dep

get:
	go get github.com/nelsonsaake/go@$(VERSION)
