package main

// Archivemount - File system for accessing archives using libarchive
// Homepage: https://github.com/cybernoid/archivemount

import (
	"fmt"
	
	"os/exec"
)

func installArchivemount() {
	// Método 1: Descargar y extraer .tar.gz
	archivemount_tar_url := "https://slackware.uk/~urchlay/src/archivemount-0.9.1.tar.gz"
	archivemount_cmd_tar := exec.Command("curl", "-L", archivemount_tar_url, "-o", "package.tar.gz")
	err := archivemount_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	archivemount_zip_url := "https://slackware.uk/~urchlay/src/archivemount-0.9.1.zip"
	archivemount_cmd_zip := exec.Command("curl", "-L", archivemount_zip_url, "-o", "package.zip")
	err = archivemount_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	archivemount_bin_url := "https://slackware.uk/~urchlay/src/archivemount-0.9.1.bin"
	archivemount_cmd_bin := exec.Command("curl", "-L", archivemount_bin_url, "-o", "binary.bin")
	err = archivemount_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	archivemount_src_url := "https://slackware.uk/~urchlay/src/archivemount-0.9.1.src.tar.gz"
	archivemount_cmd_src := exec.Command("curl", "-L", archivemount_src_url, "-o", "source.tar.gz")
	err = archivemount_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	archivemount_cmd_direct := exec.Command("./binary")
	err = archivemount_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
}
