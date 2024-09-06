package main

// LeappCli - Cloud credentials manager cli
// Homepage: https://github.com/noovolari/leapp

import (
	"fmt"
	
	"os/exec"
)

func installLeappCli() {
	// Método 1: Descargar y extraer .tar.gz
	leappcli_tar_url := "https://registry.npmjs.org/@noovolari/leapp-cli/-/leapp-cli-0.1.65.tgz"
	leappcli_cmd_tar := exec.Command("curl", "-L", leappcli_tar_url, "-o", "package.tar.gz")
	err := leappcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leappcli_zip_url := "https://registry.npmjs.org/@noovolari/leapp-cli/-/leapp-cli-0.1.65.tgz"
	leappcli_cmd_zip := exec.Command("curl", "-L", leappcli_zip_url, "-o", "package.zip")
	err = leappcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leappcli_bin_url := "https://registry.npmjs.org/@noovolari/leapp-cli/-/leapp-cli-0.1.65.tgz"
	leappcli_cmd_bin := exec.Command("curl", "-L", leappcli_bin_url, "-o", "binary.bin")
	err = leappcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leappcli_src_url := "https://registry.npmjs.org/@noovolari/leapp-cli/-/leapp-cli-0.1.65.tgz"
	leappcli_cmd_src := exec.Command("curl", "-L", leappcli_src_url, "-o", "source.tar.gz")
	err = leappcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leappcli_cmd_direct := exec.Command("./binary")
	err = leappcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: libsecret")
	exec.Command("latte", "install", "libsecret").Run()
}
