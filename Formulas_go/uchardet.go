package main

// Uchardet - Encoding detector library
// Homepage: https://www.freedesktop.org/wiki/Software/uchardet/

import (
	"fmt"
	
	"os/exec"
)

func installUchardet() {
	// Método 1: Descargar y extraer .tar.gz
	uchardet_tar_url := "https://www.freedesktop.org/software/uchardet/releases/uchardet-0.0.8.tar.xz"
	uchardet_cmd_tar := exec.Command("curl", "-L", uchardet_tar_url, "-o", "package.tar.gz")
	err := uchardet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uchardet_zip_url := "https://www.freedesktop.org/software/uchardet/releases/uchardet-0.0.8.tar.xz"
	uchardet_cmd_zip := exec.Command("curl", "-L", uchardet_zip_url, "-o", "package.zip")
	err = uchardet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uchardet_bin_url := "https://www.freedesktop.org/software/uchardet/releases/uchardet-0.0.8.tar.xz"
	uchardet_cmd_bin := exec.Command("curl", "-L", uchardet_bin_url, "-o", "binary.bin")
	err = uchardet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uchardet_src_url := "https://www.freedesktop.org/software/uchardet/releases/uchardet-0.0.8.tar.xz"
	uchardet_cmd_src := exec.Command("curl", "-L", uchardet_src_url, "-o", "source.tar.gz")
	err = uchardet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uchardet_cmd_direct := exec.Command("./binary")
	err = uchardet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
