package main

// Telnet - User interface to the TELNET protocol
// Homepage: https://opensource.apple.com/

import (
	"fmt"
	
	"os/exec"
)

func installTelnet() {
	// Método 1: Descargar y extraer .tar.gz
	telnet_tar_url := "https://github.com/apple-oss-distributions/remote_cmds/archive/refs/tags/remote_cmds-303.141.1.tar.gz"
	telnet_cmd_tar := exec.Command("curl", "-L", telnet_tar_url, "-o", "package.tar.gz")
	err := telnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	telnet_zip_url := "https://github.com/apple-oss-distributions/remote_cmds/archive/refs/tags/remote_cmds-303.141.1.zip"
	telnet_cmd_zip := exec.Command("curl", "-L", telnet_zip_url, "-o", "package.zip")
	err = telnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	telnet_bin_url := "https://github.com/apple-oss-distributions/remote_cmds/archive/refs/tags/remote_cmds-303.141.1.bin"
	telnet_cmd_bin := exec.Command("curl", "-L", telnet_bin_url, "-o", "binary.bin")
	err = telnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	telnet_src_url := "https://github.com/apple-oss-distributions/remote_cmds/archive/refs/tags/remote_cmds-303.141.1.src.tar.gz"
	telnet_cmd_src := exec.Command("curl", "-L", telnet_src_url, "-o", "source.tar.gz")
	err = telnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	telnet_cmd_direct := exec.Command("./binary")
	err = telnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
