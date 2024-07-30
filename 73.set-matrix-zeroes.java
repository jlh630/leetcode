/*
 * @lc app=leetcode.cn id=73 lang=java
 * @lcpr version=30204
 *
 * [73] 矩阵置零
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    //哈希
    public void setZeroes(int[][] matrix) {
        int n=matrix.length;
        int m=matrix[0].length;
        boolean[] rows=new boolean[n];
        boolean[] columns=new boolean[m];
        for (int i=0;i<n;i++){
            for(int j=0;j<m;j++){
                if (matrix[i][j]==0) {
                    rows[i]=true;
                    columns[j]=true;
                }
            }
        }
        for (int i=0;i<n;i++){
            for(int j=0;j<m;j++){
                if (rows[i]||columns[j]) {
                    matrix[i][j]=0;
                } 
            }
        }
    }
}
// @lc code=end



/*
// @lcpr case=start
// [[1,1,1],[1,0,1],[1,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,2,0],[3,4,5,2],[1,3,1,5]]\n
// @lcpr case=end

 */

