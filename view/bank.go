package view

import (
	"fmt"
	"strconv"
)

func GetBankFields(m *TeaModel) *[]field {
	banks, err := m.services.bankCore.GetAllBanks()

  banksFields := []field{}

	if err != nil {
		return &banksFields
	}

	for index, bank := range banks {
		banksFields = append(banksFields, field{name: fmt.Sprintf("bank %d", index), value: strconv.FormatFloat(bank.Amount, 'f', -1, 64)})
	}
  return &banksFields
}

func BankView(m *TeaModel) string {
  var s string

  // m.fields = &GetBankFields(m)
 
	for i, field := range m.fields {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s [%s]: %s\n", cursor, field.name, field.value)
	}

  return s
}
