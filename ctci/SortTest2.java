import java.util.*;
public class SortTest2
{
	/*
	public static RankNode root=null;
	public void track(int d)
	{
		if (root==null)
		{
			root=new RankNode(d);
		}
		else
			root.insert(d);
	}
	public int getRankOfNumber(int number)
	{
		return root.getRank(number);
	}
}
class RankNode
{
	public int left_size;
	public RankNode left;
	public RankNode right;
	public int data;
	public RankNode(int d)
	{
		data=d;
	}
	public void insert(int d)
	{
		if (d<=data)
		{
			if (left!=null)
			{
				left.insert(d);
			}
			else
				left=new RankNode(d);
		}
		else
		{
			if (right!=null)
			{
				right.insert(d);
			}
			else
				right=new RankNode(d);
		}
	}
	public int getRank(int d)
	{
		if (d==data)
		{
			return left_size;
		}
		else if (d<data)
		{
			if (left==null)
			{
				return -1;
			}
			else
				return left.getRank(d);
		}
		else
		{
			if (right==null)
			{
				return left_size+1;
			}
			else
				return left_size+1+right.getRank(d);
		}
	}
	
	public static int add(Person[] persons,Person bottom,HashMap<Person,Integer> cache)
	{
		if (cache.containsKey(bottom))
		{
			return cache.get(bottom);
		}
		int max_count=1;
		for (int i=0;i<persons.length;i++ )
		{
			if (persons[i].canBeAbove(bottom))
			{
				int new_count=add(persons,persons[i],cache)+1;
				if (new_count>max_count)
				{
					max_count=new_count;
				}
			}
		}
		cache.put(bottom,max_count);
		return max_count;
	}
	public static void main(String... args)
	{
		Person[] persons={new Person(180,70),new Person(160,50),new Person(181,65),new Person(165,75),new Person(155,49)};
		int max_count=0;
		HashMap<Person,Integer> cache=new HashMap<>();
		for (int i=0;i<persons.length ;i++ )
		{
			int new_count=add(persons,persons[i],cache);
			if (new_count>max_count)
			{
				max_count=new_count;
			}
		}
		System.out.println(max_count);
	}
}
class Person
{
	public double height;
	public double weight;
	public Person(double h,double w)
	{
		height=h;
		weight=w;
	}
	public boolean canBeAbove(Person person)
	{
		return this.height>person.height&&this.weight>person.weight;
	}
	*/
	ArrayList<Person> get(ArrayList<Person> items)
	{
		Collections.sort(items);
		ArrayList<Person>[] solutions=new ArrayList[items.size()+1];
		longest(items,solutions,0);
		ArrayList<Person> best=null;
		for (int i=0;i<solutions.length ;i++ )
		{
			best=seq(best,solutions[i]);
		}
		return best;
	}
	void longest(ArrayList<Person> items,ArrayList<Person>[] solutions,int index)
	{
		if (index<0||index>=items.size())
		{
			return;
		}
		Person p=items.get(index);
		ArrayList<Person> best=null;
		for (int i=0;i<index ;i++ )
		{
			if (items.get(i).isBefore(p))
			{
				best=seq(best,solutions[i]);
			}
		}
		ArrayList<Person> solution=new ArrayList<>();
		if (best!=null)
		{
			solution.addAll(best);
		}
		solution.add(p);
		solutions[index]=solution;
		longest(items,solutions,index+1);
	}
	ArrayList<Person> seq(ArrayList<Person> p1,ArrayList<Person> p2)
	{
		if (p1==null)
		{
			return p2;
		}
		if (p2==null)
		{
			return p1;
		}
		return p1.size()>p2.size()?p1:p2;
	}
}
class Person implements Comparable
{
	public int height;
	public int weight;
	public int compareTo(Object o)
	{
		Person p=(Person)o;
		if (height==p.height)
		{
			return ((Integer)weight).compareTo(p.weight);
		}
		else
			return ((Integer)height).compareTo(p.height);
	}
	public boolean isBefore(Person p)
	{
		return height<p.height&&weight<p.weight;
	}
}
