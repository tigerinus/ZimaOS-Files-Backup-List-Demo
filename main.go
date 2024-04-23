//go:generate bash -c "mkdir -p codegen/filesbackup && go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0 -generate types,client -package filesbackup https://raw.githubusercontent.com/IceWhaleTech/IceWhale-OpenAPI/main/icewhale-files-backup/openapi.yaml > codegen/filesbackup/api.go"

package main

import (
	"fmt"
	"log"

	"github.com/IceWhaleTech/ZimaOS-Common/constants"
	commonFilesbackup "github.com/IceWhaleTech/ZimaOS-Common/filesbackup"
	"github.com/tigerinus/ZimaOS-Files-Backup-List-Demo/codegen/filesbackup"
)

func main() {
	metadataPath := commonFilesbackup.DefaultMetadataPath(constants.DefaultDataPath)
	allBackups, err := commonFilesbackup.GetAllBackups[filesbackup.FolderBackup](metadataPath)
	if err != nil {
		log.Fatalf("Error getting all backups: %v", err)
	}

	for _, backups := range allBackups {
		for _, backup := range backups {
			fmt.Println(backup.BackupFolderPath)
		}
	}
}
