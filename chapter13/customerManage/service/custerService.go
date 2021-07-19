package service

import (
	"chapterDemo/chapter13/customerManage/model"
)

//
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "12345678", "150379@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (c *CustomerService) List() []model.Customer {
	return c.customers
}

//添加用户
func (c *CustomerService) Add(customer model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}
func (c *CustomerService) Delete(id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

//
func (c *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].Id == id {
			index = i
		}
	}
	return index
}

//修改信息
func (c *CustomerService) UpDate(index int, name string, gender string, age int, phone string, email string) {
	c.customers[index].Name = name
	c.customers[index].Gender = gender
	c.customers[index].Age = age
	c.customers[index].Phone = phone
	c.customers[index].Email = email
}
