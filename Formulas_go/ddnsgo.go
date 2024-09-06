package main

// DdnsGo - Simple and easy-to-use DDNS
// Homepage: https://github.com/jeessy2/ddns-go

import (
	"fmt"
	
	"os/exec"
)

func installDdnsGo() {
	// Método 1: Descargar y extraer .tar.gz
	ddnsgo_tar_url := "https://github.com/jeessy2/ddns-go/archive/refs/tags/v6.7.0.tar.gz"
	ddnsgo_cmd_tar := exec.Command("curl", "-L", ddnsgo_tar_url, "-o", "package.tar.gz")
	err := ddnsgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddnsgo_zip_url := "https://github.com/jeessy2/ddns-go/archive/refs/tags/v6.7.0.zip"
	ddnsgo_cmd_zip := exec.Command("curl", "-L", ddnsgo_zip_url, "-o", "package.zip")
	err = ddnsgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddnsgo_bin_url := "https://github.com/jeessy2/ddns-go/archive/refs/tags/v6.7.0.bin"
	ddnsgo_cmd_bin := exec.Command("curl", "-L", ddnsgo_bin_url, "-o", "binary.bin")
	err = ddnsgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddnsgo_src_url := "https://github.com/jeessy2/ddns-go/archive/refs/tags/v6.7.0.src.tar.gz"
	ddnsgo_cmd_src := exec.Command("curl", "-L", ddnsgo_src_url, "-o", "source.tar.gz")
	err = ddnsgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddnsgo_cmd_direct := exec.Command("./binary")
	err = ddnsgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
