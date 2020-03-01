
build:
	@echo "Build application image"
	@docker image build --no-cache \
		--tag kergeodeta/angles:$version \
		--tag kergeodeta/angles:latest \
		.

test:
	@echo "Build application testing image"
	@docker image build --no-cache \
		--tag kergeodeta/angles:${version}-test \
		--file Dockerfile.test \
		.

	@echo "Run unit tests"
	@docker container run --rm kergeodeta/angles:${version}-test

push:
	@echo "Publish docker images"
	@docker image push kergeodeta/angles:$version

deploy:
	@echo "Deploy the application"


clean:
	@echo "Remove docker images"
	@docker image rm -f kergeodeta/angles:$version|| true
	@docker image rm -f kergeodeta/angles:$version || true
	@docker image rm -f kergeodeta/angles:${version-test} || true