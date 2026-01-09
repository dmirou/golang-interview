.PHONY: frun
frun:
	go run cmd/file/main.go -in ./files/in.txt -out ./files/out.txt
.PHONY: crun
crun:
	go run cmd/console/main.go
.PHONY: bench
bench:
	go test -bench=. -benchmem ./...
