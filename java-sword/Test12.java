public class Test12
{
	public static boolean search(char[][] matrix,String str)
	{
		if (matrix==null)
		{
			return false;
		}
		int rowNum=matrix.length;
		int colNum=matrix[0].length;
		if (rowNum<=0||colNum<=0)
		{
			return false;
		}
		boolean[][] booMatrix=new boolean[rowNum][colNum];
		int[][] visited=new int[rowNum][colNum];
		for (int i=0;i<rowNum ;i++ )
		{
			for (int j=0;j<colNum ;j++ )
			{
				visited[i][j]=-1;
			}
		}
		return searchHelper(matrix,0,rowNum,0,rowNum,booMatrix,visited,str,0);
	}
	public static boolean searchHelper(char[][] matrix,
		int row,int rowNum,int col,int colNum,
		boolean[][] booMatrix,int[][] visited,String str,int index)
	{
		if (index==str.length())
		{
			return true;
		}
		boolean flag=false;
		if (row<rowNum&&row>=0&&col<colNum&&col>=0
			&&!booMatrix[row][col]
			&&index>visited[row][col])
		{
			visited[row][col]=index;
			if (str.charAt(index)==matrix[row][col])
		    {
			index++;
			booMatrix[row][col]=true;
			flag=searchHelper(matrix,row+1,rowNum,col,colNum,booMatrix,visited,str,index)
				||searchHelper(matrix,row-1,rowNum,col,colNum,booMatrix,visited,str,index)
				||searchHelper(matrix,row,rowNum,col+1,colNum,booMatrix,visited,str,index)
				||searchHelper(matrix,row,rowNum,col-1,colNum,booMatrix,visited,str,index);
		    }
			if (!flag)
		    {
			booMatrix[row][col]=false;
		    }
		}
		return flag;
	}
	public static void main(String[] args) 
	{
		char[][] matrix=new char[4][4];
		matrix[0]=new char[]{'a','b','t','g'};
		matrix[1]=new char[]{'c','f','c','s'};
		matrix[2]=new char[]{'j','d','e','h'};
		matrix[3]=new char[]{'f','s','x','a'};
		String str="abfcs";
		boolean flag=search(matrix,str);
		System.out.println(flag);
	}
}
