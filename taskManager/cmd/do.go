package cmd

import (
	"cor_gophercises/taskManager/src"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do tasks",
	Long:  "Mark Task as done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errorExit("Only 1 argument should be passed to CMD")
		}

		taskId, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("Value '%v' passed to 'do' cmd does not appear to be an int.\n", args[0])
			fmt.Println(err)
			errorExit("Unable to remove value provided, exiting.")
		}

		allTasks := src.GetAll()
		activeTasks := filterActive(allTasks)

		if len(activeTasks) < taskId || taskId <= 0 {
			errorExit("Task with that number does not exist.")
		}
		fmt.Printf("Removing Task: %v \n", taskId)

		toRemove := activeTasks[taskId-1]

		toRemove.Data.DoneAt = int(time.Now().Unix())

		result := src.Update(toRemove)

		if result != nil {
			fmt.Println(err)
			errorExit("Unable to remove value provided, exiting.")
		}

		fmt.Printf("Task '%v' was removed from the list.\n", taskId)
	},
}

func filterActive(tasks []src.ListItemEntry) []src.ListItemEntry {

	filtered := make([]src.ListItemEntry, 0)
	for _, v := range tasks {
		if v.Data.DoneAt == 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
