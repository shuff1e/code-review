public class Test29 
{
	public static void print(int[][] matrix)
	{
		if (matrix==null)
		{
			return;
		}
		int rowNum=matrix.length;
		int colNum=matrix[0].length;
		if (rowNum==0||colNum==0)
		{
			return;
		}
		int layer=Math.min(rowNum>>1,colNum>>1);
		for (int i=0;i<layer ;i++)
		{
			printHelper(matrix,rowNum,colNum,i);
		}
		if (((colNum&1)==1)&&(colNum<rowNum))
		{
			for (int i=layer;i<rowNum-layer ;i++ )
			{
				System.out.print(matrix[i][layer]+" ");
			}
			System.out.println();
		}
		if (((rowNum&1)==1)&&(rowNum<colNum))
		{
			for (int i=layer;i<colNum-layer ;i++ )
			{
				System.out.print(matrix[layer][i]+" ");
			}
			System.out.println();
		}
	}
	public static void printHelper(int[][] matrix,int rowNum,int colNum,int layer)
	{
		for (int i=layer;i<colNum-layer-1 ;i++ )
		{
			System.out.print(matrix[layer][i]+" ");
		}
		for (int i=layer;i<rowNum-layer-1 ;i++ )
		{
			System.out.print(matrix[i][colNum-1-layer]+" ");
		}
		for (int i=colNum-1-layer;i>layer ;i-- )
		{
			System.out.print(matrix[rowNum-1-layer][i]+" ");
		}
		for (int i=rowNum-1-layer;i>layer ;i-- )
		{
			System.out.print(matrix[i][layer]+" ");
		}
		System.out.println();
	}
	public static void print2(int[][] matrix)
	{
		if (matrix==null)
		{
			return ;
		}
		int rowNum=matrix.length;
		int colNum=matrix[0].length;
		if (rowNum==0||colNum==0)
		{
			return;
		}
		int start=0;
		while ((colNum>2*start)&&(rowNum>2*start))
		{
			print2Helper(matrix,rowNum,colNum,start);
			start++;
		}
	}
	public static void print2Helper(int[][] matrix,int rowNum,int colNum,int start)
	{
		int endX=colNum-1-start;
		int endY=rowNum-1-start;
		for (int i=start;i<=endX ;i++ )
		{
			System.out.println(matrix[start][i]);
		}
		if (start<endY)
		{
			for (int i=start+1;i<=endY ;i++ )
		    {
			     System.out.println(matrix[i][endX]);
		    }
		}
		if (start<endY&&start<endX)
		{
			for (int i=endX-1;i>=start ;i-- )
		    {
			     System.out.println(matrix[endY][i]);
		    }
		}
		if ((start<(endY-1))&&start<endX)
		{
			for (int i=endY-1;i>=start+1 ;i-- )
		    {
			     System.out.println(matrix[i][start]);
		    }
		}
	}
	public static void main(String[] args) 
	{
		int[][] matrix=new int[5][3];
		matrix[0]=new int[]{1,2,3};
		matrix[1]=new int[]{5,6,7};
		matrix[2]=new int[]{9,10,11};
		matrix[3]=new int[]{13,14,15};
		matrix[4]=new int[]{17,18,19};
		print(matrix);
		print2(matrix);
	}
}
