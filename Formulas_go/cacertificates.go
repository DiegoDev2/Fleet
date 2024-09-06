package main

// CaCertificates - Mozilla CA certificate store
// Homepage: https://curl.se/docs/caextract.html

import (
	"fmt"
	
	"os/exec"
)

func installCaCertificates() {
	// Método 1: Descargar y extraer .tar.gz
	cacertificates_tar_url := "https://curl.se/ca/cacert-2024-07-02.pem"
	cacertificates_cmd_tar := exec.Command("curl", "-L", cacertificates_tar_url, "-o", "package.tar.gz")
	err := cacertificates_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cacertificates_zip_url := "https://curl.se/ca/cacert-2024-07-02.pem"
	cacertificates_cmd_zip := exec.Command("curl", "-L", cacertificates_zip_url, "-o", "package.zip")
	err = cacertificates_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cacertificates_bin_url := "https://curl.se/ca/cacert-2024-07-02.pem"
	cacertificates_cmd_bin := exec.Command("curl", "-L", cacertificates_bin_url, "-o", "binary.bin")
	err = cacertificates_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cacertificates_src_url := "https://curl.se/ca/cacert-2024-07-02.pem"
	cacertificates_cmd_src := exec.Command("curl", "-L", cacertificates_src_url, "-o", "source.tar.gz")
	err = cacertificates_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cacertificates_cmd_direct := exec.Command("./binary")
	err = cacertificates_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
