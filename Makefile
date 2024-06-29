install-dependencies:
	go install github.com/a-h/templ/cmd/templ@latest

generate:
	templ generate
	go generate ./...

build-docker:
	docker build -t poccomaxa/installable-bot-concept:latest .

upload-docker:
	docker push poccomaxa/installable-bot-concept:latest

test-docker:
	docker run -it --rm poccomaxa/installable-bot-concept:latest
