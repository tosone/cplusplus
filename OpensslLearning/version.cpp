#include <iostream>
#include <openssl/opensslv.h>

using namespace std;

int main(void) {
  cout << OPENSSL_VERSION_TEXT << endl;
  return 0;
}