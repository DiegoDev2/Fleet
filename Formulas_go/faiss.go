package main

// Faiss - Efficient similarity search and clustering of dense vectors
// Homepage: https://github.com/facebookresearch/faiss

import (
	"fmt"
	
	"os/exec"
)

func installFaiss() {
	// Método 1: Descargar y extraer .tar.gz
	faiss_tar_url := "https://github.com/facebookresearch/faiss/archive/refs/tags/v1.8.0.tar.gz"
	faiss_cmd_tar := exec.Command("curl", "-L", faiss_tar_url, "-o", "package.tar.gz")
	err := faiss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faiss_zip_url := "https://github.com/facebookresearch/faiss/archive/refs/tags/v1.8.0.zip"
	faiss_cmd_zip := exec.Command("curl", "-L", faiss_zip_url, "-o", "package.zip")
	err = faiss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faiss_bin_url := "https://github.com/facebookresearch/faiss/archive/refs/tags/v1.8.0.bin"
	faiss_cmd_bin := exec.Command("curl", "-L", faiss_bin_url, "-o", "binary.bin")
	err = faiss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faiss_src_url := "https://github.com/facebookresearch/faiss/archive/refs/tags/v1.8.0.src.tar.gz"
	faiss_cmd_src := exec.Command("curl", "-L", faiss_src_url, "-o", "source.tar.gz")
	err = faiss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faiss_cmd_direct := exec.Command("./binary")
	err = faiss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
