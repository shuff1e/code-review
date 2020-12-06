package main

/*

1396. 设计地铁系统
请你实现一个类 UndergroundSystem ，它支持以下 3 种方法：

1. checkIn(int id, string stationName, int t)

编号为 id 的乘客在 t 时刻进入地铁站 stationName 。
一个乘客在同一时间只能在一个地铁站进入或者离开。
2. checkOut(int id, string stationName, int t)

编号为 id 的乘客在 t 时刻离开地铁站 stationName 。
3. getAverageTime(string startStation, string endStation)

返回从地铁站 startStation 到地铁站 endStation 的平均花费时间。
平均时间计算的行程包括当前为止所有从 startStation 直接到达 endStation 的行程。
调用 getAverageTime 时，询问的路线至少包含一趟行程。
你可以假设所有对 checkIn 和 checkOut 的调用都是符合逻辑的。也就是说，如果一个顾客在 t1 时刻到达某个地铁站，那么他离开的时间 t2 一定满足 t2 > t1 。所有的事件都按时间顺序给出。



示例：

输入：
["UndergroundSystem","checkIn","checkIn","checkIn","checkOut","checkOut","checkOut","getAverageTime","getAverageTime","checkIn","getAverageTime","checkOut","getAverageTime"]
[[],[45,"Leyton",3],[32,"Paradise",8],[27,"Leyton",10],[45,"Waterloo",15],[27,"Waterloo",20],[32,"Cambridge",22],["Paradise","Cambridge"],["Leyton","Waterloo"],[10,"Leyton",24],["Leyton","Waterloo"],[10,"Waterloo",38],["Leyton","Waterloo"]]

输出：
[null,null,null,null,null,null,null,14.0,11.0,null,11.0,null,12.0]

解释：
UndergroundSystem undergroundSystem = new UndergroundSystem();
undergroundSystem.checkIn(45, "Leyton", 3);
undergroundSystem.checkIn(32, "Paradise", 8);
undergroundSystem.checkIn(27, "Leyton", 10);
undergroundSystem.checkOut(45, "Waterloo", 15);
undergroundSystem.checkOut(27, "Waterloo", 20);
undergroundSystem.checkOut(32, "Cambridge", 22);
undergroundSystem.getAverageTime("Paradise", "Cambridge");       // 返回 14.0。从 "Paradise"（时刻 8）到 "Cambridge"(时刻 22)的行程只有一趟
undergroundSystem.getAverageTime("Leyton", "Waterloo");          // 返回 11.0。总共有 2 躺从 "Leyton" 到 "Waterloo" 的行程，编号为 id=45 的乘客出发于 time=3 到达于 time=15，编号为 id=27 的乘客于 time=10 出发于 time=20 到达。所以平均时间为 ( (15-3) + (20-10) ) / 2 = 11.0
undergroundSystem.checkIn(10, "Leyton", 24);
undergroundSystem.getAverageTime("Leyton", "Waterloo");          // 返回 11.0
undergroundSystem.checkOut(10, "Waterloo", 38);
undergroundSystem.getAverageTime("Leyton", "Waterloo");          // 返回 12.0


提示：

总共最多有 20000 次操作。
1 <= id, t <= 10^6
所有的字符串包含大写字母，小写字母和数字。
1 <= stationName.length <= 10
与标准答案误差在 10^-5 以内的结果都视为正确结果。

 */

type Start struct {
	station string
	t int
}

type startEnd struct {
	start string
	end string
}

type sumAmount struct {
	sum int
	amount int
}

type UndergroundSystem struct {
	startInfo map[int]Start
	table map[startEnd]sumAmount
}


func Constructor() UndergroundSystem {
	return UndergroundSystem{
		startInfo: make(map[int]Start),
		table: make(map[startEnd]sumAmount),
	}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	this.startInfo[id] = Start{
		stationName,
		t,
	}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	start := this.startInfo[id]

	newKey := startEnd{
		start: start.station,
		end: stationName,
	}

	if newInfo,ok := this.table[newKey];ok {
		newInfo.sum += t - start.t
		newInfo.amount ++
		this.table[newKey] = newInfo
	} else {
		this.table[newKey] = sumAmount{
			t - start.t,
			1,
		}
	}
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	info := this.table[startEnd{startStation,endStation}]
	return float64(info.sum) / float64(info.amount)
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */

/*

class UndergroundSystem {
    class Start {
        String station;
        int time;

        public Start(String station, int time) {
            this.station = station;
            this.time = time;
        }
    }

    class StartEnd {
        String start;
        String end;

        public StartEnd(String start, String end) {
            this.start = start;
            this.end = end;
        }

        public int hashCode() {
            return (start + end).hashCode();
        }

        public boolean equals(Object obj2) {
            if (obj2 instanceof StartEnd) {
                StartEnd startEnd2 = (StartEnd) obj2;
                return this.start.equals(startEnd2.start) && this.end.equals(startEnd2.end);
            }
            return false;
        }
    }

    class SumAmount {
        int sum;
        int amount;

        public SumAmount(int sum, int amount) {
            this.sum = sum;
            this.amount = amount;
        }
    }

    Map<Integer, Start> startInfo;
    Map<StartEnd, SumAmount> table;

    public UndergroundSystem() {
        startInfo = new HashMap<Integer, Start>();
        table = new HashMap<StartEnd, SumAmount>();
    }

    public void checkIn(int id, String stationName, int t) {
        startInfo.put(id, new Start(stationName, t));
    }

    public void checkOut(int id, String stationName, int t) {
        Start start = startInfo.get(id);
        String startStation = start.station;
        int startTime = start.time;
        StartEnd startEnd = new StartEnd(startStation, stationName);
        SumAmount sumAmount = table.getOrDefault(startEnd, new SumAmount(0, 0));
        sumAmount.sum += t - startTime;
        sumAmount.amount++;
        table.put(startEnd, sumAmount);
    }

    public double getAverageTime(String startStation, String endStation) {
        StartEnd index = new StartEnd(startStation, endStation);
        SumAmount sumAmount = table.get(index);
        int sum = sumAmount.sum, amount = sumAmount.amount;
        return 1.0 * sum / amount;
    }
}

 */