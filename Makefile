build: ## build ntpgo
	go build -o ntpgo cmd/root.go
test_install: ## test install
	sudo ./ntpgo install --server 'cn.pool.ntp.org'
test_restart: ## test start
	sudo ./ntpgo restart
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

