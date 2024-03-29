public class BinaryTest
{
	public int  updateBits(int n,int m,int j, int i)
	{
		int allOnes=~0;
		int left=allOnes<<(j+1);
		int right=(1<<i)-1;
		int mask=left|right;
		int n_cleared=mask&n;
		int m_shifted=m<<i;
		return n_cleared|m_shifted;
	}
	public static String printBinary(double num)
	{
		if (num>=1||num<0)
		{
			return "ERROR";
		}
		StringBuilder sb=new StringBuilder();
		sb.append(".");
		while (num>0)
		{
			if (sb.length()>=32)
			{
				return "ERROR";
			}
			double r=num*2;
			if (r>=1)
			{
				sb.append(1);
				num=r-1;
			}
			else
			{
				sb.append(0);
				num=r;
			}
		}
		return sb.toString();
	}
	public static String printBinary(int num)
	{
		if (num>=1||num<0)
		{
			return "ERROR";
		}
		double frac=0.5;
		StringBuilder sb=new StringBuilder();
		sb.append(".");
		while (num>0)
		{
			if (sb.length()>=32)
			{
				return "ERROR";
			}
			if (num>=frac)
			{
				sb.append(1);
				num -=frac;
			}
			else
				sb.append(0);
			frac /=2;
		}
		return sb.toString();
	}
	/*
	public int getNext(int n)
	{
		int c=n;
		int c0=0;
		int c1=0;
		while ((c>0)&&((c&1)==0))
		{
			c0++;
			c>>=1;
		}
		while ((c&1)==1)
		{
			c1++;
			c>>=1;
		}
		int p=c0+c1;
		if (p==32||p==0)
		{
			return -1;
		}
		n |=(1<<p);
		n &=~((1<<p)-1);
		n |=(1<<(c1-1)-1);
		return n;
	}
	public int getPrev(int n)
	{
		int c=n;
		int c0=0;
		int c1=0;
		while ((c&1)==1)
		{
			c1++;
			c>>=1;
		}
		while ((c>0)&&((c&1)==0))
		{
			c0++;
			c>>=1;
		}
		int p=c0+c1;
		int mask=(1<<(p+1))-1;
		n &=~mask;
		mask=(1<<(c1+1))-1;
		n |=(mask<<c0-1);
		return n;
	}
	
	public int getNext(int n)
	{
		int c=n;
		int c0=0;
		int c1=0;
		while ((c>0)&&((c&1)==0))
		{
			c0++;
			c>>=1;
		}
		while ((c&1)==1)
		{
			c1++;
			c>>=1;
		}
		int p=c0+c1;
		if (p==32||p==0)
		{
			return -1;
		}
		return n+1<<(c0)+1<<(c1-1)-1;
	}
	public int getPrev(int n)
	{
		int c=n;
		int c0=0;
		int c1=0;
		while ((c&1)==1)
		{
			c1++;
			c>>=1;
		}
		while ((c>0)&&((c&1)==0))
		{
			c0++;
			c>>=1;
		}
		int p=c0+c1;
		return n-((1<<c1)-1)-1-(1<<(c0-1)-1); 
	}
	public int bitSwapRequired(int a,int b)
	{
		int count=0;
		for (int c=a^b;c>0 ;c>>=1 )
		{
			if ((c&1)==1)
			{
				count++;
			}
		}
		return count;
	}
	public int botSwap(int a,int b)
	{
		int count=0;
		for (int c=a^b;c!=0 ;c=c&(c-1) )
		{
			count++;
		}
		return count;
	}
	public int swapOddEven(int x)
	{
		return ((x&0xaaaaaaaa)>>1)|((x&0x55555555)<<1);
	}
	
	public int findMissing(ArrayList<Bit> array)
	{
		return findMissing(array,0);
	}
	public int findMissing(ArrayList<Bit> array,int column)
	{
		if (column>=Bit.size())
		{
			return 0;
		}
		ArrayList<Bit> zero=new ArrayList<>();
		ArrayList<Bit> one=new ArrayList<>();
		for (Bit t:array )
		{
			if (t.fetch(column)==0)
			{
				zero.add(t);
			}
			else
				one.add(t);
		}
		if (zero.size()<=one.size())
		{
			int v= findMissing(zero,column+1);
			return (v<<1)|0;
		}
		else
		{
			int v=findMissing(one,column+1);
			return (v<<1)|1;
		}
	}
	*/
	public void drawLine(byte[] screen,int x1,int x2,int y,int w)
	{
		int start_offset=x1%8;
		int first_full_byte=x1/8;
		if (start_offset!=0)
		{
			first_full_byte++;
		}
		int end_offset=x2%8;
		int last_full_byte=x2/8;
		if (end_offset!=7)
		{
			last_full_byte--;
		}
		for (int i=first_full_byte;i<=last_full_byte ;i++ )
		{
			screen[w/8*y+i]=(byte)0xff;
		}
		byte start_mask=(byte)(0xff>>start_offset);
		byte end_mask=(byte)~(0xff>>(end_offset+1));
		if (x1/8==x2/8)
		{
			byte mask=(byte)(start_mask&end_mask);
			int number=w/8*y+x1/8;
			screen[number] |=mask;
		}
		else
		{
			if (start_offset!=0)
			{
				int number=w/8*y+first_full_byte-1;
				screen[number] |=start_mask;
			}
			if (end_offset!=7)
			{
				int number=w/8*y+last_full_byte+1;
				screen[number] |=end_mask;
			}
		}
	}
}
