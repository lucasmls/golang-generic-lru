test:
	@ echo
	@ echo "Starting running tests..."
	@ echo
	@ go test -v -cover ./...

%:
	@: