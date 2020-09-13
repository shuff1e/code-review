import java.util.*;
public class Test45 
{
	public static void getMinHelper(Integer[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return;

		}
		Arrays.sort(arr,new MyComparator());
		for (int i=0;i<arr.length ;i++ )
		{
			System.out.print(arr[i]);
		}
	}
	public static void main(String[] args) 
	{
		Integer[] arr={3,32,321};
		getMinHelper(arr);
	}
}
class MyComparator implements Comparator<Integer>
{
	@Override
	public int compare(Integer o1,Integer o2)
	{
		String str1=String.valueOf(o1)+String.valueOf(o2);
		String str2=String.valueOf(o2)+String.valueOf(o1);
		return str1.compareTo(str2);
	}
}
