package algos

// Problems: 704, 278, 35
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
//	(dir, result, includemid)
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
// A state parameter can also be passed to the bisect method to
// change or accumulate state between calls to bisect.
//
// By default bisect returns the "mid" index where the bisection
// ends - but user can do other things and add things into their
// own context/state objects
func Bisect(lo, hi int, cmpfunc func(lo, hi, mid int) (newmid int, incmid bool)) int {
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
