## N-Queens problem

The N-Queens problem is a classic problem in which the objective is to place N queens on an N×N chessboard in such a way that no two queens are attacking each other. A queen can attack horizontally, vertically, and diagonally.

The solution to the N-Queens problem is to find all the possible arrangements of N queens on the chessboard such that no two queens are attacking each other. This can be done using backtracking. The algorithm starts with an empty chessboard and tries to place a queen in each column, one by one. If a queen can be placed in a column, the algorithm moves to the next column. If no queen can be placed in a column, the algorithm backtracks and tries to place the queen in the previous column again. If all the queens are placed successfully, the algorithm prints the solution. If the algorithm backtracks to the first column, it means that there is no solution.

