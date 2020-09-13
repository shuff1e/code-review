import java.util.*;
public class Test32
{
	public static void print1(BinaryTree root)
	{
		if (root==null)
		{
			return ;
		}
		Deque<BinaryTree> queue=new ArrayDeque<>();
		queue.offer(root);
		while (queue.size()!=0)
		{
			BinaryTree temp=queue.poll();
			System.out.print(temp.data+" ");
			if (temp.left!=null)
			{
				queue.offer(temp.left);
			}
			if (temp.right!=null)
			{
				queue.offer(temp.right);
			}
		}
	}
	public static void print2(BinaryTree root)
	{
		if (root==null)
		{
			return ;
		}
		int toBePrinted=1;
		int nextLevel=0;
		Deque<BinaryTree> queue=new ArrayDeque<>();
		queue.offer(root);
		while (queue.size()!=0)
		{
			BinaryTree temp=queue.poll();
			System.out.print(temp.data+" ");
			toBePrinted--;
			if (temp.left!=null)
			{
				queue.offer(temp.left);
				nextLevel++;
			}
			if (temp.right!=null)
			{
				queue.offer(temp.right);
				nextLevel++;
			}
			if (toBePrinted==0)
			{
				System.out.println();
				toBePrinted=nextLevel;
				nextLevel=0;
			}
		}
	}
	public static void print3(BinaryTree root)
	{
		if (root==null)
		{
			return ;
		}
		int toBePrinted=1;
		int nextLevel=0;
		boolean leftToRight=true;
		Deque<BinaryTree> queue=new ArrayDeque<>();
		queue.offer(root);
		while (queue.size()!=0)
		{
			BinaryTree temp=queue.poll();
			System.out.print(temp.data+" ");
			toBePrinted--;
			if (!leftToRight)
			{
				if (temp.left!=null)
			    {
				queue.offer(temp.left);
				nextLevel++;
			    }
			    if (temp.right!=null)
			    {
				queue.offer(temp.right);
				nextLevel++;
			    }
			}
			else
			{
				if (temp.right!=null)
			    {
				queue.offer(temp.right);
				nextLevel++;
			    }
				if (temp.left!=null)
			    {
				queue.offer(temp.left);
				nextLevel++;
			    }
			}
			if (toBePrinted==0)
			{
				System.out.println();
				toBePrinted=nextLevel;
				nextLevel=0;
				leftToRight=!leftToRight;
			}
		}
	}
	public static void print4(BinaryTree root)
	{
		if (root==null)
		{
			return ;
		}
		Deque[] level=new ArrayDeque[2];
		level[0]=new ArrayDeque<Integer>();
		level[1]=new ArrayDeque<Integer>();
		int cur=0;
		int next=1;
		level[0].push(root);
		while (!(level[cur].isEmpty()&&level[next].isEmpty()))
		{
			BinaryTree temp=(BinaryTree)level[cur].pop();
			System.out.print(temp.data+" ");
			if (cur==0)
			{
				if (temp.left!=null)
			    {
				level[next].push(temp.left);
			    }
			    if (temp.right!=null)
			    {
				level[next].push(temp.right);
			    }
			}
			else
			{
				if (temp.right!=null)
			    {
				level[next].push(temp.right);
			    }
				if (temp.left!=null)
			    {
				level[next].push(temp.left);
			    }
			}
			if (level[cur].isEmpty())
			{
				System.out.println();
				cur=1-cur;
				next=1-next;
			}
		}
	}
	public static void main(String[] args) 
	{
		BinaryTree node1=new BinaryTree(8);
		/*
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
		//print1(node1);
		System.out.println();
		//print2(node1);
		*/
		print3(node1);
		print4(node1);
	}
}
