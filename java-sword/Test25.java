public class Test25 
{
	public static ListNode merge(ListNode head1,ListNode head2)
	{
		if (head1==null)
		{
			return head2;
		}
		if (head2==null)
		{
			return head1;
		}
		ListNode head=null;
		if (head1.data<head2.data)
		{
			head=head1;
			head.next=merge(head1.next,head2);
		}
		else
		{
			head=head2;
			head.next=merge(head1,head2.next);
		}
		return head;
	}
	public static void main(String[] args) 
	{
		ListNode head1=new ListNode(0);
		ListNode temp=head1;
		for (int i=1;i<8 ;i+=2 )
		{
			temp.next=new ListNode(i);
			temp=temp.next;
		}
		/*
		while (head1!=null)
		{
			System.out.print(head1.data+"-");
			head1=head1.next;
		}
		*/
		ListNode head2=new ListNode(0);
		temp=head2;
		for (int i=2;i<10 ;i+=2 )
		{
			temp.next=new ListNode(i);
			temp=temp.next;
		}
		head2=merge(head1,head2);
		while (head2!=null)
		{
			System.out.print(head2.data+"-");
			head2=head2.next;
		}
	}
}
