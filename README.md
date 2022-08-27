# golang_min_stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:

- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element `val` onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

You must implement a solution with `O(1)` time complexity for each function.

## Examples

**Example 1:**

```
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2

```

**Constraints:**

- `-$2^{31}$ <= val <= $2^{31}$ - 1`
- Methods `pop`, `top` and `getMin` operations will always be called on **non-empty** stacks.
- At most `3 * $10^4$` calls will be made to `push`, `pop`, `top`, and `getMin`.

## 解析

要求設計一個資料結構 MinStack

具有以下方法：

1. Constructor: 用來初始化 MinStack 結構
2. void pop():  remove stack 最上面的值
3. push(int val): 把目前值放入 stack 最上層
4. int top(): 取得目前 stack 最上面的值
5. int getMin(): 取得目前 stack 中最小值

要求這些方法都必須要在 Constant time 的時間複雜度內完成

如果使用 LinkedList , pop, push, top 可以使用以下實作來達成時間在 Constant time

建立 Node 資料結構如下：

type Node struct {

   val int

   next *Node

}

而 LinkedList 定義如下

type LinkedList struct {

   head *Node

}

1. push: 每次把新增的資料 更新為 LinkedList 的 head, 把原本的 head 改成 head.next 指標所指向的值

 ![](https://i.imgur.com/PK850bh.png)

2. pop: 更新 head 為 head.next 再回傳 popVal

![](https://i.imgur.com/ynC8TZx.png)

3.  top: 回傳 head.val

![](https://i.imgur.com/sKqOAC9.png)

以上3個都只要 Constant time 因為都是直接透過物件指標操作

困難的部份在於 getMin 如果用現在這種作法會需要 O(n) 找完整個 linkedList 才能找到

要變成 Constant time

其實只要修改 Node 結構如下

type Node struct {

   val, min int

   next *Node

}

讓每個結點紀錄當下的 min 值

每次 push 的時候 把前一個解點的 min 跟目前 val 找出最小作為當下結點的 min

![](https://i.imgur.com/UXQ86Hf.png)


透過這個值，每次存在 head.min 的值都是當下整個 stack 的最小值

因為存取 head 是 constant time 所以 getMin 也是 constant time

## 程式碼
```go
package sol

type Node struct {
	val, min int
	next     *Node
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{}
}
func (this *MinStack) Push(val int) {
	node := &Node{val: val, min: val, next: nil}
	if this.head != nil {
		if node.min > this.head.min {
			node.min = this.head.min
		}
		node.next = this.head
	}
	this.head = node
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

```
## 困難點

1. 需要理解透過多存一個 min 來達到 constant time 找到最小值

## Solve Point

- [x]  建立一個 Node 資料結構。內部具有 int 變數 val 用來儲存結點值, int 變數 min 用來儲存當下 stack 最小值, Node 指標變數 next 用來儲存下一個Node位址
- [x]  建立一個 MinStack 資料結構。內部具有 Node 指標變數 head 用來儲存目前 stack 最上方的 Node 位址