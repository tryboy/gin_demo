## 抽象工厂接口

1. ### 结构组成
   1. 产品线工厂 - 生产具有共同特征的一些了不同产品
   2. 具体产品工厂 - 多个产品线下的最总产品产出地
   3. 产品接口 - 约定产品的基本功能
   4. 组装工厂 - 根据需求返回产品线，再根据进一步需求获取产品(可理解为申请某条产品线去生产某件产品)
2. ### 经验结合
   1. 所谓的抽象，就是在工厂方法的基础上，所生产的产品不再是具体的产品，而是抽象话的产品
   2. 主体程序，原本应该获取一个实体化的产品（即返回一个产品类），变为返回一个抽象化的产品(可理解未返回了一个更具体的工厂类)，然后再用相应的需求（方法）获取相应的产品
3. ### 他山之石