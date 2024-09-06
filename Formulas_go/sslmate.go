package main

// Sslmate - Buy SSL certs from the command-line
// Homepage: https://sslmate.com

import (
	"fmt"
	
	"os/exec"
)

func installSslmate() {
	// Método 1: Descargar y extraer .tar.gz
	sslmate_tar_url := "https://packages.sslmate.com/other/sslmate-1.9.1.tar.gz"
	sslmate_cmd_tar := exec.Command("curl", "-L", sslmate_tar_url, "-o", "package.tar.gz")
	err := sslmate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sslmate_zip_url := "https://packages.sslmate.com/other/sslmate-1.9.1.zip"
	sslmate_cmd_zip := exec.Command("curl", "-L", sslmate_zip_url, "-o", "package.zip")
	err = sslmate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sslmate_bin_url := "https://packages.sslmate.com/other/sslmate-1.9.1.bin"
	sslmate_cmd_bin := exec.Command("curl", "-L", sslmate_bin_url, "-o", "binary.bin")
	err = sslmate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sslmate_src_url := "https://packages.sslmate.com/other/sslmate-1.9.1.src.tar.gz"
	sslmate_cmd_src := exec.Command("curl", "-L", sslmate_src_url, "-o", "source.tar.gz")
	err = sslmate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sslmate_cmd_direct := exec.Command("./binary")
	err = sslmate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
