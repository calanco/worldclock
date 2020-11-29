.SHELLFLAGS: -ec

# Run WorldClock to check some city times
.PHONY: run
run:
	@go run main.go --cities Europe/Berlin,Europe/Rome,America/New_York