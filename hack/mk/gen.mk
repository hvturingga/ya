gen:
	@ent generate ./ent/schema
	@for dir in hack/gen/*/main.go; do \
		go run $$dir; \
	done