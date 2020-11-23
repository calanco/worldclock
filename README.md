# WordClock
Tool to get the current time of any capital city of the world.
Usage of __goroutines__ to query __worldtimeapi.org__ API server.

## Usage
The capital cities to get the current time from have to be passed in the _capitals_ parameter.
That parameter is a comma separated list of cities expressed as `Continent/Capital`.

```bash
    -capitals string
    Insert the capitals to get time from [Continent1/Capital1, Continent2/Capital2, etc]
```

### Example
The following command:
```
    go run main.go --capitals Europe/Berlin,Europe/Rome,America/New_York
```
produces the following output:

![alt Output example](/example.png)
