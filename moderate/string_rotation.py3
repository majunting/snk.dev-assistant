import fileinput

for line in fileinput.input():
    s, f, p = [list(i) for i in line.rstrip('\n').split(",")], "False", 0
    try:
        i = s[1].index(s[0][0])
    except (ValueError, IndexError):
        print(f)
        continue
    while i < len(s[1]):
        i += p
        if s[1][i:] + s[1][:i] == s[0]:
            f = "True"
            break
        try:
            p, i = i + 1, s[1][(i+1):].index(s[0][0])
        except ValueError:
            break
    print(f)
