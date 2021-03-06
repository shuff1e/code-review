import java.util.*;
public class Test35 
{
	/*
	public static void cloneFirst(ListNode head1,HashMap<ListNode,ListNode> map)
	{
		ListNode head2=new ListNode(head1.data);
		ListNode temp1=head1.next;
		ListNode temp2=head2;
		map.put(head1,head2);
		while (temp1!=null)
		{
			temp2.next=new ListNode(temp1.data);
			temp2=temp2.next;
			map.put(temp1,temp2);
			temp1=temp1.next;
		}
	}
	public static ListNode cloneSecond(ListNode head,HashMap<ListNode,ListNode> map)
	{
		ListNode head2=map.get(head);
		while (head!=null)
		{
			map.get(head).random=map.get(head.random);
			head=head.next;
		}
		return head2;
	}
	public static ListNode cloneNodes(ListNode head)
	{
		HashMap<ListNode,ListNode> map=new HashMap<>();
		cloneFirst(head,map);
		return cloneSecond(head,map);
	}
	
    public static ListNode cloneFirst(ListNode head1)
	{
		ListNode head2=new ListNode(head1.data);
		ListNode temp1=head1.next;
		ListNode temp2=head2;
		while (temp1!=null)
		{
			temp2.next=new ListNode(temp1.data);
			
			temp2=temp2.next;
			temp1=temp1.next;
		}
		return head2;
	}
	public static ListNode cloneSecond(ListNode head1,ListNode head2)
	{
		ListNode temp1=head1;
		ListNode temp2=head2;
		while (temp1!=null)
		{
			int k=distance(temp1,head1);
			connect(temp2,head2,k);
			temp1=temp1.next;
			temp2=temp2.next;
		}
		return head2;
	}
	public static int distance(ListNode cur,ListNode head)
	{
		int distance=-1;
		int count=-1;
		while (head!=null)
		{
			count++;
			if (head==cur.random)
			{
				distance=count;
				break;
			}
			head=head.next;
		}
		return distance;
	}
	public static void connect(ListNode cur,ListNode head,int k)
	{
		if (k<0)
		{
			return;
		}
		for (int i=0;i<k ;i++ )
		{
			head=head.next;
		}
		cur.random=head;
	}
	public static ListNode cloneNodes(ListNode head1)
	{
		ListNode head2=cloneFirst(head1);
		return cloneSecond(head1,head2);
	}
	*/
	public static void cloneFirst(ListNode head)
	{
		ListNode cur=head;
		while (cur!=null)
		{
			ListNode temp=new ListNode(cur.data);
			temp.next=cur.next;
			cur.next=temp;
			cur=temp.next;
		}
	}
	public static void cloneSecond(ListNode head)
	{
		ListNode cur=head;
		while (cur!=null)
		{
			if (cur.random!=null)
			{
				ListNode cloned=cur.next;
			    cloned.random=cur.random.next;
			}
			cur=cur.next.next;
		}
	}
	public static ListNode split(ListNode head)
	{
		ListNode head1=head;
		ListNode head2=head.next;
		ListNode temp1=head1;
		ListNode temp2=head2;
		temp1.next=temp2.next;
		temp1=temp1.next;
		while (temp1!=null)
		{
			temp2.next=temp1.next;
			temp2=temp2.next;
			temp1.next=temp2.next;
			temp1=temp1.next;
		}
		return head2;
	}
	public static ListNode cloneNodes(ListNode head)
	{
		if (head==null)
		{
			return null;
		}
		cloneFirst(head);
		cloneSecond(head);
		return split(head);
	}
	public static void main(String[] args) 
	{
		ListNode n0=new ListNode(0);
		ListNode n1=new ListNode(1);
		ListNode n2=new ListNode(2);
		ListNode n3=new ListNode(3);
		ListNode n4=new ListNode(4);
		ListNode n5=new ListNode(5);
		ListNode n6=new ListNode(6);
		ListNode n7=new ListNode(7);
		ListNode n8=new ListNode(8);
		n0.random=n0;
		n2.random=n4;
		n7.random=n3;
		n3.random=n7;
		n0.next=n1;
		n1.next=n2;
		n2.next=n3;
		n3.next=n4;
		n4.next=n5;
		n5.next=n6;
		n6.next=n7;
		n7.next=n8;
		ListNode head=cloneNodes(n0);
		ListNode temp=head;
		while (temp!=null)
		{
			if (temp.data==0)
			{
				System.out.println(temp.random.data);
			}
			temp=temp.next;
		}
	}
}
