package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"to-do-list/internal/controller"

	"to-do-list/internal/storage"

	"github.com/spf13/cobra"
)

type Controller interface {
	Create(name, description, deadline string) error
	List() error
}

var rootCmd cobra.Command

var taskName string
var taskDescription string
var taskDeadline string

func main() {
	var ctrl Controller

	rootCmd = cobra.Command{
		Use:     "root",
		Version: "v1.0",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			const name = "db/storage.json"

			file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
			if errors.Is(err, os.ErrNotExist) {
				if err := os.MkdirAll(filepath.Dir(name), os.ModePerm); err != nil {
					log.Println(err)
				}

				file, err = os.Create(name)
				if err != nil {
					log.Fatalln(err)
				}
			} else if err != nil {
				return err
			}

			ctrl = controller.NewController(storage.NewStorage(file))

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("start")
			return nil
		},
	}
	rootCmd.AddCommand(&cobra.Command{
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
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("list")

			return ctrl.List()
		},
	})

	// Optional
	rootCmd.PersistentFlags().StringVarP(&taskName, "name", "n", "", "Название задачи")
	rootCmd.PersistentFlags().StringVarP(&taskDescription, "description", "d", "", "Описание задачи")
	rootCmd.PersistentFlags().StringVar(&taskDeadline, "deadline", "", "Время окончания задачи. Формат: 2006-01-02")

	err := rootCmd.Execute()
	if err != nil {
		// Required arguments are missing, etc
		log.Fatalln("error", err)
	}

	log.Println(taskName, taskDescription, taskDeadline)
}
