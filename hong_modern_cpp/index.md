# 홍정모의 따라하며 배우는 C++


> `Modern c++`을 학습하고 기억할만한 요소들을 정리합니다.
<!--more-->


---
<details>
<summary><strong><u>1. Terminology</u></strong></summary>

## 0. Terminology

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
  - c.f) Most vexing parse: syntax가 일관성이 없어서, 	uniform initialization이 도입됨. (C++ 11)
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

- `::` : scope resolution operator
</details>  

---

## 1. Basic C++

#### namespace

- using namespace
- namespace끼리 nested하게 사용가능하다.

#### Macro (preprocessor)

Build 타임전, 즉 컴파일 타임에 처리된다. (preprocessor)

- `#ifdef`, `#ifndef`

다음과 같은 경우는 multi platform 즉 여러 os 타입에 따라서 build를 다르게 해주고 싶을 떄, 사용한다. 혹은 gpu 버전에 따라서 버저닝 하고 싶을 때 사용한다.

## 2. Basic types

#### Void type
void 자체는 메모리가 할당되지 않기 때문에 선언이 불가하다.
```cpp
void my_void; // (x)
```

하지만 void의 포인터 타입은 메모리 address가 있기 때문에 선언이 가능하다.

```cpp
void *my_void; // (o)

int i = 123;
float f = 123.456f;

my_void = (void*)&i;
my_void = (void*)&f;
```

또한 모든 포인터 address의 사이즈는 같기 때문에 int, float, void 타입 상관없이 동일한 변수에 assign이 가능하다.

- `(void*)` void pointer type
- `&i`, `&f`: int와 float 변수의 address 값

#### Char type

- casting style

```cpp
// c-style casting
cout << (char)65 << endl;
cout << (int)'A' << endl;

// c++ style casting
cout << char(65) << endl;
cout << int('A') << endl;
```

- static cast

```cpp
cout << static_cast<char>(65) << endl;
cout << static_cast<int>('A') << endl;
```

static_cast의 경우, 일반적으로 primitive type들에 대해서 **compile 타임에 형변환에 대한 타입 오류를 잡아주고 싶을 때 사용합니다.** 

- string buffer

string operator는 buffer에 임시로 저장되기 떄문에, cin으로 받아들인 값의 경우 cout을 하지 않더라도 buffer에 임시로 보관되어집니다.

```cpp
char c1;
cin >> c1;
cout << static_cast<int>(c1) << endl;

cin >> c1;
cout << static_cast<int>(c1) << endl;
```

```bash
$ abc
97
98
```

- `\n` vs `endl` vs `std::flush`
	- `\n`: new line하라.
	- `endl`: buffer에 있는 모든 것들을 출력한 뒤, new line하라.
	- `std::flush`: 줄바꿈 없이 buffer에 있는 것들을 모두 쏟아내라.

#### Literal constants
> 리터럴 상수

```cpp
unsigned int n = 5u;
long n2 = 5L;
double d = 6.0e-10;
```

- decimal, ocatal, hexa

```cpp
int x = 012; // 8진수
cout << x << endl; // 10

int y = 0xF; // 16진수
cout << y<< endl; // 15
```

c++14 이후 부터 `binary literal`이 가능해졌다. 또한 literal 사이에 quota(`'`)를 넣어주게 되면 `'`를 무시해주기 떄문에, 편하게 구분이 가능하게 되었다.

```cpp
int x = 0b1010;
cout << x << endl; // 10

int x = 0b1010'1111'1010; // with quota
cout << x << endl; // 10
```

#### Symbolic Constants
> C++ 11 constexpr


```cpp
// Both is allowed
const double gravity { 9.8 };
double const gravity2 { 9.8 };  

cout << gravity << endl;
```

`const`는 보통은 앞에 붙인다. pointer ref를 배우게 되면 
const의 순서에 따라서 의미상 차이를 가지게 된다.

- runtime constants (<-> compile time constants)

```cpp

const int compile_time_const(123); // compile time

int num;
cin >> num;

const int runtime_const(num); // runtime
```

**c++ 11 부터는 runtime const와 compile-time const를 구분해주기 위해서 constexpr이 도입되었다.** 

- `constexpr`: 컴파일 타임에 initialize되는 상수를 뜻함

```cpp
constexpr int compile_time_const(123); // compile time

int num;
cin >> num;

const int runtime_const(num); // runtime
```

또한 constatns들은 일반적으로 하나의 파일에 몰아서 사용한다.

- `MY_CONSTANTS.h`
```cpp
#pragma once

namespace constants
{
	constexpr double pi(3.141592);
	constexpr double avogadro(6.22123e23);
	constexpr double gravity(9.8);
}
```

```cpp
#include <iostream>
#include "MY_CONSTANTS.h"

using namespace std;

int main()
{
	cout << int(constants::pi) << endl;
	return 0;
}

```

## 4. Variable
#### variable scope 

```cpp
using namespace std;

int main() {
    const int apple = 5;
    
    {
        cout << apple << endl;	 	// 5
        int apple = 1;
        cout << apple << endl;    // 1
    }
    
    cout << apple << endl; 				// 5
    return 0;
}
```

const를 사용하더라도 중괄호 안에서 변수는 새롭게 할당되기 때문에 할당이 가능하다.

1. Global variable
	- `cout << ::value << endl;`
2. Static  variable
3. Internal Linkage: `static int g_x;`
4. External Linkage
	- `int g_x;`
	- `extern int g_x;`
	- `extern const int g_x;`

#### `Static variable in a Function`

- os로 부터 메모리를 빌려와서, program lifetime 동안 재사용된다.
- **선언된 scope 블록 안에 제한된다.** 즉 scope를 벗어난 공간에서, 해당 variable을 참조할 수 없다.( Global과의 차이)

> *It gets allocated for the lifetime of the program. Even if the function is called multiple times, space for the static variable is allocated only once and the value of variable in the previous call gets carried through the next function call.*

```cpp
#include <iostream>

using namespace std;

void doSomething()
{
    static int a = 1; // It is called only once.
    ++a;
    cout << a << endl;
}

int main()
{
    doSomething(); // 2
    doSomething(); // 3
    doSomething(); // 4
    doSomething(); // 5
    return 0;
}
```

**디버깅할 때, 함수가 몇번 호출되는지 확인하고 싶을 때 유용하게 사용가능하다.**

> 전역변수 vs Static variable
>> static은 접근 scope안에서만 할당이 가능한데 반하여, 전역변수는 실수로 다른 scope에서 할당을 하게 되면 원치 않은 결과를 만들어낼 수 있다.

#### Linkage

local variable은 해당 소스코드(모듈)에서만 사용되므로, linkage 시켜주지 않는다.
<br/>

#### `Extern`
> global 

```cpp
// forward declaration
extern void doSomething();
extern int a;

int main()
{
	...
}
```

참고로 extern은 생략 가능하다.

{{< admonition warning "상수 메모리 낭비">}}
*header 파일에 const를 선언 및 할당까지 한 뒤, 외부 .cpp 파일들에서 이를 include 시키게 되면, 신기하게 모듈별로 const의 주소가 다르게 나온다. **즉, 메모리 낭비가 생긴다. 이를 방지하기 위해서는 header에는 signature를 넣어주고, extern const의 할당은 .cpp파일에서 하게되면 된다.***
{{< /admonition >}}

- `MyConstants.h`
```cpp
#pragma once

namespace constants
{
	extern const double pi;
	extern const double avogadro;
	extern const double gravity;
}
```

- MyConstants.cpp
```cpp
#include <iostream>

namespace constants
{
	extern const double pi(3.141592);
	extern const double avogadro(6.22123e23);
	extern const double gravity(9.8);
}
```

- helloworld.cpp
```cpp
#include <iostream>
#include "CONSTS.h"

using namespace std;

void doSomething();

int main()
{
	cout << int(constants::pi) << " " << &constants::pi << endl; // 3 0x104037d70
	doSomething(); // 3 0x104037d70

	return 0;
}
```

- helloworld2.cpp

```cpp
#include <iostream>
#include "CONSTS.h"

using namespace std;

void doSomething()
{
	cout << int(constants::pi) << " " << &constants::pi << endl;
}
```

## `Using`
- scope를 최대한 작게 가져가는게 좋다.
- 가능하면 .cpp에서 사용하는 것이 좋다.
- 전역 사용만큼은 무조건 피해라.

```cpp
namespace a
{
	int dup_int(10);
}

namespace b
{
	int dup_int(20);
}

int main()
{
	using namespace std;

	{
		using namespace a;
		cout << dup_int << endl; // 10
	}
	{
		using namespace b;
		cout << dup_int << endl; // 20
	}
}
```

## `Auto`
> Type inference

- 함수의 return type에 대해서도 auto를 사용할 수 있다.

```cpp
auto add(int x, int y)
{
	return x + y;
}
```

- `trailing return type`: 친절하게 설명을 위해서 사용

```cpp
auto add(int x, int y) -> int;
auto add(double x, double y) -> double;

auto add(int x, int y) -> int
{
	return x + y;
}
```

## std::string



## Function Parameter
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

## Pointer
### Pointer and Const
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

