package main

// KertishDfs - Kertish FileSystem and Cluster Administration CLI
// Homepage: https://github.com/freakmaxi/kertish-dfs

import (
	"fmt"
	
	"os/exec"
)

func installKertishDfs() {
	// Método 1: Descargar y extraer .tar.gz
	kertishdfs_tar_url := "https://github.com/freakmaxi/kertish-dfs/archive/refs/tags/v22.2.0147.tar.gz"
	kertishdfs_cmd_tar := exec.Command("curl", "-L", kertishdfs_tar_url, "-o", "package.tar.gz")
	err := kertishdfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kertishdfs_zip_url := "https://github.com/freakmaxi/kertish-dfs/archive/refs/tags/v22.2.0147.zip"
	kertishdfs_cmd_zip := exec.Command("curl", "-L", kertishdfs_zip_url, "-o", "package.zip")
	err = kertishdfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kertishdfs_bin_url := "https://github.com/freakmaxi/kertish-dfs/archive/refs/tags/v22.2.0147.bin"
	kertishdfs_cmd_bin := exec.Command("curl", "-L", kertishdfs_bin_url, "-o", "binary.bin")
	err = kertishdfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kertishdfs_src_url := "https://github.com/freakmaxi/kertish-dfs/archive/refs/tags/v22.2.0147.src.tar.gz"
	kertishdfs_cmd_src := exec.Command("curl", "-L", kertishdfs_src_url, "-o", "source.tar.gz")
	err = kertishdfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kertishdfs_cmd_direct := exec.Command("./binary")
	err = kertishdfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
