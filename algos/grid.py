
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
