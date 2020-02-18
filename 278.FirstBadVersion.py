# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        if isBadVersion(1):
            return 1
        i = 1
        while i < n:
            k = (i + n) // 2
            if isBadVersion(k):
                if not isBadVersion(k - 1):
                    return k
                else:
                    n = k - 1
            else:
                if isBadVersion(k + 1):
                    return k + 1
                else:
                    i = k + 1