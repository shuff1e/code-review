public class ListNode 
{
	public int data;
	public ListNode next;
	public ListNode prev;
	public ListNode random;
	public ListNode()
	{}
	public ListNode(int d)
	{
		data=d;
	}
	public boolean equals(Object obj)
	{
		if (this==obj)
		{
			return true;
		}
		if (obj!=null&&obj.getClass()==ListNode.class)
		{
			ListNode l=(ListNode)obj;
			return this.data==l.data;
		}
		return false;
	}
	public int hashCode()
	{
		return this.data;
	}
}
