# Problem Statement

Given two arrays, `A` and `B`, of the same length (`n`) representing permutations of numbers from `1` to `n`,  create an array `C` where each element `C[i]` represents the count of common numbers that appear in both `A` and `B` up to and including the index `i`.

## Examples

### Example 1

**Input:**  

    A = [2, 3, 1]
    B = [3, 1, 2]

**Output:**  

    C = [0, 1, 3]

**Explanation:**  

At index 0, there are no common numbers before it.  
At index 1, the number '3' has appeared in both A and B before index 1.  
At index 2, the numbers '1', '2', and '3' have all appeared previously in both arrays.

----

**Input:**  

    A = [1, 3, 2, 4]
    B = [3, 1, 2, 4]

**Output:**  

    C = [0, 2, 3, 4]

**Explanation:**  

At index 0, no common numbers.  
At index 1, both '1' and '3' have appeared in both arrays prior.  
At index 2, numbers '1', '2', '3' are common.  
At index 3, all numbers (1 through 4) have been seen in both arrays.  