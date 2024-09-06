package main

// Direvent - Monitors events in the file system directories
// Homepage: https://www.gnu.org.ua/software/direvent/direvent.html

import (
	"fmt"
	
	"os/exec"
)

func installDirevent() {
	// Método 1: Descargar y extraer .tar.gz
	direvent_tar_url := "https://ftp.gnu.org/gnu/direvent/direvent-5.4.tar.gz"
	direvent_cmd_tar := exec.Command("curl", "-L", direvent_tar_url, "-o", "package.tar.gz")
	err := direvent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	direvent_zip_url := "https://ftp.gnu.org/gnu/direvent/direvent-5.4.zip"
	direvent_cmd_zip := exec.Command("curl", "-L", direvent_zip_url, "-o", "package.zip")
	err = direvent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	direvent_bin_url := "https://ftp.gnu.org/gnu/direvent/direvent-5.4.bin"
	direvent_cmd_bin := exec.Command("curl", "-L", direvent_bin_url, "-o", "binary.bin")
	err = direvent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	direvent_src_url := "https://ftp.gnu.org/gnu/direvent/direvent-5.4.src.tar.gz"
	direvent_cmd_src := exec.Command("curl", "-L", direvent_src_url, "-o", "source.tar.gz")
	err = direvent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	direvent_cmd_direct := exec.Command("./binary")
	err = direvent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
