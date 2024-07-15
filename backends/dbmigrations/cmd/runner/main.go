// This is copied from https://github.com/pressly/goose/blob/bfd4286c0fda61ce69e54a272fdf90e72b301aa5/examples/go-migrations/main.go

package main

import (
	"context"
	"flag"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"

	_ "dbmigrations"
)

var (
	flags   = flag.NewFlagSet("goose", flag.ExitOnError)
	dir     = flags.String("dir", ".", "directory with migration files")
	dialect = flags.String("dialect", "mysql", "dialect for the migration")
)

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("goose: failed to parse flags: %v\n", err)
	}
	args := flags.Args()

	if len(args) < 3 {
		flags.Usage()
		return
	}

	if dialect == nil {
		log.Fatal("goose: missing required flag -dialect")
	}
	goose.SetBaseFS(nil)
	if err := goose.SetDialect(*dialect); err != nil {
		log.Fatalf("goose: failed to set dialect: %v\n", err)
	}

	dbstring, command := args[0], args[1]

	db, err := goose.OpenDBWithDriver("mysql", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
