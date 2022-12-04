using namespace std;

#include <iostream>
#include <fstream>
#include <string>

int main() {
    ifstream testInput ("./sample-input.txt");
    ifstream realInput ("./input.txt");
    string input;

    int increments = 0;
    getline(realInput, input);
    int lastLine = stoi(input);

    while (getline(realInput, input)) {
      int line = stoi(input);
      if (line > lastLine) {
        increments++;
      }
      lastLine = line;
    }

    cout << "Challenge result: " << increments << endl;
    return 0;
}