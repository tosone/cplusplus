#include <fstream>
#include <iostream>

using namespace std;

int main() {
  streampos begin, end;
  ifstream myfile("two.txt", ios::binary);
  begin = myfile.tellg();
  myfile.seekg(0, ios::end);
  end = myfile.tellg();
  myfile.close();
  cout << "size is: " << (end - begin) << " bytes.\n";
  return 0;
}
