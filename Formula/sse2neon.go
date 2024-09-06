package main

// Sse2neon - Translator from Intel SSE intrinsics to Arm/Aarch64 NEON implementation
// Homepage: https://github.com/DLTcollab/sse2neon

import (
	"fmt"
	
	"os/exec"
)

func installSse2neon() {
	// Método 1: Descargar y extraer .tar.gz
	sse2neon_tar_url := "https://github.com/DLTcollab/sse2neon/archive/refs/tags/v1.7.0.tar.gz"
	sse2neon_cmd_tar := exec.Command("curl", "-L", sse2neon_tar_url, "-o", "package.tar.gz")
	err := sse2neon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sse2neon_zip_url := "https://github.com/DLTcollab/sse2neon/archive/refs/tags/v1.7.0.zip"
	sse2neon_cmd_zip := exec.Command("curl", "-L", sse2neon_zip_url, "-o", "package.zip")
	err = sse2neon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sse2neon_bin_url := "https://github.com/DLTcollab/sse2neon/archive/refs/tags/v1.7.0.bin"
	sse2neon_cmd_bin := exec.Command("curl", "-L", sse2neon_bin_url, "-o", "binary.bin")
	err = sse2neon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sse2neon_src_url := "https://github.com/DLTcollab/sse2neon/archive/refs/tags/v1.7.0.src.tar.gz"
	sse2neon_cmd_src := exec.Command("curl", "-L", sse2neon_src_url, "-o", "source.tar.gz")
	err = sse2neon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sse2neon_cmd_direct := exec.Command("./binary")
	err = sse2neon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
