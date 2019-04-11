#include <vector>
#include <iostream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
//    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };

class Solution {
public:
    int minDepth(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }
        vector<TreeNode*> temp;
        temp.push_back(root);
        int d = 0;
        while (temp.size() > 0) {
            d++;
            float size = temp.size();
            for(int i=0; i<size; i++) {
                if (temp[0]->left == NULL && temp[0]->right == NULL) {
                    return d;
                }
                if (temp[0]->left != NULL) {
                    temp.push_back(temp[0]->left);
                }
                if (temp[0]->right != NULL) {
                    temp.push_back(temp[0]->right);
                }
                temp.erase(temp.begin());
            }
        }
        return d;
    }
};

int main() {

    TreeNode a = {2, NULL, NULL};
//    TreeNode b = {3, NULL, NULL};
//    TreeNode b = {3, NULL, NULL};
//    TreeNode b = {3, NULL, NULL};
    TreeNode root = {1, &a, NULL};

    Solution s;
    int d = s.minDepth(&root);
    cout<< d <<endl;
    return 0;
}