import sys

x1, y1, x2, y2 = map(int, input().split())

latice = [
    (1, 2),
    (2, 1),
    (1, -2),
    (2, -1),
    (-1, 2),
    (-2, 1),
    (-1, -2),
    (-2, -1),
]

for xa, ya in latice:
    for xb, yb in latice:
        if (x1 + xa == x2 + xb) and (y1 + ya == y2 + yb):
            print('Yes')
            sys.exit()

print('No')
