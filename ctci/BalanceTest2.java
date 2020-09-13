public class BalanceTest2 
{
	public static int checkHeight(TreeNode root)
	{
		if (root==null)
		{
			return 0;
		}
		int left=checkHeight(root.left);
		if (left==-1)
		{
			return -1;
		}
		int right=checkHeight(root.right);
		if (right==-1)
		{
			return -1;
		}
		int diff=Math.abs(left-right);
		if (diff>1)
		{
			return -1;
		}
		else
			return Math.max(left,right)+1;
	}
	public static boolean isBalanced(TreeNode root)
	{
		if (checkHeight(root)==-1)
		{
			return false;
		}
		else
			return true;
	}
	public static void main(String[] args) 
	{
		System.out.println("Hello World!");
	}
}
