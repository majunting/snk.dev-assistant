import fileinput, math

for line in fileinput.input():
    s = line.split()
    n = int(math.sqrt(len(s)))
    print(' '.join([s[j*n+i] for i in range(n) for j in range(n-1, -1, -1)]))
