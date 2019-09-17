#ifndef __DECLARATION_H__
#define __DECLARATION_H__

#include "interval.h"

class Intervals
{
private:
    interval **a;
    int count;
    int size;
    void resize_intervals();
    std::string to_string();
public:
    Intervals();
    Intervals(const Intervals&);
    virtual ~Intervals();
/*----------------------------*/
    int getSize();
    Intervals get_range(int size);
    interval& operator[](int i);
    Intervals& operator= (Intervals obj);
    void add(interval * i);
    std::string contains(int x);
    friend std::ostream& operator<< (std::ostream& os, Intervals& a);
};

#endif