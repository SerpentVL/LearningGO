/*  -- Программа угадывает число, задуманное пользователем --
    --
    Метод:          бинарный поиск
    Диапазон чисел: задается параметром командной строки
    По умолчанию:   1 - 30
    --
*/
#include <iostream>
#include <string>

using namespace std;

int main()
{
    int number = 30;
    cout << "Привет, Анюта!" << endl;
    cout << "--------------" << endl;
    cout << "Задумай число от 1 до " << number << "." << endl;
    cout << "Я его угадаю. " << endl;
    cout << "--------------" << endl;
    cout << "Я буду задавать вопросы, а ты будешь мне отвечать на клавиатуре. " << endl;
    cout << "По-английски. " << endl;
    cout << "Если твой ответ - да,  то набирай на клавиатуре yes " << endl;
    cout << "Если твой ответ - нет, то набирай на клавиатуре no " << endl;
    cout << "И потом нажми Enter" << endl;
    cout << "Ты поняла ?" << endl;
    string answer;
    cin >> answer;

    if(answer == "yes"){
        cout << "\nПопилили!" << endl;
    }



    return 0;
}
