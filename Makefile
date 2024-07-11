.PHONY: ncommit run get dep run-dep

VERSION = v0.0.58

ncommit:
	cls
	git add .
	git commit -m "patch: tw-selectable-text"
	git push origin main 

run:	
	cls
	cd tw-selectable-text && go run .

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
