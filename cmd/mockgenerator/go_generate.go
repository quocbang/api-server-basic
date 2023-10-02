package main

// use mockery to generate mock file
// install mock toll
//   - go install github.com/vektra/mockery/v2@latest

//go:generate mockery --name DataManager --filename mock_dm.go --dir ../../database --output ../../database/mocks
//go:generate mockery --name Services --filename mock_service.go --dir ../../database --output ../../database/mocks
//go:generate mockery --name Users --filename mock_user.go --dir ../../database --output ../../database/mocks
//go:generate mockery --name Tasks --filename mock_task.go --dir ../../database --output ../../database/mocks
