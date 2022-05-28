#include <iostream>
#include <vector>
#include <math.h>
#include <iomanip>
#include <unistd.h>
using namespace std;

//初始化
void init_os_and_user_memory()
{
}

class MemTabItem
{
public:
    int mem_tab_item_ID;   //内存块ID
    uintptr_t begin_ptr;   //内存块起始地址
    int mem_tab_item_size; //内存块大小
};

class MemTab
{
public:
    std::vector<MemTabItem> mem_tab_items;       //数据项
    uintptr_t mem_tab_min;                       //最低地址
    uintptr_t mem_tab_max;                       //最高地址
    virtual int memory_alloc(uint memory_size);  //内存分配
    void memory_recycle(uint memory_item_index); //内存回收
    void memory_degap();                         //内存拼接
};

//输入需分配内存大小, 返回内存表索引
int MemTab::memory_alloc(uint memory_size)
{
    std::cout << "not implemented!" << std::endl; //待具体分配算法实现
    return -1;
}
//输入需回收内存表索引
void MemTab::memory_recycle(uint memory_item_index)
{
    if (memory_item_index == 0)
    {
        cout << "0 内存不允许回收" << endl; //保证数据表不为空, 以免引发指针错误
    }
    MemTabItem new_mem_tab_item = mem_tab_items[memory_item_index];
    cout << "回收内存ID: " << new_mem_tab_item.mem_tab_item_ID << " 内存首地址: " << new_mem_tab_item.begin_ptr << endl;
    mem_tab_items.erase(mem_tab_items.begin() + memory_item_index);
}
//内存拼接
void MemTab::memory_degap()
{
    cout << "内存拼接: " << endl; //遍历内存表, copy memory
    for (int i = 0; i < this->mem_tab_items.size() - 1; i++)
    {
        cout << "memory copy from  " << mem_tab_items[i + 1].begin_ptr;
        mem_tab_items[i + 1].begin_ptr = mem_tab_items[i].begin_ptr + mem_tab_items[i].mem_tab_item_size;
        cout << " to " << mem_tab_items[i].begin_ptr + mem_tab_items[i].mem_tab_item_size << endl;
    }
}

// First Fit 内存分配算法
class MemTabFirstFit : public MemTab
{
public:
    int memory_alloc(uint memory_size);
};
int MemTabFirstFit::memory_alloc(uint memory_size)
{
    for (int i = 0; i < this->mem_tab_items.size(); i++) //遍历内存表
    {
        if (i == mem_tab_items.size() - 1) //最后一项
        {
            if (mem_tab_items[i].begin_ptr + memory_size <= this->mem_tab_max) //是否超出用户区
            {
                MemTabItem new_mem_tab_item; //创建新内存表项
                new_mem_tab_item.begin_ptr = mem_tab_items[i].begin_ptr + mem_tab_items[i].mem_tab_item_size;
                new_mem_tab_item.mem_tab_item_ID = i + 1;
                new_mem_tab_item.mem_tab_item_size = memory_size;
                mem_tab_items.insert(mem_tab_items.begin() + i + 1, new_mem_tab_item);
                cout << "分配内存ID: " << new_mem_tab_item.mem_tab_item_ID
                     << " 内存首地址: " << new_mem_tab_item.begin_ptr
                     << " 内存大小: " << new_mem_tab_item.mem_tab_item_size << endl;
                return i + 1;
            }
        }
        else
        {
            if (mem_tab_items[i].begin_ptr + mem_tab_items[i].mem_tab_item_size + memory_size <= mem_tab_items[i + 1].begin_ptr)
            { // 是否超出下个内存块的首地址
                MemTabItem new_mem_tab_item;
                new_mem_tab_item.begin_ptr = mem_tab_items[i].begin_ptr + mem_tab_items[i].mem_tab_item_size;
                new_mem_tab_item.mem_tab_item_ID = i + 1;
                new_mem_tab_item.mem_tab_item_size = memory_size;
                mem_tab_items.insert(mem_tab_items.begin() + i + 1, new_mem_tab_item);
                cout << "分配内存ID: " << new_mem_tab_item.mem_tab_item_ID
                     << " 内存首地址: " << new_mem_tab_item.begin_ptr
                     << " 内存大小: " << new_mem_tab_item.mem_tab_item_size << endl;
                return i + 1;
            }
        }
    }
    return -1;
}

// Worst Fit 内存分配算法
class MemTabWorstFit : public MemTab
{
public:
    int memory_alloc(uint memory_size);
};
int MemTabWorstFit::memory_alloc(uint memory_size)
{
    int cur_space = -1;    //当前空间大小
    int cur_position = -1; //最大的可插入位置
    for (int i = 0; i < this->mem_tab_items.size(); i++)
    {
        if (i == mem_tab_items.size() - 1) //是否为最后一项
        {
            int space_after_i = mem_tab_max - mem_tab_items[i].begin_ptr;
            if (space_after_i <= memory_size && space_after_i > cur_space) //当前空间更大则变更插入位置
            {
                cur_position = i;
            }
        }
        else
        {
            int space_after_i = mem_tab_items[i + 1].begin_ptr - mem_tab_items[i].begin_ptr;
            if (memory_size <= space_after_i && space_after_i > cur_space)
            { // 是否超出下个内存块的首地址
                cur_position = i;
            }
        }
    }
    if (cur_position != -1)
    {
        MemTabItem new_mem_tab_item;
        new_mem_tab_item.begin_ptr = mem_tab_items[cur_position].begin_ptr + mem_tab_items[cur_position].mem_tab_item_size;
        new_mem_tab_item.mem_tab_item_ID = cur_position + 1;
        new_mem_tab_item.mem_tab_item_size = memory_size;
        mem_tab_items.insert(mem_tab_items.begin() + cur_position + 1, new_mem_tab_item);
        return cur_position + 1;
    }
    return -1;
}

class Controller
{
public:
    MemTab *os_memory_tab;          //内核空间内存表
    MemTab *user_memory_tab;        //用户空间内存表
    void init_memory();             //物理空间初始化
    void handle_process_create();   //处理进程创建事件
    void handle_process_sleep();    //处理进程休眠事件
    void handle_process_activate(); //处理进程激活事件
};

void Controller::init_memory()
{
    this->os_memory_tab->mem_tab_min = 0;
    this->os_memory_tab->mem_tab_max = 128 * pow(2, 20);
    this->user_memory_tab->mem_tab_min = 128 * pow(2, 20);
    this->user_memory_tab->mem_tab_max = 512 * pow(2, 20);
    std::cout << "物理内存初始化: "
              << "os空间: " << hex << this->os_memory_tab->mem_tab_min << "~" << hex << this->os_memory_tab->mem_tab_max << std::endl;
    std::cout << "物理内存初始化: "
              << "user空间: " << hex << this->user_memory_tab->mem_tab_min << "~" << hex << this->user_memory_tab->mem_tab_max << std::endl;
    MemTabItem new_mem_tab_item = *new (MemTabItem);
    MemTabItem new_mem_tab_item2 = *new (MemTabItem);
    this->os_memory_tab->mem_tab_items.push_back(new_mem_tab_item);
    this->user_memory_tab->mem_tab_items.push_back(new_mem_tab_item2);
}
void Controller::handle_process_create()
{
    cout << "创建进程pid: " << rand() % 10000 << " 挂起, 随机分配内存" << endl;
    this->user_memory_tab->memory_alloc(rand() % 100 * 1000);
}
void Controller::handle_process_activate()
{
    cout << "进程激活, 随机申请或释放内存" << endl;
    for (int i = 0; i < 10; i++)
    {
        int x = rand();
        if (x % 2 == 0)
        {
            this->user_memory_tab->memory_alloc(rand() % 100 * 1000);
        }
        else if (this->user_memory_tab->mem_tab_items.size() > 1)
        {
            int y = x % (this->user_memory_tab->mem_tab_items.size() - 1);
            this->user_memory_tab->memory_recycle(y + 1);
        }
    }
}
void Controller::handle_process_sleep()
{
    cout << "进程挂起, 空闲时间完成拼接" << endl;
    this->user_memory_tab->memory_degap();
}
//模拟动态内存管理
int main(int argc, const char **argv)
{
    Controller controller;
    controller.user_memory_tab = new (MemTabFirstFit);
    controller.init_memory();
    cout << "First Fit分配: " << endl;
    controller.handle_process_create();
    controller.handle_process_activate();
    controller.handle_process_sleep();
    controller.handle_process_activate();
    controller.handle_process_sleep();

    controller.user_memory_tab = new (MemTabWorstFit);
    controller.init_memory();
    cout << "Worst Fit分配: " << endl;
    controller.handle_process_create();
    controller.handle_process_activate();
    controller.handle_process_sleep();
    controller.handle_process_activate();
    controller.handle_process_sleep();
    return 0;
}

//统计平均内存利用率
double statics_mem_utilization_rate(Controller controller)
{
    double mem_utilization_rate_average = 0; //平均内存利用率
    vector<double> mem_utilization_rates;    //一秒内内存利用率数组
    //每秒计算一次内存利用率, 然后求平均值.
    while (true)
    {
        vector<MemTabItem> mem_tab_items = controller.user_memory_tab->mem_tab_items;
        int sum = 0;
        //遍历内存表, 求利用内存和
        for (int i = 0; i < mem_tab_items.size(); i++)
        {
            sum += mem_tab_items[i].mem_tab_item_size;
        }
        //求当前内存利用率
        double mem_utilization_rate = sum / (512 - 128) * pow(2, 20);
        mem_utilization_rates.push_back(mem_utilization_rate);
        //休眠1秒
        sleep(1);
    }
    //求内存利用率平均值
    double rates_sum = 0;
    for (int i = 0; i < mem_utilization_rates.size(); i++)
    {
        rates_sum += mem_utilization_rates[i];
    }
    return rates_sum / mem_utilization_rates.size();
}