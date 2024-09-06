package main

// Mold - Modern Linker
// Homepage: https://github.com/rui314/mold

import (
	"fmt"
	
	"os/exec"
)

func installMold() {
	// Método 1: Descargar y extraer .tar.gz
	mold_tar_url := "https://github.com/rui314/mold/archive/refs/tags/v2.33.0.tar.gz"
	mold_cmd_tar := exec.Command("curl", "-L", mold_tar_url, "-o", "package.tar.gz")
	err := mold_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mold_zip_url := "https://github.com/rui314/mold/archive/refs/tags/v2.33.0.zip"
	mold_cmd_zip := exec.Command("curl", "-L", mold_zip_url, "-o", "package.zip")
	err = mold_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mold_bin_url := "https://github.com/rui314/mold/archive/refs/tags/v2.33.0.bin"
	mold_cmd_bin := exec.Command("curl", "-L", mold_bin_url, "-o", "binary.bin")
	err = mold_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mold_src_url := "https://github.com/rui314/mold/archive/refs/tags/v2.33.0.src.tar.gz"
	mold_cmd_src := exec.Command("curl", "-L", mold_src_url, "-o", "source.tar.gz")
	err = mold_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mold_cmd_direct := exec.Command("./binary")
	err = mold_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: blake3")
exec.Command("latte", "install", "blake3")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: mimalloc")
exec.Command("latte", "install", "mimalloc")
}
