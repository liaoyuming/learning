#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int lengthOfLastWord(string s) {
        if (s.length() == 0) {
            return 0;
        }
        string res;
        for(int i = s.length()-1; i>=0; i--) {
            if (s[i] != ' ') {
                res += s[i];
            } else if (res != "") {
                break;
            }
        };
        return res.length();
    }
};

int main()
{
    Solution s;
    cout << s.lengthOfLastWord("adffdd saa") << endl;
}