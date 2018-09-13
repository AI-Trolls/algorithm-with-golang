package main

func main() {
	// Peak finder

	// One-dimensional version
	// Proble: Find a peak if it exists

	// Straightforward algorithm
	// O(n)

	// Divide & Conquer
	// O(log n)
	// Look at n/2 position
	// if a[n/2] < a[n/2-1] then only look at left half 1 ~ n/2-1 to look for a peak
	// else if a[n/2] < a[n/2+1] then n/2+1 ~ n for a peak
	// else n/2 position is peak

	// T(n): "Work" algorithm dest on input of size n
	// = T(n/2) + O(1)
	// base case: T(1) = O(1)
	// T(n) = O(1) + ... + O(1) = O(log2N)
	//          log2n times

	// Two-dimensional version
	// Greedy Ascent algorithm
	// O(nm) complexity
	// O(n^2) if m=n

	// Divide & Conquer
	// Pick middle Column j=m/2
	// Gind global maximum on column j at (i,j)
	// Compare (i, j-1), (i, j), (i, j+1)
	// Pick left cols if (i, j-1) > (i,j)
	// If (i,j) >= (i, j-1), (i, j+1) => (i, j) is a 2-D peak
	// Solve the row problem with half the number of cols
	// When you have a single col, find the global maxximum <- done

	// T(n, m) = T(n, m/2) + O(n)
	// T(n, 1) = O(n)
	// T(n, m) = O(n) + ... + O(n) = O (nlog2m)
	//	              long2m

}
