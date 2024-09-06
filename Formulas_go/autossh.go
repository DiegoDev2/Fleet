package main

// Autossh - Automatically restart SSH sessions and tunnels
// Homepage: https://www.harding.motd.ca/autossh/

import (
	"fmt"
	
	"os/exec"
)

func installAutossh() {
	// Método 1: Descargar y extraer .tar.gz
	autossh_tar_url := "https://www.harding.motd.ca/autossh/autossh-1.4g.tgz"
	autossh_cmd_tar := exec.Command("curl", "-L", autossh_tar_url, "-o", "package.tar.gz")
	err := autossh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autossh_zip_url := "https://www.harding.motd.ca/autossh/autossh-1.4g.tgz"
	autossh_cmd_zip := exec.Command("curl", "-L", autossh_zip_url, "-o", "package.zip")
	err = autossh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autossh_bin_url := "https://www.harding.motd.ca/autossh/autossh-1.4g.tgz"
	autossh_cmd_bin := exec.Command("curl", "-L", autossh_bin_url, "-o", "binary.bin")
	err = autossh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autossh_src_url := "https://www.harding.motd.ca/autossh/autossh-1.4g.tgz"
	autossh_cmd_src := exec.Command("curl", "-L", autossh_src_url, "-o", "source.tar.gz")
	err = autossh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autossh_cmd_direct := exec.Command("./binary")
	err = autossh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
