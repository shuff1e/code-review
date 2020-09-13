/*
public static <T> void sort(T[] a,int fromIndex, int toIndex,  Comparator<? super T> c)

������һ����������������˳��ֻ���Ǵ�С�����������Ҫ�Ӵ�С����Ҫʹ�����ַ�ʽ

����ǣ������Java����ķ��ͣ�������߲��Ǻ��˽⣬������ʱ��ȥ�����������ĺ����˽⣬
��������������Ƽ����Ǳ��飬��������ϸ�Ľ��ܡ�

����ֻ��Ҫ������������ӾͿ����ˣ���ʵ���Ƕ���һ��Comparator���͵Ĳ������ѡ�
*/

import java.util.Arrays;
import java.util.Comparator;
public class ComparatorTest 
{
	public static void main(String[] args) 
	{
//ע�⣬Ҫ��ı�Ĭ�ϵ�����˳�򣬲���ʹ�û������ͣ�int,double, char��//��Ҫʹ�����Ƕ�Ӧ����
         Integer[] a = {9, 8, 7, 2, 3, 4, 1, 0, 6, 5};
//����һ���Զ�����MyComparator�Ķ���
//new MyComparator<>()�������޷��ƶ�MyComparator�����Ͳ���
        Comparator<Integer> cmp = new MyComparator();
        Arrays.sort(a, cmp);
        for(int i = 0; i < a.length; i ++) 
		{
             System.out.print(a[i] + " ");
        }
    }
}
//Comparator��һ���ӿڣ��������������Լ��������MyComparatorҪimplents�ýӿ�//������extends Comparator
class MyComparator implements Comparator<Integer>
	{
     @Override 
     public int compare(Integer o1, Integer o2) 
		 {
//���n1С��n2�����Ǿͷ�����ֵ�����n1����n2���Ǿͷ��ظ�ֵ��//�����ߵ�һ�£��Ϳ���ʵ�ַ���������
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