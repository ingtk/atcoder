import collections

inp = input()
alphabet = 'abcdefghijklmnopqrstuvwxyz'
c  = collections.Counter(inp)
res = ''
for a in alphabet:
    res += (a * c[a])

print(res)



