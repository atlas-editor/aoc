#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

int p2() {
  vector<string> l;
  string x;
  while (cin >> x) {
    l.push_back(x);
  }

  for (int i = 0; i < l.size(); i++) {
    for (int j = i + 1; j < l.size(); j++) {
      string a = l[i];
      string b = l[j];
      int c = a.size();
      int d = 0;
      for (int k = 0; k < a.size(); k++) {
        if (a[k] != b[k]) {
          d += 1;
        }
      }

      if (d == 1) {
        cout << a << '\n';
        cout << b << '\n';
        return 0;
      }
    }
  }

  return 0;
}

int p1() {

  string x;
  int res2 = 0;
  int res3 = 0;
  while (cin >> x) {
    unordered_map<char, int> m;
    for (char c : x) {
      m[c] += 1;
    }

    bool twos = false;
    bool threes = false;
    for (auto [c, count] : m) {
      if (!twos && (count == 2)) {
        res2++;
        twos = true;
      }
      if (!threes && (count == 3)) {
        res3++;
        threes = true;
      }
    }
  }

  cout << res2 * res3 << '\n';
  return 0;
}

int main() { return p2(); }
