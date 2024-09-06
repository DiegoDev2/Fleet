package main

// StripeMock - Mock HTTP server that responds like the real Stripe API
// Homepage: https://github.com/stripe/stripe-mock

import (
	"fmt"
	
	"os/exec"
)

func installStripeMock() {
	// Método 1: Descargar y extraer .tar.gz
	stripemock_tar_url := "https://github.com/stripe/stripe-mock/archive/refs/tags/v0.188.0.tar.gz"
	stripemock_cmd_tar := exec.Command("curl", "-L", stripemock_tar_url, "-o", "package.tar.gz")
	err := stripemock_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stripemock_zip_url := "https://github.com/stripe/stripe-mock/archive/refs/tags/v0.188.0.zip"
	stripemock_cmd_zip := exec.Command("curl", "-L", stripemock_zip_url, "-o", "package.zip")
	err = stripemock_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stripemock_bin_url := "https://github.com/stripe/stripe-mock/archive/refs/tags/v0.188.0.bin"
	stripemock_cmd_bin := exec.Command("curl", "-L", stripemock_bin_url, "-o", "binary.bin")
	err = stripemock_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stripemock_src_url := "https://github.com/stripe/stripe-mock/archive/refs/tags/v0.188.0.src.tar.gz"
	stripemock_cmd_src := exec.Command("curl", "-L", stripemock_src_url, "-o", "source.tar.gz")
	err = stripemock_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stripemock_cmd_direct := exec.Command("./binary")
	err = stripemock_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
