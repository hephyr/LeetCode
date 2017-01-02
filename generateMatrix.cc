class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        vector<vector<int>> result(n, vector<int>(n, 0));
        int top = 0, bottom = n-1, left = 0, right = n-1;
        int num = 1;
        int end = n * n;
        while (num <= end) {
            for (int i = left; i <= right; ++i) {
                result[top][i] = num;
                ++num;
            }
            ++top;
            for (int i = top; i <= bottom; ++i) {
                result[i][right] = num;
                ++num;
            }
            --right;
            for (int i = right; i >= left; --i) {
                result[bottom][i] = num;
                ++num;
            }
            --bottom;
            for (int i = bottom; i >= top; --i) {
                result[i][left] = num;
                ++num;
            }
            ++left;
        }
        return result;
    }
};
