#include <fstream>
#include <iostream>

using namespace std;

int main() {
  streampos size;
  char *memblock;

  ifstream file("two.txt", ios::in | ios::binary | ios::ate);
  if (file.is_open()) {
    size     = file.tellg();
    memblock = new char[size];
    file.seekg(0, ios::beg);
    file.read(memblock, size);
    file.close();

    cout << "the entire file content is in memory" << endl;

    delete[] memblock;
  } else {
    cout << "Unable to open file";
  }
  return 0;
}
