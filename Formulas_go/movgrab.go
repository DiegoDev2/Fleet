package main

// Movgrab - Downloader for youtube, dailymotion, and other video websites
// Homepage: https://github.com/ColumPaget/Movgrab

import (
	"fmt"
	
	"os/exec"
)

func installMovgrab() {
	// Método 1: Descargar y extraer .tar.gz
	movgrab_tar_url := "https://github.com/ColumPaget/Movgrab/archive/refs/tags/3.1.2.tar.gz"
	movgrab_cmd_tar := exec.Command("curl", "-L", movgrab_tar_url, "-o", "package.tar.gz")
	err := movgrab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	movgrab_zip_url := "https://github.com/ColumPaget/Movgrab/archive/refs/tags/3.1.2.zip"
	movgrab_cmd_zip := exec.Command("curl", "-L", movgrab_zip_url, "-o", "package.zip")
	err = movgrab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	movgrab_bin_url := "https://github.com/ColumPaget/Movgrab/archive/refs/tags/3.1.2.bin"
	movgrab_cmd_bin := exec.Command("curl", "-L", movgrab_bin_url, "-o", "binary.bin")
	err = movgrab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	movgrab_src_url := "https://github.com/ColumPaget/Movgrab/archive/refs/tags/3.1.2.src.tar.gz"
	movgrab_cmd_src := exec.Command("curl", "-L", movgrab_src_url, "-o", "source.tar.gz")
	err = movgrab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	movgrab_cmd_direct := exec.Command("./binary")
	err = movgrab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libressl")
exec.Command("latte", "install", "libressl")
}
