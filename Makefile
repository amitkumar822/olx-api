# phony actual go ko bata hai ki ye command hai for example: build, run (.PHONY: build run)
.PHONY: build run

# ye build command go ko binary banata hai jo default me bin folder me banta hai
build:
	@go build -o bin/api ./cmd/api

# run command acutal go ko run karta hai and run: build (run command ke baad build command iss liye hai ki direct run karte hi saath me build + run ho jaye)
run: build
	@./bin/api