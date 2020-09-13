import java.util.*;
public class Test 
{
	public static void main(String[] args)
	{
		ArrayList<Person> list=new ArrayList<>();
		list.add(new Person(1));
		list.add(new Person(2));
		list.add(new Person(3));
		//System.out.println(a);
		Person p=new Person(1);
		list.clone();
		System.out.println(p instanceof Cloneable);
		//System.out.println(Boolean.toBinaryString(true));
		System.out.println(Integer.toBinaryString(1));
	}
}
class Person
{
	public int age;
	public Person(int a)
	{
		age=a;
	}
}