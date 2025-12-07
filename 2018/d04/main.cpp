#include <iostream>
#include <regex>
#include <string>
#include <tuple>
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
  string l;
  pair<tuple<int, int, int, int, int>, string> t;
  vector<pair<tuple<int, int, int, int, int>, string>> v;
  while (getline(cin, l)) {
    vector<int> nums = ints(l);
    int y = nums[0];
    int m = nums[1];
    int d = nums[2];
    int h = nums[3];
    int min = nums[4];

    string ss = l.substr(l.find("] ") + 1);
    t = {{y, m, d, h, min}, ss};
    v.push_back(t);
  }

  sort(v.begin(), v.end(),
       [](const auto &a, const auto &b) { return a.first < b.first; });

  for (auto &[dt, s] : v) {
    cout << get<4>(dt) << ' ' << s << endl;
  }

  return 0;
}
