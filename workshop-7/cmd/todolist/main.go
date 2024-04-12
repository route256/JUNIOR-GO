package main

import (
	"database/sql"
	"log"
	"time"

	"to-do-list/internal/controller"
	"to-do-list/internal/datasource/cache"
	"to-do-list/internal/datasource/database"

	"github.com/allegro/bigcache/v3"
	"github.com/spf13/cobra"
)

type Controller interface {
	Create(name, description, deadline string) error
	Update(id int, name, description, deadline string, status int) error
	List() error
}

var rootCmd cobra.Command

var taskID int
var taskName string
var taskDescription string
var taskDeadline string
var taskStatus int

func main() {
	var ctrl Controller
	var db *sql.DB

	rootCmd = cobra.Command{
		Use:     "root",
		Version: "v1.0",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			db, err = sql.Open("sqlite3", "db/todo.db")
			if err != nil {
				return err
			}

			bcache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
			if err != nil {
				return err
			}

			ctrl = controller.NewController(
				cache.NewClient(
					bcache,
					database.NewClient(db),
				),
			)

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("start")
			return nil
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return db.Close()
		},
	}
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "create",
			Short: "",
			Long:  "",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("create")

				name, err := cmd.Flags().GetString("name")
				if err != nil {
					return err
				}

				description, err := cmd.Flags().GetString("description")
				if err != nil {
					return err
				}

				deadline, err := cmd.Flags().GetString("deadline")
				if err != nil {
					return err
				}

				return ctrl.Create(name, description, deadline)
			},
		},
	)
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "update",
			Short: "",
			Long:  "",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("update")

				id, err := cmd.Flags().GetInt("id")
				if err != nil {
					return err
				}

				name, err := cmd.Flags().GetString("name")
				if err != nil {
					return err
				}

				description, err := cmd.Flags().GetString("description")
				if err != nil {
					return err
				}

				deadline, err := cmd.Flags().GetString("deadline")
				if err != nil {
					return err
				}

				status, err := cmd.Flags().GetInt("status")
				if err != nil {
					return err
				}

				return ctrl.Update(id, name, description, deadline, status)
			},
		},
	)
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "list",
			Short: "",
			Long:  "",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("list")

				return ctrl.List()
			},
		},
	)

	// Optional
	rootCmd.PersistentFlags().IntVar(
		&taskID,
		"id",
		0,
		"Идентификатор задачи",
	)
	rootCmd.PersistentFlags().StringVarP(
		&taskName,
		"name",
		"n",
		"",
		"Название задачи",
	)
	rootCmd.PersistentFlags().StringVarP(
		&taskDescription,
		"description",
		"d",
		"",
		"Описание задачи",
	)
	rootCmd.PersistentFlags().StringVar(
		&taskDeadline,
		"deadline",
		"",
		"Время окончания задачи. Формат: 2006-01-02",
	)
	rootCmd.PersistentFlags().IntVarP(
		&taskStatus,
		"status",
		"s",
		0,
		"Статус задачи",
	)

	err := rootCmd.Execute()
	if err != nil {
		// Required arguments are missing, etc
		log.Fatalln("error", err)
	}

	log.Println(taskName, taskDescription, taskDeadline)
}
