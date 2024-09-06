package main

// Envchain - Secure your credentials in environment variables
// Homepage: https://github.com/sorah/envchain

import (
	"fmt"
	
	"os/exec"
)

func installEnvchain() {
	// Método 1: Descargar y extraer .tar.gz
	envchain_tar_url := "https://github.com/sorah/envchain/archive/refs/tags/v1.1.0.tar.gz"
	envchain_cmd_tar := exec.Command("curl", "-L", envchain_tar_url, "-o", "package.tar.gz")
	err := envchain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	envchain_zip_url := "https://github.com/sorah/envchain/archive/refs/tags/v1.1.0.zip"
	envchain_cmd_zip := exec.Command("curl", "-L", envchain_zip_url, "-o", "package.zip")
	err = envchain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	envchain_bin_url := "https://github.com/sorah/envchain/archive/refs/tags/v1.1.0.bin"
	envchain_cmd_bin := exec.Command("curl", "-L", envchain_bin_url, "-o", "binary.bin")
	err = envchain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	envchain_src_url := "https://github.com/sorah/envchain/archive/refs/tags/v1.1.0.src.tar.gz"
	envchain_cmd_src := exec.Command("curl", "-L", envchain_src_url, "-o", "source.tar.gz")
	err = envchain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	envchain_cmd_direct := exec.Command("./binary")
	err = envchain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsecret")
exec.Command("latte", "install", "libsecret")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
