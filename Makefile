.SHELLFLAGS: -ec

# Run WorldClock to check the time of some cities
.PHONY: run
run:
	@go run main.go --cities Europe/Berlin,Europe/Rome,America/New_York
