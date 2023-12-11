package processing

import (
	"log"

	"github.com/ilolss/outbox-pattern/internal/task"
)

// берет задачу из kafka и выводит данные в консоль
func Processing(task *task.Task) {
	log.Println("name: ", task.Name, "description: ", task.Description)
}
