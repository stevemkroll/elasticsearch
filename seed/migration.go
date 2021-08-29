package seed

import (
	"elasticsearch/employee"
	"elasticsearch/task"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func migration(db *gorm.DB) error {
	log.Printf("running migration...\n")
	if db = db.Exec("DROP TABLE IF EXISTS tasks,employees"); db.Error != nil {
		return db.Error
	}
	log.Printf("creating task table...\n")
	if err := db.AutoMigrate(&task.Task{}); err != nil {
		return err
	}
	log.Printf("creating employee table...\n")
	if err := db.AutoMigrate(&employee.Employee{}); err != nil {
		return err
	}
	return nil
}

func seedEmployees(db *gorm.DB) error {
	log.Printf("creating employee records...\n")
	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("employee_%d@gmint.io", i)
		phone := fmt.Sprintf("111111111%d", i)
		emp, err = employee.New(email, phone)
		if err != nil {
			return err
		}
		if db := db.Create(&emp); db.Error != nil {
			return db.Error
		}
		employeeList = append(employeeList, emp)
		log.Printf("%+v... OK\n", emp.ID)
	}
	return nil
}

func seedTasks(db *gorm.DB) error {
	log.Printf("creating task records...\n")
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("task_%d", i)
		tsk, err = task.New(name)
		if err != nil {
			return err
		}
		if db := db.Create(&tsk); db.Error != nil {
			return db.Error
		}
		taskList = append(taskList, tsk)
		log.Printf("%+v... OK\n", tsk.ID)
	}
	return nil
}

func assignTasks(db *gorm.DB) error {
	for i := range employeeList {
		emp := employee.Employee{}
		if db := db.Raw(`SELECT * FROM employees WHERE id=?`, employeeList[i].ID).Scan(&emp); db.Error != nil {
			return err
		}
		switch {
		case i%3 == 1:
			emp.AssignedTo = append(emp.AssignedTo, taskList[i].Name, taskList[i-1].Name)
		default:
			emp.AssignedTo = append(emp.AssignedTo, taskList[i].Name)
		}
		if db := db.Save(&emp); db.Error != nil {
			return err
		}
		employeeList[i] = emp
	}
	return nil
}
