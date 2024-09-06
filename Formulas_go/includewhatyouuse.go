package main

// IncludeWhatYouUse - Tool to analyze #includes in C and C++ source files
// Homepage: https://include-what-you-use.org/

import (
	"fmt"
	
	"os/exec"
)

func installIncludeWhatYouUse() {
	// Método 1: Descargar y extraer .tar.gz
	includewhatyouuse_tar_url := "https://include-what-you-use.org/downloads/include-what-you-use-0.22.src.tar.gz"
	includewhatyouuse_cmd_tar := exec.Command("curl", "-L", includewhatyouuse_tar_url, "-o", "package.tar.gz")
	err := includewhatyouuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	includewhatyouuse_zip_url := "https://include-what-you-use.org/downloads/include-what-you-use-0.22.src.zip"
	includewhatyouuse_cmd_zip := exec.Command("curl", "-L", includewhatyouuse_zip_url, "-o", "package.zip")
	err = includewhatyouuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	includewhatyouuse_bin_url := "https://include-what-you-use.org/downloads/include-what-you-use-0.22.src.bin"
	includewhatyouuse_cmd_bin := exec.Command("curl", "-L", includewhatyouuse_bin_url, "-o", "binary.bin")
	err = includewhatyouuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	includewhatyouuse_src_url := "https://include-what-you-use.org/downloads/include-what-you-use-0.22.src.src.tar.gz"
	includewhatyouuse_cmd_src := exec.Command("curl", "-L", includewhatyouuse_src_url, "-o", "source.tar.gz")
	err = includewhatyouuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	includewhatyouuse_cmd_direct := exec.Command("./binary")
	err = includewhatyouuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
