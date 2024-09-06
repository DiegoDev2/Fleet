package main

// Libchaos - Advanced library for randomization, hashing and statistical analysis
// Homepage: https://github.com/maciejczyzewski/libchaos

import (
	"fmt"
	
	"os/exec"
)

func installLibchaos() {
	// Método 1: Descargar y extraer .tar.gz
	libchaos_tar_url := "https://github.com/maciejczyzewski/libchaos/releases/download/v1.0/libchaos-1.0.tar.gz"
	libchaos_cmd_tar := exec.Command("curl", "-L", libchaos_tar_url, "-o", "package.tar.gz")
	err := libchaos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libchaos_zip_url := "https://github.com/maciejczyzewski/libchaos/releases/download/v1.0/libchaos-1.0.zip"
	libchaos_cmd_zip := exec.Command("curl", "-L", libchaos_zip_url, "-o", "package.zip")
	err = libchaos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libchaos_bin_url := "https://github.com/maciejczyzewski/libchaos/releases/download/v1.0/libchaos-1.0.bin"
	libchaos_cmd_bin := exec.Command("curl", "-L", libchaos_bin_url, "-o", "binary.bin")
	err = libchaos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libchaos_src_url := "https://github.com/maciejczyzewski/libchaos/releases/download/v1.0/libchaos-1.0.src.tar.gz"
	libchaos_cmd_src := exec.Command("curl", "-L", libchaos_src_url, "-o", "source.tar.gz")
	err = libchaos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libchaos_cmd_direct := exec.Command("./binary")
	err = libchaos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
