#include<iostream>
#include<cstring>
#include<vector>
#include <unistd.h>
#include <iterator>

using namespace std;

const int TIME_SLICE = 1500;



class ITask{
public:
    virtual void run()=0;
    int pid;
};

class Process: public ITask{
public:
    int status;
    void run();
    void stop();
};
void Process::run(){
        std::cout << "process " <<this->pid<<" is running..." << std::endl;
}
void Process::stop(){
    std::cout << "process "<<this->pid <<" stops"<< std::endl;
}

class ISchedule {
public:
    virtual void inter_call() = 0;
    virtual void call(ITask* task)=0;
    virtual void set_clock(int time) = 0;
};

class Schedule : public ISchedule{
public:
    void set_clock(int time);
    void inter_call();
    void call(ITask* task);
    void schedule_task();
};
void Schedule::set_clock(int time){
        std::cout << "after "<<time<<"ms another task runs" << std::endl;
}

void Schedule::call(ITask* task){
    task->run();
}
void Schedule::schedule_task(){
    std::cout << "schedule_task" << std::endl;
}

void Schedule::inter_call(){
        std::cout << "called by the clock interupt..." << std::endl;
}

class ScheduleWithTaskList : public Schedule{
public:
    vector<ITask *> task_queue;
    vector<ITask *>::iterator it;
    void add_task(ITask* task);
    void remove_task(vector<ITask *>::iterator &it);
    void show_task_queue();
};
void ScheduleWithTaskList::add_task(ITask* task){
    this->task_queue.push_back(task);
}
void ScheduleWithTaskList::remove_task(vector<ITask*>::iterator &ite){
            if(*ite==*(this->task_queue.end()-1)){
                ite=this->task_queue.begin(); 
                this->task_queue.pop_back();
                // this->it = this->task_queue.begin();
            }else{
                this->task_queue.erase(ite);
            }
}
void ScheduleWithTaskList::show_task_queue(){
    std::cout << "task_queue: " << std::endl;
    for (ITask* item: this->task_queue){
        std::cout << item->pid;
    }
    cout<<endl;
}
class ScheduleTimeSlice : public ScheduleWithTaskList{
public:
    void call();
    void schedule_task();
};
void ScheduleTimeSlice::call(){
    if(this->task_queue.begin() == this->task_queue.end()){
        std::cout << "no process" << std::endl;
    }else{
        ITask* task = *it;
        this->Schedule::call(task);
        it++;
    }
}
void ScheduleTimeSlice::schedule_task(){
    while(true){
        this->call();
        sleep(TIME_SLICE/10000);
        int i;
        Process* px;
        px = new Process;
        px->pid = ++i;
        if(i <= 3){
            this->add_task(px);
        }
        this->inter_call();
        if(rand()%2 ==0 && ((Process*)*(this->it))->pid!=0){
            std::cout << "erase: "<<((Process*)(*(this->it)))->pid << std::endl;
            this->task_queue.end();
            if(*(this->it)==*(this->task_queue.end()-1)){
                this->it -1;
                this->task_queue.pop_back();
                this->it = this->task_queue.begin();
            }else{
                this->task_queue.erase(this->it);
            }
        }
        std::cout << "Process in queue: " << std::endl;
        for (ITask* item: this->task_queue){
            std::cout << ((Process*)item)->pid << " ";
        }
    }
}

class ScheduleComeFirst :public ScheduleWithTaskList{
public:
    void schedule_task();
};

void ScheduleComeFirst::schedule_task(){
    Process p0,p1,p2;
    p0.pid=0;
    p1.pid=1;
    p2.pid=2;
    this->task_queue.push_back(&p0);
    this->task_queue.push_back(&p1);
    this->task_queue.push_back(&p2);
    this->it = this->task_queue.begin();
    while (true){
        this->show_task_queue();
        if(this->task_queue.size()==1){
            (*(this->it))->run();
            sleep(TIME_SLICE/1000);
        }else{
            this->it=this->task_queue.begin()+1;
            while(true){
                (*(this->it))->run();
                sleep(TIME_SLICE/1000);
                if (rand()%2==0){ break;}
            }
            this->remove_task(it);
        }
    }
};

class ScheduleShortFirst: public ScheduleWithTaskList{
public:
    void add_task(ITask* task);
    void schedule_task();
};
void ScheduleShortFirst::add_task(ITask* task){
    if( this->task_queue.size()==0){
        task_queue.push_back(task);
    }else{
        int pos = rand()%task_queue.size() +1;
        this->task_queue.insert(task_queue.begin()+pos, task);
        std::cout << "add task: " <<task->pid<< "at "<< pos << std::endl;
    }
    
}
void ScheduleShortFirst::schedule_task(){
    Process p0,p1,p2,p3,p4;
    p0.pid=0;
    p1.pid=1;
    p2.pid=2;
    p3.pid=3;
    p4.pid=4;
    this->add_task(&p0);
    this->add_task(&p1);
    this->add_task(&p2);
    this->add_task(&p3);
    this->add_task(&p4);

    this->it = this->task_queue.begin();
    while (true){
        this->show_task_queue();
        if(this->task_queue.size()==1){
            (*(this->it))->run();
            sleep(TIME_SLICE/1000);
        }else{
            this->it=this->task_queue.begin()+1;
            while(true){
                (*(this->it))->run();
                sleep(TIME_SLICE/1000);
                if (rand()%2==0){ break;}
            }
            this->remove_task(it);
        }
    }
}

int main(int argc, const char** argv) {
    // ScheduleWithTaskList schedule;
    // ScheduleTimeSlice schedule;
    // schedule.task_queue.reserve(100);
    // Process p1;
    // p1.pid=0;
    // Process p2;
    // p2.pid=7;
    // Process* px;
    // schedule.set_clock(TIME_SLICE);
    // schedule.add_task(&p1);
    // schedule.add_task(&p2);
    // schedule.it = schedule.task_queue.begin();
    // schedule.schedule_task();

    // ScheduleComeFirst schedule;
    // schedule.schedule_task();

    ScheduleShortFirst schedule;
    schedule.schedule_task();

    return 0;
}



