package array

var Bugs = []Bug{
	{"Out-of-bound array access", "Logic Error", 788, "Выход за пределы массива"},
	{"Out-of-bound access", "Logic Error", 788, "Выход за пределы"},
	{"Call to blocking function in critical section", "Blocking Error", 557, "Вызов блокирующей функции в критической секции"},
	{"Assignment of a non-Boolean value", "Logic Error", 681, "Присвоение не-булевого значения"},
	{"API", "Logic Error", 1006, "API"},
	{"Out-of-bound array access", "Logic Error", 788, "Выход за пределы массива"},
	{"Anti-pattern in the argument", "C String API", 398, "Анти-паттерн в аргументе"},
	{"Anti-pattern in the argument", "C String API", 398, "Анти-паттерн в аргументе"},
	{"Receiver in message expression is 'nil'", "Logic Error", 465, "В выражении, описывающем сообщение, получатель -- 'nil'"},
	{"Subscript access on an uninitialized object pointer", "Logic Error", 824, "Доступ по смещению через неинициализированный указатель на объект"},
	{"Property access on an uninitialized object pointer", "Logic Error", 824, "Доступ к свойству объекта через неинициализированный указатель на объект"},
	{"Receiver in message expression is an uninitialized value", "Logic Error", 824, "В выражении, описывающем сообщение, получатель не инициализирован"},
	{"Called C++ object pointer is null", "Logic Error", 476, "Попытка вызова по нулевому указателю на объект C++"},
	{"Called C++ object pointer is uninitialized", "Logic Error", 824, "Попытка вызова по неинициализированному указателю на объект C++"},
	{"Uninitialized argument value", "Logic Error", 824, "Неинициализированное значение аргумента"},
	{"Called function pointer is null (null dereference)", "Logic Error", 476, "Попытка вызова по нулевому указателю на функцию (разыменование нулевого указателя)"},
	{"Called function pointer is an uninitialized pointer value", "Logic Error", 824, "Попытка вызова по неинициализированному указателю на функцию"},
	{"Cast region with wrong size.", "Logic Error", 704, "Приведение типа региона к типу неподходящего размера"},
	{"Widening cast to struct type", "Logic error", 704, "Попытка преобразовать структурный тип в супертип"},
	{"Cast from non-struct type to struct type", "Logic error", 704, "Преобразование типа из неструктурного в структурный"},
	{"Mistaken dealloc", "Memory (Core Foundation/Objective-C/OSObject)", 399, "Ошибочный вызов dealloc"},
	{"Missing ivar release (leak)", "Memory (Core Foundation/Objective-C/OSObject)", 401, "Отсутствует освобождение ivar (утечка)"},
	{"Incompatible instance method return type", "Core Foundation/Objective-C", 389, "Несовместимый тип возвращаемого значения в методе экземпляра"},
	{"Potential insecure implementation-specific behavior in call 'vfork'", "Security", 242, "Потенциально небезопасное поведение, определяемое реализацией, в вызове 'vfork'"},
	{"'random' is not a secure random number generator", "Security", 338, "'random' не является безопасным генератором случайных чисел"},
	{"Potential insecure memory buffer bounds restriction in call 'strcat'", "Security", 120, "Потенциально небезопасное задание границ буфера при вызове 'strcat'"},
	{"Potential insecure memory buffer bounds restriction in call 'strcpy'", "Security", 120, "Потенциально небезопасное задание границ буфера при вызове 'strcpy'"},
	{"Insecure temporary file creation", "Security", 378, "Небезопасное создание временного файла"},
	{"Potential insecure temporary file in call 'mktemp'", "Security", 378, "Потенциально небезопасный временный файл в вызове 'mktemp'"},
	{"Potential buffer overflow in call to 'getpw'", "Security", 242, "Потенциальное переполнение буфера при вызове 'getpw'"},
	{"Potential buffer overflow in call to 'gets'", "Security", 242, "Потенциальное переполнение буфера при вызове 'gets'"},
	{"Use of deprecated function in call to 'bzero()'", "Security", 477, "Использование устаревшей функции при вызове 'bzero()'"},
	{"Use of deprecated function in call to 'bcopy()'", "Security", 477, "Использование устаревшей функции при вызове 'bcopy()'"},
	{"Use of deprecated function in call to 'bcmp()'", "Security", 477, "Использование устаревшей функции при вызове 'bcmp()'"},
	{"Potential unintended use of sizeof() on pointer type", "Logic error", 467, "Возможно непреднамеренное применение sizeof() к указателю"},
	{"Break out of jail", "Logic Error", 243, "Выход за пределы jail"},
	{"Suspicious code clone", "Code clone", 1041, "Подозрительная копия кода"},
	{"Exact code clone", "Code clone", 1041, "Точная копия кода"},
	{"Conversion", "Logic Error", 681, "Преобразование"},
	{"Destruction of a polymorphic object with no virtual destructor", "Logic error", 1079, "Попытка деструкции полиморфного объекта, не имеющего виртуального деструктора"},
	{"Dereference of undefined pointer value", "Logic Error", 824, "Разыменование указателя с неопределенным значением"},
	{"Dereference of null pointer", "Logic Error", 476, "Разыменование нулевого указателя"},
	{"Property access", "Core Foundation/Objective-C", 1006, "Получение доступа к свойству"},
	{"Division by zero", "Logic Error", 369, "Деление на ноль"},
	{"Dynamic and static type mismatch", "Type Error", 136, "Несоответствие динамического и статического типов"},
	{"Generics", "Core Foundation/Objective-C", 704, "Дженерики"},
	{"Enum cast out of range", "Logic Error", 704, "Приводимое значение выходит за пределы перечисляемого типа"},
	{"Use fixed address", "Logic Error", 587, "Использование фиксированного адреса"},
	{"GCD performance anti-pattern", "Performance", 398, "Анти-паттерн производительности GCD"},
	{"Use of Untrusted Data", "Untrusted Data", 502, "Использование ненадежных данных"},
	{"Identical expressions in conditional expression", "Logic error", 569, "Идентичные выражения в условии"},
	{"Compare of identical expressions", "Logic error", 569, "Сравнение идентичных выражений"},
	{"Identical branches", "Logic error", 1041, "Идентичные ветки"},
	{"Identical conditions", "Logic error", 1041, "Идентичные условия"},
	{"Identical conditions", "Logic error", 1041, "Идентичные условия"},
	{"Use of identical expressions", "Logic error", 1041, "Использование идентичных выражений"},
	{"Iterator invalidated", "Misuse of STL APIs", 825, "Итератор инвалидирован"},
	{"Iterator out of range", "Misuse of STL APIs", 465, "Итератор вышел за пределы"},
	{"Incomplete invalidation", "Core Foundation/Objective-C", 459, "Неполная инвалидация"},
	{"Incomplete invalidation", "Core Foundation/Objective-C", 459, "Неполная инвалидация"},
	{"Incomplete invalidation", "Core Foundation/Objective-C", 459, "Неполная инвалидация"},
	{"AST node allocates heap memory", "LLVM Conventions", 399, "Узел AST выделяет память в куче"},
	{"Plural Misuse", "Localizability Issue (Apple)", 398, "Неверное использование множественного числа"},
	{"Context Missing", "Localizability Issue (Apple)", 398, "Отсутствует контекст"},
	{"Unlocalizable string", "Localizability Issue (Apple)", 398, "Непереводимая строка"},
	{"Improper use of SecKeychain API", "API Misuse (Apple)", 399, "Неправильное использование API SecKeychain"},
	{"Improper use of 'dispatch_once'", "API Misuse (Apple)", 1006, "Неправильное использование 'dispatch_once'"},
	{"Memory leak", "Memory error", 401, "Утечка памяти"},
	{"Bad free", "Memory error", 763, "Плохой вызов free"},
	{"Use of zero allocated", "Memory error", 399, "Попытка выделить блок памяти нулевого размера"},
	{"Double delete", "Memory error", 415, "Двойной вызов delete"},
	{"Double free", "Memory error", 415, "Двойной вызов free"},
	{"Use-after-free", "Memory error", 416, "Использование после освобождения"},
	{"Offset free", "Memory error", 761, "Вызов free с неправильным адресом блока памяти"},
	{"Bad deallocator", "Memory error", 763, "Плохой деаллокатор"},
	{"Free alloca()", "Memory error", 590, "Попытка освободить при помощи free блок, созданный при помощи alloca()"},
	{"Bad free", "Memory error", 763, "Плохой вызов free"},
	{"malloc() size overflow", "Unix API", 789, "Переполнение размера при вызове malloc"},
	{"Allocator sizeof operand mismatch", "Unix API", 131, "Allocator sizeof operand mismatch"},
	{"W^X check fails, Write Exec prot flags set", "Security", 200, "Проверка W^X не пройдена, установлены оба флага - Write и Exec"},
	{"Use-after-move", "C++ move semantics", 399, "Попытка использования объекта после того, как он был перемещен (use-after-move)"},
	{"Use -drain instead of -release", "API Upgrade (Apple)", 1006, "Используйте -drain вместо -release"},
	{"CFErrorRef* null dereference", "Coding conventions (Apple)", 476, "Попытка разыменования нулевого указателя CFErrorRef*"},
	{"Bad return type when passing CFErrorRef*", "Coding conventions (Apple)", 1006, "Плохой тип возвращаемого значения при передаче CFErrorRef*"},
	{"Bad return type when passing NSError**", "Coding conventions (Apple)", 1006, "Плохой тип возвращаемого значения при передаче NSError**"},
	{"Dereference of null pointer", "Logic Error", 476, "Разыменование нулевого указателя"},
	{"Argument with 'nonnull' attribute passed null", "API", 476, "Аргумент с атрибутом 'nonnull' равен нулевому указателю"},
	{"Nullability", "Memory error", 476, "Возможность обнуления"},
	{"Suspicious number object conversion", "Logic error", 682, "Подозрительное преобразование числового объекта"},
	{"OSObject C-Style Cast", "Security", 704, "Приведение типа OSObject в C-стиле"},
	{"Nil value used as mutex for @synchronized() (no synchronization will occur)", "Logic Error", 662, "Значение nil используется как мьютекс для @synchronized (синхронизации не будет)"},
	{"Uninitialized value used as mutex for @synchronized", "Logic Error", 662, "Неинициализированное значение используется как мьютекс для @synchronized"},
	{"CFArray API", "Core Foundation/Objective-C", 465, "CFArray API"},
	{"Objective-C property misuse", "Logic error", 1006, "Неправильное использование свойства в Objective-C"},
	{"[super dealloc] should not be called more than once", "Core Foundation/Objective-C", 399, "[super dealloc] не должен вызываться более одного раза"},
	{"Unused instance variable", "Optimization", 1164, "Неиспользованная переменная в экземпляре"},
	{"Dangerous pointer arithmetic", "Logic Error", 188, "Небезопасная арифметика указателей"},
	{"Dangerous pointer arithmetic", "Logic Error", 188, "Небезопасная арифметика указателей"},
	{"Iteration of pointer-like elements", "Non-determinism", 465, "Итерирование указателеподобных элементов"},
	{"Sorting of pointer-like elements", "Non-determinism", 465, "Сортировка указателеподобных элементов"},
	{"Pointer subtraction", "Logic Error", 469, "Вычитание указателей"},
	{"Use destroyed lock", "Lock checker", 667, "Использование уничтоженного мьютекса"},
	{"Init invalid lock", "Lock checker", 667, "Инициализация невалидного мьютекса"},
	{"Destroy invalid lock", "Lock checker", 667, "Уничтожение невалидного мьютекса"},
	{"Lock order reversal", "Lock checker", 667, "Нарушение порядка блокировки"},
	{"Double unlocking", "Lock checker", 667, "Двойная разблокировка"},
	{"Double locking", "Lock checker", 667, "Двойная блокировка"},
	{"Return of pointer value outside of expected range", "Logic Error", 466, "Возвращаемое значение указателя выходит за ожидаемые пределы"},
	{"Returning null reference", "Logic Error", 465, "Возвращается нулевая ссылка"},
	{"Garbage return value", "Logic Error", 465, "Возвращаемое значение содержит мусор"},
	{"Memory leak inside autorelease pool", "Memory", 399, "Утечка памяти в пуле автоматического освобождения"},
	{"Double fclose", "Unix Stream API Error", 910, "Двойной вызов fclose"},
	{"Resource Leak", "Logic Error", 775, "Утечка ресурса"},
	{"Double fclose", "Logic Error", 910, "Двойной вызов fclose"},
	{"Illegal whence argument", "Logic Error", 628, "Недопустимое значение точки отсчёта"},
	{"NULL stream pointer", "Logic Error", 476, "Указатель на поток равен NULL"},
	{"Tainted data", "General", 19, "Помеченные данные"},
	{"Division by zero", "Logic Error", 369, "Деление на ноль"},
	{"Branch condition evaluates to a garbage value", "Logic Error", 465, "Мусорное значение в условии ветвления"},
	{"uninitialized variable captured by block", "Logic Error", 824, "Блок захватил неинициализированную переменную"},
	{"Result of operation is garbage or undefined", "Logic Error", 465, "Результат операции не определён или содержит мусор"},
	{"Array subscript is undefined", "Logic Error", 824, "Индекс массива не определен"},
	{"Unreachable code", "Dead code", 561, "Недостижимый код"},
	{"Dangerous variable-length array (VLA) declaration", "Logic Error", 789, "Небезопасное объявление массива переменной длины (VLA)"},
	{"Leaked va_list", "Memory error", 401, "Утечка va_list"},
	{"Uninitialized va_list", "Memory error", 824, "Неинициализированный va_list"},
	{"Dangerous construct in a vforked process", "Logic Error", 399, ""}}
