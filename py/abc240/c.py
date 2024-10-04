n, x = map(int, input().split())
dp = [[False] * (x+1)]
a = []
b = []
for _ in range(n):
    a1, b1 = map(int, input().split())
    a.append(a1)
    b.append(b1)
    dp.append([False] * (x+1))

dp[0][0] = True

for i in range(n):
    for j in range(x):
        if dp[i][j]:
            if j + a[i] <= x:
                dp[i+1][j + a[i]] = True
            if j + b[i] <= x:
                dp[i+1][j + b[i]] = True

if dp[n][x]:
    print('Yes')
else:
    print('No')


# vv = v.split(" ")
# n = int(vv[0])
# x = int(vv[1])
# a = []
# b = []
# for i in range(n):
#     inp = input()
#     vv = inp.split(" ")
#     a.append(int(vv[0]))
#     b.append(int(vv[1]))

# def hoge(nn, z, s):
#     if nn >= n:
#         if s == x:
#             return True
#         else:
#             return False

#     s += z[nn]

#     if s > x:
#         return False

#     if hoge(nn + 1, a, s):
#         return True

#     if hoge(nn + 1, b, s):
#         return True

#     return False

# if hoge(0, a, 0):
#     print("Yes")
# elif hoge(0, b, 0):
#     print("Yes")
# else:
#     print("No")

# n, x = map(int, input().split())
# dp = 1
# for _ in range(n):
#   a, b = map(int, input().split())
#   dp = (dp << a) | (dp << b)
# print("Yes" if ((dp >> x) & 1) == 1 else "No")

# if hoge(0, a, 0) == x:
#     print("Yes")
# elif hoge(0, b, 0) == x:
#     print("Yes")
# else:
#     print("No")
