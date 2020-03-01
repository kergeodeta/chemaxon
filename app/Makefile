
build:
	@echo "Build application image"
	@docker image build --no-cache \
		--tag kergeodeta/angles:v1.0.0 \
		.

test:
	@echo "Build application testing image"
	@docker image build --no-cache \
		--tag kergeodeta/angles:v1.0.0-test \
		--file Dockerfile.test \
		.

	@echo "Run unit tests"
	@docker container run --rm kergeodeta/angles:v1.0.0-test

push:
	@echo "Publish docker images"
	@docker image push kergeodeta/angles:v1.0.0

deploy:
	@echo "Deploy the application"

clean:
	@echo "Remove docker images"
	@docker image rm -f kergeodeta/angles:v1.0.0 || true
	@docker image rm -f kergeodeta/angles:v1.0.0-test || true