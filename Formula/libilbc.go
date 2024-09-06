package main

// Libilbc - Packaged version of iLBC codec from the WebRTC project
// Homepage: https://github.com/TimothyGu/libilbc

import (
	"fmt"
	
	"os/exec"
)

func installLibilbc() {
	// Método 1: Descargar y extraer .tar.gz
	libilbc_tar_url := "https://github.com/TimothyGu/libilbc/releases/download/v3.0.4/libilbc-3.0.4.tar.gz"
	libilbc_cmd_tar := exec.Command("curl", "-L", libilbc_tar_url, "-o", "package.tar.gz")
	err := libilbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libilbc_zip_url := "https://github.com/TimothyGu/libilbc/releases/download/v3.0.4/libilbc-3.0.4.zip"
	libilbc_cmd_zip := exec.Command("curl", "-L", libilbc_zip_url, "-o", "package.zip")
	err = libilbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libilbc_bin_url := "https://github.com/TimothyGu/libilbc/releases/download/v3.0.4/libilbc-3.0.4.bin"
	libilbc_cmd_bin := exec.Command("curl", "-L", libilbc_bin_url, "-o", "binary.bin")
	err = libilbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libilbc_src_url := "https://github.com/TimothyGu/libilbc/releases/download/v3.0.4/libilbc-3.0.4.src.tar.gz"
	libilbc_cmd_src := exec.Command("curl", "-L", libilbc_src_url, "-o", "source.tar.gz")
	err = libilbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libilbc_cmd_direct := exec.Command("./binary")
	err = libilbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
