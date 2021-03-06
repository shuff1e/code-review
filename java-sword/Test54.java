public class Test54 
{
	public static BinaryTree getKthNode(BinaryTree root,int k)
	{
		if (root==null||k==0)
		{
			return null;
		}
		Wrapper wra=new Wrapper(k);
		return KthNodeCore(root,wra);
	}
	public static BinaryTree KthNodeCore(BinaryTree root,Wrapper wra)
	{
		BinaryTree target=null;
		if (root.left!=null)
		{
			target=KthNodeCore(root.left,wra);
		}
		if (target==null)
		{
			if (wra.k==1)
			{
				target=root;
			}
			wra.k--;
		}
		if (target==null&&root.right!=null)
		{
			target=KthNodeCore(root.right,wra);
		}
		return target;
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
		////////////////////////P148
		System.out.println(getKthNode(node1,7).data);
	}
}
class Wrapper
{
	public int k;
	public Wrapper(int k)
	{
		this.k=k;
	}
}
