package main

// Checkbashisms - Checks for bashisms in shell scripts
// Homepage: https://launchpad.net/ubuntu/+source/devscripts/

import (
	"fmt"
	
	"os/exec"
)

func installCheckbashisms() {
	// Método 1: Descargar y extraer .tar.gz
	checkbashisms_tar_url := "https://deb.debian.org/debian/pool/main/d/devscripts/devscripts_2.23.7.tar.xz"
	checkbashisms_cmd_tar := exec.Command("curl", "-L", checkbashisms_tar_url, "-o", "package.tar.gz")
	err := checkbashisms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	checkbashisms_zip_url := "https://deb.debian.org/debian/pool/main/d/devscripts/devscripts_2.23.7.tar.xz"
	checkbashisms_cmd_zip := exec.Command("curl", "-L", checkbashisms_zip_url, "-o", "package.zip")
	err = checkbashisms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	checkbashisms_bin_url := "https://deb.debian.org/debian/pool/main/d/devscripts/devscripts_2.23.7.tar.xz"
	checkbashisms_cmd_bin := exec.Command("curl", "-L", checkbashisms_bin_url, "-o", "binary.bin")
	err = checkbashisms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	checkbashisms_src_url := "https://deb.debian.org/debian/pool/main/d/devscripts/devscripts_2.23.7.tar.xz"
	checkbashisms_cmd_src := exec.Command("curl", "-L", checkbashisms_src_url, "-o", "source.tar.gz")
	err = checkbashisms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	checkbashisms_cmd_direct := exec.Command("./binary")
	err = checkbashisms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
