package main

import (
    "fmt"
    "log"

    "github.com/gocql/gocql"
)

func main() {
    // connect to the cluster
    cluster := gocql.NewCluster("192.168.1.1", "192.168.1.2", "192.168.1.250")
    cluster.Keyspace = "system_schema"
    cluster.Consistency = gocql.Quorum
    session, _ := cluster.CreateSession()
    defer session.Close()

    // insert a tweet
    if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
        "me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
        log.Fatal(err)
    }

    var id gocql.UUID
    var text string

    /* Search for a specific set of records whose 'timeline' column matches
     * the value 'me'. The secondary index that we created earlier will be
     * used for optimizing the search */
    if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
        "me").Consistency(gocql.One).Scan(&id, &text); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Tweet:", id, text)

    // list all tweets
    iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
    for iter.Scan(&id, &text) {
        fmt.Println("Tweet:", id, text)
    }
    if err := iter.Close(); err != nil {
        log.Fatal(err)
    }
}