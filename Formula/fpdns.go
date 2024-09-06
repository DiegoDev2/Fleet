package main

// Fpdns - Fingerprint DNS server versions
// Homepage: https://github.com/kirei/fpdns

import (
	"fmt"
	
	"os/exec"
)

func installFpdns() {
	// Método 1: Descargar y extraer .tar.gz
	fpdns_tar_url := "https://github.com/kirei/fpdns/archive/refs/tags/20190131.tar.gz"
	fpdns_cmd_tar := exec.Command("curl", "-L", fpdns_tar_url, "-o", "package.tar.gz")
	err := fpdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fpdns_zip_url := "https://github.com/kirei/fpdns/archive/refs/tags/20190131.zip"
	fpdns_cmd_zip := exec.Command("curl", "-L", fpdns_zip_url, "-o", "package.zip")
	err = fpdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fpdns_bin_url := "https://github.com/kirei/fpdns/archive/refs/tags/20190131.bin"
	fpdns_cmd_bin := exec.Command("curl", "-L", fpdns_bin_url, "-o", "binary.bin")
	err = fpdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fpdns_src_url := "https://github.com/kirei/fpdns/archive/refs/tags/20190131.src.tar.gz"
	fpdns_cmd_src := exec.Command("curl", "-L", fpdns_src_url, "-o", "source.tar.gz")
	err = fpdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fpdns_cmd_direct := exec.Command("./binary")
	err = fpdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
