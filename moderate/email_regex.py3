import fileinput, re

email = re.compile('^((\"[0-9A-Za-z@.+-=]+\")|([0-9A-Za-z.+-=]+))@\\w*(\\w+\\.)+\\w{2,4}$')
for line in fileinput.input():
    print("true" if email.match(line) else "false")
