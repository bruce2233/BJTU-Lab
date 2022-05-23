#include<iostream>
#include<cstring>
#include<vector>
#include <unistd.h>
#include <iterator>
using namespace std;
void test2(){
    vector<int> list(100);
    while(true){
        list.push_back(1);
        std::cout << &list[0] << std::endl;
        sleep(1);
    }
}
void test3(){
    vector<int> vec;
    vec.push_back(0);
    vec.push_back(1);
    vector<int>::iterator it = vec.begin();
    it++;
    vec.erase(it);
    vec.erase(it);

}

int main()
{
//    test2();
    test3();
    // B b;
    // b.run();
}

class A{
public:
    void run(){
        cout<<"from A"<<endl;
    }
};

class B: public A{
public:
    void run(int x){
        cout<<"from B" <<endl;
    }
};


void test1(){
     vector<int> v;  //v是存放int类型变量的可变长数组，开始时没有元素
    // for (int n = 0; n<5; ++n)
    //     v.push_back(n);  //push_back成员函数在vector容器尾部添加一个元素
    vector<int>::iterator i;  //定义正向迭代器
    if (v.begin()==v.end()){
        std::cout << "zere equal" << std::endl;
    }
    for (i = v.begin(); i != v.end(); ++i) {  //用迭代器遍历容器
        cout <<"h" <<endl;
        cout << *i << " ";  //*i 就是迭代器i指向的元素
        *i *= 2;  //每个元素变为原来的2倍
    }
    cout << endl;
    //用反向迭代器遍历容器
    for (vector<int>::reverse_iterator j = v.rbegin(); j != v.rend(); ++j)
        cout << *j << " ";
}

