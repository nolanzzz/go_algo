# GoAlgo
>本仓库包含用Golang解决的Leetcode算法题题解以及笔记，学习路线主要跟随[代码随想录](https://github.com/youngyangyang04/leetcode-master)

题目以数据结构或算法类型来分类，目录如下：

1. [数组](#数组)
2. [哈希](#哈希)
3. [链表](#链表)
4. [栈与队列](#栈与队列)
5. [字符串](#字符串)
6. [二叉树](#二叉树)
7. [回溯算法](#回溯算法)
8. [贪心算法](#贪心算法)
9. [动态规划](#动态规划)
10. [排序](#排序)
11. [其他算法](#其他算法)

## 数组
- [26.删除有序数组中的重复项](Array/26删除有序数组中的重复项.md)
- [27.移除元素](Array/27移除元素.md)
- [34.在排序数组中查找元素的第一个和最后一个位置](Array/34在排序数组中查找元素的第一个和最后一个位置.md)
- [35.搜索插入位置](Array/35搜索插入位置.md)
- [51.螺旋矩阵](Array/51螺旋矩阵.md)
- [59.螺旋矩阵II](Array/59螺旋矩阵II.md)
- [69.x的平方根](Array/69x的平方根.md)
- [76.最小覆盖子串](Array/76最小覆盖子串.md)
- [209.长度最小的子数组](Array/209长度最小的子数组.md)
- [283.移动零](Array/283移动零.md)
- [367.有效的完全平方数](Array/367有效的完全平方数.md)
- [704.二分查找](Array/704二分查找.md)
- [844.比较含退格的字符串](Array/844比较含退格的字符串.md)
- [904.水果成篮](Array/904水果成篮.md)
- [977.有序数组的平方](Array/977有序数组的平方.md)
- [剑指29.顺时针打印矩阵](Array/剑指29顺时针打印矩阵.md)
- [927.三等分](Array/927三等分.md)
- [611.有效三角形的个数](Array/611有效三角形的个数.md)
- [31.下一个排列](Array/31下一个排列.md)

## 哈希
- [1.两数之和](Hash/1两数之和.md)
- [15.三数之和](Hash/15三数之和.md)
- [18.四数之和](Hash/18四数之和.md)
- [202.快乐数](Hash/202快乐数.md)
- [242.有效的字母异位词](Hash/242有效的字母异位词.md)
- [349.两个数组的交集](Hash/349两个数组的交集.md)
- [383.赎金信](Hash/383赎金信.md)
- [454.四数相加II](Hash/454四数相加II.md)
- [347.前k个高频元素](Hash/347前k个高频元素.md)
- [692.前k个高频单词](Hash/692前k个高频单词.md)
- [剑指II60.出现频率最高的k个数字](Hash/剑指II060出现频率最高的%20k%20个数字.md)

## 链表
- [19.删除链表的倒数第n个结点](LinkedList/19删除链表的倒数第n个结点.md)
- [24.两两交换链表中的结点](LinkedList/24两两交换链表中的结点.md)
- [142.环形链表II](LinkedList/142环形链表II.md)
- [203.移除链表元素](LinkedList/203移除链表元素.md)
- [206.反转链表](LinkedList/206反转链表.md)
- [707.设计链表](LinkedList/707设计链表.md)
- [面试题0207.链表相交](LinkedList/面试题0207链表相交.md)

## 栈与队列
- [20.有效的括号](StackQueue/20有效的括号.md)
- [150.逆波兰表达式求值](StackQueue/150逆波兰表达式求值.md)
- [225.用队列实现栈](StackQueue/225用队列实现栈.md)
- [232.用栈实现队列](StackQueue/232用栈实现队列.md)
- [239.滑动窗口最大值](StackQueue/239滑动窗口最大值.md)
- [1047.删除字符串中的所有相邻重复项](StackQueue/1047删除字符串中的所有相邻重复项.md)
- [71.简化路径](StackQueue/71简化路径.md)

## 字符串
- [28.实现strStr()](String/28实现strStr().md)
- [151.反转字符串里的单词](String/151反转字符串里的单词.md)
- [344.反转字符串](String/344反转字符串.md)
- [459.重复的子字符串](String/459重复的子字符串.md)
- [541.反转字符串II](String/541反转字符串II.md)
- [剑指05.替换空格](String/剑指05替换空格.md)
- [剑指58II.左旋转字符串](String/剑指58II左旋转字符串.md)

## 二叉树

- [二叉树的理论基础](Tree/二叉树理论基础.md)

### 二叉树的遍历方式
- DFS-深度优先遍历
  - 迭代法：通过栈来模拟递归
  - [144.二叉树的前序遍历](Tree/DFS/144二叉树的前序遍历.md)
  - [94.二叉树的中序遍历](Tree/DFS/94二叉树的中序遍历.md)
  - [145.二叉树的后序遍历](Tree/DFS/145二叉树的后序遍历.md)
  - [二叉树的统一迭代法](Tree/DFS/二叉树的统一迭代法.md)
- BFS-广度优先遍历
  - [102.二叉树的层序遍历](Tree/BFS/102二叉树的层序遍历.md)：通过队列模拟
  - 通过BFS解决的题目：
  - [107.二叉树的层序遍历II](Tree/BFS/107二叉树的层序遍历II.md)
  - [116.填充每个节点的下一个右侧节点指针](Tree/BFS/116填充每个节点的下一个右侧节点指针.md)
  - [117.填充每个节点的下一个右侧节点指针II](Tree/BFS/117填充每个节点的下一个右侧节点指针II.md)
  - [199.二叉树的右视图](Tree/BFS/199二叉树的右视图.md)
  - [429.N叉树的层序遍历](Tree/BFS/429N叉树的层序遍历.md)
  - [515.在每个树行中找最大值](Tree/BFS/515在每个树行中找最大值.md)
  - [559.N叉树的最大深度](Tree/BFS/559N叉树的最大深度.md)
  - [637.二叉树的层平均值](Tree/BFS/637二叉树的层平均值.md)

### 求二叉树的属性
- [101.二叉树：是否对称](Tree/101对称二叉树.md)，[100.二叉树：是否相同](Tree/100相同的树.md)
  - 递归：后序，比较的是根节点的左子树和右子树是不是相互反转
  - 迭代：使用栈/队列将两个节点顺序放入容器中进行比较
- [572.二叉树：是否为另一棵树的子树](Tree/572另一棵树的子树.md)
  - 递归：查看是否是相同的树，递归
- [104.二叉树：求最大深度](Tree/BFS/104二叉树的最大深度.md)
  - 递归：后序，求根节点最大高度就是最大深度，通过递归函数的返回值计算树的高度（左树和右树的最大高度+1）
  - 迭代：层序，层数就是最大深度
- [111.二叉树：求最小深度](Tree/BFS/111二叉树的最小深度.md)
  - 递归：后序，求根节点的最小高度就是最小深度，注意最小深度的定义
  - 迭代：层序
- [222.二叉树：求有多少个节点](Tree/222完全二叉树的节点个数.md)
  - 递归：后序，通过递归函数的返回值计算节点数量
  - 递归2：如果是完全二叉树，左右指针分别遍历至最后一层，利用特性来计算
  - 迭代：层序
- [110.二叉树：是否平衡](Tree/110平衡二叉树.md)
  - 递归：后序，先求左右子树的高度，差值大于1时直接返回-1
- [257.二叉树：找所有路径](Tree/257二叉树的所有路径.md)
  - 递归：前序，path和res数组分别记录目前path和最终结果。注意回溯的思想
- [404.二叉树：求左叶子之和](Tree/404左叶子之和.md)
  - 递归：后序，注意判断左叶子的条件
  - 迭代：直接用栈模拟后序
- [513.二叉树：找左下角的值](Tree/513找树左下角的值.md)
  - 递归：需要记录最大深度，来判断是否为最后一层
  - 迭代：层序遍历找到最后一行的第一个值
- [112.二叉树：求路径总和](Tree/112路径总和.md)，[路径总和II](Tree/113路径总和II.md)
  - 递归：顺序无所谓，递归函数返回值为bool是为了搜索一边，没有返回值是搜索整棵树

### 二叉树的修改与改造
- [226.反转二叉树](Tree/226反转二叉树.md)
  - 递归：前序，交换左右孩子
  - 迭代：前序或层序都行
- [106.中序与后序构造二叉树](Tree/106从中序与后序遍历序列构造二叉树.md)，[105.前序与中序构造二叉树](Tree/105从前序与中序遍历序列构造二叉树.md)
  - 递归：前序，以后序遍历序列的最后一个节点为root，寻找前/中序和后序中的切割点分别构造左右区间
- [654.构造最大的二叉树](Tree/654最大二叉树.md)
  - 递归：前序，找到数组的最大值作为切割点，构造左右区间
- [617.合并两个二叉树](Tree/617合并二叉树.md)
  - 递归：前序，同时操作两个节点，可以直接用其中一个节点而不用新建root
  - 迭代：层序，每次往队列中放入两个节点

### 求二叉搜索树的属性
- BST的遍历不需要额外存储空间，因为无需回溯
- [700.二叉搜索树中的搜索](Tree/700二叉搜索树中的搜索.md)
  - 递归：按大小判断递归
  - 迭代：按大小判断遍历
- [98.是不是二叉搜索树](Tree/98验证二叉搜索树.md)
  - 递归1：中序遍历，判断生成的数组是否为一个递增有序数组
  - 递归2：无需数组，递归时更新min和max值来判断
- [530.求二叉搜索树的最小绝对差](Tree/530二叉搜索树的最小绝对值.md)
  - 递归：中序，用pre来记录前一个结点方便计算差值，双指针
  - 迭代：中序，相同逻辑
- [501.求二叉搜索树中的众数](Tree/501二叉搜索树中的众数.md)
  - 递归：中序，count和maxCount,更新maxCount时清空已存在的res
  - 迭代：中序，相同逻辑
- [538.把二叉搜索树转换为累加树](Tree/538把二叉搜索树转换为累加树.md)
  - 递归：中序（右中左），双指针pre来操作累加
  - 迭代：栈模拟中序，相同逻辑

### 二叉树公共祖先问题
- [236.二叉树的最近公共祖先](Tree/236二叉树的最近公共祖先.md)
  - 递归：后序，然后判断几种情况
- [235.二叉搜索树的最近公共祖先](Tree/235二叉搜索树的最近公共祖先.md)
  - 递归：根据root的数值比较决定递归方向
  - 迭代：直接指针遍历，与递归相同逻辑

### 二叉搜索树的修改与改造
- [701.二叉搜索树中的插入操作](Tree/701二叉搜索树中的插入操作.md)
  - 递归：找到null值时插入
  - 迭代：遍历找到位置插入
- [450.二叉搜索树中的删除操作](Tree/450删除二叉搜索树中的节点.md)
  - 递归：前序，重平衡时用右树的最小值作为新的root值
- [669.修剪二叉搜索树](Tree/669修剪二叉搜索树.md)
  - 递归：前序
- [108.构造二叉搜索树](Tree/108将有序数组转换为二叉搜索树.md)
  - 递归：前序，用数组中间节点作为root进行区间分隔

## 回溯算法

- [回溯算法理论基础](Backtracking/回溯算法理论基础.md)

### 组合问题
- 组合问题其实也是一种子集问题
- [77.组合](Backtracking/77组合.md)
  - 回溯模板，res，path，start递归
- [216.组合总和III](Backtracking/216组合总和III.md)
  - 回溯模板，res，path，start递归
  - 还需要sum记录总和
- [17.电话号码的字母组合](Backtracking/17电话号码的字母组合.md)
  - 回溯模板，res，path，index递归
  - 数组来储存数字与字母的映射
- [39.组合总和](Backtracking/39组合总和.md)
  - 回溯模板，res，path，index递归
  - 由于可重复，回溯时用i而不是i+1
  - 若先排序数组，可以剪枝
- [40.组合总和II](Backtracking/40组合总和II.md)
  - 回溯模板，res，path，index递归
  - 不可重复，回溯时用i+1
  - 需要去重操作，首先要排序数组
  - 然后递归逻辑中每次要去重（candidates[i] == candidates[i-1])
  - 可以剪枝

### 分割问题
- 分隔与组合相似，组合选择变成了切割点
- [131.分割回文串](Backtracking/131分隔回文串.md)
  - 回溯模板，res，path，start递归
  - 判断回文，不符合时continue
  - 终止条件为到达末尾
- [93.复原IP地址](Backtracking/93复原IP地址.md)
  - 回溯模板，res，path，start递归
  - 判断每段地址是否合法
  - 终止条件为1.到达末尾；2.path有四段地址

### 子集问题
- 求取子集问题，不需要任何剪枝！因为子集就是要遍历整棵树
- [78.子集](Backtracking/78子集.md)
  - 回溯模板，res，path，start递归
  - 每次递归先保存当前path至res
  - 终止条件可以省略，因为for循环遍历至最后一个数字时自然结束
- [90.子集II](Backtracking/90子集II.md)
  - 与78.子集一个套路
  - 需要去重操作，首先要排序数组
  - 然后递归逻辑中每次要去重
- [491.递增子序列](Backtracking/491递增子序列.md)
  - 回溯模板，res，path，start递归
  - 根据题目条件，当path有两个元素时即可保存，保存后继续往后走
  - 需用used来记录本层用过的数字，以免重复搜索同样的数字开头的子序列

### 排列问题
- 排列问题是有序的，所以每次循环都要从0开始，不需要start参数
- 去重时需要用used来记录
- [46.全排列](Backtracking/46全排列.md)
  - 回溯模板，res，path
  - used记录出现过的元素，同样要回溯
- [47.全排列II](Backtracking/47全排列II.md)
  - 回溯模板，res，path
  - used记录出现过的元素，同样要回溯
  - 需要去重：
    1. 已经搜索过以第i个数字开头的排列，跳过
    2. 重复的数字，规定按从左到右的顺序搜索，还没有搜索过第i-1个数时跳过

### 重新安排行程（图论额外拓展）
- [332.重新安排行程](Backtracking/332重新安排行程.md)
  - 回溯模板，res，不需要path了，因为此时res相当于path，找到一条符合要求的path时便结束了
  - 需要一个map来记录起点和终点的映射
  - map需要根据终点排序，这样最终的结果才是根据字典排序的
  - 终止条件：当res的路径数等于tickets数加一时结束

### 棋盘问题
- [51.N皇后](Backtracking/51N皇后.md)
  - 回溯模板，board来记录棋盘状态，row记录当前的行
  - 递归遍历行
  - 递归里的for遍历列
  - 终止条件：当row = n时，完成了所有行
  - 检查是否合法时，查看列、45度和135度上斜线即可
- [37.解数独](Backtracking/37解数独.md)
  - 回溯模板，board记录棋盘状态
  - 需要遍历所有行所有列，以及1-9的值，找到当前cell合适的值直接返回true
  - 没有合适的值返回false
  - 检查是否合法时，查看行、列、以及九宫格

## 贪心算法

- [贪心算法理论基础](Greedy/贪心算法理论基础.md)
### 简单题
- [455.分发饼干](Greedy/455分发饼干.md)
  - 局部最优：大饼干喂给胃口大的孩子，不造成饼干尺寸的浪费
  - 全局最优：喂饱尽可能多的孩子
  - 排序饼干和孩子胃口
  - 双指针
- [1005.K次取反后最大化的数组和](Greedy/1005K次取反后最大化的数组和.md)
  - 局部最优：将绝对值最大的负数变为正，从而最大化当前数值
  - 整体最优：将尽可能多的负数取反为正，或将最小的正数取反，最大化总和
  - 第一步 按绝对值从大到小排序
  - 第二步 从头开始遍历，遇到负数便取反操作，k--
  - 第三步 如k还大于0，则继续操作取反绝对值最小的那个数
  - 第四部 求和
- [860.柠檬水找零](Greedy/860柠檬水找零.md)
  - 局部最优：碰到20的账单，优先消耗10元面值，再消耗5元，因为5元能用来找零的场景更多
  - 全局最优：完成全部找零
- [908.最小差值I](Greedy/908最小差值I.md)
  - 找出nums中的最小值和最大值
  - 它们最多可以分别增加和减少k
  - 最小差值则是(max-k)-(min+k)，若得负数则返回0

### 中等题
- [376.摆动序列](Greedy/376摆动序列.md)
  - 无需删除操作，只用记录符合条件的节点数，也就是数组的峰值数量
  - 记录currDiff和preDiff用来比较
- [738.单调递增的数字](Greedy/738单调递增的数字.md)
  - 需从后往前遍历，然后一个flag记录从哪里开始赋值为9，例：987 -> flag为1 -> 899
  - 局部最优：当遇到num[i-1] > num[i]的情况，将num[i-1]减1，num[i]赋值为9，可以保证这两位变成最大的单调递增整数
  - 全局最优：得到小于等于N的最大单增整数

#### 贪心解决股票问题
- [122.买卖股票的最佳时机II](Greedy/122买卖股票的最佳时机II.md)
  - 把利润分解为每天为单位的维度，而不是从0天到第n天整体去考虑
  - 局部最优：收集每天的正利润，也就是只收集当利润为正时
  - 全局最优：求得最大利润
- [714.买卖股票的最佳时机含手续费](Greedy/714买卖股票的最佳时机含手续费.md)
  - 局部最优：最低值买，最高值卖（算上手续费后依然盈利）
  - 全局最优：利润最大
  - 股票问题常规解法是用动态规划

#### 两个维度权衡问题
在出现两个维度相互影响时，不要一起考虑，先确定一个维度，再确定另一个维度
- [135.分发糖果](Greedy/135分发糖果.md)
  - 两次贪心
    - 第一次：从左到右遍历，确保评分高的右孩子比左孩子的糖多
    - 第二次：从右到左遍历，确保评分高的左孩子比右孩子的糖多（贪心：既比左边多也比右边多）
  - 全局最优：中间的孩子比两边的孩子糖多
- [406.根据身高重建队列](Greedy/406根据身高重建队列.md)
  - 先身高h从高到矮排序
  - 局部最优：优先按身高高的people的k来插入，插入后的people满足队列属性
  - 全局最优：全部插入，并符合属性

### 贪心难题

#### 贪心解决区间问题
各种覆盖各种去重
- [55.跳跃游戏](Greedy/55跳跃游戏.md)
  - 判断最大的跳跃范围是否能大于等于数组最后一位
  - 局部最优：每次取最大的跳跃范围
  - 全局最优：最后得到整体能跳到的最大范围
- [45.跳跃游戏II](Greedy/45跳跃游戏II.md)
  - 需要统计两个覆盖范围：当前这一步的最大覆盖 和 下一步的最大覆盖
  - 局部最优：当前可移动距离尽量多走，如果还没到终点，步数再加一
  - 整体最优：每一步都尽可能多走，达到最小步数
- [452.用最少数量的箭引爆气球](Greedy/452用最少数量的箭引爆气球.md)
  - 按左边界（起点）排序，从左向右遍历
  - 局部最优：当气球重叠时，一起射，所用弓箭最少
  - 全局最优：把所有气球射爆所用的箭最少
- [435.无重叠区间](Greedy/435无重叠区间.md)
  - 按右边界（终点）排序，从左向右遍历
  - 局部最优：优先选右边界小的区间，所以从左向右遍历，留给下一个区间的空间更大，从而尽量避免交叉
  - 全局最优：选取最多的非交叉区间
- [763.划分字母区间](Greedy/763.划分字母区间.md)
  - 统计每个字母的最远边界
  - 从头遍历，更新当前遍历到的字母的最远边界，如果最远边界与当前坐标重合，则为切割点
  - 不算是严格的贪心，因为不存在局部最优
- [56.合并区间](Greedy/56合并区间.md)
  - 按左边界排序，从左向右遍历
  - 局部最优：每次合并都取最大的右边界，可以合并更多的区间
  - 整体最优：合并所有重叠的区间

#### 其他难题
- [53.最大子数组和](Greedy/53最大子数组和.md)
  - 遍历每次将nums[i]更新为当前子数组的最大和，并判断是否大于max
  - 局部最优：当前连续和为负数的时候立刻放弃（不跟新nums[i]），从下一个元素重新计算连续和，因为负数加上下一个元素只会越来越小
  - 全局最优：选取最大连续和
- [134.加油站](Greedy/134加油站.md)
  - curSum记录当前起始点start的累加，totalSum记录全部的累加
  - 遍历，curSum小于0时起始点start加1
  - 若totalSum小于0，则不可能跑完一圈
  - 局部最优：curSum小于0时探索下一个起始点
  - 全局最优：找到可以跑一圈的起始点
- [968.监控二叉树](Greedy/968监控二叉树.md)
  - 交叉类难题
  - DFS，从叶节点向根节点递归遍历
  - 每次递归返回状态码：0：没有覆盖，1：有摄像头，2：已覆盖
  - 局部最优：让叶节点的父节点安装摄像头，每隔两个节点放一个摄像头
  - 全局最优：摄像头最少
  
## 动态规划

- [动态规划理论基础](dp/动态规划理论基础.md)

- [509.斐波那契数](dp/509斐波那契数.md)
  - dp[i]：第i个斐波那契数
  - 递推公式：dp[i] = dp[i-1] + dp[i-2]
  - 初始化：dp[0] = 1, dp[1] = 1
  - 遍历顺序：从前往后
- [70.爬楼梯](dp/70爬楼梯.md)
  - dp[i]：爬到第i阶有几种方法
  - 递推公式：dp[i] = dp[i-1] + dp[i-2]
  - 初始化：dp[0]不用考虑，dp[1] = 1, dp[2] = 2
  - 遍历顺序：从前往后
- [746.使用最小花费爬楼梯](dp/746使用最小花费爬楼梯.md)
  - dp[i]：爬到第i阶的最小花费
  - dp[0] = cost[0], dp[1] = cost[1]
- [62.不同路径](dp/62不同路径.md)
  - dp[i][j]：到达ixj网格有多少条不同的路径
  - 初始化：所有行的第一列 和 所有列的第一行 都为1
  - 两层遍历
- [63.不同路径II](dp/63不同路径II.md)
  - dp[i][j]：到达ixj网格有多少条不同的路径
  - 初始化：所有行的第一列 和 所有列的第一行 （没有障碍的） 都为1
  - 两层遍历
- [343.整数拆分](dp/343整数拆分.md)
  - dp[i]：整数i的最大拆分乘积
  - dp[2]=1
  - 两层遍历，每次用j * (i-j) 和 j * dp[i-j]的较大值与已经dp[i]比较取最大值
- [96.不同的二叉搜索树](dp/96不同的二叉搜索树.md)
  - dp[i]:1到i为节点组成的二叉搜索树的个数为dp[i]
  - dp[0] = 1
  - 两层遍历，每次要累加j从1到i作为根节点的树数量

### 背包问题
#### 0-1背包
- [0-1背包理论基础](dp/0-1背包理论基础.md)
- [416.分割等和子集](dp/416分割等和子集.md)
  - 一维数组
    - dp[j]：容量为j的背包能放的最大数值和
    - 初始化：全部为0
    - 两层遍历（j要倒序）
    - dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
  - 二维数组
    - dp[i][j]：从0-i下标的数值里任选放入容量为j的书包的最大数值和
    - 初始化：dp[i][0]=0; 当j>=nums[0]时dp[0][j]=nums[0]
    - 两层遍历
    - dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])

## 排序
- [快速排序](Sorting/快速排序.md)

## 其他算法
- [KMP算法](SpecialAlgorithms/kmp.go)