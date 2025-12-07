#include <iostream>
#include <map>
#include <regex>
#include <string>
#include <utility>
#include <vector>
using namespace std;

vector<int> ints(const string &s) {
  regex re = regex("-?\\d+");

  sregex_iterator it = sregex_iterator(s.begin(), s.end(), re);
  sregex_iterator end;

  vector<int> nums;
  for (; it != end; it++) {
    nums.push_back(stoi((*it).str()));
  }

  return nums;
}

int main() {
  map<pair<int, int>, int> m;
  string l;
  while (getline(cin, l)) {
    vector<int> nums = ints(l);
    int c = nums[1];
    int r = nums[2];
    int w = nums[3];
    int h = nums[4];
    for (int i = r; i < r + h; i++) {
      for (int j = c; j < c + w; j++) {
        m[{i, j}]++;
      }
    }
  }

  int r = 0;
  for (auto [p, i] : m) {
    if (i > 1) {
      r++;
    }
  }
  cout << r << endl;
  return 0;
}
