package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sletkov/backend/wb-task-0/internal/app/models"
	"sletkov/backend/wb-task-0/internal/app/store/postgres"

	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
)

func Start(config *Config) error {
	fmt.Println("Start server ...")

	//Connect to nats-streaming server
	sc, err := stan.Connect(config.StanClusterId, config.StanClientId, stan.NatsURL(config.StanUrl))

	if err != nil {
		fmt.Println("Can't connect to stan: ", err)
	}

	//Connect to db
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		fmt.Println("Can't connect to postgres: ", err)
	}

	defer db.Close()

	//Ping db
	if err := db.Ping(); err != nil {
		fmt.Println("Can't ping Db: ", err)
	}

	store := postgres.New(db)
	srv := newServer(store)

	//Subscribe to a channel
	sub, err := sc.Subscribe("orders",
		func(msg *stan.Msg) {
			fmt.Println("Received msg: ", string(msg.Data))

			//Create order and write it into db
			var order models.Order

			err := json.Unmarshal(msg.Data, &order)

			if err != nil {
				fmt.Println("Cant unmarshal order: ", err)
			}

			if err := order.Validate(); err != nil {
				fmt.Println(err)
				return
			}

			err = store.Order().Create(&order)

			if err != nil {
				fmt.Println("Cant create order: ", err)
			}

			//Add order in cache
			srv.cache.Add(order.Id, order, cache.NoExpiration)
		},
		stan.DeliverAllAvailable())

	if err != nil {
		fmt.Println("Can't subscribe to nats-streaming channel: ", err)
	}

	defer sub.Close()

	//Recovering cache from db in case of server's fall out
	defer func() {
		if r := recover(); r != nil {
			orders, err := store.Order().GetAll()

			if err != nil {
				fmt.Println("Can't recover cache from db: ", err)
			}

			for _, order := range orders {
				srv.cache.Add(order.Id, order, cache.NoExpiration)
			}
		}
	}()

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDb(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
