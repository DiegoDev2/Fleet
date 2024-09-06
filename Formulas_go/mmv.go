package main

// Mmv - Move, copy, append, and link multiple files
// Homepage: https://github.com/rrthomas/mmv

import (
	"fmt"
	
	"os/exec"
)

func installMmv() {
	// Método 1: Descargar y extraer .tar.gz
	mmv_tar_url := "https://github.com/rrthomas/mmv/releases/download/v2.7/mmv-2.7.tar.gz"
	mmv_cmd_tar := exec.Command("curl", "-L", mmv_tar_url, "-o", "package.tar.gz")
	err := mmv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmv_zip_url := "https://github.com/rrthomas/mmv/releases/download/v2.7/mmv-2.7.zip"
	mmv_cmd_zip := exec.Command("curl", "-L", mmv_zip_url, "-o", "package.zip")
	err = mmv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmv_bin_url := "https://github.com/rrthomas/mmv/releases/download/v2.7/mmv-2.7.bin"
	mmv_cmd_bin := exec.Command("curl", "-L", mmv_bin_url, "-o", "binary.bin")
	err = mmv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmv_src_url := "https://github.com/rrthomas/mmv/releases/download/v2.7/mmv-2.7.src.tar.gz"
	mmv_cmd_src := exec.Command("curl", "-L", mmv_src_url, "-o", "source.tar.gz")
	err = mmv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmv_cmd_direct := exec.Command("./binary")
	err = mmv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
}
