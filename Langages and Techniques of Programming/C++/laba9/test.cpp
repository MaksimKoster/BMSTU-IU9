#include <iostream>
#include "SimMatrix.h"
#include <random>

using namespace std;

int main(int argc, char const *argv[])
{
    const int n = 5;
    std::default_random_engine gen;
    std::uniform_int_distribution<int> distribution(-9,9);
    
    SimMatrix<int, n, n> mtr;
    cout << mtr << "\n";
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j <= i; j++)
        {
            mtr[i][j] = distribution(gen);
        }
    }
    cout << "m1: \n" << mtr << "m2 is filled by \"1\" \n";

    SimMatrix<int, n, n> mtr2(1);
    SimMatrix<int, n, n> res;
    cout << "m1 + m2 \n";
    res = mtr + mtr2;
    cout << res << "\n";
    res = mtr - mtr2;
    cout << "m1 - m2\n" << res << "\n";
    SimMatrix<int, n, n> mtr3;
    mtr3 = mtr * 2;
    cout << "m1 * 2\n" << mtr3 << "\n";
    mtr3 = mtr * mtr2;
    cout << mtr3 ;
    return 0;
}
