package main

// ParallelHashmap - Family of header-only, fast, memory-friendly C++ hashmap and btree containers
// Homepage: https://greg7mdp.github.io/parallel-hashmap/

import (
	"fmt"
	
	"os/exec"
)

func installParallelHashmap() {
	// Método 1: Descargar y extraer .tar.gz
	parallelhashmap_tar_url := "https://github.com/greg7mdp/parallel-hashmap/archive/refs/tags/v1.3.12.tar.gz"
	parallelhashmap_cmd_tar := exec.Command("curl", "-L", parallelhashmap_tar_url, "-o", "package.tar.gz")
	err := parallelhashmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parallelhashmap_zip_url := "https://github.com/greg7mdp/parallel-hashmap/archive/refs/tags/v1.3.12.zip"
	parallelhashmap_cmd_zip := exec.Command("curl", "-L", parallelhashmap_zip_url, "-o", "package.zip")
	err = parallelhashmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parallelhashmap_bin_url := "https://github.com/greg7mdp/parallel-hashmap/archive/refs/tags/v1.3.12.bin"
	parallelhashmap_cmd_bin := exec.Command("curl", "-L", parallelhashmap_bin_url, "-o", "binary.bin")
	err = parallelhashmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parallelhashmap_src_url := "https://github.com/greg7mdp/parallel-hashmap/archive/refs/tags/v1.3.12.src.tar.gz"
	parallelhashmap_cmd_src := exec.Command("curl", "-L", parallelhashmap_src_url, "-o", "source.tar.gz")
	err = parallelhashmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parallelhashmap_cmd_direct := exec.Command("./binary")
	err = parallelhashmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
