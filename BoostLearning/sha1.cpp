#include <boost/uuid/detail/sha1.hpp>
#include <cstdio>
#include <iostream>
#include <string>

std::string get_sha1(const std::string &p_arg) {
  boost::uuids::detail::sha1 sha1;
  sha1.process_bytes(p_arg.data(), p_arg.size());
  unsigned hash[5] = {0};
  sha1.get_digest(hash);
  char buf[41] = {0};
  for (int i = 0; i < 5; i++) {
    std::sprintf(buf + (i << 3), "%08x", hash[i]);
  }
  return std::string(buf);
}

int main(int argc, char const *argv[]) {
  std::cout << get_sha1("tosone") << std::endl;
  return 0;
}
