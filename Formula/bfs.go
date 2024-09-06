package main

// Bfs - Breadth-first version of find
// Homepage: https://tavianator.com/projects/bfs.html

import (
	"fmt"
	
	"os/exec"
)

func installBfs() {
	// Método 1: Descargar y extraer .tar.gz
	bfs_tar_url := "https://github.com/tavianator/bfs/archive/refs/tags/4.0.1.tar.gz"
	bfs_cmd_tar := exec.Command("curl", "-L", bfs_tar_url, "-o", "package.tar.gz")
	err := bfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bfs_zip_url := "https://github.com/tavianator/bfs/archive/refs/tags/4.0.1.zip"
	bfs_cmd_zip := exec.Command("curl", "-L", bfs_zip_url, "-o", "package.zip")
	err = bfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bfs_bin_url := "https://github.com/tavianator/bfs/archive/refs/tags/4.0.1.bin"
	bfs_cmd_bin := exec.Command("curl", "-L", bfs_bin_url, "-o", "binary.bin")
	err = bfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bfs_src_url := "https://github.com/tavianator/bfs/archive/refs/tags/4.0.1.src.tar.gz"
	bfs_cmd_src := exec.Command("curl", "-L", bfs_src_url, "-o", "source.tar.gz")
	err = bfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bfs_cmd_direct := exec.Command("./binary")
	err = bfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: oniguruma")
	exec.Command("latte", "install", "oniguruma").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: acl")
	exec.Command("latte", "install", "acl").Run()
	fmt.Println("Instalando dependencia: libcap")
	exec.Command("latte", "install", "libcap").Run()
	fmt.Println("Instalando dependencia: liburing")
	exec.Command("latte", "install", "liburing").Run()
}
