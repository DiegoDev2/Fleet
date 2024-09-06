package main

// Kcov - Code coverage tester for compiled programs, Python, and shell scripts
// Homepage: https://simonkagstrom.github.io/kcov/

import (
	"fmt"
	
	"os/exec"
)

func installKcov() {
	// Método 1: Descargar y extraer .tar.gz
	kcov_tar_url := "https://github.com/SimonKagstrom/kcov/archive/refs/tags/v43.tar.gz"
	kcov_cmd_tar := exec.Command("curl", "-L", kcov_tar_url, "-o", "package.tar.gz")
	err := kcov_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kcov_zip_url := "https://github.com/SimonKagstrom/kcov/archive/refs/tags/v43.zip"
	kcov_cmd_zip := exec.Command("curl", "-L", kcov_zip_url, "-o", "package.zip")
	err = kcov_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kcov_bin_url := "https://github.com/SimonKagstrom/kcov/archive/refs/tags/v43.bin"
	kcov_cmd_bin := exec.Command("curl", "-L", kcov_bin_url, "-o", "binary.bin")
	err = kcov_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kcov_src_url := "https://github.com/SimonKagstrom/kcov/archive/refs/tags/v43.src.tar.gz"
	kcov_cmd_src := exec.Command("curl", "-L", kcov_src_url, "-o", "source.tar.gz")
	err = kcov_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kcov_cmd_direct := exec.Command("./binary")
	err = kcov_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: dwarfutils")
	exec.Command("latte", "install", "dwarfutils").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
}
