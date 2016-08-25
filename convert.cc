class Solution {
public:
    string convert(string s, int numRows) {
        if (numRows >= s.size() || numRows == 1)    return s;
        vector<string> temp(numRows);
        int row = 0;
        int step;
        for (auto c : s) {
            if (row == 0)   step = 1;
            if (row == numRows - 1)    step = -1;
            temp[row] += c;
            row += step;
        }
        string str;
        for (auto a : temp) {
            str += a;
        }
        return str;
    }
};