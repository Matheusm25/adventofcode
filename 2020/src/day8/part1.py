input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read().splitlines()

acc = 0
going_to_loop = False
executed_commands = []
index = 0

while not going_to_loop:
    if index not in executed_commands:
        executed_commands.append(index)
        [command, arg] = input[index].split(' ')
        
        if command == 'acc':
            acc += int(arg)
            index += 1
        elif command == 'jmp':
            index += int(arg)
        else:
            index += 1
    else:
        going_to_loop = True
        
print(acc)