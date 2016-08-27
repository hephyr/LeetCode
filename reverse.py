class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        #define INT_MAX     2147483647
        #define INT_MIN     (-INT_MAX - 1)
        """
        temp = str(x)
        if '-' in temp:
            temp = '-' + temp[:0:-1]
        else:
            temp = temp[::-1]
        result = int(temp)
        if result > 2147483647 or result < -2147483648:
            return 0
        else :
            return result