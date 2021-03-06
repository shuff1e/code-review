import java.util.*;
public class EightQueen
{
	int GRID_SIZE=8;
	void placeQueens(int row,Integer[] columns,ArrayList<Integer[]> results)
	{
		if (row==GRID_SIZE)
		{
			results.add(columns.clone());
		}
		else
		{
			for (int col=0;col<GRID_SIZE ;col++ )
			{
				if (checkValid(columns,row,col))
				{
					columns[row]=col;
					placeQueens(row+1,columns,results);
				}
			}
		}
	}
	boolean checkValid(Integer[] columns,int row,int col)
	{
		for (int row2=0;row2<row ;row2++ )
		{
			if (col==columns[row2])
			{
				return false;
			}
			int rowDistance=row-row2;
			int colDistance=Math.abs(col-columns[row2]);
			if (rowDistance==colDistance)
			{
				return false;
			}
		}
		return true;
	}
	public ArrayList<Box> create(Box[] boxes,Box bottom)
	{
		int max_height=0;
		ArrayList<Box> max_stack=null;
		for (int i=0;i<boxes.length ;i++ )
		{
			if (boxes[i].canBeAbove(bottom))
			{
				ArrayList<Box> new_stack=create(boxes,boxes[i]);
				int new_height=height(new_stack);
				if (new_height>max_height)
				{
					max_height=new_height;
					max_stack=new_stack;
				}
			}
		}
		if (max_stack==null)
		{
			max_stack=new ArrayList<Box>();
		}
		if (bottom!=null)
		{
			max_stack.add(0,bottom);
		}
		return max_stack;
	}
	public ArrayList<Box> createDP(Box[] boxes,Box bottom,HashMap<Box,ArrayList<Box>> cache)
	{
		if (bottom!=null&&cache.containsKey(bottom))
		{
			return cache.get(bottom);
		}
		int max_height=0;
		ArrayList<Box> max_stack=null;
		for (int i=0;i<boxes.length ;i++ )
		{
			if (boxes[i].canBeAbove(bottom))
			{
				ArrayList<Box> new_stack=createDP(boxes,boxes[i],cache);
				int new_height=height(new_stack);
				if (new_height>max_height)
				{
					max_height=new_height;
					max_stack=new_stack;
				}
			}
		}
		if (max_stack==null)
		{
			max_stack=new ArrayList<Box>();
		}
		if (bottom!=null)
		{
			max_stack.add(0,bottom);
		}
		cache.put(bottom,max_stack);
		return (ArrayList<Box>)max_stack.clone();
	}
	public int f(String exp,boolean result,int s,int e,HashMap<String,Integer> cache)
	{
		if (s==e)
		{
			if (exp.charAt(s)==1&&result)
			{
				return 1;
			}
			if (exp.charAt(s)==0&&!result)
			{
				return 1;
			}
			return 0;
		}
		String key=""+result+s+e;
		if (cache.containsKey(key))
		{
			return cache.get(key);
		}
		int c=0;
		if (result)
		{
			for (int i=s+1;i<=e ;i+=2 )
			{
				if (exp.charAt(i)=='|')
				{
					c+=f(exp,true,s,i-1,cache)*f(exp,true,i+1,e,cache);
					c+=f(exp,true,s,i-1,cache)*f(exp,false,i+1,e,cache);
					c+=f(exp,false,s,i-1,cache)*f(exp,true,i+1,e,cache);
				}
				else if (exp.charAt(i)=='&')
				{
					c+=f(exp,true,s,i-1,cache)*f(exp,true,i+1,e,cache);
				}
				else if (exp.charAt(i)=='^')
				{
					c+=f(exp,true,s,i-1,cache)*f(exp,false,i+1,e,cache);
					c+=f(exp,false,s,i-1,cache)*f(exp,true,i+1,e,cache);
				}
			}
		}
		else
		{
			for (int i=s+1;i<=e ;i+=2 )
			{
				if (exp.charAt(i)=='|')
				{
					c+=f(exp,false,s,i-1,cache)*f(exp,false,i+1,e,cache);
				}
				else if (exp.charAt(i)=='&')
				{
					c+=f(exp,false,s,i-1,cache)*f(exp,false,i+1,e,cache);
					c+=f(exp,true,s,i-1,cache)*f(exp,false,i+1,e,cache);
					c+=f(exp,false,s,i-1,cache)*f(exp,true,i+1,e,cache);
				}
				else if (exp.charAt(i)=='^')
				{
					c+=f(exp,true,s,i-1,cache)*f(exp,true,i+1,e,cache);
					c+=f(exp,false,s,i-1,cache)*f(exp,false,i+1,e,cache);
				}
			}
		}
	}
}
