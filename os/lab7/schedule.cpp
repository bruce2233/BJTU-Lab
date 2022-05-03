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
};

class Process: public ITask{
public:
    int pid;
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
};
void Schedule::set_clock(int time){
        std::cout << "after "<<time<<"ms another task runs" << std::endl;
}

void Schedule::call(ITask* task){
    task->run();
}

void Schedule::inter_call(){
        std::cout << "called by the clock interupt..." << std::endl;
}

class ScheduleWithTaskList : public Schedule{
public:
    vector<ITask *> task_queue;
    vector<ITask *>::iterator it;
    void add_task(ITask* task);
};
void ScheduleWithTaskList::add_task(ITask* task){
    this->task_queue.push_back(task);
}

class ScheduleTimeSlice : public ScheduleWithTaskList{
public:
    void call();
};
void ScheduleTimeSlice::call(){
    if(this->task_queue.begin() == this->task_queue.end()){
        std::cout << "no process" << std::endl;
    }else{
        if(it == task_queue.end()){
            it=task_queue.begin();
        }
        ITask* task = *it;
        this->Schedule::call(task);
        it++;
    }
}

int main(int argc, const char** argv) {
    // ScheduleWithTaskList schedule;
    ScheduleTimeSlice schedule;
    schedule.task_queue.reserve(100);
    Process p1;
    p1.pid=0;
    Process p2;
    p2.pid=7;
    Process* px;
    schedule.set_clock(TIME_SLICE);
    schedule.add_task(&p1);
    schedule.add_task(&p2);
    schedule.it = schedule.task_queue.begin();
    int i=0;
    while(true){
        schedule.call();
        sleep(TIME_SLICE/1000);
        px = new Process;
        px->pid = i++;
        if(i <= 10){
            schedule.add_task(px);
        }
        schedule.inter_call();
        if(rand()%2 ==0 && (Process*)(*(schedule.it))->pid!=0){
            schedule.task_queue.erase(schedule.it);
        }
        std::cout << "Process in queue: " << std::endl;
        for (ITask* item: schedule.task_queue){
            std::cout << ((Process*)item)->pid << " ";
        }
    }
    return 0;
}
