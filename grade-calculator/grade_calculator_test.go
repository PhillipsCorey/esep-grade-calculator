package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 40, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestString(t *testing.T) {
	expected_value := "assignment"
	
	actual_value := Assignment.String()
	if expected_value != actual_value {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 71, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator(false)

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 68, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestCallingPFWhileFalse(t *testing.T) {
	expected_value := true

	gradeCalculator := NewGradeCalculator(false)

	// if calling ReturnPF while not using PF, we will return true if numericalGrade >= 60
	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 71, Essay)

	actual_value := gradeCalculator.isPassingPF()

	if expected_value != actual_value {
		t.Errorf("Expected isPassingPF to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestCallingPFWhilePass(t *testing.T) {
	expected_value := true

	gradeCalculator := NewGradeCalculator(true)
	gradeCalculator.markPassing()

	actual_value := gradeCalculator.isPassingPF()

	if expected_value != actual_value {
		t.Errorf("Expected isPassingPF to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestCallingPFWhileFail(t *testing.T) {
	expected_value := true

	gradeCalculator := NewGradeCalculator(true)
	gradeCalculator.markPassing()
	gradeCalculator.markFailing()

	actual_value := gradeCalculator.isPassingPF()

	if expected_value != actual_value {
		t.Errorf("Expected isPassingPF to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestCallingPFWhileDefault(t *testing.T) {
	expected_value := false // students are considered failing by default in my PF

	gradeCalculator := NewGradeCalculator(true)

	actual_value := gradeCalculator.isPassingPF()

	if expected_value != actual_value {
		t.Errorf("Expected isPassingPF to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestNumericWhileUsingPF(t *testing.T) {
	expected_value := "USING PF" // GetFinalGrade returns warning if PF is active, regardless of numeric grades

	gradeCalculator := NewGradeCalculator(true)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected isPassingPF to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}