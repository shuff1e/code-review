import java.util.*;
//ð������ѡ�����򣬲�������ϣ������
//�������򣬹鲢����
//��������Ͱ����
//������
public class SortTest
{
	//ð������
	//���ѭ���Ӻ���ǰlen-1->0���ڲ�ѭ����ǰ����0->i-1
	//O(n^2)
	public static void bubbleSort(int[] arr)
	{
		int len=arr.length;
		for (int i=len-1;i>=0 ;i-- )
		{
			for (int j=0;j<=i-1 ;j++ )
			{
				if (arr[j]>arr[j+1])
				{
					arr[j]=arr[j+1]+(arr[j+1]=arr[j])*0;
				}
			}
		}
	}
	public static void bubbleSortBetter(int[] arr)
	{
		int len=arr.length;
		boolean flag=true;
		for (int i=len-1;i>=0&&flag ;i-- )
		{
			flag=false;
			for (int j=0;j<=i-1 ;j++ )
			{
				if (arr[j]>arr[j+1])
				{
					flag=true;
					arr[j]=arr[j+1]+(arr[j+1]=arr[j])*0;
				}
			}
		}
	}
	//ѡ������
	//���ѭ����ǰ����0->len-2���ڲ�ѭ����ǰ����i+1->len-1
	public static void selectSort(int[] arr)
	{
		int len=arr.length;
		for (int i=0;i<len-1 ;i++ )
		{
			int min=i;
			for (int j=i+1;j<len ;j++ )
			{
				if (arr[j]<arr[min])
				{
					min=j;
				}
			}
			arr[i]=arr[min]+(arr[min]=arr[i])*0;
		}
	}
	//��������
	//���ѭ����ǰ����0->len-1���ڲ�ѭ���Ӻ���ǰi-1->0
	public static void insertSort(int[] arr)
	{
		int len=arr.length;
		for (int i=0;i<len ;i++ )
		{
			int j=i;
			int temp=arr[j];
			while (j>0&&temp<arr[j-1])
			{
				arr[j]=arr[j-1];
				j--;
			}
			arr[j]=temp;
		}
	}
	public static void shellSort(int[] arr)
	{
		int len=arr.length;
		for (int gap=len/2;gap>0 ;gap/=2 )
		{
			for (int i=0;i<gap ;i++ )
			{
				for (int j=i+gap;j<len ;j+=gap )
				{
					int temp=arr[j];
					int k=j;
					while (k>=gap&&arr[k-gap]>temp)
					{
						arr[k]=arr[k-gap];
						k-=gap;
					}
					arr[k]=temp;
				}
			}
		}
	}
	public static void mergeSort(int[] arr)
	{
		if (arr==null||arr.length==0)
		{
			return;
		}
		int[] temp=new int[arr.length];
		mergeSortHelper(arr,temp,0,arr.length-1);
	}
	//�鲢����
	//�ȷֿ��ٺϲ�
	//��Ҫ��merge()��һ��
	public static void mergeSortHelper(int[] arr,int[] temp,int left,int right)
	{
		if (right>left)
		{
			int mid=(left+right)/2;
			mergeSortHelper(arr,temp,left,mid);
			mergeSortHelper(arr,temp,mid+1,right);
			merge(arr,temp,left,mid+1,right);
		}
	}
	public static void merge(int[] arr,int[] temp,int left,int mid,int right)
	{
		int size=right-left+1;
		int leftEnd=mid-1;
		int tempPos=left;
		while (left<=leftEnd&&mid<=right)
		{
			if (arr[left]<=arr[mid])
			{
				temp[tempPos++]=arr[left++];
			}
			else
			{
				temp[tempPos++]=arr[mid++];
			}
		}
		while (left<=leftEnd)
		{
			temp[tempPos++]=arr[left++];
		}
		while (mid<=right)
		{
			temp[tempPos++]=arr[mid++];
		}
		for (int i=0;i<size ;i++ )
		{
			arr[right]=temp[right];
			right--;
		}
	}
	public static void quickSort(int[] arr)
	{
		int len=arr.length;
		quickSortHelper(arr,0,len-1);
	}
	//��������
	//���ҵ�pivot���ٿ���ʣ�µĲ���
	//��Ҫ�����ҵ�pivot
	public static void quickSortHelper(int[] arr,int low,int high)
	{
		if (high>low)
		{
			int pivot=partition(arr,low,high);
			quickSortHelper(arr,low,pivot-1);
			quickSortHelper(arr,pivot+1,high);
		}
	}
	public static int partition(int[] arr,int low,int high)
	{
		int pivot=arr[low];
		int left=low;
		int right=high;
		while (left<right)
		{
			while (left<right&&arr[left]<=pivot)
			{
				left++;
			}
			while (arr[right]>pivot)
			{
				right--;
			}
			if (left<right)
			{
				swap(arr,left,right);
			}
		}
		swap(arr,low,right);
		return right;
	}
	public static void swap(int[] arr,int left,int right)
	{
		int temp=arr[left];
		arr[left]=arr[right];
		arr[right]=temp;
	}
	//��������
	//helper[i]��key��input[i]��value��¼�����м�������input[i]С��Ҳ����output[]�е�index
	public static void countSort(int[] input)
	{
		int size=findMax(input);
		int[] output=new int[input.length];
		int[] helper=new int[size+1];
		for (int i=0;i<input.length ;i++ )
		{
			helper[input[i]]++;
		}
		for (int i=1;i<helper.length ;i++ )
		{
			helper[i]=helper[i]+helper[i-1];
		}
		for (int i=input.length-1;i>=0 ;i-- )
		{
			output[helper[input[i]]-1]=input[i];
			helper[input[i]]--;
		}
		for (int i=0;i<output.length ;i++ )
		{
			input[i]=output[i];
		}
	}
	//Ͱ����
	//buckets[i]�е�key��¼����arr[i]��value��arr[i]�ĸ���
	public static void bucketSort(int[] arr)
	{
		int size=findMax(arr);
		int[] buckets=new int[size+1];
		for (int i=0;i<arr.length ;i++ )
		{
			buckets[arr[i]]++;
		}
		int index=0;
		for (int i=0;i<buckets.length ;i++ )
		{
			for (int j=buckets[i];j>0 ;j-- )
			{
				arr[index++]=i;
			}
		}
	}
	public static int findMax(int[] arr)
	{
		int max=-1;
		for (int i=0;i<arr.length ;i++ )
		{
			max=Math.max(max,arr[i]);
		}
		return max;
	}
	//������
	//�ѻ���������͸
	public static void heapify(int[] arr,int index,int size)
	{
		int left=2*index+1;
		int right=2*index+2;
		int best=index;
		while (left<size)
		{
			if (arr[left]>arr[index])
			{
				best=left;
			}
			if (right<size&&arr[right]>arr[best])
			{
				best=right;
			}
			if (best!=index)
			{
				swap(arr,best,index);
			}
			else
				break;
			//����֮�󣬿��ܻ�Υ���ѵ�����
			//�����Ҫѭ���Ĺ���
			index=best;
			left=2*index+1;
			right=2*index+2;
		}
	}
	public static void buildHeap(int[] arr,int size)
	{
		//sizeΪ�ѵĴ�С
		//size/2-1Ϊ��һ����Ҷ�ڵ������
		//�����һ����Ҷ�ڵ㿪ʼ�����������
		for (int i=size/2-1;i>=0 ;i-- )
		{
			heapify(arr,i,size);
		}
	}
	public static void heapSort(int[] arr)
	{
		int size=arr.length;
		//����һ���󶥶�
		buildHeap(arr,size);
		for (int i=size-1;i>=0 ;i-- )
		{
			//���󶥶ѶѶ���ֵ�������
			swap(arr,0,i);
			//�ѻ�
			heapify(arr,0,--size);
		}
	}
	public static void main(String[] args) 
	{
		int times=50000;
		for (int t=0;t<times ;t++ )
		{
			int[] arr=new int[(int)Math.random()*100+10];
			for (int i=0;i<arr.length ;i++ )
			{
				arr[i]=(int)(Math.random()*100+10);
			}
			int[] arr2=Arrays.copyOf(arr,arr.length);
			int[] arr3=Arrays.copyOf(arr,arr.length);
			int[] arr4=Arrays.copyOf(arr,arr.length);
			int[] arr5=Arrays.copyOf(arr,arr.length);
			int[] arr6=Arrays.copyOf(arr,arr.length);
			int[] arr7=Arrays.copyOf(arr,arr.length);
			int[] arr8=Arrays.copyOf(arr,arr.length);
			int[] arr9=Arrays.copyOf(arr,arr.length);
			bubbleSort(arr);
			selectSort(arr2);
			insertSort(arr3);
			shellSort(arr4);
			quickSort(arr5);
			mergeSort(arr6);
			countSort(arr7);
			bucketSort(arr8);
			heapSort(arr9);
			if (!Arrays.toString(arr).equals(Arrays.toString(arr2)))
			{
				System.out.println("Fuck12");
			}
			if (!Arrays.toString(arr).equals(Arrays.toString(arr3)))
			{
				System.out.println("Fuck23");
			}
			if (!Arrays.toString(arr).equals(Arrays.toString(arr4)))
			{
				System.out.println("Fuck34");
			}
			if (!Arrays.toString(arr).equals(Arrays.toString(arr5)))
			{
				System.out.println("Fuck45");
			}
			if (!Arrays.toString(arr).equals(Arrays.toString(arr6)))
			{
				System.out.println("Fuck56");
			}
			if (!Arrays.toString(arr).equals(Arrays.toString(arr7)))
			{
				System.out.println("Fuck67");
			}
			if (!Arrays.toString(arr).equals(Arrays.toString(arr8)))
			{
				System.out.println("Fuck78");
			}
			if (!Arrays.toString(arr).equals(Arrays.toString(arr9)))
			{
				System.out.println("Fuck89");
			}
		}
	}
}
