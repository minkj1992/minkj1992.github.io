# 홍정모의 따라하며 배우는 C++


> `Modern c++`을 학습하고 기억할만한 요소들을 정리합니다.
<!--more-->


---
<details>
<summary><strong><u>1. Terminology</u></strong></summary>

## 1. Terminology

<h4>1.0.1. Basic C++</h4>

- prototype
- forward declaration
- header guards (=include guards)
  - 중복될 경우 한번만 include해라 (#pragma once)
- #pragma once
  - 기 정의된 preprocessor의 일종
- macro
- conditional compilation
  - macro의 #ifdef, #ifndef, #else, #endif
- fundamental data types (=primitivate data types)
- `auto`는 데이터 타입을 자동으로 컴파일 타임에 찾아준다.

<h4>1.0.2. Variable and Fundamental types</h4>

- `initialization`
  - `copy initialization` int a = 3;
  - `direct initialization` int a(3);
  - `uniform initialization`  int a{ 3 };
- `Fixed-width Integers`
- scientific notation
- `inf`: infinite
- `nan`: not a number
- `ind`: indeterminate
- `literal constants`
- `symbolic constants`
  - constexpr(c++ 11): 컴파일 타임에 값이 완전히 결정되는 상수
  - const: 컴파일 타임 / 런타임에 값이 결정되는 모든 상수. (constexpr 포함)

<h4>1.0.3. Variable Scope and Extra types</h4>

- Scoped Enumerations (`Enum Class`)
- `type aliases`
  - `typedef`
  - `using`
- `struct`
  - `member selection operator` = .
  - memory `padding`
  - 최적화를 위해서 member들의 순서를 고려해야 한다. (e.g short type 2byte는 2바이트가 뒤에 padding된다.)

<h4>1.0.4. Matrix, String, Pointer, Reference</h4>

- nullptr(null pointer)
- `void pointer` == `generic pointer`
- reference variable
</details>  

---

## 2. Function Parameter
- function parameter
```cpp
#include <iostream>
#include <array>

using namespace std;

bool isEven(const int &number)
{
	return (number % 2 == 0) ? true : false;
}

bool isOdd(const int &number)
{
	return !isEven(number);
}

void printNumbers(array<int, 10> &arr, bool (*validator)(const int &))
{
	for (int v : arr)
		if (validator(v))
			cout << v << ' ';
	cout << endl;
}

int main()
{
	array<int, 10> my_arr{0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
	printNumbers(my_arr, isEven); // 0 2 4 6 8
	printNumbers(my_arr, isOdd);  // 1 3 5 7 9
	return 0;
}
```

- function parameter with `using` or `typedef`

```cpp
#include <iostream>
#include <array>

using namespace std;
// typedef bool (*validator_fnc)(const int &);
using validator_fnc = bool (*)(const int &);

bool isEven(const int &number)
{
	return (number % 2 == 0) ? true : false;
}

bool isOdd(const int &number)
{
	return !isEven(number);
}

void printNumbers(array<int, 10> &arr, validator_fnc validator)
{
	for (int v : arr)
		if (validator(v))
			cout << v << ' ';
	cout << endl;
}

int main()
{
	array<int, 10> my_arr{0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
	printNumbers(my_arr, isEven); // 0 2 4 6 8
	printNumbers(my_arr, isOdd);  // 1 3 5 7 9
	return 0;
}
```

- function parameter with std::function and default parameter
```cpp
#include <iostream>
#include <array>
#include <functional>

using namespace std;
using validator_fnc = function<bool(const int &)>;

bool isEven(const int &number)
{
	return (number % 2 == 0) ? true : false;
}

bool isOdd(const int &number)
{
	return !isEven(number);
}

void printNumbers(array<int, 10> &arr, validator_fnc validator = isEven)
{
	for (int v : arr)
		if (validator(v))
			cout << v << ' ';
	cout << endl;
}

int main()
{
	array<int, 10> my_arr{0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
	printNumbers(my_arr);		 // 0 2 4 6 8
	printNumbers(my_arr, isOdd); // 1 3 5 7 9
	return 0;
}
```

매우 파이썬스럽게, 가장 깔끔해 보인다.

## 3. Matrix
- double pointer
```cpp
#include <iostream>
using namespace std;

void printMatrix(int **matrix, const int row, const int col)
{
	for (int r = 0; r < row; ++r)
	{
		for (int c = 0; c < col; ++c)
			cout << matrix[r][c] << " ";
		cout << endl;
	}
}

int main()
{
	const int row = 3;
	const int col = 5;
	const int values[row][col] =
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
		};

	// init
	int **matrix = new int *[row];
	for (int r = 0; r < row; ++r)
		matrix[r] = new int[col];

	// assign
	for (int r = 0; r < row; ++r)
		for (int c = 0; c < col; ++c)
			matrix[r][c] = values[r][c];

	printMatrix(matrix, row, col);

	// delete
	for (int r = 0; r < row; ++r)
		delete[] matrix[r];
	delete[] matrix;
	return 0;
}
```

- single pointer
```cpp
#include <iostream>
using namespace std;

void printMatrix(int *matrix, const int row, const int col)
{
	for (int r = 0; r < row; ++r)
	{
		for (int c = 0; c < col; ++c)
			cout << matrix[(col * r) + c] << " ";
		cout << endl;
	}
}

int main()
{
	const int row = 3;
	const int col = 5;
	const int values[row][col] =
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
		};

	// init
	int *matrix = new int[row * col];

	// assign
	for (int r = 0; r < row; ++r)
		for (int c = 0; c < col; ++c)
			matrix[(col * r) + c] = values[r][c];

	printMatrix(matrix, row, col);

	// delete
	delete[] matrix;

	return 0;
}
```

## 4. Pointer
### 4.1. Pointer and Const
```cpp
{
	using namespace std;

	const int value = 6;
	const int new_value = 7;
	const int *ptr_1 = &value;
	// 6 0x7ffeefb45158 0x7ffeefb45148
	cout << *ptr_1 << ' ' << ptr_1 << ' ' << &ptr_1 << endl; 
	// *ptr_1 = new_value; (x)
	ptr_1 = &new_value;
	// 7 0x7ffeefb45154 0x7ffeefb45148
	cout << *ptr_1 << ' ' << ptr_1 << ' ' << &ptr_1 << endl; 
}
```
`const int`를 가리키고 있는 mutable한 포인터를 의미한다.
포인터가 가리키는 값이 const int이기 때문에 dereference해서 값을 바꿀 수 없다. 
하지만 포인터 그 자체는 const하지 않기 때문에 새로운 주소값을 넣을 수 있다.

```cpp
{
	using namespace std;

	int value = 6;
	int new_value = 7;
	int *const ptr_value = &value;
	
	// 6 0x7ffee173e158 0x7ffee173e148
	cout << *ptr_value << ' ' << ptr_value << ' ' << &ptr_value << endl; 
	*ptr_value = new_value;
    // ptr_value = &new_value; (x)

	// 7 0x7ffee173e158 0x7ffee173e148
	cout << *ptr_value << ' ' << ptr_value << ' ' << &ptr_value << endl; 
}
```

`int`를 가리키는 `*const` 포인터.
포인터는 const이기 때문에 assign이 불가하지만, 포인터가 가리키는 값은 const하지 않기 때문에 변경 가능하다.
즉 포인터를 dereference(`*`)해서 값 대입 가능하다.

> c.f) int &ref(레퍼런스)와 int *const ptr는 기능이 같다.

```cpp
{
    using namespace std;

	int value = 6;
	int new_value = 7;
	const int *ptr_value = &value;
	// 6 0x7ffee6efb158 0x7ffee6efb148
	cout << *ptr_value << ' ' << ptr_value << ' ' << &ptr_value << endl; 

	value = new_value;
	// 7 0x7ffee6efb158 0x7ffee6efb148
	cout << *ptr_value << ' ' << ptr_value << ' ' << &ptr_value << endl; 

	// *ptr_value = new_value; (x)
	ptr_value = &new_value;
	// 7 0x7ffee6efb154 0x7ffee6efb148
	cout << *ptr_value << ' ' << ptr_value << ' ' << &ptr_value << endl; 
}
```
가리키는 값이 const한 int인 포인터.
value 자체는 const가 아니기 때문에 assign 가능하다.
value가 const가 아니지만, 포인터는 값을 const하게 처리하기 때문에 dereference가 불가능하다.
포인터 자체는 const하지 않기 때문에 주소 할당이 가능하다.

```c++
{
	using namespace std;

	const int value = 6;
	const int new_value = 7;
	const int *const ptr_value = &value;

	// ptr_value = &new_value; (x)
	// *ptr_value = new_value; (x)
}
```
`const int`를 가리키는 `*const` 포인터.
pointer value assign과 dereference를 통한 assign 둘 모두 불가하다.

## 5. Reference
### 5.1. Reference and Const
- reference variable은 변수의 별명이다. (주소, 값 모두 같다.)
- 파라미터로 넘겨줄 경우, 다른 함수에서 변수를 변경가능하다.

```cpp
#include <iostream>

void doSomething(const int value, const int &ref)
{
	using namespace std;

	cout << value << ' ' << &value << ' ' << ref << ' ' << &ref << endl;
}

int main()
{
	int a = 5; // 5 0x7ffeeaf4e158
	doSomething(a, a); // 5 0x7ffeeaf4e13c 5 0x7ffeeaf4e158
	doSomething(a, 5); // 5 0x7ffeeaf4e13c 5 0x7ffeeaf4e154

	return 0;
}
```
파라미터에 변수로 받는 것과, reference로 받는 것은 무슨 차이가 있을까?
파라미터를 변수로 받으면 value copy가 일어난다. (비효율) 반면 reference를 사용한다면 실제 원래 변수와 같은 주소를 가지게 된다.

또한 `const int &ref`처럼 const reference를 사용한다면 

1. immutable하면서도 reference로 파라미터를 받아 효율을 추구할 수 있다.
2. literal의 주소를 기억할 수 없기 때문에, reference variable 대입에는 lvalue가 들어가야 하지만, const reference타입은 literal을 받을 수 있다. (이 경우 literal을 위한 주소가 할당 된다.)

