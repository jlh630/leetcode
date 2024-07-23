/*
 * @lc app=leetcode.cn id=79 lang=java
 * @lcpr version=30204
 *
 * [79] 单词搜索
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
class Solution {
    public boolean exist(char[][] board, String word) {
        if (board.length*board[0].length<word.length()) {
            return false;
        }
        char start=word.charAt(0);
        for(int i=0;i<board.length;i++){

            for(int j=0;j<board[i].length;j++){
                if (start==board[i][j]) {
                    boolean [][] road=new boolean[board.length][board[0].length];
                     if(search(road,board, i, j, word, 0)){
                        return true;
                     }
                }
            }
        }
        return false;
    }

    public boolean search(boolean[][] road,char[][] board,int i,int j,String word,int index){
        
        if (board[i][j] == word.charAt(index)) {
            road[i][j]=true;

            if (index==word.length()-1) {
                return true;
            }
            //上
        if (i!=0&&road[i-1][j]!=true) {
            if (search(road,board, i-1, j, word, index+1)) {
                return true;
            }
             
        }
        //下
        if (i!=board.length-1&&road[i+1][j]!=true) {
            if ( search(road,board, i+1, j, word, index+1)) {
                return true;
            } 
        }
        //左
        if (j!=0&&road[i][j-1]!=true){
            if( search(road,board, i, j-1, word, index+1)){
                return true;
            }
        }
        //右
        if (j!=board[0].length-1&&road[i][j+1]!=true) {
            if (search(road,board, i, j+1, word, index+1)) {
                return true;
            }
            
        }
        road[i][j]=false;
        }else{
            return false;
        }
        return false;
    }
}
// @lc code=end



/*
// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"SEE"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCB"\n
// @lcpr case=end

 */

