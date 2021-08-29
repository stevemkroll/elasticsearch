package seed

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func clusterReady(readyChan chan string) {
	log.Printf("waiting for cluster...\n")
	uri := "http://localhost:9200/_cluster/health"
	for {
		res, err := http.Get(uri)
		if err == nil && res.StatusCode == http.StatusOK {
			readyChan <- "cluster... OK"
		}
		time.Sleep(time.Second)
	}
}

func removeIndices() error {
	log.Printf("removing indices...\n")
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:9200/_all", nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	log.Printf("remove indices... OK\n")
	return nil
}

func createIndices() error {
	log.Printf("creating indices...\n")
	req, err := http.NewRequest(http.MethodPut, "http://localhost:9200/employee", nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	log.Printf("employee index... OK\n")

	req, err = http.NewRequest(http.MethodPut, "http://localhost:9200/task", nil)
	if err != nil {
		return err
	}
	res, err = client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	log.Printf("task index... OK\n")
	return nil
}

func createDocuments() error {
	log.Printf("creating documents...\n")
	for _, i := range employeeList {
		b, err := json.Marshal(i)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:9200/employee/_doc/%s", i.ID), bytes.NewBufferString(string(b)))
		if err != nil {
			return err
		}
		req.Header.Add("content-type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusCreated {
			return errors.New(res.Status)
		}
		log.Printf("%s... OK\n", i.ID)
	}
	for _, i := range taskList {
		b, err := json.Marshal(i)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:9200/task/_doc/%s", i.ID), bytes.NewBuffer(b))
		if err != nil {
			return err
		}
		req.Header.Add("content-type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusCreated {
			return errors.New(res.Status)
		}
		log.Printf("%s... OK\n", i.ID)
	}
	return nil
}
