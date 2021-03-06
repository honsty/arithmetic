排序算法


------
* [1.交换排序](#1) 
 * [1.1 冒泡排序](#1.1)
 * [1.2 快速排序](#1.2)
* [2.选择排序](#2)
 * [2.1直接选择排序](#2.1)
 * [2.2堆排序](#2.2)
    
<h2 id="1"> 交换排序 </h2>

#### <h3 id="1.1">冒泡排序</h3>
    冒泡排序的思想，是把小的浮上来，大的沉下去。
数组：
![cmd-markdown-logo](/img/bubble.png)


>* 第一步：拿40跟20比，40大于20，不用交换
>* 第二步：拿20跟30比，20小于30，交换
>* 第三步：拿20跟10比，20大于10，不用交换
>* 第四步：拿10跟50比，10小于50，交换

#### <h3 id="1.2">快速排序</h3>
    通过第一遍的遍历(让left和right重合)来找到数组的切割点
数组：

![cmd-markdown-logo](/img/quick.png)

>* 第一步：首先我们从数组的left位置取出(20)作为基准(base)参照物
>* 第二步：从数组的right位置向前找，一直找到比(base)小的数，
    如果找到，将此数赋给left位置(也就是将10赋给20)
    此时数组为：10 ,40 ,50,10,60
    left和right指针分别为前后的10
>* 第三步：从数组的left位置向后找，一直找到比base大的数，
    如果找到，将此数赋给right位置(也就是40赋给10)
    此时数组为：10，40，50，40，60
    left和right指针分别为前后的40
>* 第四步：重复二、三步骤，直到left和right指针重合，
    最后将base插入到40的位置
    此时数组为：10,20,50,40,60,到此完成一次排序
>* 第五步：此时20已经潜到数组的内部，20的左侧一数组都比20小，20的右侧作为一组数都比20大
    以20为切入点对左右两边数按照一、二、三、四步骤进行
   
##<h2 id="2">选择排序</h2>
#### <h3 id="2.1">直接选择排序</h3>
    直接选择排序的思想：找到最小的数，放在第一位，然后找到第二小的放到第二位，依次类推。
    
#### <h3 id="2.2">堆排序</h3>
    可按二叉树进行理解,堆排序有两种情况
![cmd-markdown-logo](/img/heap/heap1.png)

>*  大根堆：父节点比左右子节点都大
>*  小根堆：父节点比左右子节点都小
    
##### 第一：构建大根堆
    如图：
![cmd-markdown-logo](/img/heap/heap2.png)
    这是一个无序的堆，怎么才能构建大根堆呢？
> * 我们发现，这个堆中有2个父节点(2,,3) 
> * 比较2这个父节点的子节点(4,5),发现5大
> * 将较大的子节点(5)跟父节点(2)交换，至此3的左子节点构建完毕
    如图  ![cmd-markdown-logo](/img/heap/heap3.png)
> * 比较第二个父节点(3)下面的左右子节点(5,1),发现5大
> * 父节点(3)与左子节点(5)进行交换
    最后构造的堆如下  ![cmd-markdown-logo](/img/heap/heap4.png)

##### 第二：输出大根堆
    将堆顶(5)与堆尾(2)进行交换，然后将(5)剔除根堆，由于堆顶现在是(2),所以破坏了根堆，必须重新构造，构造完之后又会出现最大值，再次交换和剔除，最后就是我们要的效果了。

