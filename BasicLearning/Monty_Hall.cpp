#include <algorithm>
#include <cstdlib>
#include <ctime>
#include <iostream>
#include <vector>

using namespace std;

int main() {
  srand(time(NULL));

  unsigned long long stay   = 0;
  unsigned long long change = 0;
  while (true) {
    vector<bool> boxes(3, false);
    boxes[0] = true; // Place a car
    random_shuffle(boxes.begin(), boxes.end());

    if (boxes[rand() % 3]) {
      stay++;
      cout << "STAY wins\n";
    } else {
      change++;
      cout << "CHANGE wins\n";
    }
    cout << "STAY: " << int((double)stay / (stay + change) * 100 + 0.5) << "%; "
         << "CHANGE: " << int((double)change / (stay + change) * 100 + 0.5) << "%\n"
         << endl;
  }
}
