#include <iostream>
#include "declaration.h"

using namespace std;

int main()
{
    Intervals a;
    a.add(new interval(1, 3));
    a.add(new interval(3, 4));
    a.add(new interval(5, 9));
    a.add(new interval(10, 20));
    cout << "Intervals:\n" << a << "\n";

    cout << "Size:" << a.getSize() << "\n";
    cout << "Second interval of range: " << a[2] << "\n";
    cout << "Does this intervals contain 10? \n - " << a.contains(10) << "\n";
    cout << "Does this intervals contain 27? \n - " << a.contains(27) << "\n";
    Intervals range = a.get_range(2);
    cout << "Intervals, bigger then 2: " << range << "\n";

    Intervals b;
    b.add(new interval(5, 7));
    cout << "Creating new interval, " << b;
    a = b;
    cout << "a = b: \n";
    cout << "a: " << a;
    cout << "b: " << b;
}