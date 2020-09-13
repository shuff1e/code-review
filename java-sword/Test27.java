import java.util.*;
public class Test27 
{
	public static void mirrorRecursive(BinaryTree node,BinaryTree left,BinaryTree right)
	{
		if (left==null&&right==null)
		{
			return ;
		}
		if (node.left!=null)
		{
			mirrorRecursive(node.left,node.left.left,node.left.right);
		}
		if (node.right!=null)
		{
			mirrorRecursive(node.right,node.right.left,node.right.right);
		}
		BinaryTree temp=node.left;
		node.left=node.right;
		node.right=temp;
	}
	public static void mirror(BinaryTree node)
	{
		mirrorRecursive(node,node.left,node.right);
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
		mirror(node1);
		BinaryTreeTraversing.DLR(node1);
		System.out.println();
		BinaryTreeTraversing.LDR(node1);
	}
}
