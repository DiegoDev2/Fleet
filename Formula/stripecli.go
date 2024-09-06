package main

// StripeCli - Command-line tool for Stripe
// Homepage: https://stripe.com/docs/stripe-cli

import (
	"fmt"
	
	"os/exec"
)

func installStripeCli() {
	// Método 1: Descargar y extraer .tar.gz
	stripecli_tar_url := "https://github.com/stripe/stripe-cli/archive/refs/tags/v1.21.4.tar.gz"
	stripecli_cmd_tar := exec.Command("curl", "-L", stripecli_tar_url, "-o", "package.tar.gz")
	err := stripecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stripecli_zip_url := "https://github.com/stripe/stripe-cli/archive/refs/tags/v1.21.4.zip"
	stripecli_cmd_zip := exec.Command("curl", "-L", stripecli_zip_url, "-o", "package.zip")
	err = stripecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stripecli_bin_url := "https://github.com/stripe/stripe-cli/archive/refs/tags/v1.21.4.bin"
	stripecli_cmd_bin := exec.Command("curl", "-L", stripecli_bin_url, "-o", "binary.bin")
	err = stripecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stripecli_src_url := "https://github.com/stripe/stripe-cli/archive/refs/tags/v1.21.4.src.tar.gz"
	stripecli_cmd_src := exec.Command("curl", "-L", stripecli_src_url, "-o", "source.tar.gz")
	err = stripecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stripecli_cmd_direct := exec.Command("./binary")
	err = stripecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
