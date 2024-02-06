OUTPUT_DIR=./output


all: goexpl


goexpl:
	go build -o $(OUTPUT_DIR)/goexpl main.go

clean:
	rm output/*
