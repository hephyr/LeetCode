class Solution {
public:
    bool isValid(string s) {
        stack<char> characters;
        string ch1 = "({[";
        string ch2 = ")}]";
        map<char, char> m = {{'(', ')'}, {'{', '}'}, {'[', ']'}};
        for (auto c : s) {
            if (ch1.find(c) != std::string::npos) {
                characters.push(c);
            } else if (ch2.find(c) != std::string::npos) {
                if (characters.size() == 0) {
                    return false;
                }
                char i = characters.top();
                characters.pop();
                if (m[i] != c) {
                    return false;
                }
            }
        }
        if (characters.size() > 0) {
            return false;
        }
        return true;
    }
};