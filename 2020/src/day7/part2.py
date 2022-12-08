import re

input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

rules = input.split('\n')

bag_name = 'shiny gold'

bag_rules = {}

    
for rule in rules:
    [bag, contains] = re.sub(r" (bags|bag)", '', rule).split(' contain ')
    inside_bags = contains[:-1].split(', ')
    inside_rule = []
    for inside_bag in inside_bags:
        [count, *name] = inside_bag.split(' ')
        if count.isdigit():
            inside_rule.append({ 'count': int(count), 'name': ' '.join(name) })
    bag_rules[bag] = inside_rule
    

def get_all_bags(bag):
    if bag in bag_rules:
        bag_names = list(map(lambda b: b['name'], bag_rules[bag]))
        inside_bags = list(filter(lambda b: b in bag_rules, bag_names))
        
        bag_count = 0
        
        for inside_bag in bag_rules[bag]:
            bag_count += inside_bag['count']
            [bags, count] = get_all_bags(inside_bag['name'])
            inside_bags.extend(bags)
            bag_count += count * inside_bag['count']
        return [inside_bags, bag_count]
    else:
        return [[], 0]
    
print(get_all_bags(bag_name)[1])
