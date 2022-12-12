input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = sample_input_file.read().splitlines()
numbers = list(map(lambda num: int(num), input))
numbers.sort()


continue_run = True

current_jolt = 0
possibilities = 0

runs = 0

def run_path(path):
    options = list(filter(lambda num: num - current_jolt <= 3, path))
    if options[0] === numbers[-1]:
        runs += 1
        return
    for i in range(len(options)):
        current_jolt = numbers[i]
        run_path(numbers[path.index(path[i]) + 1:])
        
while continue_run:
    run_path(numbers)
    possibilities += 1