import java.util.*;
public class SortTest 
{
	public static void merge(int[] a,int[] b,int lastA,int lastB)
	{
		int indexA=lastA-1;
		int indexB=lastB-1;
		int indexMerged=lastA+lastB-1;
		while (indexA>=0&&indexB>=0)
		{
			if (a[indexA]>b[indexB])
			{
				a[indexMerged--]=a[indexA--];
			}
			else
				a[indexMerged--]=b[indexB--];
		}
		while (indexB>=0)
		{
			a[indexMerged--]=b[indexB--];
		}
	}
	/*
	public class AnagramComparator implements Comparator<String>
	{
		public int compare(String s1,String s2)
		{
			return sortChars(s1).compareTo(sortChars(s2));
		}
	}
	
	public String sortChars(String s)
	{
		char[] c=s.toCharArray();
		Arrays.sort(c);
		return new String(c);
	}
	public void sort(String[] array)
	{
		HashMap<String,LinkedList<String>> map=new HashMap<>();
		for (String s:array )
		{
			String key=sortChars(s);
			if (!map.containsKey(key))
			{
				map.put(key,new LinkedList<String>());
			}
				LinkedList<String> anagrams=map.get(key);
				anagrams.push(s);
		}
		int index=0;
		for (String key:map.keySet() )
		{
			LinkedList<String> list=map.get(key);
			for (String t:list )
			{
				array[index]=t;
				index++;
			}
		}
	}
	
	public int search(int[] a,int left,int right,int x)
	{
		int mid=left+(right-left)/2;
		if (x==a[mid])
		{
			return mid;
		}
		if (right<left)
		{
			return -1;
		}
		if (a[left]<a[mid])
		{
			if (x>=a[left]&&x<=a[mid])
			{
				return search(a,left,mid-1,x);
			}
			else
			{
				return search(a,mid+1,right,x);
			}
		}
		else if (a[mid]<a[right])
		{
			if (x>=a[mid]&&x<=a[right])
			{
				return search(a,mid+1,right,x);
			}
			else
				return search(a,left,mid-1,x);
		}
		else if (a[left]==a[mid])
		{
			if (a[mid]!=a[right])
			{
				return search(a,mid+1,right,x);
			}
			else
			{
				int result=search(a,left,mid-1,x);
				if (result==-1)
				{
					return search(a,mid+1,right,x);
				}
				else
					return result;
			}
		}
		return -1;
	}
	
	public int searchR(String[] strings,String str,int first,int last)
	{
		if (first>last)
		{
			return -1;
		}
		int mid=(first+last)/2;
		if (strings[mid].equals(""))
		{
			int left=mid-1;
			int right=mid+1;
			while (true)
			{
				if (left<first&&right>last)
				{
					return -1;
				}
				else if (right<=last&&!strings[right].equals(""))
				{
					mid=right;
					break;
				}
				else if (left>=first&&!strings[left].equals(""))
				{
					mid=left;
					break;
				}
				right++;
				left--;
			}
		}
		if (str.equals(strings[mid]))
		{
			return mid;
		}
		else if (strings[mid].compareTo(str)>0)
		{
			return searchR(strings,str,first,mid-1);
		}
		else
		{
			return searchR(strings,str,mid+1,last);
		}
	}
	public int searchIte(String[] strings,String str)
	{
		int first=0;
		int last=strings.length-1;
		int left=first;
		int right=last;
		while (true)
		{
			int mid=(left+right)/2;
			if (strings[mid].equals(""))
		  {
			left=mid-1;
			right=mid+1;
			while (true)
			{
				if (left<first&&right>last)
				{
					return -1;
				}
				else if (right<=last&&!strings[right].equals(""))
				{
					mid=right;
					break;
				}
				else if (left>=first&&!strings[left].equals(""))
				{
					mid=left;
					break;
				}
				right++;
				left--;
			}
		  }
		  if (str.equals(strings[mid]))
		{
			return mid;
		}
		else if (strings[mid].compareTo(str)>0)
		{
			right=mid-1;
		}
		else
		{
			left=mid+1;
		}
		}
	}
	
	public static boolean findElement(int[][] matrix,int elem)
	{
		int row=0;
		int col=matrix[0].length-1;
		while (row<matrix.length&&col>=0)
		{
			if (matrix[row][col]==elem)
			{
				return true;
			}
			else if (matrix[row][col]>elem)
			{
				col--;
			}
			else
			{
				row++;
			}
		}
		return true;
	}
	*/
	public Coordinate findElement(int[][] matrix,Coordinate origin,Coordinate dest,int x)
	{
		if (!origin.inbounds(matrix)||!dest.inbounds(matrix))
		{
			return null;
		}
		if (matrix[origin.row][origin.column]==x)
		{
			return origin;
		}
		else if (!origin.isBefore(dest))
		{
			return null;
		}
		Coordinate start=(Coordinate)origin.clone();
		int diagDest=Math.min(dest.row-origin.row,dest.column-origin.column);
		Coordinate end=new Coordinate(strat.row+diagDest,start.column+diagDest);
		Coordinate p=new Coordinate(0,0);
		while (start.isBefore(end))
		{
			p.setToAverage(start,end);
			if (x>matrix[p.row][p.column])
			{
				start.row=p.row+1;
				start.column=p.column+1;
			}
			else
			{
				end.row=p.row-1;
				end.column=p.column-1;
			}
		}
	}
	public Coordinate partitionAndSearch(int[][] matrix,Coordinate origin,Coordinate dest,Coordinate pivot)
	{
		Coordinate lowerLeftOrigin=new Coordinate(pivot.row,origin.column);
		Coordinate lowerLeftDest=new Coordinate(dest.row,pivot.column-1);
		Coordinate upperRightOrigin=new Coordinate(origin.row,pivot.column);
		Coordinate upperRightDest=new Coordinate(pivot.row-1,dest.column);
		Coordinate lowerLeft=findElement(matrix,lowerLeftOrigin,lowerRightOrigin);
		if (lowerLeft==null)
		{
			return findElement(matrix,upperRightOrigin,upperRightDest,elem);
		}
		return lowerLeft;
	}
}
class Coordinate implements Cloneable
{
	public int row;
	public int column;
	public Coordinate(int r,int c)
	{
		row=r;
		column=c;
	}
	public boolean inbounds(int[][] matrix)
	{
		return row>=0&&column>=0&&
			row<matrix.length&&col<matrix[0].length;
	}
	public boolean isBefore(Coordinate p)
	{
		return row<=p.row&&column<=p.column;
	}
	public Object clone()
	{
		return new Coordinate(row,column);
	}
	public void setToAverage(Coordinate min,Coordinate max)
	{
		row=(min.row+max.row)/2;
		column=(min.column+max.column)/2;
	}
}
