package response

import (
	"fmt"

	"github.com/handarudwiki/models"
)

type ClassroomTaskResponse struct {
	ID          int                `json:"id"`
	ClassroomID int                `json:"classroom_id"`
	Classroom   *ClassroomResponse `json:"classroom,omitempty"`
	TaskID      int                `json:"task_id"`
	Task        *TaskResponse      `json:"task,omitempty"`
	TeacherID   int                `json:"teacher_id"`
	Teacher     *UserResponse      `json:"teacher,omitempty"`
}

func ToClassroomTaskResponse(classroomTask *models.ClassroomTask) ClassroomTaskResponse {

	// fmt.Println(classroomTask.ID)

	var classroom *ClassroomResponse
	var task *TaskResponse
	var user *UserResponse

	// Periksa apakah ClassroomTask.Classroom ada
	if classroomTask.Classroom != nil {
		fmt.Println("Classroom")
		fmt.Println("Ini test : ", classroomTask.Classroom.Name)
		classroomResp := ToClassroomResponse(classroomTask.Classroom)
		classroom = &classroomResp
	}

	// Periksa apakah ClassroomTask.Task ada
	if classroomTask.Task != nil {
		fmt.Println("Task")
		fmt.Println("Ini test : ", classroomTask.Task.Title)
		taskResp := ToTaskResponse(classroomTask.Task)
		task = &taskResp
	}

	// Periksa apakah ClassroomTask.Teacher ada
	fmt.Println("Teacher")
	fmt.Println("Ini test : ", classroomTask.Teacher)
	if classroomTask.Teacher != nil {
		fmt.Println("Teacher")
		fmt.Println("Ini test : ", classroomTask.Teacher.Name)
		userResp := ToUserResponse(classroomTask.Teacher)
		user = &userResp
	}

	return ClassroomTaskResponse{
		ID:          int(classroomTask.ID),
		ClassroomID: int(classroomTask.ClassroomID),
		Classroom:   classroom, // Nilainya bisa nil atau pointer ke object
		TaskID:      int(classroomTask.TaskID),
		Task:        task, // Nilainya bisa nil atau pointer ke object
		TeacherID:   int(classroomTask.TeacherID),
		Teacher:     user, // Nilainya bisa nil atau pointer ke object
	}
}

func ToClassroomTaskResponseSlice(classroomTasks []*models.ClassroomTask) []ClassroomTaskResponse {
	var classroomTaskResponseSlice []ClassroomTaskResponse
	for _, classroomTask := range classroomTasks {
		classroomTaskResponseSlice = append(classroomTaskResponseSlice, ToClassroomTaskResponse(classroomTask))
	}
	return classroomTaskResponseSlice
}
