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
                
                if k > 0:
                    left_visible_seat = None
                    left_seats = spaces[j][0:k]
                    left_seats.reverse()
                    for seat in left_seats:
                        if seat == '#':
                            left_visible_seat = '#'
                            break
                        elif seat == 'L':
                            break
                    if left_visible_seat == '#':
                        adjacent_seats += 1
                    
                if k < len(spaces[j]) - 1:
                    rigth_visible_seat = None
                    rigth_seats = spaces[j][k + 1:]
                    for seat in rigth_seats:
                        if seat == '#':
                            rigth_visible_seat = '#'
                            break
                        elif seat == 'L':
                            break
                    if rigth_visible_seat == '#':
                        adjacent_seats += 1
                    
                if j > 0:
                    if k > 0:
                        upper_left_diagonal_visible_seat = None
                        check = [j -1, k - 1]
                        while True:
                            if check[0] < 0 or check[1] < 0:
                                break
                            if spaces[check[0]][check[1]] == '#':
                                upper_left_diagonal_visible_seat = '#'
                                break
                            if spaces[check[0]][check[1]] == 'L':
                                upper_left_diagonal_visible_seat = 'L'
                                break
                            check[0] -= 1
                            check[1] -= 1
                        if upper_left_diagonal_visible_seat == '#':
                            adjacent_seats += 1

                    upper_visible_seat = None
                    check = j - 1
                    while True:
                        if check < 0:
                            break
                        if spaces[check][k] == '#':
                            upper_visible_seat = '#'
                            break
                        if spaces[check][k] == 'L':
                            upper_visible_seat = 'L'
                            break
                        check -= 1
                    if upper_visible_seat == '#':
                        adjacent_seats += 1
                        
                    if k < len(spaces[j]) - 1:
                        upper_rigth_diagonal_visible_seat = None
                        check = [j - 1, k + 1]
                        while True:
                            if check[0] < 0 or check[1] == len(spaces[j]):
                                break
                            if spaces[check[0]][check[1]] == '#':
                                upper_rigth_diagonal_visible_seat = '#'
                                break
                            if spaces[check[0]][check[1]] == 'L':
                                upper_rigth_diagonal_visible_seat = 'L'
                                break
                            check[0] -= 1
                            check[1] += 1
                        if upper_rigth_diagonal_visible_seat == '#':
                            adjacent_seats += 1
                
                if j < len(spaces) - 1:
                    if k > 0:
                        lower_left_diagonal_visible_seat = None
                        check = [j + 1, k - 1]
                        while True:
                            if check[0] == len(spaces) or check[1] < 0:
                                break
                            if spaces[check[0]][check[1]] == '#':
                                lower_left_diagonal_visible_seat = '#'
                                break
                            if spaces[check[0]][check[1]] == 'L':
                                lower_left_diagonal_visible_seat = 'L'
                                break
                            check[0] += 1
                            check[1] -= 1
                        if lower_left_diagonal_visible_seat == '#':
                            adjacent_seats += 1

                    lower_visible_seat = None
                    check = j + 1
                    while True:
                        if check == len(spaces):
                            break
                        if spaces[check][k] == '#':
                            lower_visible_seat = '#'
                            break
                        if spaces[check][k] == 'L':
                            lower_visible_seat = 'L'
                            break
                        check += 1
                    if lower_visible_seat == '#':
                        adjacent_seats += 1
                        
                    if k < len(spaces[j]) - 1:
                        lower_rigth_diagonal_visible_seat = None
                        check = [j + 1, k + 1]
                        while True:
                            if check[0] == len(spaces) or check[1] == len(spaces[j]):
                                break
                            if spaces[check[0]][check[1]] == '#':
                                lower_rigth_diagonal_visible_seat = '#'
                                break
                            if spaces[check[0]][check[1]] == 'L':
                                lower_rigth_diagonal_visible_seat = 'L'
                                break
                            check[0] += 1
                            check[1] += 1
                        if lower_rigth_diagonal_visible_seat == '#':
                            adjacent_seats += 1
                        
                if adjacent_seats >= 5 and spaces[j][k] == '#':
                    new_spaces[j][k] = 'L'
                elif adjacent_seats == 0 and spaces[j][k] == 'L':
                    new_spaces[j][k] = '#'
                elif spaces[j][k] == '#':
                    new_spaces[j][k] = '#'
    if new_spaces == spaces:
        print(sum(map(lambda x: x.count('#'), spaces)))
        break
    spaces = new_spaces
    