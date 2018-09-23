#include <fstream>
#include <iostream>
#include <string>

using namespace std;

int main() {
  string line;
  ifstream one("two.txt");
  if (one.is_open()) {
    while (getline(one, line)) {
      cout << line << '\n';
    }
    one.close();
  } else {
    cout << "Unable to open file";
  }
  return 0;
}
