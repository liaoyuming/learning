#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int a = 0;
        int b = 0;
        int i = 0;
        while (a < m && b < n) {
            if (nums1[i] > nums2[b]) {
                nums1.insert(nums1.begin()+i, nums2[b]);
                nums1.pop_back();
                b++;
            } else {
                a++;
            }
            i++;
        }
        while (b < n) {
            nums1[i] = nums2[b];
            b++;
            i++;
        }
    }
};

int main()
{
    Solution s;
    vector<int> num1 = {1, 9, 0, 0,0};
    vector<int> num2 = {2,5,6};
    s.merge(num1, 1, num2, 3);
    for (auto i=0; i<num1.size(); i++) {
        cout << num1[i] << endl;
    }
}