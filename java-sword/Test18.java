public class Test18 
{
	public static void delete(ListNode head,ListNode target)
	{
		if (head==null||target==null)
		{
			return;
		}
		if (target.next!=null)
		{
			ListNode temp=target.next;
			target.data=temp.data;
			target.next=temp.next;
			temp.next=null;
		}
		else if (head==target)
		{
			head=null;
			target=null;
		}
		else
		{
			ListNode temp=head;
			while (temp.next!=target)
			{
				temp=temp.next;
			}
			temp.next=null;
		}
	}
	public static void main(String[] args) 
	{
		ListNode head=new ListNode(3);
		ListNode temp=head;
		for (int i=5;i<8 ;i++ )
		{
			temp.next=new ListNode(3);
			temp=temp.next;
		}
		for (int i=1;i<3 ;i++ )
		{
			temp.next=new ListNode(3);
			temp=temp.next;
		}
		for (int i=1;i<3 ;i++ )
		{
			temp.next=new ListNode(3);
			temp=temp.next;
		}
		temp=head;
		while (temp!=null)
		{
			System.out.print(temp.data+"-");
			temp=temp.next;
		}
		System.out.println();
		head=deleteDuplicate(head);
		while (head!=null)
		{
			System.out.print(head.data+"-");
			head=head.next;
		}
	}
	public static ListNode deleteDuplicate(ListNode head)
	{
		if (head==null)
		{
			return head;
		}
		ListNode pre=null;
		ListNode cur=head;
		while (cur!=null)
		{
			ListNode next=cur.next;
			boolean needDelete=false;
			if ((next!=null)&&(next.data==cur.data))
			{
				needDelete=true;
			}
			if (!needDelete)
			{
				pre=cur;
				cur=cur.next;
			}
			else
			{
				int value=cur.data;
				while ((next!=null)&&(next.data==value))
				{
					next=next.next;
				}
				if (pre!=null)
				{
					pre.next=next;
				    cur=pre.next;
				}
				else
				{
					head=next;
					cur=next;
				}
			}
		}
		return head;
	}
}
