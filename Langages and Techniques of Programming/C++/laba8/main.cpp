#include "AntidiagMtr.h"
#include <random>

using namespace std;

int main()
{
    std::default_random_engine gen;
    std::uniform_int_distribution<int> distribution(-9,9);
    const int n = 6;
    /*-----*/
    antidiagonal_matrix<double, n> a;
    antidiagonal_matrix<double, n> b;
    for (int i = 0; i < n; i++)
    {
        a.set(i, distribution(gen));
        b.set(i, distribution(gen));
    }
    antidiagonal_matrix<double, 6> c;
    c = a + b;
    std::cout << a << "+\n" << b << "=\n" << c;
    std::cout << "|a| = " << a.determinate() << "\n";
    antidiagonal_matrix<double, n> aNorm;
    aNorm = a.normalise();
    std::cout << "a normalised: \n";
    std::cout << aNorm;

    std::cout << "\n\n";
    
    antidiagonal_matrix<int, 5> intM;
    for (int i = 0; i < 5; i++)
    {
        intM.set(i, distribution(gen));
    }
    cout << intM;
    return 0;
}


