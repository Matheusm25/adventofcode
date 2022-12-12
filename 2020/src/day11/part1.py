input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read().splitlines()
spaces = list(map(lambda x: list(x), input))

while True:
    new_spaces = list(map(lambda x: list(x), input))
    
    for j in range(len(spaces)):
        for k in range(len(spaces[j])):
            if spaces[j][k] == '.':
                new_spaces[j][k] = '.'
                continue
            else:
                adjacent_seats = 0
                if k > 0 and spaces[j][k -1] == '#':
                    adjacent_seats += 1
                if k < len(spaces[j]) - 1 and spaces[j][k + 1] == '#':
                    adjacent_seats += 1
                    
                if j > 0:
                    if k > 0 and spaces[j - 1][k -1] == '#':
                        adjacent_seats += 1
                    if spaces[j - 1][k] == '#':
                        adjacent_seats += 1
                    if k < len(spaces[j]) - 1 and spaces[j - 1][k + 1] == '#':
                        adjacent_seats += 1
                
                if j < len(spaces) - 1:
                    if k > 0 and spaces[j + 1][k -1] == '#':
                        adjacent_seats += 1
                    if spaces[j + 1][k] == '#':
                        adjacent_seats += 1
                    if k < len(spaces[j]) - 1 and spaces[j + 1][k + 1] == '#':
                        adjacent_seats += 1
                        
                if adjacent_seats >= 4 and spaces[j][k] == '#':
                    new_spaces[j][k] = 'L'
                elif adjacent_seats == 0 and spaces[j][k] == 'L':
                    new_spaces[j][k] = '#'
                elif spaces[j][k] == '#':
                    new_spaces[j][k] = '#'
    if new_spaces == spaces:
        print(sum(map(lambda x: x.count('#'), spaces)))
        break
    spaces = new_spaces
    