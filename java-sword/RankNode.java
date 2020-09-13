public class RankNode 
{
	public int left_size=0;
	public RankNode left,right;
	public int data;
	public RankNode(int data)
	{
		this.data=data;
	}
	public void insert(int d)
	{
		if (d<=data)
		{
			if (left!=null)
			{
				left.insert(d);
			}
			else
				left=new RankNode(d);
			left_size++;
		}
		else
		{
			if (right!=null)
			{
				right.insert(d);
			}
			else
				right=new RankNode(d);
		}
	}
	public int getRank(int d)
	{
		if (d==data)
		{
			return left_size;
		}
		else if (d<data)
		{
			if (left==null)
			{
				return 0;
			}
			else
				return left.getRank(d);
		}
		else
		{
			int right_rank=right==null?0:right.getRank(d);
			return left_size+1+right_rank;
		}
	}
	public static void main(String[] args) 
	{
		int[] arr=new int[]{8,5,6,4};
		RankNode root=new RankNode(arr[0]);
		for (int i=1;i<arr.length ;i++ )
		{
			root.insert(arr[i]);
		}
		System.out.println(root.getRank(7));
	}
}
