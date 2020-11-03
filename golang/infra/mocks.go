package infra

import (
	"time"

	"github.com/adrianosiq/teste-backends/golang/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockString struct {
	mock.Mock
}

func (m MockString) Fields() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

func (m MockString) Split(input string) []string {
	args := m.Called(input)
	return args.Get(0).([]string)
}

func (m MockString) SelectType(schema string, action string, fields []string) ([]string, error) {
	args := m.Called(schema, action, fields)
	return args.Get(0).([]string), args.Error(1)
}

func (m MockString) CreateJSON(fields []string, values []string) ([]byte, error) {
	args := m.Called(fields, values)
	return args.Get(0).([]byte), args.Error(1)
}

var MockMatchedBy = func(x interface{}) bool { return true }

var MockFieldsExp = []string{
	"event_id", "event_schema", "event_action", "event_timestamp", "proposal_id",
	"proposal_loan_value", "proposal_number_of_monthly_installments",
	"warranty_id", "warranty_value", "warranty_province",
	"proponent_id", "proponent_name", "proponent_age", "proponent_monthly_income", "proponent_is_main",
}

var MockSelectedFields = []string{
	"event_id", "event_schema", "event_action", "event_timestamp", "proposal_id",
	"proposal_loan_value", "proposal_number_of_monthly_installments",
}

var MockStringToJSON = []string{
	"c2d06c4f-e1dc-4b2a-af61-ba15bc6d8610,proposal,created,2019-11-11T13:26:04Z,bd6abe95-7c44-41a4-92d0-edf4978c9f4e,684397.0,72",
}

var MockStringSplit = []string{
	"c2d06c4f-e1dc-4b2a-af61-ba15bc6d8610",
	"proposal",
	"created",
	"2019-11-11T13:26:04Z",
	"bd6abe95-7c44-41a4-92d0-edf4978c9f4e",
	"684397.0",
	"72",
}

var parseEventTimestamp, _ = time.Parse(time.RFC3339, "2019-11-11T13:26:04Z")

var MockEventExp = []domain.Event{
	{
		EventID:                             uuid.MustParse("c2d06c4f-e1dc-4b2a-af61-ba15bc6d8610"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      parseEventTimestamp,
		ProposalID:                          uuid.MustParse("bd6abe95-7c44-41a4-92d0-edf4978c9f4e"),
		ProposalLoanValue:                   684397.0,
		ProposalNumberOfMonthlyInstallments: 72,
	},
}

//MockEvents is resposible for data testing events
var MockEvents = []domain.Event{
	{
		EventID:                             uuid.MustParse("c283491e-c546-4468-a264-ceba753c9de0"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      time.Now(),
		ProposalID:                          uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		ProposalLoanValue:                   50000.0,
		ProposalNumberOfMonthlyInstallments: 180,
	},
	{
		EventID:          uuid.MustParse("36cedc9f-f636-47bc-b27e-ef3a86b50616"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		WarrantyID:       uuid.MustParse("aa795eff-ecbb-4de6-915e-92bb91beebf5"),
		WarrantyValue:    5679086.17,
		WarrantyProvince: "BA",
	},
	{
		EventID:          uuid.MustParse("c3768df0-6a82-4350-a193-20b6296afa47"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		WarrantyID:       uuid.MustParse("87e1a079-3d20-4255-adfa-aae8b3ba0a06"),
		WarrantyValue:    6139723.75,
		WarrantyProvince: "ES",
	},
	{
		EventID:                uuid.MustParse("f27566eb-4df6-4879-be4f-cc8b8fc5f1ce"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		ProponentID:            uuid.MustParse("2f384cfb-5b25-44f2-8989-0321ab687cd7"),
		ProponentName:          "Ms. Daryl Pfannerstill",
		ProponentAge:           18,
		ProponentMonthlyIncome: 1120.0,
		ProponentIsMain:        true,
	},
	{
		EventID:                uuid.MustParse("700e70c5-8ff1-4902-a3b9-0c546f7d7a0b"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
		ProponentID:            uuid.MustParse("e1042bdc-28f4-47f5-92fe-dcc55a12e88e"),
		ProponentName:          "Kerry Collier",
		ProponentAge:           19,
		ProponentMonthlyIncome: 61400.0,
		ProponentIsMain:        false,
	},
	{
		EventID:                             uuid.MustParse("41a8d96b-b454-434b-9f55-02dabb99cb74"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      time.Now(),
		ProposalID:                          uuid.MustParse("56d54d2e-0cde-4aec-89c8-5108d488d262"),
		ProposalLoanValue:                   50000.0,
		ProposalNumberOfMonthlyInstallments: 180,
	},
	{
		EventID:          uuid.MustParse("10cf5ee6-6498-4e36-ab48-ada80d466ad7"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("56d54d2e-0cde-4aec-89c8-5108d488d262"),
		WarrantyID:       uuid.MustParse("a29db9e2-c52c-4a1c-a39f-5ecaf026557a"),
		WarrantyValue:    3236698.01,
		WarrantyProvince: "BA",
	},
	{
		EventID:          uuid.MustParse("eedce8ae-d4e7-48e6-88eb-65ad98a2fa6b"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("56d54d2e-0cde-4aec-89c8-5108d488d262"),
		WarrantyID:       uuid.MustParse("81312022-095f-4e3a-a550-2f4986e5667b"),
		WarrantyValue:    2892918.65,
		WarrantyProvince: "ES",
	},
	{
		EventID:                uuid.MustParse("dd05eeee-3f1d-4860-8d26-0ed018f6185e"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("56d54d2e-0cde-4aec-89c8-5108d488d262"),
		ProponentID:            uuid.MustParse("b5237edc-6071-41d9-ba0d-055fce616f84"),
		ProponentName:          "Eloy Goodwin",
		ProponentAge:           24,
		ProponentMonthlyIncome: 1112.0,
		ProponentIsMain:        true,
	},
	{
		EventID:                uuid.MustParse("bed74014-5691-4b39-a77c-9bd13a86afbf"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("56d54d2e-0cde-4aec-89c8-5108d488d262"),
		ProponentID:            uuid.MustParse("1e310124-2314-4658-88e0-431f3324ef4b"),
		ProponentName:          "Del Schneider",
		ProponentAge:           20,
		ProponentMonthlyIncome: 87981.42,
		ProponentIsMain:        false,
	},
	{
		EventID:                             uuid.MustParse("86d0ba39-389b-41c7-a17d-ce2ab3727d1d"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      time.Now(),
		ProposalID:                          uuid.MustParse("1b548c88-7109-4aed-a199-757cbe56294e"),
		ProposalLoanValue:                   50000.0,
		ProposalNumberOfMonthlyInstallments: 180,
	},
	{
		EventID:          uuid.MustParse("5fc848bb-17a3-46de-81bc-99eaeadcc15e"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("1b548c88-7109-4aed-a199-757cbe56294e"),
		WarrantyID:       uuid.MustParse("a057d397-42ce-4c45-b48b-f2dfedee1757"),
		WarrantyValue:    6920045.92,
		WarrantyProvince: "GO",
	},
	{
		EventID:          uuid.MustParse("c72d30cb-1c20-4592-afb2-4a2296deaa31"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("1b548c88-7109-4aed-a199-757cbe56294e"),
		WarrantyID:       uuid.MustParse("ed68a36a-8e7a-4c4e-829f-1bf1f03f1cb0"),
		WarrantyValue:    7791279.21,
		WarrantyProvince: "DF",
	},
	{
		EventID:                uuid.MustParse("4cca2a57-4e30-4f21-818e-53fb6e1bcd21"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("1b548c88-7109-4aed-a199-757cbe56294e"),
		ProponentID:            uuid.MustParse("e20db52a-05b1-4827-bda9-dd86c0cd13e9"),
		ProponentName:          "Morgan Luettgen",
		ProponentAge:           55,
		ProponentMonthlyIncome: 500.0,
		ProponentIsMain:        true,
	},
	{
		EventID:                uuid.MustParse("068f5b3c-165e-4078-9ddc-732e2581459f"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("1b548c88-7109-4aed-a199-757cbe56294e"),
		ProponentID:            uuid.MustParse("93ed2942-f9e8-4238-b94e-e5e955a9db81"),
		ProponentName:          "Fredericka Effertz",
		ProponentAge:           61,
		ProponentMonthlyIncome: 352195.52,
		ProponentIsMain:        false,
	},
	{
		EventID:                             uuid.MustParse("bce1e4e7-1bad-4ef2-8a12-6bb044b0f3fb"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      time.Now(),
		ProposalID:                          uuid.MustParse("713a33d1-c246-4287-b770-42b2219f6c34"),
		ProposalLoanValue:                   50000.0,
		ProposalNumberOfMonthlyInstallments: 180,
	},
	{
		EventID:          uuid.MustParse("0036e0ee-06cb-4add-8291-a0fad480d923"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("713a33d1-c246-4287-b770-42b2219f6c34"),
		WarrantyID:       uuid.MustParse("d3490ce2-16f7-4e2c-a544-52b403bf7cf9"),
		WarrantyValue:    7897100.18,
		WarrantyProvince: "DF",
	},
	{
		EventID:          uuid.MustParse("0c0f7761-7d62-4d01-8c5d-9f3bf6cf61f9"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("713a33d1-c246-4287-b770-42b2219f6c34"),
		WarrantyID:       uuid.MustParse("c169069a-7ea7-4674-84c5-52771559c30e"),
		WarrantyValue:    7539192.57,
		WarrantyProvince: "GO",
	},
	{
		EventID:                uuid.MustParse("8ec8a5d8-c4fb-4155-ab4a-0d073a326bb1"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("713a33d1-c246-4287-b770-42b2219f6c34"),
		ProponentID:            uuid.MustParse("b5708629-8d15-4edc-b312-dd80e540edcb"),
		ProponentName:          "Cory Rice",
		ProponentAge:           18,
		ProponentMonthlyIncome: 500.0,
		ProponentIsMain:        true,
	},
	{
		EventID:                uuid.MustParse("701d1f12-2c39-4424-b4a4-f65d685d6739"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("713a33d1-c246-4287-b770-42b2219f6c34"),
		ProponentID:            uuid.MustParse("9a3655fc-5c78-404b-83f4-93361e56e300"),
		ProponentName:          "Santos Gottlieb",
		ProponentAge:           20,
		ProponentMonthlyIncome: 405801.49,
		ProponentIsMain:        false,
	},
	{
		EventID:                             uuid.MustParse("2584d9f0-2647-41c0-89ec-459f1c3a0f84"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      time.Now(),
		ProposalID:                          uuid.MustParse("dd6011d2-7202-42e5-8dc1-372316af38f8"),
		ProposalLoanValue:                   50000.0,
		ProposalNumberOfMonthlyInstallments: 180,
	},
	{
		EventID:          uuid.MustParse("6c9226ad-9896-4b66-8983-fb7a13a5133d"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("dd6011d2-7202-42e5-8dc1-372316af38f8"),
		WarrantyID:       uuid.MustParse("316ee9c9-ae10-4b19-bd4a-3c1a066e90fe"),
		WarrantyValue:    3148988.79,
		WarrantyProvince: "DF",
	},
	{
		EventID:          uuid.MustParse("607ffe90-c6a2-4e9b-83f7-c5ad7deacd70"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("dd6011d2-7202-42e5-8dc1-372316af38f8"),
		WarrantyID:       uuid.MustParse("f9fe6ae1-9094-4367-ae37-655c94c1a454"),
		WarrantyValue:    2861268.2,
		WarrantyProvince: "GO",
	},
	{
		EventID:                uuid.MustParse("a86882bc-d65d-4efc-929e-685efd33c81b"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("dd6011d2-7202-42e5-8dc1-372316af38f8"),
		ProponentID:            uuid.MustParse("b800674a-cf26-47ff-85e7-88f52bd6586e"),
		ProponentName:          "Anna Koepp",
		ProponentAge:           24,
		ProponentMonthlyIncome: 500.0,
		ProponentIsMain:        true,
	},
	{
		EventID:                uuid.MustParse("a66f374e-73ef-427f-ada1-7edebcc53bdc"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("dd6011d2-7202-42e5-8dc1-372316af38f8"),
		ProponentID:            uuid.MustParse("66d533ba-a158-4a18-9082-24820ade05e4"),
		ProponentName:          "Neoma Bruen",
		ProponentAge:           30,
		ProponentMonthlyIncome: 106064.86,
		ProponentIsMain:        false,
	},
	{
		EventID:                             uuid.MustParse("e058782a-d520-470b-ab71-7aaf138d57ce"),
		EventSchema:                         "proposal",
		EventAction:                         "created",
		EventTimestamp:                      time.Now(),
		ProposalID:                          uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
		ProposalLoanValue:                   50000.0,
		ProposalNumberOfMonthlyInstallments: 180,
	},
	{
		EventID:          uuid.MustParse("0a413116-47a9-45a6-ac0d-d47cfda91292"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
		WarrantyID:       uuid.MustParse("b3006fec-f934-49c2-9248-9f033de5ac37"),
		WarrantyValue:    2753912.59,
		WarrantyProvince: "DF",
	},
	{
		EventID:          uuid.MustParse("bf94d992-1c27-4118-9fa2-48badbf7062e"),
		EventSchema:      "warranty",
		EventAction:      "added",
		EventTimestamp:   time.Now(),
		ProposalID:       uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
		WarrantyID:       uuid.MustParse("1bcaafa1-dadc-46db-93bb-3d0f553b1753"),
		WarrantyValue:    3628485.48,
		WarrantyProvince: "GO",
	},
	{
		EventID:                uuid.MustParse("cf1f047a-4eb0-4b78-887b-0a45fa2cc9da"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
		ProponentID:            uuid.MustParse("9be533ae-0075-4d2e-91a2-3eef197d1457"),
		ProponentName:          "Crissy Spinka",
		ProponentAge:           55,
		ProponentMonthlyIncome: 200.0,
		ProponentIsMain:        true,
	},
	{
		EventID:                uuid.MustParse("8a824bf3-b251-4231-84a3-2117131e40ea"),
		EventSchema:            "proponent",
		EventAction:            "added",
		EventTimestamp:         time.Now(),
		ProposalID:             uuid.MustParse("924166e0-1d9b-4cf4-8d68-f59a80060844"),
		ProponentID:            uuid.MustParse("1af45f04-68ee-4aba-b178-63e3eaadbe24"),
		ProponentName:          "Teddy Senger",
		ProponentAge:           48,
		ProponentMonthlyIncome: 52370.85,
		ProponentIsMain:        false,
	},
}

//MockApprovedExp is responsible response process expect
var MockApprovedExp = []domain.ProposalApproved{
	{
		ProposalID: uuid.MustParse("dad6498a-78ad-4a4f-b1e1-34e1f240000c"),
	},
	{
		ProposalID: uuid.MustParse("56d54d2e-0cde-4aec-89c8-5108d488d262"),
	},
}
