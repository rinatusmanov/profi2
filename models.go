package main

import (
	"time"
)

//Сведений о заказе, включая контактные данные (телефон или e-mail) заказчика, объем ремонтируемого помещения и перечень необходимых работ
//(черновые работы, отделочные работы, работа электрика, монтажные работы). Ответственный за ввод информации – пользователь с ролью «Клиент».
//Сведений о статусе заказа («Новый» (по умолчанию), «На согласовании», «Открыт», «Закрыт»).
//Ответственный за ввод информации – пользователь с ролью «Владелец».
//Сведений о сотрудниках фирмы, включая их паспортные данные, квалификацию (перечень выполняемых работ).
//Ответственный за ввод информации – пользователь с ролью «Владелец».
//Сведений о назначении сотрудника фирмы на требуемый вид работы отдельного заказа с указанием временного интервала выполнения работы
//с наименьшей единицей выполнения в один день. Следует учитывать, что один сотрудник в один день может выполнять
//только один вид работ на одном заказе. Ответственный за ввод информации – пользователь с ролью «Владелец».
//Справочной информации о количестве сотрудников, необходимом в день для выполнения каждой работы из перечня работ,
//и общем количестве дней выполнения каждой работы в зависимости от категории помещения, определяемой объемом помещения
//(«Малое помещение», «Среднее помещение», «Большое помещение»). Ответственный за ввод информации – пользователь с ролью «Владелец».


type Status struct {
	StatusID          uint `gorm:"primarykey"`
	StatusCreatedAt   time.Time
	StatusUpdatedAt   time.Time
	StatusDeletedAt   *time.Time `gorm:"index"`
	StatusName        string
	StatusDescription string
}

type WorkType struct {
	WorkID          uint `gorm:"primarykey"`
	WorkCreatedAt   time.Time
	WorkUpdatedAt   time.Time
	WorkDeletedAt   *time.Time `gorm:"index"`
	WorkName        string
	WorkDescription string
}

type Customer struct {
	CustomerID        uint `gorm:"primarykey"`
	CustomerCreatedAt time.Time
	CustomerUpdatedAt time.Time
	CustomerDeletedAt *time.Time `gorm:"index"`
	CustomerPhone     string
	CustomerEMail     string
}

type Room struct {
	RoomID          uint `gorm:"primarykey"`
	RoomCreatedAt   time.Time
	RoomUpdatedAt   time.Time
	RoomDeletedAt   *time.Time `gorm:"index"`
	RoomName        string
	RoomAddress     string
	RoomDescription string
	RoomVolume      float64
}

type RoomType struct {
	RoomTypeID         uint `gorm:"primarykey"`
	RoomTypeCreatedAt  time.Time
	RoomTypeUpdatedAt  time.Time
	RoomTypeDeletedAt  *time.Time `gorm:"index"`
	RoomTypeName       string
	RoomTypeFromVolume float64
	RoomTypeToVolume   float64
}

type Order struct {
	OrderID         uint `gorm:"primarykey"`
	OrderCreatedAt  time.Time
	OrderUpdatedAt  time.Time
	OrderDeletedAt  *time.Time `gorm:"index"`
	OrderCustomerID uint
	OrderCustomer   Customer
	OrderWork       []WorkType `gorm:"many2many:order_works;"`
	OrderRoomID     uint
	OrderRoom       Room
	OrderGannets    []OrderGantt `gorm:"foreignKey:OrderID"`
	StatusID        uint
	Status          Status
}

type OrderGantt struct {
	OrderGanttID        uint `gorm:"primarykey"`
	OrderGanttCreatedAt time.Time
	OrderGanttUpdatedAt time.Time
	OrderGanttDeletedAt *time.Time `gorm:"index"`
	OrderID             uint
	OrderGanttFromDate  time.Time
	OrderGanttToDate    time.Time
	EmployeeID          uint
	Employee            Employee
}

type Employee struct {
	EmployeeID        uint `gorm:"primarykey"`
	EmployeeCreatedAt time.Time
	EmployeeUpdatedAt time.Time
	EmployeeDeletedAt *time.Time `gorm:"index"`
	EmployeeWork      []WorkType `gorm:"many2many:employee_works;"`
	EmployeeDocuments []Document `gorm:"foreignKey:EmployeeID"`
}

type Document struct {
	DocumentID        uint `gorm:"primarykey"`
	DocumentCreatedAt time.Time
	DocumentUpdatedAt time.Time
	DocumentDeletedAt *time.Time `gorm:"index"`
	DocumentInfo      string
	EmployeeID        uint
}
