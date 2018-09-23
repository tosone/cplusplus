#include <fstream>
#include <iostream>

using namespace std;

int main() {
  ofstream one;
  one.open("one.txt");
  one << "Writing this to a file.\n";
  one.close();

  ofstream two("two.txt");
  if (two.is_open()) {
    two << "This is a line.\n";
    two << "This is another line.\n";
    two.close();
  } else
    cout << "Unable to open file";
  return 0;
}
