// 设开始的为K，遍历第一次，得到（k-n），n是链表长度

// 再从头开始，走（n-k）步，就得到倒数第k个节点
public class Test52 
{
	public static ListNode getCommon(ListNode node1,ListNode node2)
	{
		int len1=getLength(node1);
		int len2=getLength(node2);
		int lengthDiff=len1-len2;
		ListNode longList=node1;
		ListNode shortList=node2;
		if (len1<len2)
		{
	        lengthDiff=len2-len1;
		    longList=node2;
		    shortList=node1;
		}
		for (int i=0;i<lengthDiff ;i++ )
		{
			longList=longList.next;
		}
		while (longList!=null&&shortList!=null&&longList!=shortList)
		{
			longList=longList.next;
			shortList=shortList.next;
		}
		return longList;
	}
	public static void main(String[] args) 
	{
		ListNode head=new ListNode(0);
		ListNode temp=head;
		ListNode helper=null;
		for (int i=1;i<3 ;i++ )
		{
			temp.next=new ListNode(i);
			temp=temp.next;
		}
		ListNode head2=new ListNode(0);
		ListNode temp2=head2;
		for (int i=4;i<10 ;i++ )
		{
			temp2.next=temp.next=new ListNode(i);
			temp=temp.next;
			temp2=temp2.next;
		}
		System.out.println(getCommon(head,head2).data);
	}
	public static int getLength(ListNode node)
	{
		int count=0;
		while (node!=null)
		{
			count++;
			node=node.next;
		}
		return count;
	}
}
