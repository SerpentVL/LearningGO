package array

// Это - база переводов для типа сообщения.
// см. также ../trace_translator/messages.go - там переводятся сами сообщения.
// но если сообщение совпадает с типом, то переводы сообщения берутся из этой базы bugs

var Bugs = []Bug{
	{"0.7", "Out-of-bound array access", "Logic Error", 788, "Выход за пределы массива"},                                              // ArrayBoundChecker.cpp:69, "Access out-of-bound array element (buffer overflow)"
	{"0.7", "Out-of-bound access", "Logic Error", 788, "Выход за пределы"},                                                            // ArrayBoundCheckerV2.cpp:250 + complicate
	{"0.7", "Call to blocking function in critical section", "Blocking Error", 557, "Вызов блокирующей функции в критической секции"}, // BlockInCriticalSectionChecker.cpp:81
	{"0.3", "Assignment of a non-Boolean value", "Logic Error", 681, "Присвоение не-булевого значения"},                               // BoolAssignmentChecker.cpp:36
	{"0.5", "API", "Logic Error", 1006, "API"},
	{"0.7", "Out-of-bound array access", "Logic Error", 788, "Выход за пределы массива"},
	{"0.3", "Anti-pattern in the argument", "C String API", 398, "Анти-паттерн в аргументе"},
	{"0.3", "Anti-pattern in the argument", "C String API", 398, "Анти-паттерн в аргументе"},
	{"0.5", "Receiver in message expression is 'nil'", "Logic Error", 465, "В выражении, описывающем сообщение, получатель -- 'nil'"},                                       // CallAndMessageChecker.cpp:509
	{"0.7", "Subscript access on an uninitialized object pointer", "Logic Error", 824, "Доступ по смещению через неинициализированный указатель на объект"},                 // CallAndMessageChecker.cpp:478
	{"0.7", "Property access on an uninitialized object pointer", "Logic Error", 824, "Доступ к свойству объекта через неинициализированный указатель на объект"},           // CallAndMessageChecker.cpp:472
	{"0.7", "Receiver in message expression is an uninitialized value", "Logic Error", 824, "В выражении, описывающем сообщение, получатель не инициализирован"},            // CallAndMessageChecker.cpp:465
	{"0.8", "Called C++ object pointer is null", "Logic Error", 476, "Попытка вызова по нулевому указателю на объект C++"},                                                  // CallAndMessageChecker.cpp:391
	{"0.7", "Called C++ object pointer is uninitialized", "Logic Error", 824, "Попытка вызова по неинициализированному указателю на объект C++"},                            // CallAndMessageChecker.cpp:379
	{"0.5", "Uninitialized argument value", "Logic Error", 824, "Неинициализированное значение аргумента"},                                                                  // CallAndMessageChecker.cpp:355
	{"0.8", "Called function pointer is null (null dereference)", "Logic Error", 476, "Попытка вызова по нулевому указателю на функцию (разыменование нулевого указателя)"}, // CallAndMessageChecker.cpp:335
	{"0.7", "Called function pointer is an uninitialized pointer value", "Logic Error", 824, "Попытка вызова по неинициализированному указателю на функцию"},                // CallAndMessageChecker.cpp:324
	{"0.5", "Cast region with wrong size.", "Logic Error", 704, "Приведение типа региона к типу неподходящего размера"},
	{"0.5", "Widening cast to struct type", "Logic error", 704, "Попытка преобразовать структурный тип в супертип"},
	{"0.5", "Cast from non-struct type to struct type", "Logic error", 704, "Преобразование типа из неструктурного в структурный"},
	{"0.5", "Mistaken dealloc", "Memory (Core Foundation/Objective-C/OSObject)", 399, "Ошибочный вызов dealloc"},
	{"0.7", "Missing ivar release (leak)", "Memory (Core Foundation/Objective-C/OSObject)", 401, "Отсутствует освобождение ivar (утечка)"}, // CheckObjCDealloc.cpp:757
	{"0.6", "Incompatible instance method return type", "Core Foundation/Objective-C", 389, "Несовместимый тип возвращаемого значения в методе экземпляра"},
	{"0.7", "Potential insecure implementation-specific behavior in call 'vfork'", "Security", 242, "Потенциально небезопасное поведение, определяемое реализацией, в вызове 'vfork'"}, // CheckSecuritySyntaxOnly.cpp:934
	{"0.7", "'random' is not a secure random number generator", "Security", 338, "'random' не является безопасным генератором случайных чисел"},                                        // CheckSecuritySyntaxOnly.cpp:914
	{"1.0", "Potential insecure memory buffer bounds restriction in call 'strcat'", "Security", 120, "Потенциально небезопасное задание границ буфера при вызове 'strcat'"},            // CheckSecuritySyntaxOnly.cpp:699
	{"1.0", "Potential insecure memory buffer bounds restriction in call 'strcpy'", "Security", 120, "Потенциально небезопасное задание границ буфера при вызове 'strcpy'"},            // CheckSecuritySyntaxOnly.cpp:727
	{"0.7", "Insecure temporary file creation", "Security", 378, "Небезопасное создание временного файла"},                                                                             // CheckSecuritySyntaxOnly.cpp:666
	{"0.7", "Potential insecure temporary file in call 'mktemp'", "Security", 378, "Потенциально небезопасный временный файл в вызове 'mktemp'"},                                       // CheckSecuritySyntaxOnly.cpp:581
	{"0.3", "Potential buffer overflow in call to 'getpw'", "Security", 242, "Потенциальное переполнение буфера при вызове 'getpw'"},                                                   // CheckSecuritySyntaxOnly.cpp:539
	{"0.3", "Potential buffer overflow in call to 'gets'", "Security", 242, "Потенциальное переполнение буфера при вызове 'gets'"},                                                     // CheckSecuritySyntaxOnly.cpp:499
	{"0.2", "Use of deprecated function in call to 'bzero()'", "Security", 477, "Использование устаревшей функции при вызове 'bzero()'"},                                               // CheckSecuritySyntaxOnly.cpp:462
	{"0.2", "Use of deprecated function in call to 'bcopy()'", "Security", 477, "Использование устаревшей функции при вызове 'bcopy()'"},                                               // CheckSecuritySyntaxOnly.cpp:420
	{"0.2", "Use of deprecated function in call to 'bcmp()'", "Security", 477, "Использование устаревшей функции при вызове 'bcmp()'"},                                                 // CheckSecuritySyntaxOnly.cpp:379
	{"0.5", "Potential unintended use of sizeof() on pointer type", "Logic error", 467, "Возможно непреднамеренное применение sizeof() к указателю"},
	{"0.8", "Break out of jail", "Logic Error", 243, "Выход за пределы jail"},
	{"0.3", "Suspicious code clone", "Code clone", 1041, "Подозрительная копия кода"},
	{"0.3", "Exact code clone", "Code clone", 1041, "Точная копия кода"},
	{"0.3", "Conversion", "Logic Error", 681, "Преобразование"},
	{"0.5", "Destruction of a polymorphic object with no virtual destructor", "Logic error", 1079, "Попытка деструкции полиморфного объекта, не имеющего виртуального деструктора"},
	{"0.7", "Dereference of undefined pointer value", "Logic Error", 824, "Разыменование указателя с неопределенным значением"}, // DereferenceChecker.cpp:201
	{"0.8", "Dereference of null pointer", "Logic Error", 476, "Разыменование нулевого указателя"},                              // DereferenceChecker.cpp:129 NonNullParamChecker.cpp:211
	{"0.3", "Property access", "Core Foundation/Objective-C", 1006, "Получение доступа к свойству"},
	{"1.0", "Division by zero", "Logic Error", 369, "Деление на ноль"}, // DivZeroChecker.cpp:48
	{"0.5", "Dynamic and static type mismatch", "Type Error", 136, "Несоответствие динамического и статического типов"},
	{"0.5", "Generics", "Core Foundation/Objective-C", 704, "Дженерики"},
	{"0.5", "Enum cast out of range", "Logic Error", 704, "Приводимое значение выходит за пределы перечисляемого типа"}, // EnumCastOutOfRangeChecker.cpp:83
	{"0.1", "Use fixed address", "Logic Error", 587, "Использование фиксированного адреса"},                             // FixedAddressChecker.cpp:54
	{"0.2", "GCD performance anti-pattern", "Performance", 398, "Анти-паттерн производительности GCD"},                  // GCDAntipatternChecker.cpp:195
	{"0.5", "Use of Untrusted Data", "Untrusted Data", 502, "Использование ненадежных данных"},
	{"0.2", "Identical expressions in conditional expression", "Logic error", 569, "Идентичные выражения в условии"}, // IdenticalExprChecker.cpp:285
	{"0.2", "Compare of identical expressions", "Logic error", 569, "Сравнение идентичных выражений"},                // IdenticalExprChecker.cpp:264
	{"0.3", "Identical branches", "Logic error", 1041, "Идентичные ветки"},                                           // IdenticalExprChecker.cpp:173
	{"0.3", "Identical conditions", "Logic error", 1041, "Идентичные условия"},                                       // IdenticalExprChecker.cpp:120
	{"0.3", "Identical conditions", "Logic error", 1041, "Идентичные условия"},                                       // IdenticalExprChecker.cpp:143
	{"0.3", "Use of identical expressions", "Logic error", 1041, "Использование идентичных выражений"},
	{"0.5", "Iterator invalidated", "Misuse of STL APIs", 825, "Итератор инвалидирован"},
	{"0.5", "Iterator out of range", "Misuse of STL APIs", 465, "Итератор вышел за пределы"},
	{"0.7", "Incomplete invalidation", "Core Foundation/Objective-C", 459, "Неполная инвалидация"},               // IvarInvalidationChecker.cpp:546
	{"0.7", "Incomplete invalidation", "Core Foundation/Objective-C", 459, "Неполная инвалидация"},               // IvarInvalidationChecker.cpp:564
	{"0.7", "Incomplete invalidation", "Core Foundation/Objective-C", 459, "Неполная инвалидация"},               // IvarInvalidationChecker.cpp:569
	{"0.5", "AST node allocates heap memory", "LLVM Conventions", 399, "Узел AST выделяет память в куче"},        // LLVMConventionsChecker.cpp:287
	{"0.2", "Plural Misuse", "Localizability Issue (Apple)", 398, "Неверное использование множественного числа"}, // LocalizationChecker.cpp:1387
	{"0.2", "Context Missing", "Localizability Issue (Apple)", 398, "Отсутствует контекст"},                      // LocalizationChecker.cpp:1189
	{"0.2", "Unlocalizable string", "Localizability Issue (Apple)", 398, "Непереводимая строка"},
	{"0.7", "Improper use of SecKeychain API", "API Misuse (Apple)", 399, "Неправильное использование API SecKeychain"},
	{"0.3", "Improper use of 'dispatch_once'", "API Misuse (Apple)", 1006, "Неправильное использование 'dispatch_once'"},
	{"0.7", "Memory leak", "Memory error", 401, "Утечка памяти"},
	{"0.7", "Bad free", "Memory error", 763, "Плохой вызов free"},
	{"0.5", "Use of zero allocated", "Memory error", 399, "Попытка выделить блок памяти нулевого размера"},
	{"0.2", "Double delete", "Memory error", 415, "Двойной вызов delete"},
	{"0.7", "Double free", "Memory error", 415, "Двойной вызов free"},
	{"0.7", "Use-after-free", "Memory error", 416, "Использование после освобождения"},
	{"0.7", "Offset free", "Memory error", 761, "Вызов free с неправильным адресом блока памяти"},
	{"0.7", "Bad deallocator", "Memory error", 763, "Плохой деаллокатор"},
	{"0.2", "Free alloca()", "Memory error", 590, "Попытка освободить при помощи free блок, созданный при помощи alloca()"},
	{"0.7", "Bad free", "Memory error", 763, "Плохой вызов free"},
	{"0.7", "malloc() size overflow", "Unix API", 789, "Переполнение размера при вызове malloc"},                                         // MallocOverflowSecurityChecker.cpp:286
	{"0.5", "Allocator sizeof operand mismatch", "Unix API", 131, "Аргумент sizeof не совпадает с типом указателя на выделяемую память"}, // MallocSizeofChecker.cpp:240
	{"0.8", "W^X check fails, Write Exec prot flags set", "Security", 200, "Проверка W^X не пройдена, установлены оба флага - Write и Exec"},
	{"0.5", "Use-after-move", "C++ move semantics", 399, "Попытка использования объекта после того, как он был перемещен (use-after-move)"},
	{"0.3", "Use -drain instead of -release", "API Upgrade (Apple)", 1006, "Используйте -drain вместо -release"},
	{"0.8", "CFErrorRef* null dereference", "Coding conventions (Apple)", 476, "Попытка разыменования нулевого указателя CFErrorRef*"},
	{"0.3", "Bad return type when passing CFErrorRef*", "Coding conventions (Apple)", 1006, "Плохой тип возвращаемого значения при передаче CFErrorRef*"}, // NSErrorChecker.cpp:123
	{"0.3", "Bad return type when passing NSError**", "Coding conventions (Apple)", 1006, "Плохой тип возвращаемого значения при передаче NSError**"},     // NSErrorChecker.cpp:76
	{"0.8", "Dereference of null pointer", "Logic Error", 476, "Разыменование нулевого указателя"},
	{"0.7", "Argument with 'nonnull' attribute passed null", "API", 476, "Аргумент с атрибутом 'nonnull' равен нулевому указателю"},
	{"0.6", "Nullability", "Memory error", 476, "Возможность обнуления"},
	{"0.4", "Suspicious number object conversion", "Logic error", 682, "Подозрительное преобразование числового объекта"},                                                                          // NumberObjectConversionChecker.cpp:188
	{"0.5", "OSObject C-Style Cast", "Security", 704, "Приведение типа OSObject в C-стиле"},                                                                                                        // PointerIterationChecker.cpp:52
	{"0.7", "Nil value used as mutex for @synchronized() (no synchronization will occur)", "Logic Error", 662, "Значение nil используется как мьютекс для @synchronized (синхронизации не будет)"}, // ObjCAtSyncChecker.cpp:70
	{"0.7", "Uninitialized value used as mutex for @synchronized", "Logic Error", 662, "Неинициализированное значение используется как мьютекс для @synchronized"},                                 // ObjCAtSyncChecker.cpp:47
	{"0.5", "CFArray API", "Core Foundation/Objective-C", 465, "CFArray API"},
	{"0.3", "Objective-C property misuse", "Logic error", 1006, "Неправильное использование свойства в Objective-C"},
	{"0.5", "[super dealloc] should not be called more than once", "Core Foundation/Objective-C", 399, "[super dealloc] не должен вызываться более одного раза"},
	{"0.1", "Unused instance variable", "Optimization", 1164, "Неиспользованная переменная в экземпляре"}, // ObjCUnusedIVarsChecker.cpp:165
	{"0.6", "Dangerous pointer arithmetic", "Logic Error", 188, "Небезопасная арифметика указателей"},     // PointerArithChecker.cpp:172
	{"0.6", "Dangerous pointer arithmetic", "Logic Error", 188, "Небезопасная арифметика указателей"},     // PointerArithChecker.cpp:195
	{"0.5", "Iteration of pointer-like elements", "Non-determinism", 465, "Итерирование указателеподобных элементов"},
	{"0.5", "Sorting of pointer-like elements", "Non-determinism", 465, "Сортировка указателеподобных элементов"}, // PointerSortingChecker.cpp:52
	{"0.9", "Pointer subtraction", "Logic Error", 469, "Вычитание указателей"},                                    // PointerSubChecker.cpp:63
	{"0.7", "Use destroyed lock", "Lock checker", 667, "Использование уничтоженного мьютекса"},
	{"0.7", "Init invalid lock", "Lock checker", 667, "Инициализация невалидного мьютекса"},
	{"0.7", "Destroy invalid lock", "Lock checker", 667, "Уничтожение невалидного мьютекса"},
	{"0.7", "Lock order reversal", "Lock checker", 667, "Нарушение порядка блокировки"},
	{"0.7", "Double unlocking", "Lock checker", 667, "Двойная разблокировка"},
	{"0.7", "Double locking", "Lock checker", 667, "Двойная блокировка"},
	{"0.7", "Return of pointer value outside of expected range", "Logic Error", 466, "Возвращаемое значение указателя выходит за ожидаемые пределы"}, // ReturnPointerRangeChecker.cpp:72
	{"0.5", "Returning null reference", "Logic Error", 465, "Возвращается нулевая ссылка"},                                                           // ReturnUndefChecker.cpp:116
	{"0.5", "Garbage return value", "Logic Error", 465, "Возвращаемое значение содержит мусор"},                                                      // ReturnUndefChecker.cpp:98
	{"0.7", "Memory leak inside autorelease pool", "Memory", 399, "Утечка памяти в пуле автоматического освобождения"},                               // RunLoopAutoreleaseLeakChecker.cpp:116
	{"0.1", "Double fclose", "Unix Stream API Error", 910, "Двойной вызов fclose"},                                                                   // StreamChecker.cpp:310
	{"0.6", "Resource Leak", "Logic Error", 775, "Утечка ресурса"},                                                                                   // StreamChecker.cpp:343
	{"0.1", "Double fclose", "Logic Error", 910, "Двойной вызов fclose"},
	{"0.3", "Illegal whence argument", "Logic Error", 628, "Недопустимое значение точки отсчёта"}, // StreamChecker.cpp:284
	{"0.8", "NULL stream pointer", "Logic Error", 476, "Указатель на поток равен NULL"},           // StreamChecker.cpp:255
	{"0.5", "Tainted data", "General", 19, "Помеченные данные"},
	{"1.0", "Division by zero", "Logic Error", 369, "Деление на ноль"},                                                               // TestAfterDivZeroChecker.cpp:168
	{"0.5", "Branch condition evaluates to a garbage value", "Logic Error", 465, "Мусорное значение в условии ветвления"},            // UndefBranchChecker.cpp:68
	{"0.7", "uninitialized variable captured by block", "Logic Error", 824, "Блок захватил неинициализированную переменную"},         // UndefCapturedBlockVarChecker.cpp:77
	{"0.5", "Result of operation is garbage or undefined", "Logic Error", 465, "Результат операции не определён или содержит мусор"}, // UndefResultChecker.cpp:94
	{"0.7", "Array subscript is undefined", "Logic Error", 824, "Индекс массива не определен"},                                       // UndefinedArraySubscriptChecker.cpp:52
	{"0.2", "Unreachable code", "Dead code", 561, "Недостижимый код"},                                                                // RunLoopAutoreleaseLeakChecker.cpp:116
	{"0.5", "Dangerous variable-length array (VLA) declaration", "Logic Error", 789, "Небезопасное объявление массива переменной длины (VLA)"},
	{"0.7", "Leaked va_list", "Memory error", 401, "Утечка va_list"},
	{"0.7", "Uninitialized va_list", "Memory error", 824, "Неинициализированный va_list"},
	{"0.7", "Dangerous construct in a vforked process", "Logic Error", 399, "Небезопасная конструкция в процессе, вызванном vfork()"},

	{"0.3", "Dead initialization", "Dead store", 563, "Неиспользуемая инициализация"},             // DeadStoresChecker.cpp
	{"0.3", "Dead increment", "Dead store", 563, "Неиспользуемый инкремент"},                      // DeadStoresChecker.cpp
	{"0.3", "Dead assignment", "Dead store", 563, "Неиспользуемое присваивание"},                  // DeadStoresChecker.cpp:247
	{"0.3", "Dead nested assignment", "Dead store", 563, "Неиспользуемое вложенное присваивание"}, // DeadStoresChecker.cpp:255

	{"0.3", "Possible loss of sign/precision.", "Logic Error", 681, "Возможнаяя потеря точности/знака числа"},                              // ConversionChecker.cpp:121
	{"0.8", "No call of chdir(\"/\") immediately after chroot", "Break out of jail", 243, "Нет вызова chdir(\"/\") сразу после chroot"},    // ChrootChecker.cpp:127
	{"0.4", "Uninitialized fields", "Logic Error", 456, "Неинициализированные переменные"},                                                 // UninitializedObject/UninitializedObjectChecker.cpp:48
	{"0.5", "Sum of expressions causes overflow.", "API", 190, "Сумма выражений вызывает переполнение"},                                    // CStringChecker.cpp:640
	{"0.4", "Argument is not a null-terminated string.", "Unix API", 686, "Аргумент не является нуль-терминированной строкой"},             // CStringChecker.cpp:623
	{"0.7", "Null pointer argument in call to byte string function", "Unix API", 688, "Нулевой указатель как аргумент при вызове функции"}, // CStringChecker.cpp:583

	{"0.3", "Floating point variable used as loop counter", "Security", 254, "Переменная с плавающей запятой используется как счетчик цикла"}, // CheckSecuritySyntaxOnly.cpp:338
	{"0.8", "Missing call to superclass", "Logic Error", 388, "Отсутствует вызов суперкласса"},                                                // ObjCMissingSuperCallChecker.cpp:208
	{"0.5", "Assigned value is garbage or undefined", "Logic Error", 465, "Присвоенное значение содержит мусор или неопределено"},             // UndefinedAssignmentChecker:56
	//	{"0.7", "Extra ivar release", "Logic Error", 0, ""} // CheckObjCDealloc.cpp:761 // понятия не имею как перевести и какую cwe присвоить
}
