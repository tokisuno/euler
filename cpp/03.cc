/*
 *  Largest Prime Factor
 *
 *  The prime factorsof 13195 are 5, 7, 13, 29
 *  What is the largest prime factor of:
 *    600'851'475'143
 */

#include <iostream>
#include <vector>

std::vector<int> findFactor(int n) {
    std::vector<int> factors;
    for (int i = 1; i <= n; i++) {
        if (n % i == 0) {
            factors.push_back(n);
        }
    }
    return factors;
}

bool isPrime(int n) {
    for (int i = 0; i < n; i++) {
        if (findFactor(n).size() == 2)
            return true;
    }
    return false;
}

int main() {
    int num { 13195 };

    std::vector<int> factors { findFactor(num) };

    for (int i = 1; i <= findFactor:

    std::cout << "The largest prime is: " << std::endl;

    return 0;
}
