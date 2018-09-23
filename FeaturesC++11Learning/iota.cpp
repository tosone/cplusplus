#include <iostream>
#include <numeric>

using namespace std;

int main() {
  int arr[5]  = {0};
  char str[3] = {0};
  iota(arr, arr + 5, 10);  // changes arr to {10, 11, 12, 13, 14}
  iota(str, str + 3, 'a'); // {'a', 'b', 'c'}
  cout << str << endl;
  for (auto i : arr) { cout << i << endl; }
}
