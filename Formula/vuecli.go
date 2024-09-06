package main

// VueCli - Standard Tooling for Vue.js Development
// Homepage: https://cli.vuejs.org/

import (
	"fmt"
	
	"os/exec"
)

func installVueCli() {
	// Método 1: Descargar y extraer .tar.gz
	vuecli_tar_url := "https://registry.npmjs.org/@vue/cli/-/cli-5.0.8.tgz"
	vuecli_cmd_tar := exec.Command("curl", "-L", vuecli_tar_url, "-o", "package.tar.gz")
	err := vuecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vuecli_zip_url := "https://registry.npmjs.org/@vue/cli/-/cli-5.0.8.tgz"
	vuecli_cmd_zip := exec.Command("curl", "-L", vuecli_zip_url, "-o", "package.zip")
	err = vuecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vuecli_bin_url := "https://registry.npmjs.org/@vue/cli/-/cli-5.0.8.tgz"
	vuecli_cmd_bin := exec.Command("curl", "-L", vuecli_bin_url, "-o", "binary.bin")
	err = vuecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vuecli_src_url := "https://registry.npmjs.org/@vue/cli/-/cli-5.0.8.tgz"
	vuecli_cmd_src := exec.Command("curl", "-L", vuecli_src_url, "-o", "source.tar.gz")
	err = vuecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vuecli_cmd_direct := exec.Command("./binary")
	err = vuecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: terminal-notifier")
	exec.Command("latte", "install", "terminal-notifier").Run()
}
