/*
 * @lc app=leetcode.cn id=54 lang=java
 * @lcpr version=30204
 *
 * [54] 螺旋矩阵
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start

import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res=new ArrayList<>();
        int maxLen= matrix.length * matrix[0].length;
        boolean[][] road=new boolean[matrix.length][matrix[0].length];
        road[0][0]=true;
        res.add(matrix[0][0]);
        int[] currij=new int[]{0,0};
        while (res.size()!=maxLen) {
        
        insetrt(matrix,res,road,currij[0],currij[1],0,1,currij);  
        insetrt(matrix,res,road,currij[0],currij[1],+1,0,currij);
        insetrt(matrix,res,road,currij[0],currij[1],0,-1,currij);
        insetrt(matrix,res,road,currij[0],currij[1],-1,0,currij);
        }

        return res;
    }

    public void insetrt(int[][] matrix, List<Integer> res,boolean[][] road,int curri,int currj,int addi,int addj,int[] currij){
        curri+=addi;
        currj+=addj;
        while (curri>=0&&curri<matrix.length&&currj>=0&&currj<matrix[0].length&&road[curri][currj]!=true ) {
            res.add(matrix[curri][currj]);
            road[curri][currj]=true;
            curri+=addi;
            currj+=addj;
        }
        currij[0]=curri-addi;
        currij[1]=currj-addj;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

 */

