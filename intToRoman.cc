#include <iostream>
#include <string>

using namespace std;

string getRoman(int num, string a, string b, string c) {
    switch(num) {
        case 0: return ""; break;
        case 1: return a; break;
        case 2: return a+a; break;
        case 3: return a+a+a; break;
        case 4: return a+b; break;
        case 5: return b; break;
        case 6: return b+a; break;
        case 7: return b+a+a; break;
        case 8: return b+a+a+a; break;
        case 9: return a+c; break;
    }
    return "";
}
string intToRoman(int num) {
    string result;
    int o = num % 10;
    num /= 10;
    int oo = num % 10;
    num /= 10;
    int h = num % 10;
    num /= 10;
    int t = num % 10;
    string temp;
    temp = getRoman(t, "M", "", "");
    result += temp;
    temp = getRoman(h, "C", "D", "M");
    result += temp;
    temp = getRoman(oo, "X", "L", "C");
    result += temp;
    temp = getRoman(o, "I", "V", "X");
    result += temp;
    return result;
}

int main() {
    cout << intToRoman(999) << endl;
    return 0;
}