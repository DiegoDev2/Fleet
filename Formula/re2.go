package main

// Re2 - Alternative to backtracking PCRE-style regular expression engines
// Homepage: https://github.com/google/re2

import (
	"fmt"
	
	"os/exec"
)

func installRe2() {
	// Método 1: Descargar y extraer .tar.gz
	re2_tar_url := "https://github.com/google/re2/archive/refs/tags/2024-07-02.tar.gz"
	re2_cmd_tar := exec.Command("curl", "-L", re2_tar_url, "-o", "package.tar.gz")
	err := re2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	re2_zip_url := "https://github.com/google/re2/archive/refs/tags/2024-07-02.zip"
	re2_cmd_zip := exec.Command("curl", "-L", re2_zip_url, "-o", "package.zip")
	err = re2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	re2_bin_url := "https://github.com/google/re2/archive/refs/tags/2024-07-02.bin"
	re2_cmd_bin := exec.Command("curl", "-L", re2_bin_url, "-o", "binary.bin")
	err = re2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	re2_src_url := "https://github.com/google/re2/archive/refs/tags/2024-07-02.src.tar.gz"
	re2_cmd_src := exec.Command("curl", "-L", re2_src_url, "-o", "source.tar.gz")
	err = re2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	re2_cmd_direct := exec.Command("./binary")
	err = re2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
}
