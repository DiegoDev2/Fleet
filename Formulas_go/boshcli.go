package main

// BoshCli - Cloud Foundry BOSH CLI v2
// Homepage: https://bosh.io/docs/cli-v2/

import (
	"fmt"
	
	"os/exec"
)

func installBoshCli() {
	// Método 1: Descargar y extraer .tar.gz
	boshcli_tar_url := "https://github.com/cloudfoundry/bosh-cli/archive/refs/tags/v7.7.1.tar.gz"
	boshcli_cmd_tar := exec.Command("curl", "-L", boshcli_tar_url, "-o", "package.tar.gz")
	err := boshcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boshcli_zip_url := "https://github.com/cloudfoundry/bosh-cli/archive/refs/tags/v7.7.1.zip"
	boshcli_cmd_zip := exec.Command("curl", "-L", boshcli_zip_url, "-o", "package.zip")
	err = boshcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boshcli_bin_url := "https://github.com/cloudfoundry/bosh-cli/archive/refs/tags/v7.7.1.bin"
	boshcli_cmd_bin := exec.Command("curl", "-L", boshcli_bin_url, "-o", "binary.bin")
	err = boshcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boshcli_src_url := "https://github.com/cloudfoundry/bosh-cli/archive/refs/tags/v7.7.1.src.tar.gz"
	boshcli_cmd_src := exec.Command("curl", "-L", boshcli_src_url, "-o", "source.tar.gz")
	err = boshcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boshcli_cmd_direct := exec.Command("./binary")
	err = boshcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
