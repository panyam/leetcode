package algos

// Problems: 704, 278, 35, 69
// https://leetcode.com/problems/first-bad-version/
//
// We have a generic bisect method we should use for all
// bisection problems - Keep this in some kind of template
//
// We pass this with lo, hi and a "checker" function
// lo is usually "0" or start of the array index
// hi is usually the index of the last element - len(input) - 1
// checker func takes the current lo, hi and a mid and returns
//
//	the following:
//
//	(dir, includemid)
//
// dir can be 0, -1 or 1.
// if dir is 0 then we stop the bisection there itself
// if dir > 0 then we are bisecting on the "right half" and lo,hi are updated to lo -> mid - 1
// if dir < 0 then we are bisecting on the "left half" and lo,hi are updated to mid + 1,hi
//
// With includemid - if it is True then when we bisect in left or
// right we also include "mid" in the new range so we dont skip
// "mid".  This is useful in some problems.
//
// By default bisect returns the "mid" index where the bisection ends
func BisectAdv(lo, hi int, cmpfunc func(lo, hi, mid int) (newmid int, incmid bool)) int {
	for lo <= hi {
		// mid = ((lo + hi) / 2)
		mid := lo + (hi-lo)/2
		dir, incmid := cmpfunc(lo, hi, mid)
		if dir == 0 {
			return mid
		} else if dir < 0 { // left
			hi = mid - 1
			if incmid {
				hi += 1
			}
		} else { // right side
			lo = mid + 1
			if incmid {
				lo -= 1
			}
		}
	}
	return -1
}

// Simpler Bisection
//
// Just provide lo, hi and checker function that returns a new low, high and when to stop
// When stop == true, the low and hi can be used to return values that the caller should return.
//
// When doing the bisection always think of the following cases:

// target < F[lo]			// most commonly lo = 0
// target > F[hi]			// most commonly hi == N - 1
// target == F[mid]		// most often - success case

// target != F[mid]		// problems mainly here combined with bottom two
// lo == hi
// lo == hi + 1  AND F[lo] < target < F[hi]		<--- case of insertion
func Bisect(lo, hi int, cmpfunc func(lo, hi, mid int) (newlow, newhi int, stop bool)) int {
	stop := false
	for lo <= hi && !stop {
		// mid = ((lo + hi) / 2)
		mid := lo + (hi-lo)/2
		lo, hi, stop = cmpfunc(lo, hi, mid)
		if stop {
			return lo
		}
	}
	return -1
}
