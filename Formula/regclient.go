package main

// Regclient - Docker and OCI Registry Client in Go and tooling using those libraries
// Homepage: https://github.com/regclient/regclient

import (
	"fmt"
	
	"os/exec"
)

func installRegclient() {
	// Método 1: Descargar y extraer .tar.gz
	regclient_tar_url := "https://github.com/regclient/regclient/archive/refs/tags/v0.7.1.tar.gz"
	regclient_cmd_tar := exec.Command("curl", "-L", regclient_tar_url, "-o", "package.tar.gz")
	err := regclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	regclient_zip_url := "https://github.com/regclient/regclient/archive/refs/tags/v0.7.1.zip"
	regclient_cmd_zip := exec.Command("curl", "-L", regclient_zip_url, "-o", "package.zip")
	err = regclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	regclient_bin_url := "https://github.com/regclient/regclient/archive/refs/tags/v0.7.1.bin"
	regclient_cmd_bin := exec.Command("curl", "-L", regclient_bin_url, "-o", "binary.bin")
	err = regclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	regclient_src_url := "https://github.com/regclient/regclient/archive/refs/tags/v0.7.1.src.tar.gz"
	regclient_cmd_src := exec.Command("curl", "-L", regclient_src_url, "-o", "source.tar.gz")
	err = regclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	regclient_cmd_direct := exec.Command("./binary")
	err = regclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
