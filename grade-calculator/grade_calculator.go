package esepunittests

type GradeCalculator struct {
	allAssignments []Grade
	usingPF bool
	PFstate bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(usingPF bool) *GradeCalculator {
	return &GradeCalculator{
		allAssignments: make([]Grade, 0),
		usingPF: usingPF,
		PFstate: false,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.allAssignments = append(gc.allAssignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})
	case Exam:
		gc.allAssignments = append(gc.allAssignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.allAssignments = append(gc.allAssignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.filterOneTypeFromList(Assignment))
	exam_average := computeAverage(gc.filterOneTypeFromList(Exam))
	essay_average := computeAverage(gc.filterOneTypeFromList(Essay))

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func (gc *GradeCalculator) filterOneTypeFromList(t GradeType) []Grade {
	var output []Grade 
	
	for _, gradeObj := range gc.allAssignments {
		if gradeObj.Type == t {
			output = append(output, gradeObj)
		}
	}

	return output
}

func computeAverage(grades []Grade) float64 {
	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return float64(sum) / float64(len(grades))
}
