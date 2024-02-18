.PHONY: git git-reset

git:
	@echo "git for dev..."

git-reset:
	git fetch origin main && git reset --hard origin/main