/************************
 * Natural numbers that are below 10 that are multiples of 3 or 5
 * 3, 5, 6, 9
 * Sum of 3, 5, 6, 9 = 23
 *
 * Find the sum of all the multiples of 3 or 5 below 1000
 *****************************/

#include <iostream>
#include <numeric>
#include <vector>

int main() {
    std::vector<int> multiplesOf3;
    std::vector<int> multiplesOf5;
    int sum3{};
    int sum5{};

    for (int i = 0; i < 1000; i++) {
        if (i % 3 == 0) {
            multiplesOf3.push_back(i);
        }
    };

    for (int j = 0; j < 1000; j++) {
        if (j % 5 == 0) {
            if (j % 3 == 0) {
                continue;
            } else {
                multiplesOf5.push_back(j);
            }
        }
    }

    sum3 = std::accumulate(multiplesOf3.begin(), multiplesOf3.end(), 0);
    sum5 = std::accumulate(multiplesOf5.begin(), multiplesOf5.end(), 0);

    std::cout << sum3 + sum5;
}
