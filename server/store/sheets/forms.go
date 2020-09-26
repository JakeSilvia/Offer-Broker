package sheets

import "backup/server/entities"

const (
	FormSheet = "forms"
)

func (s *Sheet) GetForms() ([]*entities.Form, error) {
	s.Sheet = FormSheet
	s.StartRange = "A2"
	s.StopRange = "Z99999999"

	resp, err := s.doReadRequest()
	if err != nil {
		return nil, err
	}

	forms := make([]*entities.Form, 0)
	for _, formValue := range resp.Values {
		form := &entities.Form{}
		form.ToForm(formValue)
		forms = append(forms, form)
	}

	return forms, nil
}
func (s *Sheet) SaveForm(form *entities.Form) error {
	s.Sheet = FormSheet
	s.StartRange = "A2"
	s.StopRange = "append"
	update := &SheetResp{
		Values:         [][]string{form.ToList()},
		MajorDimension: "ROWS",
	}

	return s.doAppendRequest(update)
}
