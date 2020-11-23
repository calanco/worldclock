.SHELLFLAGS: -ec

# Run WorldClock to check some capital times
.PHONY: run
run:
	@go run main.go --capitals Europe/Berlin,Europe/Rome,America/New_York