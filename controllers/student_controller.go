package controllers

import (
	"CRUDProject/dao"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/students", dao.CreateStudentHandler).Methods("POST")
	r.HandleFunc("/students/{id}", dao.GetStudentByIDHandler).Methods("GET")
	r.HandleFunc("/students", dao.GetStudentsHandler).Methods("GET")
	r.HandleFunc("/students/{id}", dao.UpdateStudentHandler).Methods("PUT")
	r.HandleFunc("/students/{id}", dao.DeleteStudentHandler).Methods("DELETE")

	r.HandleFunc("/teachers", dao.CreateTeacherHandler).Methods("POST")
	r.HandleFunc("/teachers/{id}", dao.GetTeacherByIDHandler).Methods("GET")
	r.HandleFunc("/teachers", dao.GetTeachersHandler).Methods("GET")
	r.HandleFunc("/teachers/{id}", dao.UpdateTeacherHandler).Methods("PUT")
	r.HandleFunc("/teachers/{id}", dao.DeleteTeacherHandler).Methods("DELETE")

	r.HandleFunc("/classes", dao.CreateClassHandler).Methods("POST")
	r.HandleFunc("/classes/{id}", dao.GetClassByIDHandler).Methods("GET")
	r.HandleFunc("/classes", dao.GetClassesHandler).Methods("GET")
	r.HandleFunc("/classes/{id}", dao.UpdateClassHandler).Methods("PUT")
	r.HandleFunc("/classes/{id}", dao.DeleteClassHandler).Methods("DELETE")
	r.HandleFunc("/classes/register", dao.RegisterClassHandler).Methods("POST")

	r.HandleFunc("/subjects", dao.CreateSubjectHandler).Methods("POST")
	r.HandleFunc("/subjects/{id}", dao.GetSubjectByIDHandler).Methods("GET")
	r.HandleFunc("/subjects", dao.GetSubjectsHandler).Methods("GET")
	r.HandleFunc("/subjects/{id}", dao.UpdateSubjectHandler).Methods("PUT")
	r.HandleFunc("/subjects/{id}", dao.DeleteSubjectHandler).Methods("DELETE")

	return r
}
