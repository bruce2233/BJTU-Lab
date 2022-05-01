#include<iostream>

int main(){
    int b =1;
    int &a = b;
    a = 2;
    std::cout << b << std::endl;
}