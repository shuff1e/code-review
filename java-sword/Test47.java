public class Test47 
{
	public static int getSum(int[][] arr,int row,int col)
	{
		if (row>=arr.length||col>=arr[0].length)
		{
			return 0;
		}
		int num1=arr[row][col]+getSum(arr,row,col+1);
		int num2=arr[row][col]+getSum(arr,row+1,col);
		return num1>num2?num1:num2;
	}
	public static int getSum(int[][] arr)
	{
		if (arr==null||arr.length==0||arr[0].length==0)
		{
			return 0;
		}
		return getSum(arr,0,0);
	}
	public static void main(String[] args) 
	{
		int[][] matrix=new int[1][4];
		matrix[0]=new int[]{1,10,3,8};
		/*
		matrix[1]=new int[]{12,2,9,6};
		matrix[2]=new int[]{5,7,4,11};
		matrix[3]=new int[]{3,7,16,5};
		*/
		matrix=null;
		System.out.println(getSum(matrix));
		System.out.println(getMaxValue(matrix));
		System.out.println(getMaxValueBetter(matrix));
	}
	public static int getMaxValue(int[][] arr)
	{
		if (arr==null||arr.length==0||arr[0].length==0)
		{
			return 0;
		}
		int row=arr.length;
		int col=arr[0].length;
		int[][] helper=new int[row][col];
		for (int i=0;i<row ;i++ )
		{
			for (int j=0;j<col ;j++ )
			{
				int num1=0;
				int num2=0;
				if (i>0)
				{
					num1=helper[i-1][j];
				}
				if (j>0)
				{
					num2=helper[i][j-1];
				}
				helper[i][j]=arr[i][j]+Math.max(num1,num2);
			}
		}
		return helper[row-1][col-1];
	}
	public static int getMaxValueBetter(int[][] arr)
	{
		if (arr==null||arr.length==0||arr[0].length==0)
		{
			return 0;
		}
		int row=arr.length;
		int col=arr[0].length;
		int[] helper=new int[col];
		for (int i=0;i<row ;i++ )
		{
			for (int j=0;j<col ;j++ )
			{
				int num1=0;
				int num2=0;
				if (i>0)
				{
					num1=helper[j];
				}
				if (j>0)
				{
					num2=helper[j-1];
				}
				helper[j]=arr[i][j]+Math.max(num1,num2);
			}
		}
		return helper[col-1];
	} 
}
