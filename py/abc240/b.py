import collections

n = int(input())
a = list(map(int, input().split()))

counter = collections.Counter(a)
print(len(counter.keys()))
