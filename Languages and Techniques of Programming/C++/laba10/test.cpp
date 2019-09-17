#include <iostream>
#include "sequence.h"

int main(){
    std::vector<int > a;
    a.emplace_back(6);
    a.emplace_back(10);
    Sequence<int, 2> sq;
    sq.add(a);
    std::vector<int> b, c;
    b.emplace_back(5); b.emplace_back(6);
    c.emplace_back(2); c.emplace_back(3);
    sq.add(b); sq.add(c);
    for (auto it = sq.begin(); it != sq.end(); )
    {
        std::cout << *it++ << " ";
    }
    std::cout << endl;
    for (auto it = sq.end(); it != sq.begin(); )
    {
        std::cout << *--it << " ";
    }
}