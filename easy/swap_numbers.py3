import fileinput

for line in fileinput.input():
    print(' '.join([i[-1] + i[1:-1] + i[0] for i in line.split()]))
