/**
 * Definition for an interval.
 * struct Interval {
 *     int start;
 *     int end;
 *     Interval() : start(0), end(0) {}
 *     Interval(int s, int e) : start(s), end(e) {}
 * };
 */
class Solution {
public:
    vector<Interval> merge(vector<Interval>& intervals) {
        int n = intervals.size();
        vector<Interval> result;
        vector<int> temp(n, 1);
        for (int i = 0; i < n; ++i) {
            if (temp[i] == 0)  continue;
            for (int j = i+1; j < n; ++j) {
                if (temp[j] == 0)  continue;
                if ((intervals[i].end >= intervals[j].start && intervals[i].start <= intervals[j].start)||
                    (intervals[j].end >= intervals[i].start && intervals[j].start <= intervals[i].start)) {
                    intervals[i].start = (intervals[i].start < intervals[j].start) ? intervals[i].start : intervals[j].start;
                    intervals[i].end = (intervals[i].end > intervals[j].end) ? intervals[i].end : intervals[j].end;
                    temp[j] = 0;
                    j = i;
                }
            }
        }
        for (int i = 0; i < n; ++i) {
            if (temp[i] == 0)  continue;
            result.push_back(intervals[i]);
        }
        return result;
    }
};
