public class Test6 
{
	static void printHelper(ListNode node)
	{
		if (node.next==null)
		{
			System.out.println(node.data);
			return;
		}
		printHelper(node.next);
		System.out.println(node.data);
	}
	static void print(ListNode node)
	{
		if (node==null)
		{
			return;
		}
		printHelper(node);
	}
	public static void main(String[] args) 
	{
		ListNode node=new ListNode(1);
		ListNode head=node;
	    for (int i=2;i<11 ;i++ )
	    {
			ListNode temp=new ListNode(i);
			node.next=temp;
			node=temp;
	    }
		print(head);
	}
}
