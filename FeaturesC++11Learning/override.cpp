#include <iostream>

struct base {
  virtual void foo() = 0;
};

struct derived : base {
  virtual void foo() override {
    std::cout << __PRETTY_FUNCTION__ << std::endl;
  }
};

int main() {
  base *override = new derived();
  override->foo();
  return 0;
}
