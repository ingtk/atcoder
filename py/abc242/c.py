digit = int(input())

dp = [[]]

num = [0, 2, 3, 3, 3, 3, 3, 3, 3, 2]

for n in range(2, digit):
    new_num = [0] * 10
    for k in range(1, 10):
        if 2 <= k and k <= 8:
            new_num[k] = num[k] + num[k - 1] + num[k + 1]
        elif k == 1:
            new_num[k] = num[k] + num[k + 1]
        elif k == 9:
            new_num[k] = num[k] + num[k - 1]
        new_num[k] %= 998244353
    num = new_num

print(sum([num[k] for k in range(1, 10)]) % 998244353)
# a = 10000

# print(a % 998244353)
