public class Test22 
{
	public static ListNode search(ListNode head,int k)
	{
		if (head==null||k==0)
		{
			return null;
		}
		ListNode temp=head;
		for (int i=0;i<k-1 ;i++ )
		{
			temp=temp.next;
			if (temp==null)
			{
				return null;
			}
		}
		while (temp.next!=null)
		{
			temp=temp.next;
			head=head.next;
		}
		return head;
	}
	public static void main(String[] args) 
	{
		ListNode head=new ListNode(0);
		ListNode temp=head;
		for (int i=1;i<10 ;i++ )
		{
			temp.next=new ListNode(i);
			temp=temp.next;
		}
		temp=head;
		while (temp!=null)
		{
			System.out.println(temp.data);
			temp=temp.next;
		}
		System.out.println(search(head,1).data);
	}
}
