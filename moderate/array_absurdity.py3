import fileinput

for line in fileinput.input():
    st = line.rstrip('\n').split(";")
    s, r = sorted(st[1].split(",")), ""
    for i in range(0, len(s)-1):
        if s[i] == s[i+1]:
            r = s[i]
            break
    print(r)
