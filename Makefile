install-hooks:
	ln -sf ../../scripts/pre-commit.sh .git/hooks/pre-commit
	chmod +x scripts/pre-commit.sh