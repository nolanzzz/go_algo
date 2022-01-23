# GoAlgo
本仓库包含用Golang解决的Leetcode算法题题解以及笔记，学习路线主要跟随[代码随想录](https://github.com/youngyangyang04/leetcode-master)<br>
题目以数据结构或算法类型来分类，目录如下：

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

## 哈希
- [1.两数之和](Hash/1两数之和.md)
- [15.三数之和](Hash/15三数之和.md)
- [18.四数之和](Hash/18四数之和.md)
- [202.快乐数](Hash/202快乐数.md)
- [242.有效的字母异位词](Hash/242有效的字母异位词.md)
- [349.两个数组的交集](Hash/349两个数组的交集.md)
- [383.赎金信](Hash/383赎金信.md)
- [454.四数相加II](Hash/454四数相加II.md)

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
- [347.前k个高频元素](StackQueue/347前k个高频元素.md)
- [1047.删除字符串中的所有相邻重复项](StackQueue/1047删除字符串中的所有相邻重复项.md)

## 字符串
- [28.实现strStr()](String/28实现strStr().md)
- [151.反转字符串里的单词](String/151反转字符串里的单词.md)
- [344.反转字符串](String/344反转字符串.md)
- [459.重复的子字符串](String/459重复的子字符串.md)
- [541.反转字符串II](String/541反转字符串II.md)
- [剑指05.替换空格](String/剑指05替换空格.md)
- [剑指58II.左旋转字符串](String/剑指58II左旋转字符串.md)

## 二叉树

### 二叉树的理论基础
- [理论基础](Tree/Tree.md)

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

### 回溯算法理论基础
- [理论基础](Backtracking/notes.md)

### 组合
- [77.组合](Backtracking/77组合.md)
  - 回溯模板，res，path，start递归
- [216.组合总和III](Backtracking/216组合总和III.md)
  - 回溯模板，res，path，start递归，还需要sum记录总和
- [17.电话号码的字母组合](Backtracking/17电话号码的字母组合.md)
  - 回溯模板，res，path，index递归，数组来储存数字与字母的映射

## 排序
- [快速排序](Sorting/快速排序.md)

## 其他算法
- [KMP算法](SpecialAlgorithms/kmp.go)