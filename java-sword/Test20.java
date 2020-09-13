public class Test20 
{
	public static Result scanInteger(char[] ch,Result result)
	{
		if ((result.index<ch.length)&&
			(ch[result.index]=='+'||ch[result.index]=='-'))
		{
			result.index++;
		}
		return scanUnsignedInteger(ch,result);
	}
	public static Result scanUnsignedInteger(char[] ch,Result result)
	{
		int before=result.index;
		while (result.index<ch.length
			&&ch[result.index]>='0'
		    &&ch[result.index]<='9')
		{
				result.index++;
		}
		result.numeric=result.index>before;
		return result;
	}
	public static boolean isNumeric(char[] ch)
	{
		if (ch==null||ch.length==0)
		{
			return false;
		}
		Result result=new Result(false,0);
		result=scanInteger(ch,result);
		if ((result.index<ch.length)&&(ch[result.index]=='.'))
		{
			result.index++;
			result.numeric =(scanUnsignedInteger(ch,result).numeric)||result.numeric;
		}
		if ((result.index<ch.length)&&
			(ch[result.index]=='e'||ch[result.index]=='E'))
		{
			result.index++;
			result.numeric=scanInteger(ch,result).numeric&&result.numeric;
		}
		return result.numeric&&(result.index==ch.length);
	}
	public static void main(String[] args) 
	{
		String str="-1";
		System.out.println(isNumeric(str.toCharArray()));
	}
}
class Result
{
	public boolean numeric;
	public int index;
	public Result(boolean n,int i)
	{
		numeric=n;
		index=i;
	}
}
