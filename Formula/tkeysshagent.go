package main

// TkeySshAgent - SSH agent for use with the TKey security stick
// Homepage: https://tillitis.se/

import (
	"fmt"
	
	"os/exec"
)

func installTkeySshAgent() {
	// Método 1: Descargar y extraer .tar.gz
	tkeysshagent_tar_url := "https://github.com/tillitis/tkey-ssh-agent/archive/refs/tags/v1.0.0.tar.gz"
	tkeysshagent_cmd_tar := exec.Command("curl", "-L", tkeysshagent_tar_url, "-o", "package.tar.gz")
	err := tkeysshagent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tkeysshagent_zip_url := "https://github.com/tillitis/tkey-ssh-agent/archive/refs/tags/v1.0.0.zip"
	tkeysshagent_cmd_zip := exec.Command("curl", "-L", tkeysshagent_zip_url, "-o", "package.zip")
	err = tkeysshagent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tkeysshagent_bin_url := "https://github.com/tillitis/tkey-ssh-agent/archive/refs/tags/v1.0.0.bin"
	tkeysshagent_cmd_bin := exec.Command("curl", "-L", tkeysshagent_bin_url, "-o", "binary.bin")
	err = tkeysshagent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tkeysshagent_src_url := "https://github.com/tillitis/tkey-ssh-agent/archive/refs/tags/v1.0.0.src.tar.gz"
	tkeysshagent_cmd_src := exec.Command("curl", "-L", tkeysshagent_src_url, "-o", "source.tar.gz")
	err = tkeysshagent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tkeysshagent_cmd_direct := exec.Command("./binary")
	err = tkeysshagent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pinentry-mac")
	exec.Command("latte", "install", "pinentry-mac").Run()
	fmt.Println("Instalando dependencia: pinentry")
	exec.Command("latte", "install", "pinentry").Run()
}
