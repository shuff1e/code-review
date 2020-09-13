public class Test23 
{
	public static ListNode meetingNode(ListNode head)
	{
		if (head==null)
		{
			return null;
		}
		ListNode slow=head;
		ListNode fast=head;
		while (fast!=null&&fast.next!=null)
		{
			slow=slow.next;
			fast=fast.next.next;
			if (slow==fast)
			{
				return slow;
			}
		}
		return null;
	}
	public static ListNode findEntrance(ListNode head)
	{
		ListNode meet=meetingNode(head);
		if (meet==null)
		{
			return null;
		}
		int count=1;
		ListNode meetHelper=meet;
		while (meet.next!=meetHelper)
		{
			count++;
			meet=meet.next;
		}
		ListNode fast=head;
		for (int i=0;i<count ;i++ )
		{
			fast=fast.next;
		}
		ListNode slow=head;
		while (slow!=fast)
		{
			slow=slow.next;
			fast=fast.next;
		}
		return slow;
	}
	public static void main(String[] args) 
	{
		ListNode head=new ListNode(0);
		ListNode temp=head;
		ListNode helper=null;
		for (int i=1;i<6 ;i++ )
		{
			temp.next=new ListNode(i);
			temp=temp.next;
			if (i==2)
			{
				//helper=temp;
			}
		}
		temp.next=helper;
		temp=head;
		/*
		for (int i=0;i<15;i++)
		{
			System.out.print(temp.data+"-");
			temp=temp.next;
		}
		*/
		System.out.println(findEntrance(head));
	}
}
