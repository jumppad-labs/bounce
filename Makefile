build:
	go build -o dist/bounce

install_local:
	sudo mv dist/bounce /usr/local/bin/