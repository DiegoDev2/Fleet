package main

// Pwnat - Proxy server that works behind a NAT
// Homepage: https://samy.pl/pwnat/

import (
	"fmt"
	
	"os/exec"
)

func installPwnat() {
	// Método 1: Descargar y extraer .tar.gz
	pwnat_tar_url := "https://github.com/samyk/pwnat/archive/refs/tags/v0.3.0.tar.gz"
	pwnat_cmd_tar := exec.Command("curl", "-L", pwnat_tar_url, "-o", "package.tar.gz")
	err := pwnat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pwnat_zip_url := "https://github.com/samyk/pwnat/archive/refs/tags/v0.3.0.zip"
	pwnat_cmd_zip := exec.Command("curl", "-L", pwnat_zip_url, "-o", "package.zip")
	err = pwnat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pwnat_bin_url := "https://github.com/samyk/pwnat/archive/refs/tags/v0.3.0.bin"
	pwnat_cmd_bin := exec.Command("curl", "-L", pwnat_bin_url, "-o", "binary.bin")
	err = pwnat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pwnat_src_url := "https://github.com/samyk/pwnat/archive/refs/tags/v0.3.0.src.tar.gz"
	pwnat_cmd_src := exec.Command("curl", "-L", pwnat_src_url, "-o", "source.tar.gz")
	err = pwnat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pwnat_cmd_direct := exec.Command("./binary")
	err = pwnat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
