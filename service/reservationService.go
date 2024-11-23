package service

import (
	"mindguard/dao"
	"mindguard/model"
	"time"
)

type ReservationService struct {
}

// 获取相应的预约
func (rs *ReservationService) GetRes(teacher_id, student_id int64) *model.Reservation {
	// 从数据库获取相应的预约
	reservation := dao.GetRes(teacher_id, student_id)
	return reservation
}

// 获取所有与学生有关的预约
func (rs *ReservationService) GetStudentReservations(student_id int64) []model.User {
	// 从数据库 获取学生预约
	reservations := dao.QueryReservationAboutStudent(student_id)

	// 根据查询的所有预约 通过orderedid 查询所有学生
	var teachers []model.User

	for _, reservation := range reservations {
		teacher := dao.QueryUserById(int(reservation.OrderedId))
		teachers = append(teachers, *teacher)
	}
	return teachers
}

// 老师接受预约
func (rs *ReservationService) AcceptOrder(student_id, teacher_id int64) bool {
	// 从数据库 修改预约
	return dao.UpdateTeacherReservation(student_id, teacher_id)
}

// 获取所有与老师有关的预约
func (rs *ReservationService) GetTeacherReservations(teacherId int64) []model.User {
	// 从数据库获取所有与老师有关的预约
	reservations := dao.QueryReservationsAboutTeacher(teacherId)
	if reservations == nil {
		return nil
	}

	// 根据查询的所有预约 通过orderid 查询所有学生
	var students []model.User

	for _, reservation := range reservations {
		student := dao.QueryUserById(int(reservation.OrderId))
		students = append(students, *student)
	}
	return students

}

// 取消预约
func (rs *ReservationService) CancelOrder(appointmentId int64) bool {
	// 从数据库删除预约
	return dao.DeleteReservation(appointmentId)
}

// 获取所有预约
func (rs *ReservationService) GetAllReservations() []model.Reservation {
	// 从数据库获取所有预约
	reservations := dao.GetAllReservations()
	return reservations
}

// 发起预约
func (rs *ReservationService) PostOrder(orderId, orderedeId int64) *model.Reservation {
	// 从数据库添加预约
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	reservation := model.Reservation{
		OrderId:         orderId,
		OrderedId:       orderedeId,
		AppointmentTime: currentTime,
		Status:          "预约中",
	}
	re := dao.InsertReservation(&reservation)
	if re != nil {
		return re
	}

	return nil
}

// 获取所有教师
func (rs *ReservationService) GetAllTeachers() []model.User {
	// 从数据库查询所有在线教师
	teachers := dao.GetAllTeachers()
	return teachers
}
