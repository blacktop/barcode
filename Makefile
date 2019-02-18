
.PHONY: snapshot
snapshot: ## Create a snapshot
	goreleaser --snapshot --rm-dist