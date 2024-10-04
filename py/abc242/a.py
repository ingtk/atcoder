inp = input()
v=[int(i) for i in inp.split(' ')]
a = v[0]
b = v[1]
c = v[2]
x = v[3]
ans = 0
if x <= a:
    ans = 1
elif x <= b:
    ans = c / (b - a)

print(ans)

