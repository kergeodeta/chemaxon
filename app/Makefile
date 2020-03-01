
build:
	@echo "Build application image"
	@docker image build --no-cache \
		--tag kergeodeta/angles:v1.0.0 \
		--tag kergeodeta/angles:latest \
		.

	@echo "Build application testing image"
	@docker image build --no-cache \
		--tag kergeodeta/angles:v1.0.0-test \
		--file Dockerfile.test \
		.
test:
	@echo "Run unit tests"
	@docker container run --rm kergeodeta/angles:v1.0.0-test
