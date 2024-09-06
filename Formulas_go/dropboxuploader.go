package main

// DropboxUploader - Bash script for interacting with Dropbox
// Homepage: https://www.andreafabrizi.it/2016/01/01/Dropbox-Uploader/

import (
	"fmt"
	
	"os/exec"
)

func installDropboxUploader() {
	// Método 1: Descargar y extraer .tar.gz
	dropboxuploader_tar_url := "https://github.com/andreafabrizi/Dropbox-Uploader/archive/refs/tags/1.0.tar.gz"
	dropboxuploader_cmd_tar := exec.Command("curl", "-L", dropboxuploader_tar_url, "-o", "package.tar.gz")
	err := dropboxuploader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dropboxuploader_zip_url := "https://github.com/andreafabrizi/Dropbox-Uploader/archive/refs/tags/1.0.zip"
	dropboxuploader_cmd_zip := exec.Command("curl", "-L", dropboxuploader_zip_url, "-o", "package.zip")
	err = dropboxuploader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dropboxuploader_bin_url := "https://github.com/andreafabrizi/Dropbox-Uploader/archive/refs/tags/1.0.bin"
	dropboxuploader_cmd_bin := exec.Command("curl", "-L", dropboxuploader_bin_url, "-o", "binary.bin")
	err = dropboxuploader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dropboxuploader_src_url := "https://github.com/andreafabrizi/Dropbox-Uploader/archive/refs/tags/1.0.src.tar.gz"
	dropboxuploader_cmd_src := exec.Command("curl", "-L", dropboxuploader_src_url, "-o", "source.tar.gz")
	err = dropboxuploader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dropboxuploader_cmd_direct := exec.Command("./binary")
	err = dropboxuploader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
