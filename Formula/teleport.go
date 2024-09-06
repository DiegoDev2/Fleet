package main

// Teleport - Modern SSH server for teams managing distributed infrastructure
// Homepage: https://goteleport.com/

import (
	"fmt"
	
	"os/exec"
)

func installTeleport() {
	// Método 1: Descargar y extraer .tar.gz
	teleport_tar_url := "https://github.com/gravitational/teleport/archive/refs/tags/v14.3.3.tar.gz"
	teleport_cmd_tar := exec.Command("curl", "-L", teleport_tar_url, "-o", "package.tar.gz")
	err := teleport_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	teleport_zip_url := "https://github.com/gravitational/teleport/archive/refs/tags/v14.3.3.zip"
	teleport_cmd_zip := exec.Command("curl", "-L", teleport_zip_url, "-o", "package.zip")
	err = teleport_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	teleport_bin_url := "https://github.com/gravitational/teleport/archive/refs/tags/v14.3.3.bin"
	teleport_cmd_bin := exec.Command("curl", "-L", teleport_bin_url, "-o", "binary.bin")
	err = teleport_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	teleport_src_url := "https://github.com/gravitational/teleport/archive/refs/tags/v14.3.3.src.tar.gz"
	teleport_cmd_src := exec.Command("curl", "-L", teleport_src_url, "-o", "source.tar.gz")
	err = teleport_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	teleport_cmd_direct := exec.Command("./binary")
	err = teleport_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
	exec.Command("latte", "install", "go@1.22").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: yarn")
	exec.Command("latte", "install", "yarn").Run()
	fmt.Println("Instalando dependencia: libfido2")
	exec.Command("latte", "install", "libfido2").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
