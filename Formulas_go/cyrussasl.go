package main

// CyrusSasl - Simple Authentication and Security Layer
// Homepage: https://www.cyrusimap.org/sasl/

import (
	"fmt"
	
	"os/exec"
)

func installCyrusSasl() {
	// Método 1: Descargar y extraer .tar.gz
	cyrussasl_tar_url := "https://github.com/cyrusimap/cyrus-sasl/releases/download/cyrus-sasl-2.1.28/cyrus-sasl-2.1.28.tar.gz"
	cyrussasl_cmd_tar := exec.Command("curl", "-L", cyrussasl_tar_url, "-o", "package.tar.gz")
	err := cyrussasl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cyrussasl_zip_url := "https://github.com/cyrusimap/cyrus-sasl/releases/download/cyrus-sasl-2.1.28/cyrus-sasl-2.1.28.zip"
	cyrussasl_cmd_zip := exec.Command("curl", "-L", cyrussasl_zip_url, "-o", "package.zip")
	err = cyrussasl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cyrussasl_bin_url := "https://github.com/cyrusimap/cyrus-sasl/releases/download/cyrus-sasl-2.1.28/cyrus-sasl-2.1.28.bin"
	cyrussasl_cmd_bin := exec.Command("curl", "-L", cyrussasl_bin_url, "-o", "binary.bin")
	err = cyrussasl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cyrussasl_src_url := "https://github.com/cyrusimap/cyrus-sasl/releases/download/cyrus-sasl-2.1.28/cyrus-sasl-2.1.28.src.tar.gz"
	cyrussasl_cmd_src := exec.Command("curl", "-L", cyrussasl_src_url, "-o", "source.tar.gz")
	err = cyrussasl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cyrussasl_cmd_direct := exec.Command("./binary")
	err = cyrussasl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: krb5")
exec.Command("latte", "install", "krb5")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
