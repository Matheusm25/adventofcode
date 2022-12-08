import re

input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

rules = input.split('\n')

bag_name = 'shiny gold'
count = 0

bag_rules = {}
    
for rule in rules:
    [bag, contains] = re.sub(r" (bags|bag)", '', rule).split(' contain ')
    inside_bags = contains[:-1].split(', ')
    bag_rules[bag] = list(map(lambda inside_bag: ' '.join(inside_bag.split(' ')[1:]), inside_bags))
    
def get_all_bags(bag):
    if bag in bag_rules:
        inside_bags = list(filter(lambda b: b in bag_rules, bag_rules[bag]))
        for inside_bag in bag_rules[bag]:
            inside_bags.extend(get_all_bags(inside_bag))
        return inside_bags
    else:
        return []        

for rule in rules:
    [bag, contains] = re.sub(r" (bags|bag)", '', rule).split(' contain ') if ' contain ' in rule else [rule, '']
    inside_bags = get_all_bags(bag)
    count += bag_name in inside_bags

print(count)