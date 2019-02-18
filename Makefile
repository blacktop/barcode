.PHONY: snapshot
snapshot: ## Create a snapshot
	goreleaser --snapshot --rm-dist

.PHONY: bump
bump: ## Incriment version patch number
	scripts/bump_version.sh -p $(shell cat .version) > .version
	git commit -am "bumping version to $(shell cat .version)"
	git push

.PHONY: release
release: bump ## Create a new release from the VERSION
	scripts/make_release.sh $(shell cat .version)
	goreleaser --rm-dist

.PHONY: destroy
destroy: ## Remove release from the VERSION
	rm -rf dist
	git tag -d ${VERSION}
	git push origin :refs/tags/${VERSION}

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := all