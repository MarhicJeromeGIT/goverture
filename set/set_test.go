package set

import "testing"

func TestNewSet(t *testing.T) {
	s := NewSet()
	if s == nil {
		t.Errorf("Expected new set to be created")
	}
}

func TestElements(t *testing.T) {
	s := NewSet()
	if len(s.Elements()) != 0 {
		t.Errorf("Expected no elements in the set")
	}
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if len(s.Elements()) != 3 {
		t.Errorf("Expected 3 elements in the set")
	}
	// assert what are the elements
	expected_elements := []int{1, 2, 3}
	for i, val := range s.Elements() {
		if val != expected_elements[i] {
			t.Errorf("Expected element %v but got %v", expected_elements[i], val)
		}
	}
}

func TestAdd(t *testing.T) {
    s := NewSet()
	if s.Contains(1) {
		t.Errorf("Did not expect set to contain 1")
	}
    s.Add(1)
    if !s.Contains(1) {
        t.Errorf("Expected 1 to be added to the set")
    }
}

func TestRemove(t *testing.T) {
	s := NewSet()
	s.Remove(2)
	s.Add(2)
	s.Remove(2)
	if s.Contains(2) {
		t.Errorf("Did not expect set to contain 2")
	}
}

func TestContains(t *testing.T) {
    s := NewSet()
    s.Add(2)
    if !s.Contains(2) {
        t.Errorf("Expected set to contain 2")
    }
    if s.Contains(3) {
        t.Errorf("Did not expect set to contain 3")
    }
}

func TestString(t *testing.T) {
	s := NewSet()
	if s.String() != "{}" {
		t.Errorf("Expected string representation to be {} but got %v", s.String())
	}
	s.Add(2)
	s.Add(3)
	if s.String() != "{2, 3}" {
		t.Errorf("Expected string representation to be {2, 3} but got %v", s.String())
	}
}