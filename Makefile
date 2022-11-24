fibonacci-art:
	echo "Fibonacci Art"

build: 
	go build -o bin/fibonacci-art main.go

clean:
	rm -rf *.png *.jpg bin