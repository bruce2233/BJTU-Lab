#include<iostream>
#include<cstring>

int main(){
    Process process1;
    Process process2;

}

class Controller {
    virtual void interCall() = 0;
};

class Task{
    virtual void run()=0;
};

class Process: public Task{
    int pid;
    int status;
    void run(){
        std::cout << "a process is running..." << std::endl;
    }
};
