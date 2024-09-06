package main

// Sk - Fuzzy Finder in rust!
// Homepage: https://github.com/lotabout/skim

import (
	"fmt"
	
	"os/exec"
)

func installSk() {
	// Método 1: Descargar y extraer .tar.gz
	sk_tar_url := "https://github.com/lotabout/skim/archive/refs/tags/v0.10.4.tar.gz"
	sk_cmd_tar := exec.Command("curl", "-L", sk_tar_url, "-o", "package.tar.gz")
	err := sk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sk_zip_url := "https://github.com/lotabout/skim/archive/refs/tags/v0.10.4.zip"
	sk_cmd_zip := exec.Command("curl", "-L", sk_zip_url, "-o", "package.zip")
	err = sk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sk_bin_url := "https://github.com/lotabout/skim/archive/refs/tags/v0.10.4.bin"
	sk_cmd_bin := exec.Command("curl", "-L", sk_bin_url, "-o", "binary.bin")
	err = sk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sk_src_url := "https://github.com/lotabout/skim/archive/refs/tags/v0.10.4.src.tar.gz"
	sk_cmd_src := exec.Command("curl", "-L", sk_src_url, "-o", "source.tar.gz")
	err = sk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sk_cmd_direct := exec.Command("./binary")
	err = sk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
