package main

// VirustotalCli - Command-line interface for VirusTotal
// Homepage: https://github.com/VirusTotal/vt-cli

import (
	"fmt"
	
	"os/exec"
)

func installVirustotalCli() {
	// Método 1: Descargar y extraer .tar.gz
	virustotalcli_tar_url := "https://github.com/VirusTotal/vt-cli/archive/refs/tags/1.0.1.tar.gz"
	virustotalcli_cmd_tar := exec.Command("curl", "-L", virustotalcli_tar_url, "-o", "package.tar.gz")
	err := virustotalcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virustotalcli_zip_url := "https://github.com/VirusTotal/vt-cli/archive/refs/tags/1.0.1.zip"
	virustotalcli_cmd_zip := exec.Command("curl", "-L", virustotalcli_zip_url, "-o", "package.zip")
	err = virustotalcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virustotalcli_bin_url := "https://github.com/VirusTotal/vt-cli/archive/refs/tags/1.0.1.bin"
	virustotalcli_cmd_bin := exec.Command("curl", "-L", virustotalcli_bin_url, "-o", "binary.bin")
	err = virustotalcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virustotalcli_src_url := "https://github.com/VirusTotal/vt-cli/archive/refs/tags/1.0.1.src.tar.gz"
	virustotalcli_cmd_src := exec.Command("curl", "-L", virustotalcli_src_url, "-o", "source.tar.gz")
	err = virustotalcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virustotalcli_cmd_direct := exec.Command("./binary")
	err = virustotalcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
