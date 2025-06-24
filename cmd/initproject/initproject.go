package initproject

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/flazhgrowth/fg-tamagochi/pkg/db/entity"
	projecttemplates "github.com/flazhgrowth/fg-tamagochi/project_templates"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	commands := &cobra.Command{
		Use:   "init",
		Short: "initialize app structure",
		Run:   initAppStructures,
	}
	commands.Flags().String("packagename", "", "(go.mod) package name")

	return commands
}

func checkToolsAvailibility() {
	toolsCmds := map[string]string{
		"swag":    "github.com/swaggo/swag/cmd/swag@latest",
		"migrate": "installFrom;https://github.com/golang-migrate/migrate/tree/master/cmd/migrate",
		"wire":    "installFrom;https://github.com/google/wire",
	}
	for key, toolCmd := range toolsCmds {
		execCmd := exec.Command(key)
		if key == "migrate" {
			execCmd = exec.Command(key, "-version")
		}
		_, err := execCmd.CombinedOutput()
		if err != nil {
			log.Error().Msgf("error on checking %s availibility: %s", key, err.Error())
			if strings.HasPrefix(toolCmd, "installFrom;") {
				raws := strings.Split(toolCmd, ";")
				fmt.Printf("get %s from %s\n", key, raws[1])
				continue
			}
			fmt.Printf("attempting on installing %s from %s\n", key, toolCmd)
			execCmd = exec.Command("go", "install", toolCmd)
			if _, err := execCmd.CombinedOutput(); err != nil {
				panic(err)
			}
		}
		fmt.Println(key, " available")
	}
}

func initAppStructures(cmd *cobra.Command, args []string) {
	fmt.Println("checking available tools (swaggo, go migrate, wire)")
	checkToolsAvailibility()

	packagename, err := cmd.Flags().GetString("packagename")
	if err != nil {
		panic(err)
	}

	defaultEntities := map[string]entity.ProjectSchema{
		"account": {
			Implementations: entity.ProjectImplementationsSchema{
				Transport: map[string]entity.DirectoryElementsSchema{
					"acccountapi": []string{"account.go", "type.go"},
				},
				Usecase: map[string]entity.DirectoryElementsSchema{
					"accountuc": {"account.go", "type.go"},
				},
				Repository: map[string]entity.DirectoryElementsSchema{
					"accountrepo": {"account.go", "type.go"},
				},
			},
			Entities: map[string]entity.DirectoryElementsSchema{
				"account": {
					"account.go", "api.go", "usecase.go", "repository.go", "type.go", "database.go",
				},
			},
			Database: entity.DirectoryElementsSchema{
				"migrations", "seeder",
			},
		},
	}

	internalDir := entity.DirName("./internal")
	if err := mkDir(internalDir.Val()); err != nil { // make dir implementations
		log.Error().Msgf("error on creating ./internal dir %s", err.Error())
		return
	}
	// copyFile("./project_templates/main.go.templ", "./main.go.templ")
	if err := projecttemplates.MaingoTemplate.WriteTo("./main.go", nil); err != nil {
		log.Error().Msgf("error on creating main.go from template: %s", err.Error())
		return
	}

	for ent, selectedEntity := range defaultEntities {
		// init parent dir
		implementationsDirName := internalDir.EndWith("/implementations") // dir ./internal/implementations name
		entitiesDirName := internalDir.EndWith("/entity")                 // dir ./internal/entity name
		databaseDirName := internalDir.EndWith("/database")               // dir ./internal/database name
		if err := mkDir(implementationsDirName.Val()); err != nil {       // make dir ./internal/implementations
			log.Error().Msgf("error on creating implementations dir %s", err.Error())
			return
		}
		if err := mkDir(entitiesDirName.Val()); err != nil { // make dir ./internal/entities
			log.Error().Msgf("error on creating entities dir %s", err.Error())
			return
		}
		if err := mkDir(databaseDirName.Val()); err != nil { // make dir ./internal/database
			log.Error().Msgf("error on creating database dir %s", err.Error())
			return
		}

		// transport dir and its elements
		for dirName, dirElems := range selectedEntity.Implementations.Transport {
			currentDir := implementationsDirName.EndWith("/transport") // ./internal/implementations/transport
			if err := mkDir(currentDir.Val()); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", currentDir.Val(), err.Error())
				return
			}
			currentDir = currentDir.EndWith(dirName) // ./internal/implementations/transport/accountapi
			if err := mkDir(currentDir.Val()); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", currentDir.Val(), err.Error())
				return
			}

			for _, dirElem := range dirElems {
				filepath := currentDir.EndWith(dirElem).Val()
				selectedTemplate := projecttemplates.TransportImplEmptyTemplate
				if strings.Contains(dirElem, ent) {
					selectedTemplate = projecttemplates.TransportImplTemplate
				}
				if err := selectedTemplate.WriteTo(filepath, map[string]any{
					"entity":       ent,
					"entity_title": strings.Title(ent),
					"packagename":  packagename,
				}); err != nil {
					log.Error().Msgf("error on creating %s file: %s", filepath, err.Error())
					return
				}
			}
		}

		// usecase dir and its elements
		for dirName, dirElems := range selectedEntity.Implementations.Usecase {
			currentDir := implementationsDirName.EndWith("/usecase") // ./internal/implementations/usecase
			if err := mkDir(currentDir.Val()); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", currentDir.Val(), err.Error())
				return
			}
			currentDir = currentDir.EndWith(dirName) // ./internal/implementations/usecase/accountuc
			if err := mkDir(currentDir.Val()); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", currentDir.Val(), err.Error())
				return
			}

			for _, dirElem := range dirElems {
				filepath := currentDir.EndWith(dirElem).Val()
				selectedTemplate := projecttemplates.UsecaseImplEmptyTemplate
				if strings.Contains(dirElem, ent) {
					selectedTemplate = projecttemplates.UsecaseImplTemplate
				}
				if err := selectedTemplate.WriteTo(filepath, map[string]any{
					"entity":       ent,
					"entity_title": strings.Title(ent),
					"packagename":  packagename,
				}); err != nil {
					log.Error().Msgf("error on creating %s file: %s", filepath, err.Error())
					return
				}
			}
		}

		// repository dir and its elements
		for dirName, dirElems := range selectedEntity.Implementations.Repository {
			currentDir := implementationsDirName.EndWith("/repository") // ./internal/implementations/repository
			if err := mkDir(currentDir.Val()); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", currentDir.Val(), err.Error())
				return
			}
			if err := mkDir(currentDir.EndWith("db").Val()); err != nil { // ./internal/implementations/repository/db
				log.Error().Msgf("error on creating %s dir: %s", currentDir.EndWith("db").Val(), err.Error())
				return
			}
			currentDir = currentDir.EndWith("db").EndWith(dirName) // ./internal/implementations/repository/db/accountrepo
			if err := mkDir(currentDir.Val()); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", currentDir.Val(), err.Error())
				return
			}

			for _, dirElem := range dirElems {
				filepath := currentDir.EndWith(dirElem).Val()
				selectedTemplate := projecttemplates.DBRepositoryImplEmptyTemplate
				if strings.Contains(dirElem, ent) {
					selectedTemplate = projecttemplates.DBRepositoryImplTemplate
				}
				if err := selectedTemplate.WriteTo(filepath, map[string]any{
					"entity":       ent,
					"entity_title": strings.Title(ent),
					"packagename":  packagename,
				}); err != nil {
					log.Error().Msgf("error on creating %s file: %s", filepath, err.Error())
					return
				}
			}
		}

		// entities dir and its elements
		for entityDirName, entitiesDirElems := range selectedEntity.Entities {
			currentDir := entitiesDirName.EndWith(entityDirName) // ./internal/entity/account
			if err := mkDir(currentDir.Val()); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", currentDir.Val(), err.Error())
				return
			}

			for _, dirElem := range entitiesDirElems {
				filepath := currentDir.EndWith(dirElem).Val()
				selectedEntity := projecttemplates.EntityTemplate
				if strings.Contains(dirElem, "api") {
					selectedEntity = projecttemplates.EntityAPIInterfaceTemplate
				} else if strings.Contains(dirElem, "usecase") {
					selectedEntity = projecttemplates.EntityUsecaseInterfaceTemplate
				} else if strings.Contains(dirElem, "repository") {
					selectedEntity = projecttemplates.EntityRepositoryInterfaceTemplate
				}
				if err := selectedEntity.WriteTo(filepath, map[string]any{
					"entity":       ent,
					"entity_title": strings.Title(ent),
					"packagename":  packagename,
				}); err != nil {
					log.Error().Msgf("error on creating %s file: %s", filepath, err.Error())
					return
				}
			}
		}

		for _, dirName := range selectedEntity.Database {
			filepath := databaseDirName.EndWith(dirName).Val()
			if err := mkDir(filepath); err != nil {
				log.Error().Msgf("error on creating %s dir: %s", filepath, err.Error())
				return
			}
		}
	}

	// make .gitignore
	if err := projecttemplates.GitignoreTemplate.WriteTo("./.gitignore", nil); err != nil {
		log.Error().Msgf("error on creating .gitignore from template: %s", err.Error())
		return
	}
}

func copyFile(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	// Optional: Copy file permissions
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %w", err)
	}
	err = os.Chmod(dstPath, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	return nil
}

func mkDir(path string) error {
	cmd := exec.Command("mkdir", path)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
