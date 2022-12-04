using namespace std;

#include <iostream>
#include <fstream>
#include <string>

int main() {
    ifstream testInput ("./sample-input.txt");
    ifstream realInput ("./input.txt");
    string input;

    int result;
    int fileLineCount = 0;

    while (getline(realInput, input)) {
      fileLineCount++;
    }
  
    realInput.clear();
    realInput.seekg(0, ios::beg);

    string instructions[fileLineCount]; 

    for (int i = 0; i < fileLineCount; i++) {
      getline(realInput, input);
      instructions[i] = input;
    }

    int depth = 0;
    int hPosition = 0;
    int aim = 0;

    for (int i = 0; i < fileLineCount; i++) {
      string command = instructions[i].substr(0, instructions[i].find(" "));
      int value = stoi(instructions[i].substr(instructions[i].find(" ") + 1, instructions[i].length()));

      if (command == "forward") {
        hPosition += value;
        depth += aim * value;
      } else if (command == "down") {
        aim += value;
      } else {
        aim -= value;
      }
    }

    result = hPosition * depth;

    cout << "Challenge result: " << result << endl;
    return 0;
}