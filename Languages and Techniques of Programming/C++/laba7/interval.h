#ifndef __INTERVAL_H__
#define __INTERVAL_H__

#include <iostream>

class interval
{
private:
    int a, b;
public:
    interval(int, int);
    interval(interval& obj);
    bool contains(int x);
    std::string to_string();
    int size();
    interval& operator=(const interval& other);
    friend std::ostream& operator<< (std::ostream& os, interval& interval);
};


#endif //__INTERVAL_H__