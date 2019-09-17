#include "declaration.h"
#include <iostream>

Intervals::Intervals()
{
    size = 1;
    a = new interval*[size];
    count = 0;
}

Intervals::Intervals(const Intervals& obj)
{
    size = obj.size;
    count = obj.count;
    a = new interval*[size];
    std::copy(obj.a, obj.a + count, a);
}

Intervals::~Intervals()
{
    for (int i = 0; i < getSize(); i++)
    {
        delete a[i];
    }
    delete [] a;
}

int Intervals::getSize()
{
    return count;
}

void Intervals::add(interval *i)
{
    if (size < count + 1)
    {
        resize_intervals();
    }

    a[count] = i;
    count++;
}


std::string Intervals::contains(int x)
{
    for (int i = 0; i < count; ++i)
    {
        if(a[i]->contains(x))
        {
            return "true";
        }
    }
    return "false";
}

std::string Intervals::to_string(){
    std::string res = "";
    for (int i = 0; i < this->getSize(); i++)
    {
        res.append(this[i].to_string());
    }
    return res;
}

Intervals Intervals::get_range(int size)
{
    Intervals res;
    for (int i = 0; i < this->count; ++i)
    {
        if (a[i]->size() > size)
        {
            res.add(a[i]);
        }
    }
    return res;
}

void Intervals::resize_intervals()
{
    size *= 10;
    auto **new_interval = new interval*[size];
    for (int i = 0; i < size/10; ++i)
    {
        new_interval[i] = this->a[i];
    }

    delete [] this->a;
    this->a = new_interval;
}

interval& Intervals::operator[](int i)
{
    if (i < 0 && i >=size)
    {
        throw std::exception();
    }
    else
    {
        return *a[i];
    }
}
Intervals& Intervals::operator= (Intervals obj)
{
    std::swap(size, obj.size);
    std::swap(count, obj.count);
    std::swap(a, obj.a);
    return *this;
}

std::ostream& operator<< (std::ostream& os, Intervals& a)
{
    for (int i = 0; i < a.getSize(); i++)
    {
        os << a[i];
    }
    os << std::endl;

    return os;
}

//////////////////////////
//////////////////////////

bool interval::contains(int x)
{
    return x >= a && x <= b;

}

int interval::size()
{
    return b-a;
}
std::string interval::to_string()
{
    return "(" + std::to_string(a) + "; " + std::to_string(b) + ")"; 
}
interval& interval::operator=(const interval& obj)
{
    if (this != &obj) {
        a = obj.a;
        b = obj.b;
    }
    return *this;
}

std::ostream& operator<< (std::ostream& os, interval& interval)
{
    os << "[" << interval.a <<"," << interval.b << "]"<<std::endl;
    return os;
}
interval::interval(int a, int b)
{
    this->a = a;
    this->b = b;
}
interval::interval(interval& obj)
{
    a = obj.a;
    b = obj.b;
}