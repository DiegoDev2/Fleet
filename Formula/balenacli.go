package main

// BalenaCli - Command-line tool for interacting with the balenaCloud and balena API
// Homepage: https://www.balena.io/docs/reference/cli/

import (
	"fmt"
	
	"os/exec"
)

func installBalenaCli() {
	// Método 1: Descargar y extraer .tar.gz
	balenacli_tar_url := "https://registry.npmjs.org/balena-cli/-/balena-cli-19.0.3.tgz"
	balenacli_cmd_tar := exec.Command("curl", "-L", balenacli_tar_url, "-o", "package.tar.gz")
	err := balenacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	balenacli_zip_url := "https://registry.npmjs.org/balena-cli/-/balena-cli-19.0.3.tgz"
	balenacli_cmd_zip := exec.Command("curl", "-L", balenacli_zip_url, "-o", "package.zip")
	err = balenacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	balenacli_bin_url := "https://registry.npmjs.org/balena-cli/-/balena-cli-19.0.3.tgz"
	balenacli_cmd_bin := exec.Command("curl", "-L", balenacli_bin_url, "-o", "binary.bin")
	err = balenacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	balenacli_src_url := "https://registry.npmjs.org/balena-cli/-/balena-cli-19.0.3.tgz"
	balenacli_cmd_src := exec.Command("curl", "-L", balenacli_src_url, "-o", "source.tar.gz")
	err = balenacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	balenacli_cmd_direct := exec.Command("./binary")
	err = balenacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node@20")
	exec.Command("latte", "install", "node@20").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
