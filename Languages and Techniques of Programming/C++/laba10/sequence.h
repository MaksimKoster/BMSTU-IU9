// #include <iostream>
#include <iterator>
#include <vector>

using namespace std;

template <typename T, int N>
class ScalarIterator;

template <typename T, int N>
class Sequence
{
private:
    vector<vector<T> > seq;
    friend class ScalarIterator<T, N>;
public:
    Sequence(){}

    void add(vector<T> x)
    {
        if (x.size() != N) throw "incorrect dimencity";
        seq.emplace_back(x);
    }

    ScalarIterator<T, N> begin();
    ScalarIterator<T, N> end();

    vector<T>& operator[] (int n) 
    {
        return seq[n];
    }

    const vector<T>& operator[] (int n) const
    {
        return seq[n];
    }
};

template <typename T, int N>
class ScalarIterator:
    public iterator<
        bidirectional_iterator_tag,
        T,
        ptrdiff_t,
        const T*,
        const T
    >
{
private:
    Sequence<T, N> *s;
    bool is_default;
    int ind;

    bool is_end() const {
        return ind >= s->seq.size();
    }

public:
        ScalarIterator(): is_default(true) {}
        ScalarIterator(Sequence<T, N> &obj): ScalarIterator<T, N>(obj, 0) {};
        ScalarIterator(Sequence<T, N> &obj, int ind): s(&obj), ind(ind), is_default(false) {}

    bool operator== (const ScalarIterator<T, N> &other) const
    {
        return (is_default && other.is_default) ||
            (is_default && other.is_end()) ||
            (is_end() && other.is_default) ||
            (s == other.s && ind == other.ind);
    }

    bool operator!= (const ScalarIterator<T, N> &other) const
    {
        return !(*this == other);
    }

    const T operator* ()
    {
        if (is_default) throw "not initialized iterator";
        T res = 0;
        for (int i = 0; i < N; i++)
            res += (s->seq[ind])[i] * (s->seq[ind+1])[i];
        return res;
    }

    ScalarIterator<T, N>& operator++()
    {
        if (is_default) throw "not initialized iterator";
        ind++;
        return *this;
    }

    ScalarIterator<T, N> operator++(int)
    {
        ScalarIterator<T, N> tmp (*this);
        ind++;
        return tmp;
    }

    ScalarIterator<T, N>& operator--()
    {
        if (is_default) throw "not initialized iterator";
        ind--;
        return *this;
    }

    ScalarIterator<T, N> operator--(int)
    {
        ScalarIterator<T, N> tmp(*this);
        ind--;
        return tmp;
    }
};

template<typename T, int N>
ScalarIterator<T, N> Sequence<T, N>::begin()
{
    return ScalarIterator<T, N>(*this);
}

template<typename T, int N>
ScalarIterator<T, N> Sequence<T, N>::end()
{
    return ScalarIterator<T, N>(*this, this->seq.size()-1);
}
