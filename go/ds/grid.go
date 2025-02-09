package ds

import "iter"

// A basic twod array - ideally for cache in DP problems

type Grid[T any] struct {
	Rows   int
	Cols   int
	Values []T
}

func NewGrid[T any](rows, cols int) *Grid[T] {
	return &Grid[T]{
		Rows:   rows,
		Cols:   cols,
		Values: make([]T, rows*cols),
	}
}

func (t *Grid[T]) Get(r, c int) T {
	return t.Values[r*t.Cols+c]
}

func (t *Grid[T]) Set(r, c int, val T) {
	t.Values[r*t.Cols+c] = val
}

func (t *Grid[T]) SafeGet(r, c int, defaultVal T) T {
	if r >= 0 && r < t.Rows && c >= 0 && c < t.Cols {
		return t.Values[r*t.Cols+c]
	} else {
		return defaultVal
	}
}

func (t *Grid[T]) SafeSet(r, c int, val T) {
	if r >= 0 && r < t.Rows && c >= 0 && c < t.Cols {
		t.Values[r*t.Cols+c] = val
	}
}

// Go through all the 8 neighbors of a location in a grid.
func IterAllNeighbors(R, C, r, c int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for dr := -1; dr <= 1; dr += 1 {
			for dc := -1; dc <= 1; dc += 1 {
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < R && nc >= 0 && nc < C {
					if nr != r || nc != c {
						if !yield(nr, nc) {
							return
						}
					}
				}
			}
		}
	}
}

func IterSideNeighbors(R, C, r, c int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for dr := -1; dr <= 1; dr += 1 {
			for dc := -1; dc <= 1; dc += 1 {
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < R && nc >= 0 && nc < C {
					if nr*nc == 0 { // Rules out (dr,dc) in (-1, -1), (1, -1), (-1, 1), (1, 1)
						if nr != r || nc != c { // This rules out dr,dc = (0,0)
							if !yield(nr, nc) {
								return
							}
						}
					}
				}
			}
		}
	}
}

/*
def sideneigh(R, C, r, c):
    """ Return the 4 non diagonal neighbors of a location in a grid. """
    for nr,nc in [(r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)]:
        if nr >= 0 and nr < R and nc >= 0 and nc < C:
            yield nr,nc

def itermatrix(R, C):
    return ((r,c) for r in range(R) for c in range(C))

def cell_neighbors(R, C, nfunc):
    R = len(matrix)
    C = len(matrix[0])

    def neighfunc(loc, parents=None):
        r,c = loc
        for nr,nc in nfunc(R, C, r, c):
            yield nr,nc
    return neighfunc
*/
