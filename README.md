# <span style="color:magenta; size : 20px"> Linear-Stats </span>

### <span style="color:pink"> Objectives
The purpose of this project is to calculate key statistical metrics from a dataset, specifically:
- Average
- Median
- Variance
- Standard Deviation

The program must read numerical data from a file and compute the statistics mentioned above. The data file should contain one value per line.
Each line represents a value in the statistical population.

### <span style="color:pink"> Author 

- Hasnae Lamrani

### <span style="color:pink"> How to run ?

To run your program, use the following command if your project is implemented in Go:
```
$ go run . data.txt
```
After reading the file, your program should execute the calculations and print the results in the following format (example values are provided):
```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

### <span style="color:pink"> How to test ?

- After downloading the file of test.

- To be able to run this script you need to have installed locally
 [`docker-engine`](https://docs.docker.com/engine/install/)

- Run the script with `./bin/math-skills`or `./run.sh math-skills`. The `<bin/math-skills>` must be inside the `/bin` directory, then run the program with the created file *data.txt* by the previous command.