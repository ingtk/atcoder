n = int(input())
a = [int(v) for v in input().split()]

i = 0

c = []

while i < n:
    v = a[i]
    if len(c) == 0:
        c += [(v, 1)]
    else:
        bc = c[len(c) - 1]
        cnt = 1
        if v == bc[0]:
            cnt = bc[1] + 1

        if v == cnt:
            for _ in range(cnt - 1):
                c.pop()
        else:
            c += [(v, cnt)]

    print(len(c))
    i += 1

