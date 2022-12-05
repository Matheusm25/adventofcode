input_file = open("./input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

map = input.split('\n')
trees_hit = 0

mapped_length = len(map[0])
position = [0, 0]

while True:
    position[0] = position[0] + 3 if position[0] + 3 < mapped_length else position[0] + 3 - mapped_length
    position[1] += 1

    if map[position[1]][position[0]] == '#':
        trees_hit += 1
    
    if position[1] + 1 == len(map):
        break

print(trees_hit)