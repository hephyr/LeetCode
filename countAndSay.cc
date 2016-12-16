class Solution {
public:
    string countAndSay(int n) {
        string s = "1";
        string result = s;
        int count = 1;
        while (n > 1) {
            result = "";
            for (int i = 0; i < s.size(); ++i) {
                if (i+1 == s.size() || s[i] != s[i+1]) {
                        result.push_back('0' + count);
                        result.push_back(s[i]);
                        count = 1;
                } else if (s[i] == s[i+1]) {
                    ++count;
                }
            }
            s = result;
            --n;
        }
        return result;
    }
};
