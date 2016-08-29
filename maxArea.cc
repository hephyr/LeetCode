class Solution {
public:
    int maxArea(vector<int>& height) {
        int maxarea = 0;
        int temparea;
        int left = 0;
        int right = height.size() - 1;
        while (left < right) {
            temparea = (right - left) * (height[left] < height[right] ? height[left] : height[right]);
            if (temparea > maxarea) {
                maxarea = temparea;
            }
            
            if (height[left] < height[right]) {
                int temp = left;
                do {
                    ++left;
                } while (height[temp] >= height[left] && left < right);
            } else {
                int temp = right;
                do {
                    --right;
                } while (height[temp] >= height[right] && left < right);
            }
        }
        return maxarea;
    }
};