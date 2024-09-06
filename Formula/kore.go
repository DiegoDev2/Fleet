package main

// Kore - Web application framework for writing web APIs in C
// Homepage: https://kore.io/

import (
	"fmt"
	
	"os/exec"
)

func installKore() {
	// Método 1: Descargar y extraer .tar.gz
	kore_tar_url := "https://kore.io/releases/kore-4.2.3.tar.gz"
	kore_cmd_tar := exec.Command("curl", "-L", kore_tar_url, "-o", "package.tar.gz")
	err := kore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kore_zip_url := "https://kore.io/releases/kore-4.2.3.zip"
	kore_cmd_zip := exec.Command("curl", "-L", kore_zip_url, "-o", "package.zip")
	err = kore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kore_bin_url := "https://kore.io/releases/kore-4.2.3.bin"
	kore_cmd_bin := exec.Command("curl", "-L", kore_bin_url, "-o", "binary.bin")
	err = kore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kore_src_url := "https://kore.io/releases/kore-4.2.3.src.tar.gz"
	kore_cmd_src := exec.Command("curl", "-L", kore_src_url, "-o", "source.tar.gz")
	err = kore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kore_cmd_direct := exec.Command("./binary")
	err = kore_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
