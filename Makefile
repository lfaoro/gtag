APP ?= "./cmd/gtag"

install:
	go install -mod vendor "$(APP)"

build:
	@go build -mod vendor "$(APP)"

test:
	@go test -mod vendor -cover "$(APP)"

release:
	cd $(APP) && \
	goreleaser release --rm-dist --config=../../.goreleaser.yml

lc:
	@find . -type f -name "*.go" -not -path "*/vendor/*" -not -path "./docs/*" \
	 | xargs wc -l | sort