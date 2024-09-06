package main

// Micromamba - Fast Cross-Platform Package Manager
// Homepage: https://github.com/mamba-org/mamba

import (
	"fmt"
	
	"os/exec"
)

func installMicromamba() {
	// Método 1: Descargar y extraer .tar.gz
	micromamba_tar_url := "https://github.com/mamba-org/mamba/archive/refs/tags/micromamba-1.5.9.tar.gz"
	micromamba_cmd_tar := exec.Command("curl", "-L", micromamba_tar_url, "-o", "package.tar.gz")
	err := micromamba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	micromamba_zip_url := "https://github.com/mamba-org/mamba/archive/refs/tags/micromamba-1.5.9.zip"
	micromamba_cmd_zip := exec.Command("curl", "-L", micromamba_zip_url, "-o", "package.zip")
	err = micromamba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	micromamba_bin_url := "https://github.com/mamba-org/mamba/archive/refs/tags/micromamba-1.5.9.bin"
	micromamba_cmd_bin := exec.Command("curl", "-L", micromamba_bin_url, "-o", "binary.bin")
	err = micromamba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	micromamba_src_url := "https://github.com/mamba-org/mamba/archive/refs/tags/micromamba-1.5.9.src.tar.gz"
	micromamba_cmd_src := exec.Command("curl", "-L", micromamba_src_url, "-o", "source.tar.gz")
	err = micromamba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	micromamba_cmd_direct := exec.Command("./binary")
	err = micromamba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cli11")
	exec.Command("latte", "install", "cli11").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: nlohmann-json")
	exec.Command("latte", "install", "nlohmann-json").Run()
	fmt.Println("Instalando dependencia: spdlog")
	exec.Command("latte", "install", "spdlog").Run()
	fmt.Println("Instalando dependencia: tl-expected")
	exec.Command("latte", "install", "tl-expected").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: libsolv")
	exec.Command("latte", "install", "libsolv").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: reproc")
	exec.Command("latte", "install", "reproc").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: yaml-cpp")
	exec.Command("latte", "install", "yaml-cpp").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
