package main

// Chrpath - Tool to edit the rpath in ELF binaries
// Homepage: https://tracker.debian.org/pkg/chrpath

import (
	"fmt"
	
	"os/exec"
)

func installChrpath() {
	// Método 1: Descargar y extraer .tar.gz
	chrpath_tar_url := "https://deb.debian.org/debian/pool/main/c/chrpath/chrpath_0.16.orig.tar.gz"
	chrpath_cmd_tar := exec.Command("curl", "-L", chrpath_tar_url, "-o", "package.tar.gz")
	err := chrpath_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chrpath_zip_url := "https://deb.debian.org/debian/pool/main/c/chrpath/chrpath_0.16.orig.zip"
	chrpath_cmd_zip := exec.Command("curl", "-L", chrpath_zip_url, "-o", "package.zip")
	err = chrpath_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chrpath_bin_url := "https://deb.debian.org/debian/pool/main/c/chrpath/chrpath_0.16.orig.bin"
	chrpath_cmd_bin := exec.Command("curl", "-L", chrpath_bin_url, "-o", "binary.bin")
	err = chrpath_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chrpath_src_url := "https://deb.debian.org/debian/pool/main/c/chrpath/chrpath_0.16.orig.src.tar.gz"
	chrpath_cmd_src := exec.Command("curl", "-L", chrpath_src_url, "-o", "source.tar.gz")
	err = chrpath_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chrpath_cmd_direct := exec.Command("./binary")
	err = chrpath_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
