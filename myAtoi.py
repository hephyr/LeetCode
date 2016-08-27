class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        INT_MAX = 2147483647
        INT_MIN = -2147483648
        str = str.strip()
        l = str.split()
        mul = 1
        result = 0
        error_sym = False
        for c in str:
            if c in string.digits:
                n = int(c)
                result *= 10
                result += n
            elif error_sym:
                break
            elif c == '+':
                error_sym = True
            elif c == '-':
                mul = -1
                error_sym = True
            else:
                break

        result *= mul
        if result > INT_MAX:
            return INT_MAX
        elif result < INT_MIN:
            return INT_MIN
        else:
            return result