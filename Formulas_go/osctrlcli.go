package main

// OsctrlCli - Fast and efficient osquery management
// Homepage: https://osctrl.net

import (
	"fmt"
	
	"os/exec"
)

func installOsctrlCli() {
	// Método 1: Descargar y extraer .tar.gz
	osctrlcli_tar_url := "https://github.com/jmpsec/osctrl/archive/refs/tags/v0.3.8.tar.gz"
	osctrlcli_cmd_tar := exec.Command("curl", "-L", osctrlcli_tar_url, "-o", "package.tar.gz")
	err := osctrlcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osctrlcli_zip_url := "https://github.com/jmpsec/osctrl/archive/refs/tags/v0.3.8.zip"
	osctrlcli_cmd_zip := exec.Command("curl", "-L", osctrlcli_zip_url, "-o", "package.zip")
	err = osctrlcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osctrlcli_bin_url := "https://github.com/jmpsec/osctrl/archive/refs/tags/v0.3.8.bin"
	osctrlcli_cmd_bin := exec.Command("curl", "-L", osctrlcli_bin_url, "-o", "binary.bin")
	err = osctrlcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osctrlcli_src_url := "https://github.com/jmpsec/osctrl/archive/refs/tags/v0.3.8.src.tar.gz"
	osctrlcli_cmd_src := exec.Command("curl", "-L", osctrlcli_src_url, "-o", "source.tar.gz")
	err = osctrlcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osctrlcli_cmd_direct := exec.Command("./binary")
	err = osctrlcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
