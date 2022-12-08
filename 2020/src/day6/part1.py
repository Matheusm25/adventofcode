input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

groups = input.split('\n\n')

pointCount = 0

for group in groups:
    pointCount += len(set(group.replace('\n', '')))

print(pointCount)
    

    

