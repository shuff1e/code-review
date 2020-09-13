public class Test26 
{
	public static boolean match(BinaryTree t1,BinaryTree t2)
	{
		if (t1==null&&t2!=null)
		{
			return false;
		}
		if (t2==null)
		{
			return true;
		}
		if (t1.data!=t2.data)
		{
			return false;
		}
		return match(t1.left,t2.left)&&match(t1.right,t2.right);
	}
	public static boolean isSubTreeCore(BinaryTree t1,BinaryTree t2)
	{
		if (t1==null)
		{
			return false;
		}
		boolean result=false;
		if (t1.data==t2.data)
		{
			result=match(t1,t2);
		}
		if (!result)
		{
			result=isSubTreeCore(t1.left,t2)||isSubTreeCore(t1.right,t2);
		}
		return result;
	}
	public static boolean isSubTree(BinaryTree t1,BinaryTree t2)
	{
		if (t2==null)
		{
			return true;
		}
		return isSubTreeCore(t1,t2);
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
		BinaryTree node2=new BinaryTree(8);
		left=new BinaryTree(9);
		right=new BinaryTree(2);
		node2.left=left;
		node2.right=right;
		/////////////////////////
		System.out.println(isSubTree(node1,node2));
		double a=0.0000_0000_1;
		System.out.println(a);
	}
	public static boolean equal(double a,double b)
	{
		return Math.abs(a-b)<0.0000_0000_1;
	}
}
