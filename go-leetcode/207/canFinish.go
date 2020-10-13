package main

import "fmt"

/*
207. 课程表
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

示例 1:

输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
示例 2:

输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。

提示：

输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5
 */

func main() {
	a := 2
	matrix := [][]int{{1,0}}
	fmt.Println(canFinish(a,matrix))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	valid := true

	visited := make([]int,numCourses)

	edges := make([][]int,numCourses)
	for i := 0;i<numCourses;i++ {
		edges[i] = make([]int,0)
	}

	for i := 0;i<len(prerequisites);i++ {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]],prerequisites[i][0])
	}

	for i := 0;i<numCourses && valid;i++ {
		if visited[i] == 0 {
			dfs(i,edges,visited,&valid)
		}
	}
	return valid
}

func dfs(i int,edges [][]int,visited []int,valid *bool) {
	visited[i] = 1
	for j:=0;j<len(edges[i]);j++ {
		if visited[edges[i][j]] == 0 {
			dfs(edges[i][j],edges,visited,valid)
			if !(*valid) {
				return
			}
		} else if visited[edges[i][j]] == 1 {
			*valid = false
			return
		}
	}
	visited[i] = 2
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	indeg := make([]int,numCourses)

	edges := make([][]int,numCourses)
	for i := 0;i<numCourses;i++ {
		edges[i] = make([]int,0)
	}

	for i := 0;i<len(prerequisites);i++ {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]],prerequisites[i][0])
		indeg[prerequisites[i][0]] = indeg[prerequisites[i][0]] + 1
	}

	queue := []int{}
	for i := 0;i<numCourses;i++ {
		if indeg[i] == 0 {
			queue = append(queue,i)
		}
	}

	visited := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited ++
		for j := 0;j<len(edges[u]);j++ {
			indeg[edges[u][j]] = indeg[edges[u][j]] - 1
			if indeg[edges[u][j]] == 0 {
				queue = append(queue,edges[u][j])
			}
		}
	}
	return visited == numCourses
}

/*
拓扑排序（Topological Sorting）是一个有向无环图（DAG, Directed Acyclic Graph）的所有顶点的线性序列。
且该序列必须满足下面两个条件：

每个顶点出现且只出现一次。
若存在一条从顶点 A 到顶点 B 的路径，那么在序列中顶点 A 出现在顶点 B 的前面。
有向无环图（DAG）才有拓扑排序，非DAG图没有拓扑排序一说。

方法一：深度优先搜索
思路

我们可以将深度优先搜索的流程与拓扑排序的求解联系起来，用一个栈来存储所有已经搜索完成的节点。

对于一个节点 u，如果它的所有相邻节点都已经搜索完成，那么在搜索回溯到 u 的时候，u 本身也会变成一个已经搜索完成的节点。这里的「相邻节点」指的是从 u 出发通过一条有向边可以到达的所有节点。

假设我们当前搜索到了节点 u，如果它的所有相邻节点都已经搜索完成，那么这些节点都已经在栈中了，此时我们就可以把 u入栈。
可以发现，如果我们从栈顶往栈底的顺序看，由于 u 处于栈顶的位置，那么 u 出现在所有 u 的相邻节点的前面。因此对于 u 这个节点而言，它是满足拓扑排序的要求的。

这样以来，我们对图进行一遍深度优先搜索。当每个节点进行回溯的时候，我们把该节点放入栈中。最终从栈顶到栈底的序列就是一种拓扑排序。

算法

对于图中的任意一个节点，它在搜索的过程中有三种状态，即：

「未搜索」：我们还没有搜索到这个节点；

「搜索中」：我们搜索过这个节点，但还没有回溯到该节点，即该节点还没有入栈，还有相邻的节点没有搜索完成；

「已完成」：我们搜索过并且回溯过这个节点，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求。

通过上述的三种状态，我们就可以给出使用深度优先搜索得到拓扑排序的算法流程，在每一轮的搜索搜索开始时，我们任取一个「未搜索」的节点开始进行深度优先搜索。

我们将当前搜索的节点 u 标记为「搜索中」，遍历该节点的每一个相邻节点 v：

如果 v 为「未搜索」，那么我们开始搜索 v，待搜索完成回溯到 u；

如果 v 为「搜索中」，那么我们就找到了图中的一个环，因此是不存在拓扑排序的；

如果 v 为「已完成」，那么说明 v 已经在栈中了，而 u 还不在栈中，因此 u 无论何时入栈都不会影响到 (u, v) 之前的拓扑关系，以及不用进行任何操作。

当 u 的所有相邻节点都为「已完成」时，我们将 u 放入栈中，并将其标记为「已完成」。

在整个深度优先搜索的过程结束后，如果我们没有找到图中的环，那么栈中存储这所有的 n 个节点，从栈顶到栈底的顺序即为一种拓扑排序。

class Solution {
    List<List<Integer>> edges;
    int[] visited;
    boolean valid = true;

    public boolean canFinish(int numCourses, int[][] prerequisites) {
        edges = new ArrayList<List<Integer>>();
        for (int i = 0; i < numCourses; ++i) {
            edges.add(new ArrayList<Integer>());
        }
        visited = new int[numCourses];

        for (int[] info : prerequisites) {
            edges.get(info[1]).add(info[0]);
        }

        for (int i = 0; i < numCourses && valid; ++i) {
            if (visited[i] == 0) {
                dfs(i);
            }
        }
        return valid;
    }


    public void dfs(int u) {
        visited[u] = 1;
        for (int v: edges.get(u)) {
            if (visited[v] == 0) {
                dfs(v);
                if (!valid) {
                    return;
                }
            } else if (visited[v] == 1) {
                valid = false;
                return;
            }
        }
        visited[u] = 2;
    }
}

方法二: 广度优先搜索

方法一的深度优先搜索是一种「逆向思维」：最先被放入栈中的节点是在拓扑排序中最后面的节点。
我们也可以使用正向思维，顺序地生成拓扑排序，这种方法也更加直观。

我们考虑拓扑排序中最前面的节点，该节点一定不会有任何入边，也就是它没有任何的先修课程要求。当我们将一个节点加入答案中后，
我们就可以移除它的所有出边，代表着它的相邻节点少了一门先修课程的要求。
如果某个相邻节点变成了「没有任何入边的节点」，那么就代表着这门课可以开始学习了。
按照这样的流程，我们不断地将没有入边的节点加入答案，
直到答案中包含所有的节点（得到了一种拓扑排序）或者不存在没有入边的节点（图中包含环）。

上面的想法类似于广度优先搜索，因此我们可以将广度优先搜索的流程与拓扑排序的求解联系起来。

算法

我们使用一个队列来进行广度优先搜索。初始时，所有入度为 0 的节点都被放入队列中，
它们就是可以作为拓扑排序最前面的节点，并且它们之间的相对顺序是无关紧要的。

在广度优先搜索的每一步中，我们取出队首的节点 u：

我们将 u 放入答案中；

我们移除 u 的所有出边，也就是将 u 的所有相邻节点的入度减少 1。
如果某个相邻节点 v 的入度变为 0，那么我们就将 v 放入队列中。

在广度优先搜索的过程结束后。如果答案中包含了这 n 个节点，
那么我们就找到了一种拓扑排序，否则说明图中存在环，也就不存在拓扑排序了。

class Solution {
    List<List<Integer>> edges;
    int[] indeg;

    public boolean canFinish(int numCourses, int[][] prerequisites) {
        edges = new ArrayList<List<Integer>>();
        for (int i = 0; i < numCourses; ++i) {
            edges.add(new ArrayList<Integer>());
        }
        indeg = new int[numCourses];
        for (int[] info : prerequisites) {
            edges.get(info[1]).add(info[0]);
            ++indeg[info[0]];
        }


        Queue<Integer> queue = new LinkedList<Integer>();
        for (int i = 0; i < numCourses; ++i) {
            if (indeg[i] == 0) {
                queue.offer(i);
            }
        }

        int visited = 0;
        while (!queue.isEmpty()) {
            ++visited;
            int u = queue.poll();
            for (int v: edges.get(u)) {
                --indeg[v];
                if (indeg[v] == 0) {
                    queue.offer(v);
                }
            }
        }

        return visited == numCourses;
    }
}

 */