#include <iostream>
#include <map>
#include <regex>
#include <set>
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
  set<int> s;
  set<int> all;
  string l;
  while (getline(cin, l)) {
    vector<int> nums = ints(l);
    int id = nums[0];
    all.insert(id);
    int c = nums[1];
    int r = nums[2];
    int w = nums[3];
    int h = nums[4];
    for (int i = r; i < r + h; i++) {
      for (int j = c; j < c + w; j++) {
        if (m[{i, j}] != 0) {
          s.insert(m[{i, j}]);
          s.insert(id);
        }
        m[{i, j}] = id;
      }
    }
  }

  for (int id : all) {
    if (!s.count(id)) {
      cout << id << endl;
    }
  }
  return 0;
}
