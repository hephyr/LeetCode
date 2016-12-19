class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int size = matrix.size();
        vector<vector<int>> v(matrix.begin(), matrix.end());
        for (int i = 0; i < size; ++i) {
            for (int j = 0; j < size; ++j) {
                v[j][size-i-1] = matrix[i][j];
            }
        }
        matrix = v;
    }
};
