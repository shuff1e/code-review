import java.util.*;
public class BinaryTreeTraversing 
{
	public static void DLR(BinaryTree root)
	{
		if (root!=null)
		{
			System.out.print(root.data+" ");
			DLR(root.left);
			DLR(root.right);
		}
	}
	public static void LDR(BinaryTree root)
	{
		if (root!=null)
		{
			LDR(root.left);
			System.out.print(root.data+" ");
			LDR(root.right);
		}
	}
	public static void main(String[] args) 
	{
		System.out.println("Hello World!");
	}
	public static void pre(BinaryTree root)
	{
		Deque<BinaryTree> stack=new ArrayDeque<>();
		while (true)
		{
			while (root!=null)
			{
				System.out.print(root.data+" ");
				stack.push(root);
				root=root.left;
			}
			if (stack.isEmpty())
			{
				break;
			}
			root=stack.pop();
			root=root.right;
		}
	}
	public static void mid(BinaryTree root)
	{
		Deque<BinaryTree> stack=new ArrayDeque<>();
		while (true)
		{
			while (root!=null)
			{
				stack.push(root);
				root=root.left;
			}
			if (stack.isEmpty())
			{
				break;
			}
			root=stack.pop();
			System.out.print(root.data+" ");
			root=root.right;
		}
	}
	public static void post(BinaryTree root)
	{
		ArrayDeque<BinaryTree> stack=new ArrayDeque<>();
		BinaryTree tmp=null;
		while (true)
		{
			while (root!=null)
			{
				stack.push(root);
				root=root.left;
			}
			if (stack.isEmpty())
			break;
			if (stack.peek().right==null)
			{
				root=stack.pop();
				System.out.print(root.data+" ");
				if (root==stack.peek().right)
				{
					System.out.print(stack.peek().data+" ");
					tmp=stack.pop();
				}
			}
			if (!stack.isEmpty())
			{
				root=stack.peek().right;
				if (root==tmp)
				{
					System.out.print(stack.pop().data+" ");
					if (stack.isEmpty())
					{
						break;
					}
					root=stack.peek().right;
				}
			}
			else
				root=null;
		}
	}
	//post2����ȷ�ķ�����post��������
	public static void post2(BinaryTree root)
	{
		BinaryTree temp=root;
		Deque<BinaryTree> stack=new ArrayDeque<>();
		while (root!=null)
		{
			//��������ջ
			for (;root.left!=null ;root=root.left )
			{
				stack.push(root);
			}
			//��ǰ�ڵ��������������������Ѿ����
			while (root!=null&&(root.right==null||root.right==temp))
			{
				System.out.print(root.data+" ");
				temp=root;//��¼��һ��������ڵ�
				if (stack.isEmpty())
				{
					return;
				}
				root=stack.pop();
			}
			stack.push(root);
			root=root.right;
		}
	}
}