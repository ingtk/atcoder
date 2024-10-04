import sys

n, m = map(int, input().split())
a = list(map(int, input().split()))
b = list(map(int, input().split()))

am = {}
for v in a:
    if v in am:
        am[v] += 1
    else:
        am[v] = 1

for i in range(m):
    x = b[i]
    if x not in am or am[x] <= 0:
        print("No")
        sys.exit()
    am[x] -= 1

print("Yes")
