def checkRoute(input, right, down):
    trees_hit = 0
    mapped_length = len(input[0])
    position = [0, 0]

    while True:
        position[0] = position[0] + right if position[0] + right < mapped_length else position[0] + right - mapped_length
        position[1] += down

        if map[position[1]][position[0]] == '#':
            trees_hit += 1
        
        if position[1] + 1 >= len(map):
            break

    return trees_hit

input_file = open("./input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()
map = input.split('\n')

print(
    checkRoute(map, 1, 1)
    * checkRoute(map, 3, 1)
    * checkRoute(map, 5, 1)
    * checkRoute(map, 7, 1)
    * checkRoute(map, 1, 2)
)
