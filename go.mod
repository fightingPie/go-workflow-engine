module github.com/go-workflow/go-workflow

go 1.15

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/mumushuiding/util v0.0.0-20210203080010-04699a081184
	github.com/onsi/ginkgo v1.15.2 // indirect
	github.com/onsi/gomega v1.11.0 // indirect
	github.com/robfig/cron v1.2.0
)

replace (
	github.com/go-workflow/go-workflow/workflow-config => ./workflow-config
	github.com/go-workflow/go-workflow/workflow-engine/model => ./workflow-engine/model
	github.com/go-workflow/go-workflow/workflow-engine/service => ./workflow-engine/service
	github.com/go-workflow/go-workflow/workflow-router => ./workflow-router
)
