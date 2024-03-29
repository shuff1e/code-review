import java.util.*;
public class Test68 
{
	public static boolean getPath(BinaryTree root,BinaryTree target,Deque<BinaryTree> queue)
	{
		if (root==null||target==null)
		{
			return false;
		}
		queue.addLast(root);
		if (root.data==target.data)
		{
			return true;
		}
		boolean flag=getPath(root.left,target,queue);
		if (!flag)
		{
			flag=getPath(root.right,target,queue);
		}
		if (!flag)
		{
			queue.removeLast();
		}
		return flag;
	}
	public static BinaryTree findCommon(BinaryTree root,BinaryTree target1,BinaryTree target2)
	{
		Deque<BinaryTree> queue1=new ArrayDeque<>();
		Deque<BinaryTree> queue2=new ArrayDeque<>();
		getPath(root,target1,queue1);
		getPath(root,target2,queue2);
		BinaryTree result=null;
		while ((!queue1.isEmpty())&&(!queue2.isEmpty()))
		{
			if (queue1.peekFirst().data==queue2.peekFirst().data)
			{
				result=queue1.removeFirst();
				queue2.removeFirst();
			}
			else
			{
				queue1.removeFirst();
			    queue2.removeFirst();
			}
		}
		return result;
	}
	public static void main(String[] args) 
	{
		BinaryTree node1=new BinaryTree(8);
		BinaryTree left=new BinaryTree(10);
		BinaryTree right=new BinaryTree(11);
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
		System.out.println(findCommon(
			null,new BinaryTree(9),null).data);
	}
}
