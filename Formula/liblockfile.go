package main

// Liblockfile - Library providing functions to lock standard mailboxes
// Homepage: https://tracker.debian.org/pkg/liblockfile

import (
	"fmt"
	
	"os/exec"
)

func installLiblockfile() {
	// Método 1: Descargar y extraer .tar.gz
	liblockfile_tar_url := "https://deb.debian.org/debian/pool/main/libl/liblockfile/liblockfile_1.17.orig.tar.gz"
	liblockfile_cmd_tar := exec.Command("curl", "-L", liblockfile_tar_url, "-o", "package.tar.gz")
	err := liblockfile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblockfile_zip_url := "https://deb.debian.org/debian/pool/main/libl/liblockfile/liblockfile_1.17.orig.zip"
	liblockfile_cmd_zip := exec.Command("curl", "-L", liblockfile_zip_url, "-o", "package.zip")
	err = liblockfile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblockfile_bin_url := "https://deb.debian.org/debian/pool/main/libl/liblockfile/liblockfile_1.17.orig.bin"
	liblockfile_cmd_bin := exec.Command("curl", "-L", liblockfile_bin_url, "-o", "binary.bin")
	err = liblockfile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblockfile_src_url := "https://deb.debian.org/debian/pool/main/libl/liblockfile/liblockfile_1.17.orig.src.tar.gz"
	liblockfile_cmd_src := exec.Command("curl", "-L", liblockfile_src_url, "-o", "source.tar.gz")
	err = liblockfile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblockfile_cmd_direct := exec.Command("./binary")
	err = liblockfile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
