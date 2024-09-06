package main

// Overtls - Simple proxy tunnel for bypassing the GFW
// Homepage: https://github.com/ShadowsocksR-Live/overtls

import (
	"fmt"
	
	"os/exec"
)

func installOvertls() {
	// Método 1: Descargar y extraer .tar.gz
	overtls_tar_url := "https://github.com/ShadowsocksR-Live/overtls/archive/refs/tags/v0.2.33.tar.gz"
	overtls_cmd_tar := exec.Command("curl", "-L", overtls_tar_url, "-o", "package.tar.gz")
	err := overtls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	overtls_zip_url := "https://github.com/ShadowsocksR-Live/overtls/archive/refs/tags/v0.2.33.zip"
	overtls_cmd_zip := exec.Command("curl", "-L", overtls_zip_url, "-o", "package.zip")
	err = overtls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	overtls_bin_url := "https://github.com/ShadowsocksR-Live/overtls/archive/refs/tags/v0.2.33.bin"
	overtls_cmd_bin := exec.Command("curl", "-L", overtls_bin_url, "-o", "binary.bin")
	err = overtls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	overtls_src_url := "https://github.com/ShadowsocksR-Live/overtls/archive/refs/tags/v0.2.33.src.tar.gz"
	overtls_cmd_src := exec.Command("curl", "-L", overtls_src_url, "-o", "source.tar.gz")
	err = overtls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	overtls_cmd_direct := exec.Command("./binary")
	err = overtls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
