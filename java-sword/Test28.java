public class Test28 
{
	public static boolean isSymmetrical(BinaryTree root)
	{
		return isSymmetrical(root,root);
	}
	public static boolean isSymmetrical(BinaryTree root1,BinaryTree root2)
	{
		if (root1==null&&root2==null)
		{
			return true;
		}
		if (root1==null||root2==null)
		{
			return false;
		}
		if (root1.data!=root2.data)
		{
			return false;
		}
		return isSymmetrical(root1.left,root2.right)
			&&isSymmetrical(root1.right,root2.left);
	}
	public static void main(String[] args) 
	{
		BinaryTree node1=new BinaryTree(8);
		BinaryTree left=new BinaryTree(8);
		BinaryTree right=new BinaryTree(8);
		node1.left=left;
		node1.right=right;
		/*
		BinaryTree temp=new BinaryTree(9);
		left.left=temp;
		temp=new BinaryTree(2);
		left.right=temp;
		left=new BinaryTree(4);
		right=new BinaryTree(7);
		temp.left=left;
		temp.right=right;
		*/
		System.out.println(isSymmetrical(node1));
	}
}
