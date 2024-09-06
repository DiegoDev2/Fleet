package main

// Assimp - Portable library for importing many well-known 3D model formats
// Homepage: https://www.assimp.org/

import (
	"fmt"
	
	"os/exec"
)

func installAssimp() {
	// Método 1: Descargar y extraer .tar.gz
	assimp_tar_url := "https://github.com/assimp/assimp/archive/refs/tags/v5.4.3.tar.gz"
	assimp_cmd_tar := exec.Command("curl", "-L", assimp_tar_url, "-o", "package.tar.gz")
	err := assimp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	assimp_zip_url := "https://github.com/assimp/assimp/archive/refs/tags/v5.4.3.zip"
	assimp_cmd_zip := exec.Command("curl", "-L", assimp_zip_url, "-o", "package.zip")
	err = assimp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	assimp_bin_url := "https://github.com/assimp/assimp/archive/refs/tags/v5.4.3.bin"
	assimp_cmd_bin := exec.Command("curl", "-L", assimp_bin_url, "-o", "binary.bin")
	err = assimp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	assimp_src_url := "https://github.com/assimp/assimp/archive/refs/tags/v5.4.3.src.tar.gz"
	assimp_cmd_src := exec.Command("curl", "-L", assimp_src_url, "-o", "source.tar.gz")
	err = assimp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	assimp_cmd_direct := exec.Command("./binary")
	err = assimp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
