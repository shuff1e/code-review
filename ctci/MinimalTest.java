import java.util.*;
public class MinimalTest
{
	TreeNode createMinimalBST(int[] arr,int start,int end)
	{
		if (end<start)
		{
			return null;
		}
		int mid=start+(end-start)/2;
		TreeNode root=new TreeNode(arr[mid]);
		root.left=createMinimalBST(arr,start,mid-1);
		root.right=createMinimalBST(arr,mid+1,end);
		return root;
	}
	ArrayList<LinkedList<TreeNode>> create(TreeNode root)
	{
		ArrayList<LinkedList<TreeNode>> lists=new ArrayList<>();
		createLevel(root,lists,0);
		return lists;
	}
	void createLevel(TreeNode root,ArrayList<LinkedList<TreeNode>> lists,int level)
	{
		if (root==null)
		{
			return ;
		}
		LinkedList<TreeNode> list=null;
		if (lists.size()==level)
		{
			list=new LinkedList<TreeNode>();
			lists.add(list);
		}
		else
		{
			list=lists.get(level);
		}
		list.add(root);
		createLevel(root.left,lists,level+1);
		createLevel(root.right,lists,level+1);
	}
	ArrayList<LinkedList<TreeNode>> create2(TreeNode root)
	{
		ArrayList<LinkedList<TreeNode>> lists=new ArrayList<>();
		LinkedList<TreeNode> cur=new LinkedList<>();
		if (root!=null)
		{
			cur.add(root);
		}
		while (cur.size()>0)
		{
			lists.add(cur);
			LinkedList<TreeNode> parents=cur;
			cur=new LinkedList<TreeNode>();
			for (TreeNode parent:parents )
			{
				if (parent.left!=null)
				{
					cur.add(parent.left);
				}
				if (parent.right!=null)
				{
					cur.add(parent.right);
				}
			}
		}
		return lists;
	}
	/*
	public static int last_printed=Integer.MIN_VALUE;
	public static boolean checkBST(TreeNode n)
	{
		if (n==null)
		{
			return true;
		}
		if (!checkBST(n.left))
		{
			return false;
		}
		if (n.data<last_printed)
		{
			return false;
		}
		last_printed=n.data;
		if (!checkBST(n.right))
		{
			return false;
		}
		return true;
	}
	
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
	public TreeNode inorderSucc(TreeNode n)
	{
		if (n==null)
		{
			return null;
		}
		if (n.right!=null)
		{
			return leftMostChild(n.right)
		}
		else
		{
			TreeNode q=n;
			TreeNode x=q.parent;
			while (x!=null&&x.left!=q)
			{
				q=x;
				x=x.parent;
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
		boolean is_p_on_left=covers(root.left,p);
		boolean is_q_on_left=covers(root.left,q);
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
		return commonAncestorHelper(root,p,q);
	}
	
	TreeNode commonAncestorBad(TreeNode root,TreeNode p,TreeNode q)
	{
		if (root==null)
		{
			return null;
		}
		if (root==p&&root==q)
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
	*/
	public static class Result
	{
		public TreeNode node;
		public boolean isAncestor;
		public Result(TreeNode node,boolean isAncestor)
		{
			this.node=node;
			this.isAncestor=isAncestor;
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
		if (rx.isAncestor)
		{
			return rx;
		}
		Result ry=commonAncestorHelper(root.right,p,q);
		if (ry.isAncestor)
		{
			return ry;
		}
		if (rx.node!=null&&ry.node!=null)
		{
			return new Result(root,true);
		}
		else if (root==p||root==q)
		{
			boolean isAncestor=rx.node!=null||ry.node!=null?true:false;
			return new Result(root,isAncestor);
		}
		else
			return new Result(rx.node!=null?rx.node:ry.node,false);
	}
	TreeNode commonAncestor(TreeNode root,TreeNode p,TreeNode q)
	{
		Result r=commonAncestorHelper(root,p,q);
		if (r.isAncestor)
		{
			return r.node;
		}
		return null;
	}
	boolean matchTree(TreeNode r1,TreeNode r2)
	{
		if (r2==null&&r1==null)
		{
			return true;
		}
		if (r1==null||r2==null)
		{
			return false;
		}
		if (r1.data!=r2.data)
		{
			return false;
		}
		return matchTree(r1.left,r2.left)&&matchTree(r1.right,r2.right);
	}
	boolean subTree(TreeNode r1,TreeNode r2)
	{
		if (r1==null)
		{
			return false;
		}
		if (r1.data==r2.data)
		{
			if (matchTree(r1,r2))
			{
				return true;
			}
		}
		return subTree(r1.left,r2)||subTree(r1.right,r2);
	}
	boolean contains(TreeNode r1,TreeNode r2)
	{
		if (r2==null)
		{
			return true;
		}
		return subTree(r1,r2);
	}
	public void findSum(TreeNode node,int sum,int[] path,int level)
	{
		if (node==null)
		{
			return;
		}
		path[level]=node.data;
		int t=0;
		for (int i=level;i>=0 ;i-- )
		{ 
			t+=path[i];
			if (t==sum)
			{
				print(path,i,level);
			}
		}
		findSum(node.left,sum,path,level+1);
		findSum(node.right,sum,path,level+1);
		path[level]=Integer.MAX_VALUE;
	}
	public void findSum(TreeNode node,int sum)
	{
		int depth=depth(node);
		int[] path=new int[depth];
		findSum(node,sum,path,0);
	}
	public int depth(TreeNode node)
	{
		if (node==null)
		{
			return 0;
		}
		else
			return 1+Math.max(depth(node.left),depth(node.right));
	}
	public static void print(int[] path,int start,int end)
	{
		for (int i=start;i<=end ;i++ )
		{
			System.out.print(path[i]+" ");
		}
		System.out.println();
	}
	public static void main(String[] args) 
	{
		System.out.println("Hello World!");
	}
}
