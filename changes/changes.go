package changes

import (
	"os"
	"path/filepath"

	"github.com/docker/docker/builder/dockerignore"
	"github.com/docker/docker/pkg/fileutils"
	"github.com/rjeczalik/notify"
)

type Change struct {
	Operation string
	Base      string
	Path      string
}

func Partition(changes []Change) (adds []Change, removes []Change) {
	for _, c := range changes {
		switch c.Operation {
		case "add":
			adds = append(adds, c)
		case "remove":
			removes = append(removes, c)
		}
	}

	return
}

func Watch(dir string, ch chan Change) error {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	ignore, err := readDockerIgnoreRecursive(abs)
	if err != nil {
		return err
	}

	sym, err := filepath.EvalSymlinks(abs)
	if err != nil {
		return err
	}

	return watchForChanges(sym, ignore, ch)
}

func readDockerIgnoreRecursive(root string) ([]string, error) {
	ignore := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info != nil && info.Name() == ".dockerignore" {
			lines, err := readDockerIgnore(path)
			if err != nil {
				return err
			}

			// get the relative base between the root of the docker context and this dockerignore
			rel, err := filepath.Rel(root, filepath.Dir(path))
			if err != nil {
				return err
			}

			for _, line := range lines {
				// append the dockerignore lines including the relative base
				ignore = append(ignore, filepath.Join(rel, line))
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ignore, nil
}

func readDockerIgnore(file string) ([]string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	ignore, err := dockerignore.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	return ignore, nil
}

func watchForChanges(dir string, ignore []string, ch chan Change) error {
	c := make(chan notify.EventInfo, 1)

	if err := notify.Watch(filepath.Join(dir, "..."), c, notify.Write, notify.Remove); err != nil {
		return err
	}

	for event := range c {
		rel, err := filepath.Rel(dir, event.Path())
		if err != nil {
			return err
		}

		match, err := fileutils.Matches(rel, ignore)
		if err != nil {
			return err
		}

		if match {
			continue
		}

		switch event.Event() {
		case notify.Remove:
			ch <- Change{
				Operation: "remove",
				Base:      dir,
				Path:      rel,
			}
		case notify.Write:
			ch <- Change{
				Operation: "add",
				Base:      dir,
				Path:      rel,
			}
		}
	}

	return nil
}
