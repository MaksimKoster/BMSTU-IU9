
#ifndef _ANTIDIAGMTR_H_
#define _ANTIDIAGMTR_H_

#include <iostream>
#include <iomanip>

template <typename T, int N>
class antidiagonal_matrix;
template<typename T, int N>
std::ostream& operator<<(std::ostream&, const antidiagonal_matrix<T, N>&);

template <typename T, int N>
class antidiagonal_matrix
{
private:
    T a[N];
public:
    antidiagonal_matrix();
    /*-------*/
    T at(int i, int j);
    int determinate();
    void set(int, T);
    // void set(int, int, T);
    // antidiagonal_matrix& operator= (antidiagonal_matrix obj);
    antidiagonal_matrix operator+ (antidiagonal_matrix obj);
    friend std::ostream& operator<< <> (std::ostream& os, const antidiagonal_matrix& a);
};
/*
Элемент с индексом (i, j): 47

Определитель: 55

Сложение с другой матрицей: 87

Нормализация: 208
*/

template<typename T, int N>
antidiagonal_matrix<T, N>::antidiagonal_matrix(){
    for (int i = 0; i < N; i++)
    {
        a[i] = 0;
    }
}

template<typename T, int N>
T antidiagonal_matrix<T, N>::at(int i, int j)
{
    if (i < 0 || j < 0 || i > N || j > N) throw std::exception();
    if (i != j) return 0;
    return a[i];
}

template<typename T, int N>
int antidiagonal_matrix<T, N>::determinate()
{
    int res = -1;
    for(int i = 0; i < N; i++)
    {
        // res *= at(i, i);
        res *= this->a[i];
    }
    return res;
}


template<typename T, int N>
void antidiagonal_matrix<T, N>::set(int i, T k)
{
    this->a[i] = k;
}

// template<typename T, int N>
// void antidiagonal_matrix<T, N>::set(int i, int j, T k)
// {
//     if (i != j) throw std::exception();
//     this->a[i] = k;
// }

// template<typename T, int N>
// antidiagonal_matrix<T, N>& antidiagonal_matrix<T, N>::operator= (antidiagonal_matrix obj)
// {
//     std::swap(this->a, obj.a);
//     return *this;
// }

template<typename T, int N>
antidiagonal_matrix<T, N> antidiagonal_matrix<T, N>::operator+ (antidiagonal_matrix obj)
{
    antidiagonal_matrix<double, N> new_a;
    for (int i = 0; i < N; i++)
    {
        new_a.a[i] = this->a[i] + obj.a[i];
    }
    return new_a;
}

template<typename T, int N>
std::ostream& operator<< (std::ostream& os, antidiagonal_matrix<T, N>& a)
{
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < N; j++)
        {
            os << a.at(i, N-1-j) << "\t";
        }
        os << std::endl;
    }
    return os;
}

/*----------------
----spec 4 double-------
-------------------
------------------------*/

template<int N>
class antidiagonal_matrix<double, N>{
private:
    double a[N];
public:
    antidiagonal_matrix();
    /*-------*/
    double at(int i, int j);
    double determinate();
    void set(int, double);
    void set(int, int, double);
    // antidiagonal_matrix& operator= (antidiagonal_matrix obj);
    antidiagonal_matrix operator+ (antidiagonal_matrix obj);
    friend std::ostream& operator<< <> (std::ostream& os, const antidiagonal_matrix& a);
    antidiagonal_matrix<double, N> normalise();
};

template<int N>
antidiagonal_matrix<double, N>::antidiagonal_matrix(){
    for (int i = 0; i < N; i++)
    {
        a[i] = 0;
    }
}

template<int N>
double antidiagonal_matrix<double, N>::at(int i, int j)
{
    if (i < 0 || j < 0 || i > N || j > N) throw std::exception();
    if (i != j) return 0;
    return a[i];
}

template<int N>
double antidiagonal_matrix<double, N>::determinate()
{
    double res = -1;
    for(int i = 0; i < N; i++)
    {
        res *= this->a[i];
    }
    return res;
}


template<int N>
void antidiagonal_matrix<double, N>::set(int i, double k)
{
    this->a[i] = k;
}

template<int N>
void antidiagonal_matrix<double, N>::set(int i, int j, double k)
{
    if (i != j) throw std::exception();
    this->a[i] = k;
}

// template<int N>
// antidiagonal_matrix<double, N>& antidiagonal_matrix<double, N>::operator= (antidiagonal_matrix obj)
// {
//     std::swap(this->a, obj.a);
//     return *this;
// }

template<int N>
antidiagonal_matrix<double, N> antidiagonal_matrix<double, N>::operator+ (antidiagonal_matrix obj)
{
    antidiagonal_matrix<double, N> new_a;
    for (int i = 0; i < N; i++)
    {
        new_a.a[i] = this->a[i] + obj.a[i];
    }
    return new_a;
}

template<int N>
std::ostream& operator<< (std::ostream& os, antidiagonal_matrix<double, N>& a)
{
    for (int i = 0; i < N; i++)
    {
        for (int j = 0; j < N; j++)
        {
            os << std::fixed << std::setprecision(3) << a.at(i, N-1-j) << "\t";
        }
        os << std::endl;
    }
    return os;
}


template<int N>
antidiagonal_matrix<double, N> antidiagonal_matrix<double, N>::normalise()
{
    double d = determinate();
    antidiagonal_matrix<double, N> b;
    for (int i = 0; i < N; i++)
    {
        b.a[i] = this->a[i] / d;
    }
    return b;
}


#endif //_ANTIDIAGMTR_H_