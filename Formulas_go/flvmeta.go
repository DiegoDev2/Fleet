package main

// Flvmeta - Manipulate Adobe flash video files (FLV)
// Homepage: https://flvmeta.com/

import (
	"fmt"
	
	"os/exec"
)

func installFlvmeta() {
	// Método 1: Descargar y extraer .tar.gz
	flvmeta_tar_url := "https://github.com/noirotm/flvmeta/archive/refs/tags/v1.2.2.tar.gz"
	flvmeta_cmd_tar := exec.Command("curl", "-L", flvmeta_tar_url, "-o", "package.tar.gz")
	err := flvmeta_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flvmeta_zip_url := "https://github.com/noirotm/flvmeta/archive/refs/tags/v1.2.2.zip"
	flvmeta_cmd_zip := exec.Command("curl", "-L", flvmeta_zip_url, "-o", "package.zip")
	err = flvmeta_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flvmeta_bin_url := "https://github.com/noirotm/flvmeta/archive/refs/tags/v1.2.2.bin"
	flvmeta_cmd_bin := exec.Command("curl", "-L", flvmeta_bin_url, "-o", "binary.bin")
	err = flvmeta_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flvmeta_src_url := "https://github.com/noirotm/flvmeta/archive/refs/tags/v1.2.2.src.tar.gz"
	flvmeta_cmd_src := exec.Command("curl", "-L", flvmeta_src_url, "-o", "source.tar.gz")
	err = flvmeta_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flvmeta_cmd_direct := exec.Command("./binary")
	err = flvmeta_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
