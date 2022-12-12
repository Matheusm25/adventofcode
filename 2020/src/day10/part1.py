input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read().splitlines()
numbers = list(map(lambda num: int(num), input))
numbers.sort()

single_jolt_difference = 0
three_jolt_difference = 1 

current_jolt = 0

for i in range(len(numbers)):
    if numbers[i] - current_jolt == 1:
        single_jolt_difference += 1
    elif numbers[i] - current_jolt == 3:
        three_jolt_difference += 1
    
    current_jolt = numbers[i]

    
print(single_jolt_difference, three_jolt_difference, single_jolt_difference * three_jolt_difference)