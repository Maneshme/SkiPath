# SkiPath
This application is to find the best ski path from a 2 dimensional array containing the elevation of the mountains.

The sample array considered here is a 5x5 array

40, 88, 71, 32, 10
21, 57, 94, 36, 8
62, 39, 25, 58, 67
48, 5, 10, 66, 39
54, 2, 0, 31, 41

The rules for finding the best path are:
1. Skier can go from a higher elevation to a lower one
2. Skier can go north, south, east, west.
3. The path found with maximum length is considered, if we get two paths with the same length we calculate the depth and choose the steepest path

In this example the longest path found are.

1. 88-57-39-25-10-5-2-0
2. 94-57-39-25-10-5-2-0

The best path is the second one as it is the steepest with a depth of 94
