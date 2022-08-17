from math import sqrt, floor
import sys

t1, t2, a1, a2 = map(int, input().split())


def is_prime(v):
    return len([k for k in range(2, floor(sqrt(v) + 1)) if (v) % k == 0]) == 0


t_is_win = True
for t in range(t1, t2+1):
    find_prime = False
    for a in range(a1, a2+1):
        prime = {}
        if is_prime(t + a):
            find_prime = True
            break
    if not find_prime:
        print('Takahashi')
        sys.exit()

print('Aoki')
