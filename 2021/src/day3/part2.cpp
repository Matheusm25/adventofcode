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

    while (getline(testInput, input)) {
      fileLineCount++;
    }
  
    testInput.clear();
    testInput.seekg(0, ios::beg);

    string bitCode[fileLineCount]; 

    for (int i = 0; i < fileLineCount; i++) {
      getline(testInput, input);
      bitCode[i] = input;
    }

    string gammaRate;
    string epsilonRate;

    for (int i = 0; i < bitCode[0].length(); i++) {
      int trueCount = 0;
      for (int j = 0; j < fileLineCount; j++) {
        if (bitCode[j][i] == '1') {
          trueCount++;
        }
      }

      if (trueCount >= fileLineCount / 2) {
        gammaRate += '1';
        epsilonRate += '0';
      } else {
        gammaRate += '0';
        epsilonRate += '1';
      }
    }

    result = stoi(gammaRate, 0, 2) * stoi(epsilonRate, 0, 2);

    cout << "Gamma rate: " << gammaRate << endl;
    cout << "Epsilon rate: " << epsilonRate << endl;


    cout << "Challenge result: " << result << endl;
    return 0;
}