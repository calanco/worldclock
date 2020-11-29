# WordClock
Tool to get the current time of any big city of the world.
Usage of __goroutines__ to query __worldtimeapi.org__ API server.

## Usage
The cities to get the current time from have to be passed in the _cities_ parameter.
That parameter is a comma separated list of cities expressed as `Continent/City`.

```bash
    -cities string
    Insert the cities to get time from [Continent1/City1, Continent2/City2, etc]
```

### Example
The following command:
```
    go run main.go --cities Europe/Berlin,Europe/Rome,America/New_York
```
produces the following output:

![alt Output example](/example.png)
