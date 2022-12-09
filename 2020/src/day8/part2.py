import re

input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read().splitlines()

def run_commands(commands):
    acc = 0
    going_to_loop = False
    executed_commands = []
    index = 0
    executed_last_command = False
    
    while not going_to_loop:
        if index not in executed_commands:
            executed_commands.append(index)
            [command, arg] = commands[index].split(' ')
            
            if command == 'acc':
                acc += int(arg)
                index += 1
            elif command == 'jmp':
                index += int(arg)
            else:
                index += 1
        else:
            going_to_loop = True
            
        if index == len(commands):
            executed_last_command = True
            break
    return acc if executed_last_command else 0

outside_index = 0
acc = 0

while acc == 0:
    if 'nop' in input[outside_index] or 'jmp' in input[outside_index]:
        commands = input.copy()
        [command, args] = commands[outside_index].split(' ')
        commands[outside_index] = '{} {}'.format('jmp' if command == 'nop' else 'nop', args)
        acc = run_commands(commands)
    outside_index += 1
    
print(acc)