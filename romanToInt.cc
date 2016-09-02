class Solution {
public:
    int romanToInt(string s) {
        map<char, int> roman = {
            {'M', 1000},
            {'D', 500},
            {'C', 100},
            {'L', 50},
            {'X', 10},
            {'V', 5},
            {'I', 1}
        };
        if (s.size() == 0) {
            return 0;
        }
        int result = roman[s[0]];
        for (int i = 1; i < s.size(); ++i) {
            int prev = roman[s[i - 1]];
            int curr = roman[s[i]];
            if (prev < curr) {
                result += (curr - 2*prev);//add prev before curr, so subtract 2*prev
            } else {
                result += curr;
            }
        }
        return result;
    }
};