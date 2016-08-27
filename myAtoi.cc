class Solution {
public:
    int myAtoi(string str) {
        #define INT_MAX 2147483647
        #define INT_MIN (-INT_MAX - 1)
        int n;
        int result = 0;
        int temp = 1;
        string num = " +-0123456789";
        bool symbol = true;
        bool space = true;
        for (auto &c : str) {
            if (c == '+' || c == '-') {
                space = false;
                if (!symbol)    return 0;
                symbol = false;
                if (c == '-') {
                    temp = -1;
                }
                c = '0';
            }
            
            if ((num.find(c) == string::npos))    break;
            else if (c == ' ') {
                if (space) {
                    continue;
                }
                else
                    break;
            }
            n = c - '0';
            if (long(result) * 10 + n > INT_MAX && temp == 1)    return INT_MAX;
            else if (-(long(result) * 10 + n) < INT_MIN && temp == -1)   return INT_MIN;
            result *= 10;
            result += n;
            space = false;
        }
        result *= temp;
        return result;
    }
};