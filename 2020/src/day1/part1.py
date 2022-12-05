import os.path

input_file = open("./input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

numbers = input.split('\n')

for i in range(len(numbers)):
    for j in range(i, len(numbers)):
        if int(numbers[i]) + int(numbers[j]) == 2020:
            print(int(numbers[i]) * int(numbers[j]))

