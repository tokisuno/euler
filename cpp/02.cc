/*
 *  Considering the terms in the fib sequence whose values don't exceed 4mil:
 *  Find the same of the even-valued terms
 *
 */

#include <iostream>
#include <numeric>
#include <vector>

bool isEven(int n) {
    return n % 2;
}

std::vector<int> generateEvenFibs() {
    std::vector<int> fibs;

    int a { 0 };
    int b { 1 };
    int c{};

    for (int i = 2; i <= 1000; i++) {
        c = a + b;
        a = b;
        b = c;

        if (b >= 4e6) {
            break;
        }
        if (isEven(b)) {
            fibs.push_back(b);
        }
    }
    return fibs;
}

int main() {
    std::vector<int> evenFibs { generateEvenFibs() };

    for (const int& i : evenFibs)
        std::cout << "num: " << i << '\n';

    int size { std::accumulate(evenFibs.begin(), evenFibs.end(), 0) + 1 };
    std::cout << size << '\n';
}
