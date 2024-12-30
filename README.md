# Go Off-by-One Error Example

This repository demonstrates a common off-by-one error in Go when iterating over arrays.  The provided code attempts to access an array element beyond its bounds, leading to a runtime panic.  The solution shows how to correctly iterate within the array's valid index range.

## Bug Description

The original code has a loop that runs from `i = 0` to `i <= 5`.  This means it attempts to access `a[5]`, which is out of bounds for a 5-element array (valid indices are 0-4).

## Solution

The corrected code adjusts the loop condition to `i < 5`, ensuring that the loop iterates only within the valid index range of the array.