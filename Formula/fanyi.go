package main

// Fanyi - Chinese and English translate tool in your command-line
// Homepage: https://github.com/afc163/fanyi

import (
	"fmt"
	
	"os/exec"
)

func installFanyi() {
	// Método 1: Descargar y extraer .tar.gz
	fanyi_tar_url := "https://registry.npmjs.org/fanyi/-/fanyi-8.0.3.tgz"
	fanyi_cmd_tar := exec.Command("curl", "-L", fanyi_tar_url, "-o", "package.tar.gz")
	err := fanyi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fanyi_zip_url := "https://registry.npmjs.org/fanyi/-/fanyi-8.0.3.tgz"
	fanyi_cmd_zip := exec.Command("curl", "-L", fanyi_zip_url, "-o", "package.zip")
	err = fanyi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fanyi_bin_url := "https://registry.npmjs.org/fanyi/-/fanyi-8.0.3.tgz"
	fanyi_cmd_bin := exec.Command("curl", "-L", fanyi_bin_url, "-o", "binary.bin")
	err = fanyi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fanyi_src_url := "https://registry.npmjs.org/fanyi/-/fanyi-8.0.3.tgz"
	fanyi_cmd_src := exec.Command("curl", "-L", fanyi_src_url, "-o", "source.tar.gz")
	err = fanyi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fanyi_cmd_direct := exec.Command("./binary")
	err = fanyi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: macos-term-size")
	exec.Command("latte", "install", "macos-term-size").Run()
}
