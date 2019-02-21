package models

import (
	am "backend/models/cmsmodels/assignmentmodels"
	cm "backend/models/cmsmodels/coursemodels"
	sm "backend/models/cmsmodels/submissionmodels"
	um "backend/models/usermodels"
)

type (
	Assignment am.MongoAssignment
	Course     cm.MongoCourse
	User       um.MongoUser
	Submission sm.MongoSubmission
)

func NewMongoAssignmentInterface() *am.AssignmentInterface {
	return am.New()
}

func NewMongoCourseInterface() *cm.CourseInterface {
	return cm.New()
}

func NewMongoUserInterface() *um.UserInterface {
	return um.New()
}

func NewMongoSubmissionInterface() *sm.SubmissionInterface {
	return sm.New()
}
