package main

// Poac - Package manager and build system for C++
// Homepage: https://github.com/poac-dev/poac

import (
	"fmt"
	
	"os/exec"
)

func installPoac() {
	// Método 1: Descargar y extraer .tar.gz
	poac_tar_url := "https://github.com/poac-dev/poac/archive/refs/tags/0.10.1.tar.gz"
	poac_cmd_tar := exec.Command("curl", "-L", poac_tar_url, "-o", "package.tar.gz")
	err := poac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	poac_zip_url := "https://github.com/poac-dev/poac/archive/refs/tags/0.10.1.zip"
	poac_cmd_zip := exec.Command("curl", "-L", poac_zip_url, "-o", "package.zip")
	err = poac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	poac_bin_url := "https://github.com/poac-dev/poac/archive/refs/tags/0.10.1.bin"
	poac_cmd_bin := exec.Command("curl", "-L", poac_bin_url, "-o", "binary.bin")
	err = poac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	poac_src_url := "https://github.com/poac-dev/poac/archive/refs/tags/0.10.1.src.tar.gz"
	poac_cmd_src := exec.Command("curl", "-L", poac_src_url, "-o", "source.tar.gz")
	err = poac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	poac_cmd_direct := exec.Command("./binary")
	err = poac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nlohmann-json")
exec.Command("latte", "install", "nlohmann-json")
	fmt.Println("Instalando dependencia: toml11")
exec.Command("latte", "install", "toml11")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: libgit2")
exec.Command("latte", "install", "libgit2")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
