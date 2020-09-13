public class Test55 
{
	public static int depth = 0;
	// 初始的level为1
	public static int getTreeDepth2(BinaryTree root,int level) {
		if (root == null) {
			return level
		}
		depth = Math.max(depth,level);
		getTreeDepth2(root.left,level+1);
		getTreeDepth2(root.right,level+1)
	}

	public static int getTreeDepth(BinaryTree root)
	{
		if (root==null)
		{
			return 0;
		}
		int left=getTreeDepth(root.left);
		int right=getTreeDepth(root.right);
		return left>right?(left+1):(right+1);
	}
	public static void main(String[] args) 
	{
		BinaryTree node1=new BinaryTree(8);
		BinaryTree left=new BinaryTree(8);
		BinaryTree right=new BinaryTree(7);
		node1.left=left;
		node1.right=right;
		BinaryTree temp=new BinaryTree(9);
		left.left=temp;
		temp=new BinaryTree(2);
		left.right=temp;
		left=new BinaryTree(4);
		right=new BinaryTree(7);
		temp.left=left;
		temp.right=right;
		////////////////////////P148
		BinaryTree root=new BinaryTree(8);
		System.out.println(getTreeDepth(root));
		System.out.println(isBalanced(node1));
		System.out.println(isBalanced(node1,new Wrapper(0)));
	}
	public static boolean isBalanced(BinaryTree root)
	{
		if (root==null)
		{
			return true;
		}
		int left=getTreeDepth(root.left);
		int right=getTreeDepth(root.right);
		if ((int)Math.abs(left-right)>1)
		{
			return false;
		}
		return isBalanced(root.left)&&isBalanced(root.right);
	}
	// 每个子树返回给父子树两个信息，一个是自己是否是平衡的，另一个是自己的高度
	// 父子树再返回给自己的父子树，这2个信息
	// 后序遍历
	public static boolean isBalanced(BinaryTree root,Wrapper wra)
	{
		if (root==null)
		{
			wra.k=0;
			return true;
		}
		Wrapper left=new Wrapper(0);
		Wrapper right=new Wrapper(0);
		if (isBalanced(root.left,left)&&isBalanced(root.right,right))
		{
			int diff=(int)Math.abs(left.k-right.k);
			if (diff<=1)
			{
				wra.k=Math.max(left.k,right.k)+1;
				return true;
			}
		}
		return false;
	}
}
class Wrapper
{
	public int k;
	public Wrapper(int k)
	{
		this.k=k;
	}
}