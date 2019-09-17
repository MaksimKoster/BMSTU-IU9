#ifndef _SIMMTR_H_
#define _SIMMTR_H_

#include <iostream>

template <typename T, int M, int N>
class SimMatrix;
template<typename T, int M, int N>
std::ostream& operator<<(std::ostream&, const SimMatrix<T, M, N>&);


template<typename T, int M, int N>
class SimMatrix
{
private:
    T a[N * (N + 1) / 2];
    int k;
    class Row{
    private:
        SimMatrix *m;
        int i;
    public:
        Row(SimMatrix *m, int i)
        {
            this->m = m;
            this->i = i;
        }
        T& operator[] (int j)
        {
            return m->a[i > j ? i*N-(i-1)*i/2+j-i : j*N-(j-1)*j/2+i-j];
        }
        const T& operator[] (int j) const
        {
            return m->a[i > j ? i*N-(i-1)*i/2+j-i : j*N-(j-1)*j/2+i-j];
        }
    };
public:
    SimMatrix();
    SimMatrix(int c);
    //**************
    Row operator[] (int i)
    {
        return Row(this, i);
    }
    const Row &operator[] (int i) const
    {
        return Row(this, i);
    }
    //**************
    SimMatrix operator+ (SimMatrix obj); // v
    SimMatrix operator- (SimMatrix obj); // v
    SimMatrix operator* (SimMatrix obj); // v
    SimMatrix operator* (T c); // v
    friend std::ostream& operator<< <> (std::ostream& os, const SimMatrix& a); // v
};

template<typename T, int M, int N>
SimMatrix<T, M, N>::SimMatrix(){
    if (M != N) throw std::exception();
    k = N * (N + 1) / 2;
    for (int i = 0; i < k; i++)
    {
        a[i] = 0;
    }
}

template<typename T, int M, int N>
SimMatrix<T, M, N>::SimMatrix(int c){
    if (M != N) throw std::exception();
    k = N * (N + 1) / 2;
    for (int i = 0; i <= k; i++)
    {
        a[i] = c;
    }
}

template<typename T, int M, int N>
SimMatrix<T, M, N> SimMatrix<T, M, N>::operator+(SimMatrix obj)
{
    SimMatrix<T, M, N> res;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j <= i; j++)
        {
            res[i][j] = (*this)[i][j] + obj[i][j];
        }
    }
    return res;
}

template<typename T, int M, int N>
SimMatrix<T, M, N> SimMatrix<T, M, N>::operator-(SimMatrix obj)
{
    SimMatrix<T, M, N> res;
    res = (*this) + obj * (-1);
    return res;
}

template<typename T, int M, int N>
SimMatrix<T, M, N> SimMatrix<T, M, N>::operator*(T c)
{
    SimMatrix<T, M, N> res;
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j <= i; j++)
        {
            res[i][j] = (*this)[i][j] * c;
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
        for (int j = 0; j < N; j++)
        {
            res[i][j] = 0;
            for (int k = 0; k < N; k++)
            {
                res[i][j] = res[i][j] + (*this)[i][k] * m[k][j];
            }
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

//a[i>=j ? i*N-(i-1)*i/2+j-i : j*N-(j-1)*j/2+i-j]
//i>=j ? i*N-(i-1)*i/2+j-i : j*N-(j-1)*j/2+i-j

#endif //_SIMMTR_H_