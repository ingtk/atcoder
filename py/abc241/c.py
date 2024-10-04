import sys

n = int(input())
s = []
for i in range(n):
    s.append(input())

b = 6


def check(box):
    for i in range(b):
        # horizontal
        if len(["#" for k in range(b) if box[i][k] == "#"]) >= 4:
            return True
        # vertical
        if len(["#" for k in range(b) if box[k][i] == "#"]) >= 4:
            return True

    if len(["#" for k in range(b) if box[k][k] == "#"]) >= 4:
        return True
    if len(["#" for k in range(b) if box[k][b - k - 1] == "#"]) >= 4:
        return True


for y in range(n - b + 1):
    for x in range(n - b + 1):
        box = [k[x : x + b] for k in s[y : y + b]]
        if check(box):
            print("Yes")
            sys.exit()

print("No")


# def vcheck(i, j, cnt=0, dot_cnt=0):
#     if dot_cnt > 2:
#         return False

#     if cnt == 6:
#         return True

#     if cnt < 6:
#         if s[i+cnt][j] == '.':
#             dot_cnt+=1
#         cnt+=1
#         return vcheck(i, j, cnt, dot_cnt)

# def hcheck(i, j, cnt=0, dot_cnt=0):
#     if dot_cnt > 2:
#         return False

#     if cnt == 6:
#         return True

#     if cnt < 6:
#         if s[i][j+cnt] == '.':
#             dot_cnt+=1
#         cnt+=1
#         return hcheck(i, j, cnt, dot_cnt)

# def scheck(i, j, cnt=0, dot_cnt=0):
#     if dot_cnt > 2:
#         return False

#     if cnt == 6:
#         return True

#     if cnt < 6:
#         if s[i+cnt][j+cnt] == '.':
#             dot_cnt+=1
#         cnt+=1
#         return scheck(i, j, cnt, dot_cnt)

# def rscheck(i, j, cnt=0, dot_cnt=0):
#     if dot_cnt > 2:
#         return False

#     if cnt == 6:
#         return True

#     if cnt < 6:
#         if s[n-1-i-cnt][j+cnt] == '.':
#             dot_cnt+=1
#         cnt+=1
#         return rscheck(i, j, cnt, dot_cnt)
