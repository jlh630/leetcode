/*
 * @lc app=leetcode.cn id=48 lang=java
 * @lcpr version=30204
 *
 * [48] 旋转图像
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public void rotate(int[][] matrix) {
        int n=matrix.length;
        //上下镜像
        for(int i=0;i<n/2;i++){
            for(int j=0;j<n;j++){
                int tmp=matrix[i][j];
                matrix[i][j]=matrix[n-1-i][j];
                matrix[n-1-i][j]=tmp;
            }
        }
        //斜线镜像
        for(int i=0;i<n;i++){
            for(int j=0;j<i;j++){
                    if (i==j) {
                        continue;
                    }
                int tmp=matrix[i][j];
                matrix[i][j]=matrix[j][i];
                matrix[j][i]=tmp;
            }
        }
    }
}
// @lc code=end



/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]\n
// @lcpr case=end

 */

