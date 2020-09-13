import java.util.*;
public class Test7 
{
	public static BinaryTree create(int[] preArr,int[] midArr)
	{
		if (preArr==null||midArr==null||preArr.length==0||midArr.length==0)
		{
			return null;
		}
		return createHelper(preArr,0,preArr.length-1,midArr,0,midArr.length-1);
	}
	public static BinaryTree 
		createHelper(int[] preArr,int preStart,int preEnd,int[] midArr,int midStart,int midEnd)
	{
		BinaryTree root=new BinaryTree();
		//根节点就是前序遍历中开头的节点
		root.data=preArr[preStart];
		if (preStart==preEnd)
		{
			if (midStart==midEnd&&preArr[preStart]==midArr[midStart])
			{
				return root;
			}
			else
			{
				throw new RuntimeException("input valid");
			}
		}
		//找到中序遍历中根节点的位置:temp
		int temp=midStart;
		while (temp<=midEnd&&midArr[temp]!=preArr[preStart])
		{
			temp++;
		}
		if (temp>midEnd)
		{
			throw new RuntimeException("input valid");
		}
		int leftLen=temp-midStart;
		if (leftLen>0)
		{
			root.left=createHelper(preArr,preStart+1,preStart+leftLen,midArr,midStart,temp-1);
		}
		if (leftLen<midEnd-midStart)
		{
			root.right=createHelper(preArr,preStart+leftLen+1,preEnd,midArr,temp+1,midEnd);
		}
			return root;
	}
	public static void DFS(BinaryTree root)
	{
		if (root==null)
		{
			return;
		}
		System.out.print(root.data+" ");
		DFS(root.left);
		DFS(root.right);
	}
	public static void BFS(BinaryTree root)
	{
		if (root==null)
		{
			return;
		}
		Deque<BinaryTree> queue=new ArrayDeque<>();
		queue.offer(root);
		while (!queue.isEmpty())
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
	public static void main(String[] args) 
	{
		int[] preArr={1,2,4,7,3,5,6,8};
		int[] midArr={4,7,2,1,5,3,8,6};
		BinaryTree t=create(preArr,midArr);
		t=new BinaryTree(4);
		//DFS(t);
		System.out.println();
		//DFSNoRecursive(t);
		System.out.println();
		//BFS(t);
		BinaryTreeTraversing.pre(t);
		System.out.println();
		BinaryTreeTraversing.mid(t);
		System.out.println();
		BinaryTreeTraversing.post2(t);
	}
	public static void DFSNoRecursive(BinaryTree root)
	{
		Deque<BinaryTree> stack=new ArrayDeque<>();
		stack.push(root);
		System.out.print(root.data+" ");
		root.visited=true;
		while (!stack.isEmpty())
		{
			BinaryTree temp=stack.peek();
			BinaryTree helper=DFSNoRecursiveHelper(temp);
			if (helper!=null)
			{
				System.out.print(helper.data+" ");
				helper.visited=true;
				stack.push(helper);
			}
			else
				stack.pop();
		}
	}
	public static BinaryTree DFSNoRecursiveHelper(BinaryTree root)
	{
		if (root.left!=null&&root.left.visited==false)
		{
			return root.left;
		}
		else if (root.right!=null&&root.right.visited==false)
		{
			return root.right;
		}
		return null;
	}
}
