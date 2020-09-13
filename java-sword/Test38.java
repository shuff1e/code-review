import java.util.*;
public class Test38
{
    // 给定字符串abc
    // 对于字符串ab，字符c可以插在ab的任意位置
    public static void makeArray(String[] strArr)
    {
        HashSet<String> set=new HashSet<>();
        set.add(strArr[0]);
        int index=1;
        while (index<strArr.length)
        {
            Iterator it=set.iterator();
            set=new HashSet<String>();
            while (it.hasNext())
            {
                String str=(String)it.next();
                StringBuilder sb=new StringBuilder(str);
                for (int i=0;i<=str.length() ;i++ )
                {
                    sb.insert(i,strArr[index]);
                    set.add(sb.toString());
                    sb.delete(i,i+1);
                }
            }
            index++;
        }
        for (Object str :set )
        {
            System.out.println((String)str);
        }
    }
    public static void permutationHelper(char[] arr,int index)
    {
        if (index==arr.length)
        {
            print(arr);
        }
        else
        {
            for (int i=index;i<arr.length ;i++ )
            {
                char temp=arr[i];
                arr[i]=arr[index];
                arr[index]=temp;

                permutationHelper(arr,index+1);

                temp=arr[index];
                arr[index]=arr[i];
                arr[i]=temp;
            }
        }
    }
    public static void print(char[] arr)
    {
        System.out.println(new String(arr));
    }
    public static void permutation(char[] arr)
    {
        if (arr==null||arr.length==0)
        {
            return;
        }
        permutationHelper(arr,0);
    }
    public static void main(String[] args)
    {
        String[] strArr=new String[]{"a","b","c"};
        makeArray(strArr);
        //System.out.println("Hello World!");
//        char[] arr=new char[]{'a','b','c','d'};
        //arr=new char[]{'a'};
//        permutation(arr);
//        for (char str:arr )
//        {
//            System.out.println(str);
//        }
//        find(arr,arr.length);
    }

    // 得到长度为N的字符中，取m个字符的取法
    public static void find(char[] arr,int m)
    {
        findNM(arr,0,new char[m],0);
    }

    public static void findNM(char[] arr,int indexN,char[] helper,int indexM)
    {
        if ((arr.length-indexN)<(helper.length-indexM))
        {
            return;
        }
        if (indexM==helper.length)
        {
            print(helper);
            return;
        }
        helper[indexM]=arr[indexN];
        findNM(arr,indexN+1,helper,indexM+1);

        findNM(arr,indexN+1,helper,indexM);
    }

    // 得到字符串所有字符的组合
    public static void find(char[] arr) {
        char[] helper = new char[arr.length];
        findNM2(arr,0,helper,0);
    }
    public static void findNM2(char[]arr,int indexN,char[] helper,int indexM) {
        if (indexN == arr.length) {
            print(helper);
            return;
        }

        // 选择这个字符
        helper[indexM] = arr[indexN];
        findNM2(arr,indexN+1,helper,indexM+1);
        // 不选择这个字符
        helper[indexM] = '0';
        findNM2(arr,indexN+1,helper,indexM);
    }
}