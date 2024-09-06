package main

// Mbelib - P25 Phase 1 and ProVoice vocoder
// Homepage: https://github.com/szechyjs/mbelib

import (
	"fmt"
	
	"os/exec"
)

func installMbelib() {
	// Método 1: Descargar y extraer .tar.gz
	mbelib_tar_url := "https://github.com/szechyjs/mbelib/archive/refs/tags/v1.3.0.tar.gz"
	mbelib_cmd_tar := exec.Command("curl", "-L", mbelib_tar_url, "-o", "package.tar.gz")
	err := mbelib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mbelib_zip_url := "https://github.com/szechyjs/mbelib/archive/refs/tags/v1.3.0.zip"
	mbelib_cmd_zip := exec.Command("curl", "-L", mbelib_zip_url, "-o", "package.zip")
	err = mbelib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mbelib_bin_url := "https://github.com/szechyjs/mbelib/archive/refs/tags/v1.3.0.bin"
	mbelib_cmd_bin := exec.Command("curl", "-L", mbelib_bin_url, "-o", "binary.bin")
	err = mbelib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mbelib_src_url := "https://github.com/szechyjs/mbelib/archive/refs/tags/v1.3.0.src.tar.gz"
	mbelib_cmd_src := exec.Command("curl", "-L", mbelib_src_url, "-o", "source.tar.gz")
	err = mbelib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mbelib_cmd_direct := exec.Command("./binary")
	err = mbelib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
