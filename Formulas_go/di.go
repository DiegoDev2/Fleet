package main

// Di - Advanced df-like disk information utility
// Homepage: https://gentoo.com/di/

import (
	"fmt"
	
	"os/exec"
)

func installDi() {
	// Método 1: Descargar y extraer .tar.gz
	di_tar_url := "https://downloads.sourceforge.net/project/diskinfo-di/di-4.53.tar.gz"
	di_cmd_tar := exec.Command("curl", "-L", di_tar_url, "-o", "package.tar.gz")
	err := di_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	di_zip_url := "https://downloads.sourceforge.net/project/diskinfo-di/di-4.53.zip"
	di_cmd_zip := exec.Command("curl", "-L", di_zip_url, "-o", "package.zip")
	err = di_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	di_bin_url := "https://downloads.sourceforge.net/project/diskinfo-di/di-4.53.bin"
	di_cmd_bin := exec.Command("curl", "-L", di_bin_url, "-o", "binary.bin")
	err = di_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	di_src_url := "https://downloads.sourceforge.net/project/diskinfo-di/di-4.53.src.tar.gz"
	di_cmd_src := exec.Command("curl", "-L", di_src_url, "-o", "source.tar.gz")
	err = di_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	di_cmd_direct := exec.Command("./binary")
	err = di_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
