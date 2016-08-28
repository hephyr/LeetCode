class Solution {
public:
    int myAtoi(string str) {
        #define INT_MAX 2147483647
        #define INT_MIN (-INT_MAX - 1)
        if (str.empty()) {
            return 0;
        }
        while(str[0] == ' ') {
            str = str.substr(1);
        }
        string nums = "0123456789";
        int mul = 1;
        int result = 0;
        bool sym_error = false;
        for (const auto &c : str) {
            if (nums.find(c) != string::npos) {
                int n = c - '0';
                if ((long(result)*10 + n) > INT_MAX && mul == 1) {
                    return INT_MAX;
                } else if (-(long(result) * 10 + n) < INT_MIN) {
                    return INT_MIN;
                }
                result *= 10;
                result += n;
            } else if (sym_error) {
                break;
            } else if (c == '+') {
                sym_error = true;
            } else if (c == '-') {
                mul = -1;
                sym_error = true;
            } else {
                break;
            }
        }
        result *= mul;
        return result;
    }
};