input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

seats = input.split('\n')

max_seat_id = 0

for seat in seats:
    row = seat[:7]
    column = seat[7:]
     
    row_range = range(127)
    
    for step in row:
        if step == 'F':
            row_range = row_range[:int(len(row_range) / 2)]
        else:
            row_range = row_range[round(len(row_range) / 2):]

    col_range = range(7)
    
    for step in column:
        if len(col_range) > 1:
            if step == 'L':
                col_range = col_range[:int(len(col_range) / 2)]
            else:
                col_range = col_range[round(len(col_range) / 2):]
        else:
            col_range = range(col_range.start) if step == 'L' else range(col_range.stop)
    
    row_id = row_range.stop * 8 + col_range.stop
    if row_id > max_seat_id:
        max_seat_id = row_id
        

print(max_seat_id)
