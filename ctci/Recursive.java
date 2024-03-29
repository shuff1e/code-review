import java.util.*;
public class Recursive
{
	public int countWays(int n)
	{
		if (n<0)
		{
			return 0;
		}
		if (n==0)
		{
			return 1;
		}
		return countWays(n-1)+countWays(n-2)+countWays(n-3);
	}
	public int countWaysDP(int n)
	{
		int[] map=new int[n+1];
		for (int i=0;i<map.length ;i++ )
		{
			map[i]=-1;
		}
		return countWaysDPHelper(n,map);
	}
	public int countWaysDPHelper(int n,int[] map)
	{
		if (n<0)
		{
			return 0;
		}
		if (n==0)
		{
			return 1;
		}
		if (map[n]!=-1)
		{
			return map[n];
		}
		return countWaysDPHelper(n-1,map)+countWaysDPHelper(n-2,map)
			+countWaysDPHelper(n-3,map);
	}
	public boolean getPath(int x,int y,ArrayList<Point> path)
	{
		Point p=new Point(x,y);
		path.add(p);
		if (x==0&&y==0)
		{
			return true;
		}
		boolean succ=false;
		if (x>=1&&isFree(x-1,y))
		{
			succ=getPath(x-1,y,path);
		}
		if (!succ&&y>=1&&isFree(x,y-1))
		{
			succ=getPath(x,y-1,path);
		}
		if (!succ)
		{
			path.remove(p);
		}
		return succ;
	}
	public boolean getPathDP(int x,int y,ArrayList<Point> path,
		HashTable<Point ,Boolean>cache)
	{
		Point p=new Point(x,y);
		if (cache.containsKey(p))
		{
			return cache.get(p);
		}
		path.add(p);
		if (x==0&&y==0)
		{
			return true;
		}
		boolean succ=false;
		if (x>=1&&isFree(x-1,y))
		{
			succ=getPathDP(x-1,y,path,cache);
		}
		if (!succ&&y>=1&&isFree(x,y-1))
		{
			succ=getPathDP(x,y-1,path,cache);
		}
		cache.put(p,succ);
		if (!succ)
		{
			path.remove(p);
		}
		return succ;
	}
	public static int magicFast(int[] array,int start,int end)
	{
		if (end<start||start<0||end>=array.length)
		{
			return -1;
		}
		int mid=start+(end-start)/2;
		if (array[mid]==mid)
		{
			return mid;
		}
		else if (array[mid]<mid)
		{
			return magicFast(array,mid+1,end);
		}
		else
			return magicFast(array,start,mid-1);
	}
	public static int magicFast2(int[] array,int start,int end)
	{
		if (end<start||start<0||end>=array.length)
		{
			return -1;
		}
		int midIndex=start+(end-start)/2;
		int midValue=array[midIndex];
		if (midValue==midIndex)
		{
			return midIndex;
		}
		int leftIndex=Math.min(midIndex-1,midValue);
		int left=magicFast2(array,start,leftIndex);
		if (left>0)
		{
			return left;
		}
		int rightIndex=Math.max(midIndex+1,midValue);
		int right=magicFast2(array,rightIndex,end);
		return right;
	}
}
