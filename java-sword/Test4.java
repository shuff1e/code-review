public class Test4 
{
	public static boolean find(int[][] arr,int x)
	{
		if (arr==null)
		{
			return false;
		}
		boolean flag=false;
		int row=0;
		int col=arr[0].length-1;
		while (row<arr.length&&col>=0)
		{
			if (arr[row][col]==x)
			{
				flag=true;
				break;
			}
			else if (arr[row][col]>x)
			{
				col--;
			}
			else
				row++;
		}
		return flag;
	}
	public static void main(String[] args) 
	{
		int[][] matrix=new int[][]{new int[]{1,2,8,9},new int[]{2,4,9,12},new int[]{4,7,10,13},new int[]{6,8,11,15}};
		System.out.println(find(matrix,16));
		String str="hello";
		String str2=str.toUpperCase();
		System.out.println(str);
		System.out.println(str2);
	}
}
