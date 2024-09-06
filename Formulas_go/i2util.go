package main

// I2util - Internet2 utility tools
// Homepage: https://github.com/perfsonar/i2util

import (
	"fmt"
	
	"os/exec"
)

func installI2util() {
	// Método 1: Descargar y extraer .tar.gz
	i2util_tar_url := "https://github.com/perfsonar/i2util/archive/refs/tags/v5.1.2.tar.gz"
	i2util_cmd_tar := exec.Command("curl", "-L", i2util_tar_url, "-o", "package.tar.gz")
	err := i2util_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	i2util_zip_url := "https://github.com/perfsonar/i2util/archive/refs/tags/v5.1.2.zip"
	i2util_cmd_zip := exec.Command("curl", "-L", i2util_zip_url, "-o", "package.zip")
	err = i2util_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	i2util_bin_url := "https://github.com/perfsonar/i2util/archive/refs/tags/v5.1.2.bin"
	i2util_cmd_bin := exec.Command("curl", "-L", i2util_bin_url, "-o", "binary.bin")
	err = i2util_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	i2util_src_url := "https://github.com/perfsonar/i2util/archive/refs/tags/v5.1.2.src.tar.gz"
	i2util_cmd_src := exec.Command("curl", "-L", i2util_src_url, "-o", "source.tar.gz")
	err = i2util_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	i2util_cmd_direct := exec.Command("./binary")
	err = i2util_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
