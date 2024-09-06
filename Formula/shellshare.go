package main

// Shellshare - Live Terminal Broadcast
// Homepage: https://shellshare.net

import (
	"fmt"
	
	"os/exec"
)

func installShellshare() {
	// Método 1: Descargar y extraer .tar.gz
	shellshare_tar_url := "https://github.com/vitorbaptista/shellshare/archive/refs/tags/v1.1.0.tar.gz"
	shellshare_cmd_tar := exec.Command("curl", "-L", shellshare_tar_url, "-o", "package.tar.gz")
	err := shellshare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shellshare_zip_url := "https://github.com/vitorbaptista/shellshare/archive/refs/tags/v1.1.0.zip"
	shellshare_cmd_zip := exec.Command("curl", "-L", shellshare_zip_url, "-o", "package.zip")
	err = shellshare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shellshare_bin_url := "https://github.com/vitorbaptista/shellshare/archive/refs/tags/v1.1.0.bin"
	shellshare_cmd_bin := exec.Command("curl", "-L", shellshare_bin_url, "-o", "binary.bin")
	err = shellshare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shellshare_src_url := "https://github.com/vitorbaptista/shellshare/archive/refs/tags/v1.1.0.src.tar.gz"
	shellshare_cmd_src := exec.Command("curl", "-L", shellshare_src_url, "-o", "source.tar.gz")
	err = shellshare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shellshare_cmd_direct := exec.Command("./binary")
	err = shellshare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
