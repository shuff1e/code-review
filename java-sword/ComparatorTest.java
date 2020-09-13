/*
public static <T> void sort(T[] a,int fromIndex, int toIndex,  Comparator<? super T> c)

上面有一个拘束，就是排列顺序只能是从小到大，如果我们要从大到小，就要使用这种方式

这里牵扯到了Java里面的泛型，如果读者不是很了解，可以暂时不去管它，如果真的很想了解，
建议查阅上面我推荐的那本书，上面有详细的介绍。

读者只需要读懂下面的例子就可以了，其实就是多了一个Comparator类型的参数而已。
*/

import java.util.Arrays;
import java.util.Comparator;
public class ComparatorTest 
{
	public static void main(String[] args) 
	{
//注意，要想改变默认的排列顺序，不能使用基本类型（int,double, char）//而要使用它们对应的类
         Integer[] a = {9, 8, 7, 2, 3, 4, 1, 0, 6, 5};
//定义一个自定义类MyComparator的对象
//new MyComparator<>()，错误：无法推断MyComparator的类型参数
        Comparator<Integer> cmp = new MyComparator();
        Arrays.sort(a, cmp);
        for(int i = 0; i < a.length; i ++) 
		{
             System.out.print(a[i] + " ");
        }
    }
}
//Comparator是一个接口，所以这里我们自己定义的类MyComparator要implents该接口//而不是extends Comparator
class MyComparator implements Comparator<Integer>
	{
     @Override 
     public int compare(Integer o1, Integer o2) 
		 {
//如果n1小于n2，我们就返回正值，如果n1大于n2我们就返回负值，//这样颠倒一下，就可以实现反向排序了
            if(o1 < o2) 
		    { 
              return 1;
            }
		    else if(o1 > o2)
			{
              return -1;
            }
		    else
			{
              return 0;   
            }  
        }    
  }