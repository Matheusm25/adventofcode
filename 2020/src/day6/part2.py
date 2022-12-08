input_file = open("input.txt", "r")
sample_input_file = open("sample-input.txt", "r")

input = input_file.read()

groups = input.split('\n\n')

pointCount = 0

for group in groups:
    lines = group.split('\n')
    points = 0
    
    for letter in lines[0]:
        allHasLetter = list(filter(lambda line: line.find(letter) >= 0, lines))
        if len(allHasLetter) == len(lines):
            points += 1

    pointCount += points        
    
    

print(pointCount)
    

    

