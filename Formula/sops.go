package main

// Sops - Editor of encrypted files
// Homepage: https://github.com/getsops/sops

import (
	"fmt"
	
	"os/exec"
)

func installSops() {
	// Método 1: Descargar y extraer .tar.gz
	sops_tar_url := "https://github.com/getsops/sops/archive/refs/tags/v3.9.0.tar.gz"
	sops_cmd_tar := exec.Command("curl", "-L", sops_tar_url, "-o", "package.tar.gz")
	err := sops_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sops_zip_url := "https://github.com/getsops/sops/archive/refs/tags/v3.9.0.zip"
	sops_cmd_zip := exec.Command("curl", "-L", sops_zip_url, "-o", "package.zip")
	err = sops_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sops_bin_url := "https://github.com/getsops/sops/archive/refs/tags/v3.9.0.bin"
	sops_cmd_bin := exec.Command("curl", "-L", sops_bin_url, "-o", "binary.bin")
	err = sops_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sops_src_url := "https://github.com/getsops/sops/archive/refs/tags/v3.9.0.src.tar.gz"
	sops_cmd_src := exec.Command("curl", "-L", sops_src_url, "-o", "source.tar.gz")
	err = sops_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sops_cmd_direct := exec.Command("./binary")
	err = sops_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
