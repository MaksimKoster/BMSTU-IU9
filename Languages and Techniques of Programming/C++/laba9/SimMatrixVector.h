#ifndef _SIMMTR_H_
#define _SIMMTR_H_

#include <iostream>
#include <vector>

template <typename T, int M, int N>
class SimMatrix;
template<typename T, int M, int N>
std::ostream& operator<<(std::ostream&, const SimMatrix<T, M, N>&);


template<typename T, int M, int N>
class SimMatrix
{
private:
    vector<vector<int> > a (n);
    
public:
    SimMatrix();
    //**************
    class Row{
    private:
        SimMatrix *m;
        int i;
    public:
        Row(SimMatrix *m, int i);
        T& operator[] (int j)
    };

    SimMatrix::Row operator[] (int i)
    {
        return Row(this, i);
    }
    //**************
    SimMatrix operator+ (SimMatrix obj); // v
    SimMatrix operator- (SimMatrix obj); // v
    SimMatrix operator* (SimMatrix obj); // v
    SimMatrix operator* (T c); // v
    Row operator[] (int i);
    friend std::ostream& operator<< <> (std::ostream& os, const SimMatrix& a); // v
};

template<typename T, int M, int N>
SimMatrix<T, M, N>::SimMatrix(){
    if (M != N) throw std::exception();
    int k = 1;
    for (int i = 0; i < N; i++)
    {
        a[i].resize(k);
        for (int j = 0; j < k; j++)
        {
            arr[i][j] = 0;
        }
        k++;
    }
}

template<typename T, int M, int N>
SimMatrix<T, M, N> SimMatrix<T, M, N>::operator+(SimMatrix obj)
{
    SimMatrix<T, M, N> res;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; i < i; j++)
        {
            res[i][j] = this->a[std::max(i, j)][std::min(i, j)] + obj[i][j];
        }
    }
    return res;
}

template<typename T, int M, int N>
SimMatrix<T, M, N> SimMatrix<T, M, N>::operator-(SimMatrix obj)
{
    SimMatrix<T, M, N> res;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; i < i; j++)
        {
            res[i][j] = this->a[i][j] - obj[i][j];
        }
    }
    return res;
}

template<typename T, int M, int N>
SimMatrix<T, M, N> SimMatrix<T, M, N>::operator*(T c)
{
    SimMatrix<T, M, N> res;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; i < i; j++)
        {
            res[i][j] = this->a[i][j] * c;
        }
    }
    return res;
}

template<typename T, int M, int N>
SimMatrix<T, M, N> SimMatrix<T, M, N>::operator*(SimMatrix m)
{
    SimMatrix<T, M, N> res;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < i; j++)
        res[i][j] = 0;
        for (int k = 0; k < N; k++)
        {
            res[i][j] = res[i][j] + this->a[std::min(i, k)][std::max(i, k)] + m[i][k];
        }
    }
    return res;
}

template<typename T, int M, int N>
std::ostream& operator<< (std::ostream& os, SimMatrix<T, M, N>& a)
{
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < N; j++)
        {
            os << a[i][j] << "\t";
        }
        os << std::endl;
    }
    return os;
}

#endif