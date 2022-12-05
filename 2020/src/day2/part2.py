import re

input_file = open("./input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

passwords_stats = input.split('\n')

valid_passwords = 0

for stats in passwords_stats:
    [rule, password] = stats.split(': ')
    [positions, letter] = rule.split(' ')
    [first, second] = positions.split('-')
    if (password[int(first) - 1] == letter or password[int(second) - 1] == letter) and not (password[int(first) - 1] == letter and password[int(second) - 1] == letter):
        valid_passwords += 1


print(valid_passwords)

