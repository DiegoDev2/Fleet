package main

// Terracognita - Reads from existing Cloud Providers and generates Terraform code
// Homepage: https://github.com/cycloidio/terracognita

import (
	"fmt"
	
	"os/exec"
)

func installTerracognita() {
	// Método 1: Descargar y extraer .tar.gz
	terracognita_tar_url := "https://github.com/cycloidio/terracognita/archive/refs/tags/v0.8.4.tar.gz"
	terracognita_cmd_tar := exec.Command("curl", "-L", terracognita_tar_url, "-o", "package.tar.gz")
	err := terracognita_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terracognita_zip_url := "https://github.com/cycloidio/terracognita/archive/refs/tags/v0.8.4.zip"
	terracognita_cmd_zip := exec.Command("curl", "-L", terracognita_zip_url, "-o", "package.zip")
	err = terracognita_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terracognita_bin_url := "https://github.com/cycloidio/terracognita/archive/refs/tags/v0.8.4.bin"
	terracognita_cmd_bin := exec.Command("curl", "-L", terracognita_bin_url, "-o", "binary.bin")
	err = terracognita_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terracognita_src_url := "https://github.com/cycloidio/terracognita/archive/refs/tags/v0.8.4.src.tar.gz"
	terracognita_cmd_src := exec.Command("curl", "-L", terracognita_src_url, "-o", "source.tar.gz")
	err = terracognita_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terracognita_cmd_direct := exec.Command("./binary")
	err = terracognita_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
