
def allneigh(R, C, r, c):
    """ Return all the 8 neighbours of a location in a grid. """
    for dr in [-1, 0, 1]:
        for dc in [-1, 0, 1]:
            nr,nc = r + dr, c + dc
            if nr >= 0 and nr < R and nc >= 0 and nc < C:
                if (nr,nc) != (r,c):
                    yield nr,nc

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
