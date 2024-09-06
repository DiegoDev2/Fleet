package main

// SafCli - CLI for the MITRE Security Automation Framework (SAF)
// Homepage: https://saf-cli.mitre.org

import (
	"fmt"
	
	"os/exec"
)

func installSafCli() {
	// Método 1: Descargar y extraer .tar.gz
	safcli_tar_url := "https://registry.npmjs.org/@mitre/saf/-/saf-1.4.12.tgz"
	safcli_cmd_tar := exec.Command("curl", "-L", safcli_tar_url, "-o", "package.tar.gz")
	err := safcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	safcli_zip_url := "https://registry.npmjs.org/@mitre/saf/-/saf-1.4.12.tgz"
	safcli_cmd_zip := exec.Command("curl", "-L", safcli_zip_url, "-o", "package.zip")
	err = safcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	safcli_bin_url := "https://registry.npmjs.org/@mitre/saf/-/saf-1.4.12.tgz"
	safcli_cmd_bin := exec.Command("curl", "-L", safcli_bin_url, "-o", "binary.bin")
	err = safcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	safcli_src_url := "https://registry.npmjs.org/@mitre/saf/-/saf-1.4.12.tgz"
	safcli_cmd_src := exec.Command("curl", "-L", safcli_src_url, "-o", "source.tar.gz")
	err = safcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	safcli_cmd_direct := exec.Command("./binary")
	err = safcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
