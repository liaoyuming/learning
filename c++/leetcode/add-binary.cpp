#include <iostream>

using namespace std;

class Solution {
public:
    string addBinary(string a, string b) {
        string sum;
        float al = a.length();
        float bl = b.length();
        int t = 0;
        for (int i=0; i<al || i<bl || t>0; i++) {
            int k = 0;
            if (i < al) {
                k += a[al-i-1] - '0';
            }
            if (i < bl) {
                k += b[bl-i-1] - '0';
            }
            k += t;
            t = k / 2;
            k = k % 2;
            sum = char(k + '0') + sum;
        }
        return sum;
    }
};

int main()
{
    Solution s;
    string r = s.addBinary("111", "1");
    cout << r << endl;
    return 0;
}