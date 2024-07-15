// This is copied from https://github.com/pressly/goose/blob/bfd4286c0fda61ce69e54a272fdf90e72b301aa5/examples/go-migrations/main.go

package main

import (
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
	flags.Parse(os.Args[1:])
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

	dbstring, command := args[1], args[2]

	db, err := goose.OpenDBWithDriver("sqlite", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
