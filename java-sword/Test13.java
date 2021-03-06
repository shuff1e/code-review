public class Test13 
{
	public static int move(int[][] matrix,int k)
	{
		if (matrix==null)
		{
			return 0;
		}
		int rowNum=matrix.length;
		int colNum=matrix[0].length;
		if (rowNum<=0||colNum<=0)
		{
			return 0;
		}
		boolean[][] visited=new boolean[rowNum][colNum];
		return moveHelper(matrix,0,0,k,visited);
	}
	public static int moveHelper(int[][] matrix,int row, int col,int k,boolean[][] visited)
	{
		int rowNum=matrix.length;
		int colNum=matrix[0].length;
		int count=0;
		if (0<=row&&row<rowNum&&0<=col&&col<colNum
			&&!visited[row][col]&&checkValid(row,col,k))
		{
			visited[row][col]=true;
			count=1+moveHelper(matrix,row-1,col,k,visited)
				+moveHelper(matrix,row+1,col,k,visited)
				+moveHelper(matrix,row,col-1,k,visited)
				+moveHelper(matrix,row,col+1,k,visited);
		}
		return count;
	}
	public static boolean checkValid(int row,int col,int k)
	{
		int count=0;
		while (row>0)
		{
			count+=row%10;
			row/=10;
		}
		while (col>0)
		{
			count+=col%10;
			col/=10;
		}
		if (count>k)
		{
			return false;
		}
		else
		{
			return true;
		}
	}
	public static void main(String[] args) 
	{
		int[][] matrix=new int[3][1];
		int k=3;
		System.out.println(move(matrix,k));
	}
}
