#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

int main() {
  unordered_set<int> s;
  vector<int> v;

  int f = 0;
  s.insert(f);

  int x;
  while (cin >> x) {
    v.push_back(x);
  }

  int i = 0;
  while (1) {
    f += v[i];
    if (s.count(f)) {
      cout << f << '\n';
      return 0;
    }
    s.insert(f);
    i = (i + 1) % v.size();
  }
}
