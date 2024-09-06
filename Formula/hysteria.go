package main

// Hysteria - Feature-packed proxy & relay tool optimized for lossy, unstable connections
// Homepage: https://hysteria.network/

import (
	"fmt"
	
	"os/exec"
)

func installHysteria() {
	// Método 1: Descargar y extraer .tar.gz
	hysteria_tar_url := "https://github.com/apernet/hysteria/archive/refs/tags/app/v2.5.1.tar.gz"
	hysteria_cmd_tar := exec.Command("curl", "-L", hysteria_tar_url, "-o", "package.tar.gz")
	err := hysteria_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hysteria_zip_url := "https://github.com/apernet/hysteria/archive/refs/tags/app/v2.5.1.zip"
	hysteria_cmd_zip := exec.Command("curl", "-L", hysteria_zip_url, "-o", "package.zip")
	err = hysteria_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hysteria_bin_url := "https://github.com/apernet/hysteria/archive/refs/tags/app/v2.5.1.bin"
	hysteria_cmd_bin := exec.Command("curl", "-L", hysteria_bin_url, "-o", "binary.bin")
	err = hysteria_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hysteria_src_url := "https://github.com/apernet/hysteria/archive/refs/tags/app/v2.5.1.src.tar.gz"
	hysteria_cmd_src := exec.Command("curl", "-L", hysteria_src_url, "-o", "source.tar.gz")
	err = hysteria_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hysteria_cmd_direct := exec.Command("./binary")
	err = hysteria_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
