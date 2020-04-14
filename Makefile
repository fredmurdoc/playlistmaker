test:
	cd tests && go test
	
build_for_raspbian:
	cd main && env GOOS=linux GOARCH=arm GOARM=5 go build && cd - && mv main/main  dist/raspbian/playlistmaker

build: clean
	cd main && go build && cd - && mv main/main dist/debian/playlistmaker

clean:
	rm -rf dist && mkdir -p dist/debian && mkdir -p dist/raspbian
