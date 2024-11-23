package dao

import (
	"fmt"
	"mindguard/model"
)

// 根据order_id ordered_id 获取预约
func GetRes(teacher_id, student_id int64) *model.Reservation {
	var reservation model.Reservation
	err := DB.Where("order_id = ? AND ordered_id = ?", student_id, teacher_id).Find(&reservation).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &reservation
}

// 获取所有与学生有关的预约
func QueryReservationAboutStudent(student_id int64) []model.Reservation {
	var reservations []model.Reservation
	err := DB.Where("order_id = ?", student_id).Find(&reservations).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return reservations
}

// 修改老师的预约
func UpdateTeacherReservation(student_id, teacher_id int64) bool {
	tx := DB.Model(&model.Reservation{}).Where("order_id = ? AND ordered_id = ?", student_id, teacher_id).
		Update("status", "预约成功")
	fmt.Println(tx)
	if tx.RowsAffected != 1 {
		return false
	}
	return true
}

// 获取所有与老师有关的预约
func QueryReservationsAboutTeacher(teacherId int64) []model.Reservation {
	var reservations []model.Reservation
	err := DB.Where("ordered_id = ?", teacherId).Find(&reservations).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return reservations
}

// 获取所有预约
func GetAllReservations() []model.Reservation {
	var reservations []model.Reservation
	err := DB.Find(&reservations).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return reservations
}

// 删除预约
func DeleteReservation(appointmentId int64) bool {
	var reservation model.Reservation
	err := DB.Where("id = ?", appointmentId).Delete(&reservation).Error
	if err != nil {
		return false
	}
	return true
}

// 添加预约
func InsertReservation(reservation *model.Reservation) *model.Reservation {
	err := DB.Create(reservation).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("reservation", reservation)
	return reservation
}

// 获取所有教师
func GetAllTeachers() []model.User {
	var users []model.User
	err := DB.Where("status = ?", "教师").Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return users
}
