import java.util.*;
public class BinaryTreeTest 
{
	public void findSum(TreeNode root,int sum,int[] path,int level)
	{
		if (root==null)
		{
			return;
		}
		path[level]=root.data;
		int t=0;
		for (int i=level;i>=0 ;i-- )
		{
			t+=path[level];
			if (t==sum)
			{
				print(path,i,level);
			}
		}
		path[level]=Integer.MIN_VALUE;
	}
	public int depth(TreeNode root)
	{
		if (root==null)
		{
			return 0;
		}
		return Math.max(depth(root.left),depth(root.right))+1;
	}
	public void print(int[] path,int start,int end)
	{
		for (int i=start;i<=end ;i++ )
		{
			System.out.print(path[i]+" ");
		}
		System.out.println();
	}
	public void findSum(TreeNode root,int sum)
	{
		int depth=depth(root);
		int[] path=new int[depth];
		findSum(root,sum,path,0);
	}
	/*
	boolean containsTree(TreeNode t1,TreeNode t2)
	{
		if (t2==null)
		{
			return true;
		}
		return subTree(t1,t2);
	}
	boolean subTree(TreeNode t1,TreeNode t2)
	{
		if (t1==null)
		{
			return false;
		}
		if (t1.data==t2.data)
		{
			if (matchTree(t1,t2))
			{
				return true;
			}
		}
		return subTree(t1.left,t2)||subTree(t1.right,t2);
	}
	boolean matchTree(TreeNode t1,TreeNode t2)
	{
		if (t1==null&&t2==null)
		{
			return true;
		}
		if (t1==null||t2==null)
		{
			return false;
		}
		if (t1.data!=t2.data)
		{
			return false;
		}
		return matchTree(t1.left,t2.left)&&matchTree(t1.right,t2.right);
	}
	boolean covers(TreeNode root,TreeNode p)
	{
		if (root==null)
		{
			return false;
		}
		if (root==p)
		{
			return true;
		}
		return covers(root.left,p)||covers(root.right,p);
	}
	TreeNode commonAncestorHelper(TreeNode root,TreeNode p,TreeNode q)
	{
		if (root==null)
		{
			return null;
		}
		if (root==p||root==q)
		{
			return root;
		}
		boolean is_p_on_left=covers(root,p);
		boolean is_q_on_left=covers(root,q);
		if (is_p_on_left!=is_q_on_left)
		{
			return root;
		}
		TreeNode child=is_p_on_left?root.left:root.right;
		return commonAncestorHelper(child,p,q);
	}
	TreeNode commonAncestor(TreeNode root,TreeNode p,TreeNode q)
	{
		if (!covers(root,p)||!covers(root,q))
		{
			return null;
		}
		return commonAncestor(root,p,q);
	}
	TreeNode commonAncestorBad(TreeNode root,TreeNode p,TreeNode q)
	{
		if (root==null)
		{
			return null;
		}
		if (root==p||root==q)
		{
			return root;
		}
		TreeNode x=commonAncestorBad(root.left,p,q);
		if (x!=null&&x!=p&&x!=q)
		{
			return x;
		}
		TreeNode y=commonAncestorBad(root.right,p,q);
		if (y!=null&&y!=p&&y!=q)
		{
			return y;
		}
		if (x!=null&&y!=null)
		{
			return root;
		}
		else if (root==p||root==q)
		{
			return root;
		}
		else
			return x==null?y:x;
	}
	
	public static class Result
	{
		public TreeNode node;
		public boolean isAnc;
		public Result(TreeNode n,boolean is)
		{
			node=n;
			isAnc=is;
		}
	}
	Result commonAncestorHelper(TreeNode root,TreeNode p,TreeNode q)
	{
		if (root==null)
		{
			return new Result(null,false);
		}
		if (root==p&&root==q)
		{
			return new Result(root,true);
		}
		Result rx=commonAncestorHelper(root.left,p,q);
		if (rx.isAnc)
		{
			return rx;
		}
		Result ry=commonAncestorHelper(root.right,p,q);
		if (ry.isAnc)
		{
			return ry;
		}
		if (rx.node!=null&&ry.node!=null)
		{
			return new Result(root,true);
		}
		else if (root==p||root==q)
		{
			boolean flag=rx.node!=null||ry.node!=null?true:false;
			return new Result(root,flag);
		}
		else
			return new Result(rx.node!=null?rx.node:ry.node,false);
	}
	public TreeNode inorderSucc(TreeNode n)
	{
		if (n==null)
		{
			return null;
		}
		if (n.right!=null)
		{
			return leftMostChild(n.right);
		}
		else
		{
			TreeNode q=n;
			TreeNode x=q.parent;
			while (x!=null&&x.left!=q)
			{
				q=x;
				x=q.parent;
			}
			return x;
		}
	}
	public TreeNode leftMostChild(TreeNode n)
	{
		if (n==null)
		{
			return null;
		}
		while (n.left!=null)
		{
			n=n.left;
		}
		return n;
	}
	
	public static boolean checkBST(TreeNode root)
	{
		int[] array=new int[size(root)];
		copyBST(root,array);
		for (int i=1;i<array.length ;i++ )
		{
			if (array[i]<=array[i-1])
			{
				return false;
			}
		}
		return true;
	}
	public static int index=0;
	public static void copyBST(TreeNode root,int[] array)
	{
		if (root==null)
		{
			return ;
		}
		copyBST(root.left,array);
		array[index++]=root.data;
		copyBST(root.right,array);
	}
	public static int size(TreeNode root)
	{
		if (root==null)
		{
			return 0;
		}
		int size=1;
		Deque<TreeNode> queue=new ArrayDeque<>();
		queue.offer(root);
		while (queue.size()>0)
		{
			TreeNode node=queue.poll();
			if (node.left!=null)
			{
				queue.offer(node.left);
				size++;
			}
			if (node.right!=null)
			{
				queue.offer(node.right);
				size++;
			}
		}
		return size;
	}
	
	public static int last_printed=Integer.MIN_VALUE;
	public static boolean checkBST(TreeNode root)
	{
		if (root==null)
		{
			return true;
		}
		if (!checkBST(root.left))
		{
			return false;
		}
		if (root.data<last_printed)
		{
			return false;
		}
		last_printed=root.data;
		if (!checkBST(root.right))
		{
			return false;
		}
		return true;
	}
	*/
	boolean checkBST(TreeNode root)
	{
		return checkBST(root,Integer.MIN_VALUE,Integer.MAX_VALUE);
	}
	boolean checkBST(TreeNode root,int min,int max)
	{
		if (root==null)
		{
			return true;
		}
		if (root.data<min||root.data>max)
		{
			return false;
		}
		if (!checkBST(root.left,min,root.data)||!checkBST(root.right,root.data,max))
		{
			return false;
		}
		return true;
	}
}
