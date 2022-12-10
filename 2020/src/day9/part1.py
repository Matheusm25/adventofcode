input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read().splitlines()
numbers = list(map(lambda num: int(num), input))

def get_possibilities(start, end):
    nums = numbers[start:end]
    possibilities = []
    
    for i in range(len(nums) - 1):
        for j in range(i + 1, len(nums)):
            possibilities.append(nums[i] + nums[j])
            
    return possibilities
            

preamble = 25

for i in range(preamble, len(numbers)):
    if not numbers[i] in get_possibilities(i - preamble, i):
        print(numbers[i])
        break