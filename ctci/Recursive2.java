import java.util.*;
public class Recursive2
{
	/*
	ArrayList<ArrayList<Integer>> getSubsets(ArrayList<Integer> set,int index)
	{
		ArrayList<ArrayList<Integer>> allSubsets=new ArrayList<ArrayList<Integer>>();
		if (index<0)
		{
			allSubsets=new ArrayList<ArrayList<Integer>>();
			allSubsets.add(new ArrayList<Integer>());
		}
		else
		{
			ArrayList<ArrayList<Integer>> moreSubsets=new ArrayList<ArrayList<Integer>>();
			allSubsets=getSubsets(set,index-1);
			for (ArrayList<Integer> temp:allSubsets )
			{
				ArrayList<Integer> newSubset= new ArrayList<>();
				newSubset.addAll(temp);
				newSubset.add(set.get(index));
				moreSubsets.add(newSubset);
			}
			allSubsets.addAll(moreSubsets);
		}
		return allSubsets;
	}
	
	ArrayList<ArrayList<Integer>> getSubsets(ArrayList<Integer> set,int index)
	{
		ArrayList<ArrayList<Integer>> allSubsets=new ArrayList<ArrayList<Integer>>();
		allSubsets.add(new ArrayList<Integer>());
		while (index<set.size())
		{
			ArrayList<ArrayList<Integer>> moreSubsets=new ArrayList<ArrayList<Integer>>();
			for (ArrayList<Integer> temp:allSubsets )
			{
				ArrayList<Integer> newSubset= new ArrayList<>();
				newSubset.addAll(temp);
				newSubset.add(set.get(index));
				moreSubsets.add(newSubset);
			}
			allSubsets.addAll(moreSubsets);
		}
		return allSubsets;
	}
	ArrayList<ArrayList<Integer>> getSubsets2(ArrayList<Integer> set)
	{
		ArrayList<ArrayList<Integer>> subsets=new ArrayList<>();
		int max=1<<set.size();
		for (int i=0;i<max ;i++ )
		{
			ArrayList<Integer> subset=convertIntToSet(i,set);
			subsets.add(subset);
		}
		return subsets;
	}
	ArrayList<Integer> convertIntToSet(int x,ArrayList<Integer> set)
	{
		int index=0;
		ArrayList<Integer> subset=new ArrayList<>();
		for (int k=x;k>0;k>>=1 )
		{
			if ((k&1)==1)
			{
				subset.add(set.get(index));
			}
			index++;
		}
		return subset;
	}
	
	public static ArrayList<String> getPerms(String str)
	{
		if (str==null)
		{
			return null;
		}
		ArrayList<String> result=new ArrayList<>();
		if (str.length()==0)
		{
			result.add("");
			return result;
		}
		char c=str.charAt(0);
		String remainder=str.substring(1);
		ArrayList<String> words=getPerms(remainder);
		for (String word:words )
		{
			for (int i=0;i<=word.length() ;i++ )
			{
				String s=insert(word,c,i);
				result.add(s);
			}
		}
		return result;
	}
	static String insert(String word,char c,int i)
	{
		String start=word.substring(0,i);
		String end=word.substring(i);
		return start+c+end;
	}
	public static Set<String> generate(int index)
	{
		Set<String> set=new HashSet<>();
		if (index==0)
		{
			set.add("");
		}
		else
		{
			Set<String> words=generate(index-1);
			for (String word:words )
			{
				for (int i=0;i<word.length() ;i++ )
				{
					if (word.charAt(i)=='(')
					{
						String s=insert(word,i);
					}
				}
				if (!set.contains("()"+word))
				{
					set.add("()"+word);
				}
			}
		}
		return set;
	}
	public static String insert(String str,int index)
	{
		String left=str.substring(0,index+1);
		String right=str.substring(index+1,str.length());
		return left+"()"+right;
	}
	
	public void add(ArrayList<String> list,int left,int right,char[] str,int count)
	{
		if (left<0||right<left)
		{
			return ;
		}
		if (left==0&&right==0)
		{
			String s=String.copyValueOf(str);
			list.add(s);
		}
		else
		{
			if (left>0)
			{
				str[count]='(';
				add(list,left-1,right,str,count+1);
			}
			if (right>left)
			{
				str[count]=')';
				add(list,left,right-1,str,count+1);
			}
		}
	}
	public ArrayList<String> generate(int count)
	{
		char[] str=new char[count*2];
		ArrayList<String> list=new ArrayList<>();
		add(list,count,count,str,0);
		return list;
	}
	
	boolean paintFill(Color[][] screen,int x,int y,Color oColor,Color nColor)
	{
		if (x<0||x>=screen[0].length||
			y<0||y>=screen.length)
		{
			return false;
		}
		if (screen[y][x]==oColor)
		{
			screen[y][x]=nColor;
			paintFill(screen,x-1,y,oColor,nColor);
			paintFill(screen,x+1,y,oColor,nColor);
			paintFill(screen,x,y-1,oColor,nColor);
			paintFill(screen,x,y+1,oColor,nColor);
		}
		return true;
	}
	boolean paintFill(Color[][] screen,int x,int y,Color nColor)
	{
		if (screen[y][x]==nColor)
		{
			return false;
		}
		return paintFill(screen,x,y,screen[y][x],nColor);
	}
	
	public int makeChange(int n,int denom)
	{
		int next_denom=0;
		switch(denom)
		{
			case 25:
				next_denom=10;
			break;
			case 10:
				next_denom=5;
			break;
			case 5:
				next_denom=1;
			break;
			case 1:
				return 1;
		}
		int ways=0;
		for (int i=0;i*denom<=n ;i++ )
		{
			ways+=makeChange(n-denom*i,next_denom);
		}
		return ways;
	}
	*/
}
enum Color
{
	Black,White,Red,Yellow,Green
}
