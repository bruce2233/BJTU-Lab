#include<iostream>
#include<cstring>
#include<vector>
#include <unistd.h>

using namespace std;

const int TIME_SLICE = 1500;

class ISchedule {
public:
    virtual void interCall() = 0;
    virtual void addTask() = 0;
    virtual void call(Process)=0;
    virtual void setClock(int) = 0;
};

class ITask{
public:
    virtual void run()=0;
};

class Process: public ITask{
public:
    int pid;
    int status;
    void run();
};

void Process::run(){
        std::cout << "process is" <<this->pid<<" running..." << std::endl;
}

class Schedule : public ISchedule{
public:
    void setClock(int time){
        std::cout << "after "<<time<<"ms another task runs" << std::endl;
    }

    void interCall(){
        std::cout << "called by the clock interupt..." << std::endl;
    }

     void addTask(){
        std::cout << "create a task..." << std::endl;
    }

    void call(Process &process){
        std::cout << "schedule process..." << process.pid << std::endl;
    }
};

class ScheduleWithTaskQueue : public Schedule{
public:
    ITask *task_queue[3];
};

int main(int argc, const char** argv) {
    ScheduleWithTaskQueue schedule;
    // ISchedule * sc;
    // sc = &swtq;
    
    Process p1;
    p1.pid=888;
    // schedule.task_queue.push_back(&p1);
    schedule.task_queue[0]= &p1;

    while(true){
        schedule.setClock(TIME_SLICE);
        schedule.addTask();
        
        sleep(TIME_SLICE/1000);
        schedule.interCall();
    }
    return 0;
}
