using namespace std;

#include <iostream>
#include <fstream>
#include <string>

int main() {
    ifstream testInput ("./sample-input.txt");
    ifstream realInput ("./input.txt");
    string input;

    int increments = 0;

    int fileLineCount = 0;

    while (getline(realInput, input)) {
      fileLineCount++;
    }
  
    realInput.clear();
    realInput.seekg(0, ios::beg);

    int fileData[fileLineCount];

    for (int i = 0; i < fileLineCount; i++) {
      getline(realInput, input);
      fileData[i] = stoi(input);
    }

    int lastSum = fileData[0] + fileData[1] + fileData[2];

    for (int i = 1; i < fileLineCount; i++) {
      int currentSum = 0;
      for (int j = 0; j < 3; j++) {
        if (i + j < fileLineCount) {
          currentSum += fileData[i + j];
        }
      }
      if (currentSum > lastSum) {
        increments++;
      }
      lastSum = currentSum;
    }

    cout << "Challenge result: " << increments << endl;
    return 0;
}