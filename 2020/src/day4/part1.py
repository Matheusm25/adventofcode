import re

input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

obrigatory_fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']

documents = input.split('\n\n')
validCount = 0

for document in documents:
    field_names = list(map(lambda field: field.split(':')[0], re.sub(r"\n", ' ', document).split(' ')))
    is_valid = True
    for field in obrigatory_fields:
        if field not in field_names:
            is_valid = False
            break
    if is_valid:
        validCount += 1
        
print(validCount)
    