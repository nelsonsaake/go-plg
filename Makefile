.PHONY: ncommit run get dep run-dep

VERSION = v0.0.75

ncommit:
	cls
	git add .
	git commit -m "wip: ws-files"
	git push origin main 

run:	
	cls
	cd ws-files && go run .

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
