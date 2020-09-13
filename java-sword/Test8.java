public class Test8 
{
	public static BinaryTree find(BinaryTree root)
	{
		if (root==null)
		{
			return null;
		}
		if (root.right!=null)
		{
			return leftMost(root.right);
		}
		else if (root.parent!=null&&root.parent.left==root)
		{
			return root.parent;
		}
		else
		{
			while (root.parent!=null&&root.parent.right==root)
			{
				root=root.parent;
			}
			return root.parent;
		}
	}
	public static BinaryTree leftMost(BinaryTree root)
	{
		while (root.left!=null)
		{
			root=root.left;
		}
		return root;
	}
	public static void main(String[] args) 
	{
		System.out.println("Hello World!");
	}
}
