public class Test24 
{
	public static ListNode reverse(ListNode head)
	{
		if (head==null)
		{
			return null;
		}
		ListNode cur=head;
		ListNode prev=null;
		ListNode next=null;
		while (cur!=null)
		{
			next=cur.next;
			cur.next=prev;
			prev=cur;
			cur=next;
		}
		return prev;
	}
	public static void main(String[] args) 
	{
		ListNode head=new ListNode(0);
		ListNode temp=head;
		for (int i=1;i<6 ;i++ )
		{
			temp.next=new ListNode(i);
			temp=temp.next;
		}
		head=new ListNode(100);
		head=reverse(head);
		while (head!=null)
		{
			System.out.print(head.data+"-");
			head=head.next;
		}
	}
}
