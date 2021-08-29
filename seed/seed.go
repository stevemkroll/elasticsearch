package seed

import (
	"elasticsearch/employee"
	"elasticsearch/task"
	"errors"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var client = &http.Client{}

var (
	ErrMigration        error = errors.New("err running migration")
	ErrSeedingEmployees error = errors.New("err seeding employees")
	ErrSeedingTasks     error = errors.New("err seeding tasks")
	ErrAssigningTasks   error = errors.New("err assigning tasks")
	ErrClusterHealth    error = errors.New("err cluster health")
	ErrClusterIndex     error = errors.New("err creating cluster index")
	ErrClusterDocument  error = errors.New("err creating cluster document")
)

var (
	employeeList []employee.Employee
	taskList     []task.Task
)
var (
	emp employee.Employee
	tsk task.Task
	err error
)

func Run(db *gorm.DB) {
	if err = migration(db); err != nil {
		log.Printf("%s\n", ErrMigration.Error())
		panic(err)
	}
	if err = seedEmployees(db); err != nil {
		log.Printf("%s\n", err.Error())
		panic(ErrSeedingEmployees)
	}
	if err = seedTasks(db); err != nil {
		log.Printf("%s\n", err.Error())
		panic(ErrSeedingTasks)
	}

	if err = assignTasks(db); err != nil {
		log.Printf("%s\n", err.Error())
		panic(ErrAssigningTasks)
	}

	readyChan := make(chan string)
	go clusterReady(readyChan)
	log.Printf("%s\n", <-readyChan)

	if err := removeIndices(); err != nil {
		log.Println(err.Error())
		panic(ErrClusterIndex)
	}

	if err := createIndices(); err != nil {
		log.Println(err.Error())
		panic(ErrClusterIndex)
	}

	if err := createDocuments(); err != nil {
		log.Println(err.Error())
		panic(ErrClusterDocument)
	}

}
