import java.io.*;
public class Test37 
{
	public static void serialize(BinaryTree root)
	{
		try(ObjectOutputStream oos=new ObjectOutputStream(
			new FileOutputStream("a.txt")))
		{
			serializeHelper(root,oos);
		}
		catch (Exception ex)
		{
			ex.printStackTrace();
		}
	}
	public static void serializeHelper(BinaryTree root,ObjectOutputStream oos)
	{
		try
		{
			if (root==null)
		    {
			oos.writeObject("$");
			return;
		    }
		    oos.writeObject(root.data);
		    serializeHelper(root.left,oos);
		    serializeHelper(root.right,oos);
		    }
		catch (Exception ex)
		{
			ex.printStackTrace();
		}
	}
	public static BinaryTree deSerialize()
	{
		try (ObjectInputStream ois=new ObjectInputStream(
			new FileInputStream("a.txt")))
		{
			BinaryTree root=deSerializeHelper(ois);
			return root;
		}
		catch (Exception ex)
		{
			ex.printStackTrace();
			return null;
		}
	}
	public static BinaryTree deSerializeHelper(ObjectInputStream ois)
	{
		Result result=readStream(ois);
		if (result.isNumber)
		{
			BinaryTree root=new BinaryTree(result.number);
			root.left=deSerializeHelper(ois);
			root.right=deSerializeHelper(ois);
			return root;
		}
		return null;
	}
	public static Result readStream(ObjectInputStream ois)
	{
		try
		{
			Object obj=ois.readObject();
		    if (obj instanceof String)
		    {
			    return new Result(0,false);
		    }
		    int number=(int)obj;
		    return new Result(number,true);
		}
		catch (Exception ex)
		{
			ex.printStackTrace();
			return null;
		}
	}
	public static BinaryTree serialTest(BinaryTree root)
	{
		serialize(root);
		return deSerialize();
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
		BinaryTreeTraversing.DLR(node1);
		System.out.println();
		//////////////////////
		node1=new BinaryTree(8);
		BinaryTree root=serialTest(node1);
		//System.out.println(root);
		BinaryTreeTraversing.DLR(root);
	}
}
class Result
{
	public int number;
	public boolean isNumber;
	public Result(int n,boolean i)
	{
		number=n;
		isNumber=i;
	}
}
