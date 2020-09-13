import java.util.*;
public class Test34 
{
	public static void findHelper(BinaryTree root,int k,int sum,
		ArrayList<BinaryTree> list)
	{
		list.add(root);
		if (root==null)
		{
			return;
		}
		sum+=root.data;
		if (sum==k)
		{
			print(list);
		}
		findHelper(root.left,k,sum,list);
		list.remove(list.size()-1);
		findHelper(root.right,k,sum,list);
		list.remove(list.size()-1);
	}
	public static void print(ArrayList<BinaryTree> list)
	{
		for (int i=0;i<list.size() ;i++ )
		{
			System.out.print(list.get(i).data+"->");
		}
		System.out.println();
	}
	public static void find(BinaryTree root,int k)
	{
		if (root==null)
		{
			return ;
		}
		findHelper(root,k,0,new ArrayList<BinaryTree>());
	}
	public static void main(String[] args) 
	{
		BinaryTree node1=new BinaryTree(8);
		BinaryTree left=new BinaryTree(8);
		BinaryTree right=new BinaryTree(10);
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
		find(node1,23);
		//System.out.println("Hello World!");
	}
}
