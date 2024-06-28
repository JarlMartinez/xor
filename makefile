clean:
	rm bin/xor

build: clean
	go build -o bin/xor main.go