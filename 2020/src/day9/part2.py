input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read().splitlines()
numbers = list(map(lambda num: int(num), input))

# part_1_result = 127 # sample result
part_1_result = 50047984

    
for i in range(len(numbers) - 1):
    acc = numbers[i]
    used = [numbers[i]]

    for j in range(i + 1, len(numbers)):
        acc += numbers[j]
        used.append(numbers[j])
        
        if acc == part_1_result or acc > part_1_result:
            break
    if acc == part_1_result:
        used.sort()
        print(used, used[0] + used[-1])
        break