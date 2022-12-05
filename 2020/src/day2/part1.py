import re

input_file = open("./input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

passwords_stats = input.split('\n')

valid_passwords = 0

for stats in passwords_stats:
    [rule, password] = stats.split(': ')
    [limits, letter] = rule.split(' ')
    [min, max] = limits.split('-')
    letterCount = len(re.sub(r"[^{}]+".format(letter), "", password))
    if letterCount >= int(min) and letterCount <= int(max):
        valid_passwords += 1


print(valid_passwords)

