package main

// CiliumCli - CLI to install, manage & troubleshoot Kubernetes clusters running Cilium
// Homepage: https://cilium.io

import (
	"fmt"
	
	"os/exec"
)

func installCiliumCli() {
	// Método 1: Descargar y extraer .tar.gz
	ciliumcli_tar_url := "https://github.com/cilium/cilium-cli/archive/refs/tags/v0.16.16.tar.gz"
	ciliumcli_cmd_tar := exec.Command("curl", "-L", ciliumcli_tar_url, "-o", "package.tar.gz")
	err := ciliumcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ciliumcli_zip_url := "https://github.com/cilium/cilium-cli/archive/refs/tags/v0.16.16.zip"
	ciliumcli_cmd_zip := exec.Command("curl", "-L", ciliumcli_zip_url, "-o", "package.zip")
	err = ciliumcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ciliumcli_bin_url := "https://github.com/cilium/cilium-cli/archive/refs/tags/v0.16.16.bin"
	ciliumcli_cmd_bin := exec.Command("curl", "-L", ciliumcli_bin_url, "-o", "binary.bin")
	err = ciliumcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ciliumcli_src_url := "https://github.com/cilium/cilium-cli/archive/refs/tags/v0.16.16.src.tar.gz"
	ciliumcli_cmd_src := exec.Command("curl", "-L", ciliumcli_src_url, "-o", "source.tar.gz")
	err = ciliumcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ciliumcli_cmd_direct := exec.Command("./binary")
	err = ciliumcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
