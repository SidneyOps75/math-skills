## PROJECT MATH-SKILLS

### Objectives
This program should be able to read from a file and print the result of each statistic mentioned below.
The statistics to be computed are :
1.  Average
2.  Median
3.  Variance
4.  Standard Deviation

The above mentioned are seperated from the main file and stored in their own functions in [skills](./skills/skills.go) file.

In other words, this program must be able to read the data present in the path passed as argument.

### How to Run

To run your program a command similar to this one will be used if your project is made in Go:
```bash
$> go run . data.txt
```

After reading the file, your program must execute each of the calculations asked above and print the results in the following manner (the following numbers are only examples):
```bash
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```


### Testing
In the [skills_test](./skills/skills_test.go) is where this program has tested each statistics computation. 

### Authors

This simple math project has been authored by [John Opiyo](https://github.com/SidneyOps75)