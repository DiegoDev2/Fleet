package main

// Sparse - Static C code analysis tool
// Homepage: https://sparse.wiki.kernel.org/

import (
	"fmt"
	
	"os/exec"
)

func installSparse() {
	// Método 1: Descargar y extraer .tar.gz
	sparse_tar_url := "https://mirrors.edge.kernel.org/pub/software/devel/sparse/dist/sparse-0.6.4.tar.xz"
	sparse_cmd_tar := exec.Command("curl", "-L", sparse_tar_url, "-o", "package.tar.gz")
	err := sparse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sparse_zip_url := "https://mirrors.edge.kernel.org/pub/software/devel/sparse/dist/sparse-0.6.4.tar.xz"
	sparse_cmd_zip := exec.Command("curl", "-L", sparse_zip_url, "-o", "package.zip")
	err = sparse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sparse_bin_url := "https://mirrors.edge.kernel.org/pub/software/devel/sparse/dist/sparse-0.6.4.tar.xz"
	sparse_cmd_bin := exec.Command("curl", "-L", sparse_bin_url, "-o", "binary.bin")
	err = sparse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sparse_src_url := "https://mirrors.edge.kernel.org/pub/software/devel/sparse/dist/sparse-0.6.4.tar.xz"
	sparse_cmd_src := exec.Command("curl", "-L", sparse_src_url, "-o", "source.tar.gz")
	err = sparse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sparse_cmd_direct := exec.Command("./binary")
	err = sparse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
