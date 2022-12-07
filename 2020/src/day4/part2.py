import re

def validate_data_by_type(type, value):
    if type == 'byr':
        return int(value) >= 1920 and int(value) <= 2002
    elif type == 'iyr':
        return int(value) >= 2010 and int(value) <= 2020
    elif type == 'eyr':
        return int(value) >= 2020 and int(value) <= 2030
    elif type == 'hgt':
        if 'cm' in value:
            value = value.split('cm')[0]
            return int(value) >= 150 and int(value) <= 193
        else:
            value = value.split('in')[0]
            return int(value) >= 59 and int(value) <= 76
    elif type == 'hcl':
        return bool(re.match(r"#[0-9A-Fa-f]{6}", value))
    elif type == 'ecl':
        return value in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']
    elif type == 'pid':
        return len(value) == 9
    else:
        return True
        

def has_all_fields(fields):
    obrigatory_fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
    field_names = list(map(lambda field: field.split(':')[0], re.sub(r"\n", ' ', document).split(' ')))
    valid = True
    
    for field in obrigatory_fields:
        if field not in field_names:
            valid = False
            break
    return valid

input_file = open("input.txt", "r")
sample_input_file = open("sample-input-part2.txt", "r")

input = input_file.read()


documents = input.split('\n\n')
validCount = 0

for document in documents:
    fields = re.sub(r"\n", ' ', document).split(' ')
    is_valid = has_all_fields(fields)
    
    for fieldData in fields:
        [field, value] = fieldData.split(':')
        if not validate_data_by_type(field, value):
            is_valid = False
            break
    if is_valid:
        validCount += 1
        
    
        
print(validCount)
    
    
