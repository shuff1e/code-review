public class BalanceTest 
{
	public static int getHeight(TreeNode root)
	{
		if (root==null)
		{
			return 0;
		}
		return Math.max(getHeight(root.left),getHeight(root.right))+1;
	}
	public static boolean isBalanced(TreeNode root)
	{
		int flag=Math.abs(getHeight(root.left)-getHeight(root.right));
		if (flag>1)
		{
			return false;
		}
		return isBalanced(root.left)&&isBalanced(root.right);
	}
	public static void main(String[] args) 
	{
		System.out.println("Hello World!");
	}
}
