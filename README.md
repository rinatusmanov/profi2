Владельцу крупной фирмы, оказывающей услуги по ремонту помещений, для организации рабочего процесса необходима информационная система, облегчающая сбор и хранение заказов, распределение сотрудников на заказы.

В основе системы должна лежать реляционная база данных, содержащая информацию, необходимую для ввода и хранения следующих данных:

Сведений о заказе, включая контактные данные (телефон или e-mail) заказчика, объем ремонтируемого помещения и перечень необходимых работ (черновые работы, отделочные работы, работа электрика, монтажные работы). Ответственный за ввод информации – пользователь с ролью «Клиент».
Сведений о статусе заказа («Новый» (по умолчанию), «На согласовании», «Открыт», «Закрыт»). Ответственный за ввод информации – пользователь с ролью «Владелец».
Сведений о сотрудниках фирмы, включая их паспортные данные, квалификацию (перечень выполняемых работ). Ответственный за ввод информации – пользователь с ролью «Владелец».
Сведений о назначении сотрудника фирмы на требуемый вид работы отдельного заказа с указанием временного интервала выполнения работы с наименьшей единицей выполнения в один день. Следует учитывать, что один сотрудник в один день может выполнять только один вид работ на одном заказе. Ответственный за ввод информации – пользователь с ролью «Владелец».
Справочной информации о количестве сотрудников, необходимом в день для выполнения каждой работы из перечня работ, и общем количестве дней выполнения каждой работы в зависимости от категории помещения, определяемой объемом помещения («Малое помещение», «Среднее помещение», «Большое помещение»). Ответственный за ввод информации – пользователь с ролью «Владелец».
Информационная система должна поддерживать вывод следующей информации для пользователей в зависимости от роли:

Для пользователя с ролью «Клиент» система выводит ориентировочную продолжительность ремонта и требуемое количество сотрудников каждой квалификации по заданному заказу.
Для пользователя с ролью «Владелец» система отображает перечень заказов со статусом, отличным от «Закрыт», с указанием длительности ремонта; также система отображает по каждому сотруднику заказы, выполняемые сотрудником в текущий момент времени, и позволяет проверить заказы, выполняемые сотрудником, на отсутствие пересечений по времени.
Необходимо:

Используя любую общепринятую нотацию, изобразить схему инфологической модели предметной области.
Используя любую общепринятую нотацию, изобразить схему даталогической модели базы данных, удовлетворяющую третьей нормальной форме, с выделением первичных и внешних ключей, типа и направления связей.
Изобразить прототипы web-страниц работы пользователей с ролями «Клиент» и «Владелец» с системой. 
Прототипы должны иллюстрировать возможности ввода и вывода доступной пользователям информации.
Описать алгоритм расчета для данного заказа временных интервалов ремонта с указанием количества требуемых сотрудников по каждой из работ исходя из доступности сотрудников для выполнения работы (незанятости на другой работе).
Описать алгоритм проверки сотрудника на занятость в другом заказе.
Используя операторы языка SQL, написать запрос для вывода списка сотрудников, у которых за последние 12 месяцев было не менее 6 заказов на ремонт помещений категории НЕ «Малое помещение».
Требования к структуре оформления решения с указанием критериев оценивания и максимального количества баллов за каждую часть решения:

1. Введение - до 2 баллов.

2. Основная часть:

2.1. Схема решения, описание процессов, описание информационных потоков–до 5 баллов.

2.2. Структура базы данных, отражающая специфику предметной области. БД должна соответствовать третьей нормальной форме, не быть перегруженной дублированием, учитывать сущности-справочники, не иметь ошибок по связям– до 5 баллов.

2.3. Прототипы визуальных интерфейсов – до 5 баллов.

2.4. Описание алгоритмов (в виде схемы алгоритмов, программного кода или псевдокода) – до 7 баллов.

2.5. Описание SQL-запрoca – до 4 баллов.

3. Заключение (выводы) – до 2 баллов.

К вопросу еще не были загружены файлы

Приложить файл с ответом.

Файл