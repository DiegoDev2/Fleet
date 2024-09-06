package main

// Botan - Cryptographic algorithms and formats library in C++
// Homepage: https://botan.randombit.net/

import (
	"fmt"
	
	"os/exec"
)

func installBotan() {
	// Método 1: Descargar y extraer .tar.gz
	botan_tar_url := "https://botan.randombit.net/releases/Botan-3.5.0.tar.xz"
	botan_cmd_tar := exec.Command("curl", "-L", botan_tar_url, "-o", "package.tar.gz")
	err := botan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	botan_zip_url := "https://botan.randombit.net/releases/Botan-3.5.0.tar.xz"
	botan_cmd_zip := exec.Command("curl", "-L", botan_zip_url, "-o", "package.zip")
	err = botan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	botan_bin_url := "https://botan.randombit.net/releases/Botan-3.5.0.tar.xz"
	botan_cmd_bin := exec.Command("curl", "-L", botan_bin_url, "-o", "binary.bin")
	err = botan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	botan_src_url := "https://botan.randombit.net/releases/Botan-3.5.0.tar.xz"
	botan_cmd_src := exec.Command("curl", "-L", botan_src_url, "-o", "source.tar.gz")
	err = botan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	botan_cmd_direct := exec.Command("./binary")
	err = botan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
