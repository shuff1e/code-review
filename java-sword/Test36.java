public class Test36 
{
	public static BinaryTree convertHelper(BinaryTree root,boolean isLeft)
	{
		if (root.left!=null)
		{
			BinaryTree leftNode=convertHelper(root.left,true);
			root.left=leftNode;
			leftNode.right=root;
		}
		if (root.right!=null)
		{
			BinaryTree rightNode=convertHelper(root.right,false);
			root.right=rightNode;
			rightNode.left=root;
		}
		return isLeft?mostRight(root):mostLeft(root);
	}
	public static BinaryTree mostLeft(BinaryTree root)
	{
		while (root.left!=null)
		{
			root=root.left;
		}
		return root;
	}
	public static BinaryTree mostRight(BinaryTree root)
	{
		while (root.right!=null)
		{
			root=root.right;
		}
		return root;
	}
	public static BinaryTree convert(BinaryTree root)
	{
		if (root==null)
		{
			return null;
		}
		return convertHelper(root,false);
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
		////////////////////////
		node1=new BinaryTree(8);
		BinaryTree head=convert(node1);
		BinaryTree helper=null;
		while (head!=null)
		{
			System.out.print(head.data+"->");
			helper=head;
			head=head.right;
		}
		System.out.println();
		head=helper;
		while (head!=null)
		{
			System.out.print(head.data+"->");
			head=head.left;
		}	
	}
}
