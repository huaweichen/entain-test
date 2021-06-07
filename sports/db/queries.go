package db

const (
	eventList = "list"
)

func getSportEventsQueries() map[string]string {
	return map[string]string{
		eventList: `
			SELECT 
				id, 
				name, 
				advertised_start_time
			FROM sports
		`,
	}
}
