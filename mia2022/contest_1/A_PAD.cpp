/*
  Jakub Galecki
  Parity Alternated Deletions Solution
*/

#include <algorithm>
#include <iostream>
#include <vector>

void
solve()
{
  int n;
  int sum = 0;
  std::vector<int> odd_numbers, even_numbers;
  std::cin >> n;

  for (int i = 0; i < n; i++) {
    int number;
    std::cin >> number;

    if (number & 1) {
      odd_numbers.push_back(number);
    } else {
      even_numbers.push_back(number);
    }

    sum += number;
  }

  std::sort(odd_numbers.begin(), odd_numbers.end(), std::greater<int>());
  std::sort(even_numbers.begin(), even_numbers.end(), std::greater<int>());

  uint16_t lenght = std::min(odd_numbers.size(), even_numbers.size());

  for (int i = 0; i < lenght; i++) {
    sum -= odd_numbers[i];
  }

  for (int i = 0; i < lenght; i++) {
    sum -= even_numbers[i];
  }

  if (odd_numbers.size() > lenght) {
    sum -= odd_numbers[lenght];
  }

  if (even_numbers.size() > lenght) {
    sum -= even_numbers[lenght];
  }

  std::cout << sum << std::endl;
}

int
main()
{
  solve();
  return 0;
}