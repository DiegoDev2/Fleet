package main

// PinboardNotesBackup - Efficiently back up the notes you've saved to Pinboard
// Homepage: https://github.com/bdesham/pinboard-notes-backup

import (
	"fmt"
	
	"os/exec"
)

func installPinboardNotesBackup() {
	// Método 1: Descargar y extraer .tar.gz
	pinboardnotesbackup_tar_url := "https://github.com/bdesham/pinboard-notes-backup/archive/refs/tags/v1.0.5.7.tar.gz"
	pinboardnotesbackup_cmd_tar := exec.Command("curl", "-L", pinboardnotesbackup_tar_url, "-o", "package.tar.gz")
	err := pinboardnotesbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pinboardnotesbackup_zip_url := "https://github.com/bdesham/pinboard-notes-backup/archive/refs/tags/v1.0.5.7.zip"
	pinboardnotesbackup_cmd_zip := exec.Command("curl", "-L", pinboardnotesbackup_zip_url, "-o", "package.zip")
	err = pinboardnotesbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pinboardnotesbackup_bin_url := "https://github.com/bdesham/pinboard-notes-backup/archive/refs/tags/v1.0.5.7.bin"
	pinboardnotesbackup_cmd_bin := exec.Command("curl", "-L", pinboardnotesbackup_bin_url, "-o", "binary.bin")
	err = pinboardnotesbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pinboardnotesbackup_src_url := "https://github.com/bdesham/pinboard-notes-backup/archive/refs/tags/v1.0.5.7.src.tar.gz"
	pinboardnotesbackup_cmd_src := exec.Command("curl", "-L", pinboardnotesbackup_src_url, "-o", "source.tar.gz")
	err = pinboardnotesbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pinboardnotesbackup_cmd_direct := exec.Command("./binary")
	err = pinboardnotesbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc")
	exec.Command("latte", "install", "ghc").Run()
}
