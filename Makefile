LS_LINT_IMAGE=lslintorg/ls-lint:1.11.2

ls-lint:
	docker run --rm -v ${PWD}:/data ${LS_LINT_IMAGE}