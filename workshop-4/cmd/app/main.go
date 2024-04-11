package main

import (
	"context"
	"encoding/json"
	"fmt"
	core "test_jr_6/internal/app/core"
	"test_jr_6/internal/app/rating"
	"test_jr_6/internal/app/task"
	"test_jr_6/internal/app/task_group"
)

type Connector struct {
}

func main() {
	ctx := context.Background()

	connector := Connector{}

	calcVariants := map[rating.Type]rating.CalcVariant{
		rating.SimpleType:   rating.SimpleVariant{},
		rating.MultiplyType: rating.MultiplyVariant{},
		rating.AdditionType: rating.AdditionVariant{},

	}

	coreService := core.NewCoreService(
		task.NewService(task.NewRepository(connector)),
		task_group.NewService(task_group.NewRepository(connector)),
		rating.NewService(rating.NewRepository(connector), calcVariants),
	)

	fmt.Println("\nСоздания группы задач")
	taskGroup, err := coreService.CreateTaskGroup(ctx, core.CreateTasksGroupRequest{
		Name:            "Группа задач 1",
		Description:     "Описание группы задач 1",
		Price:           600,
		SecondsToDecide: 900,
	})
	fmt.Println(taskGroup, err)

	fmt.Println("\nРедактирования группы задач")
	err = coreService.UpdateTaskGroup(ctx, core.UpdateTaskGroupRequest{
		ID:              4,
		Name:            "Группа задач 4",
		Description:     "Описание группы задач 4",
		Price:           100,
		SecondsToDecide: 1000,
	})
	fmt.Println(err)

	fmt.Println("\nПолучения группы задач")
	taskGroupForCustomer, _ := coreService.GetTaskGroupForCustomer(ctx, 24)
	fmt.Println(taskGroupForCustomer)

	fmt.Println("\nСоздания задач из json")
	tasks := []task.Task{
		{
			Name:        "Задание 1",
			Description: "Описание задания 1",
			TaskGroupID: 1,
			CustomerID:  2,
			Overlap:     5,
		},
		{
			Name:        "Задание 2",
			Description: "Описание задания 2",
			TaskGroupID: 1,
			CustomerID:  2,
			Overlap:     5,
		},
	}
	taskJSON, _ := json.Marshal(tasks)

	taskIDs, _ := coreService.CreateTasksFromJSON(ctx, string(taskJSON), 1)
	fmt.Println(taskIDs)

	fmt.Println("\nРедактирования задач")
	err = coreService.UpdateTasks(ctx, core.UpdateTasksRequest{
		Tasks: []core.UpdateTaskRequest{
			{
				ID:          1,
				Name:        "Задание 1",
				Description: "Описание задания 1",
				Overlap:     5,
			},
			{
				ID:          2,
				Name:        "Задание 2",
				Description: "Описание задания 2",
				Overlap:     3,
			},
		},
		TaskGroupID: 5,
		CustomerID:  123,
	})
	fmt.Println(err)

	fmt.Println("\nПолучение задач на просмотр заказчиком")
	tasksForCustomer, _ := coreService.GetTasksForCustomer(ctx, 24)
	fmt.Println(tasksForCustomer)

	fmt.Println("\nПолучение задачи на решение")
	taskForSolve, _ := coreService.GetTaskForSolve(ctx, 24)
	fmt.Println(taskForSolve)

	fmt.Println("\nРешение задачи и пересчет рейтинга")
	_ = coreService.SolveTask(ctx, 2, 23, "ответ 1")
}
