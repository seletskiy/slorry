build:
	GOPATH=$(GOPATH):`pwd` go build tracker.go
	strip ./tracker
	gzexe ./tracker && rm ./tracker~

clean:
	rm ./tracker
