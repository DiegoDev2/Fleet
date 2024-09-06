package main

// Qbs - Build tool for developing projects across multiple platforms
// Homepage: https://wiki.qt.io/Qbs

import (
	"fmt"
	
	"os/exec"
)

func installQbs() {
	// Método 1: Descargar y extraer .tar.gz
	qbs_tar_url := "https://download.qt.io/official_releases/qbs/2.4.1/qbs-src-2.4.1.tar.gz"
	qbs_cmd_tar := exec.Command("curl", "-L", qbs_tar_url, "-o", "package.tar.gz")
	err := qbs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qbs_zip_url := "https://download.qt.io/official_releases/qbs/2.4.1/qbs-src-2.4.1.zip"
	qbs_cmd_zip := exec.Command("curl", "-L", qbs_zip_url, "-o", "package.zip")
	err = qbs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qbs_bin_url := "https://download.qt.io/official_releases/qbs/2.4.1/qbs-src-2.4.1.bin"
	qbs_cmd_bin := exec.Command("curl", "-L", qbs_bin_url, "-o", "binary.bin")
	err = qbs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qbs_src_url := "https://download.qt.io/official_releases/qbs/2.4.1/qbs-src-2.4.1.src.tar.gz"
	qbs_cmd_src := exec.Command("curl", "-L", qbs_src_url, "-o", "source.tar.gz")
	err = qbs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qbs_cmd_direct := exec.Command("./binary")
	err = qbs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}
