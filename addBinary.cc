class Solution {
public:
    string addBinary(string a, string b) {
        string result;
        int temp = 0;
        while (a.size() > 0 || b.size() > 0) {
            int x, y;
            if (a.size() > 0) {
                x = a.back() - '0';
                a.pop_back();
            } else {
                x = 0;
            }
            if (b.size() > 0) {
                y = b.back() - '0';
                b.pop_back();
            } else {
                y = 0;
            }
            temp += (x + y);
            if (temp == 3) {
                temp = 1;
                result.insert(0, "1");
            } else if (temp == 2) {
                temp = 1;
                result.insert(0, "0");
            } else if (temp == 1) {
                temp = 0;
                result.insert(0, "1");
            } else if (temp == 0) {
                temp = 0;
                result.insert(0, "0");
            }
        }
        if (temp == 1) {
            result.insert(0, "1");
        }
        return result;
    }
};
