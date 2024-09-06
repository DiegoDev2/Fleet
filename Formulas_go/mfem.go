package main

// Mfem - Free, lightweight, scalable C++ library for FEM
// Homepage: https://mfem.org/

import (
	"fmt"
	
	"os/exec"
)

func installMfem() {
	// Método 1: Descargar y extraer .tar.gz
	mfem_tar_url := "https://github.com/mfem/mfem/archive/refs/tags/v4.7.tar.gz"
	mfem_cmd_tar := exec.Command("curl", "-L", mfem_tar_url, "-o", "package.tar.gz")
	err := mfem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mfem_zip_url := "https://github.com/mfem/mfem/archive/refs/tags/v4.7.zip"
	mfem_cmd_zip := exec.Command("curl", "-L", mfem_zip_url, "-o", "package.zip")
	err = mfem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mfem_bin_url := "https://github.com/mfem/mfem/archive/refs/tags/v4.7.bin"
	mfem_cmd_bin := exec.Command("curl", "-L", mfem_bin_url, "-o", "binary.bin")
	err = mfem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mfem_src_url := "https://github.com/mfem/mfem/archive/refs/tags/v4.7.src.tar.gz"
	mfem_cmd_src := exec.Command("curl", "-L", mfem_src_url, "-o", "source.tar.gz")
	err = mfem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mfem_cmd_direct := exec.Command("./binary")
	err = mfem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: hypre")
exec.Command("latte", "install", "hypre")
	fmt.Println("Instalando dependencia: metis")
exec.Command("latte", "install", "metis")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: suite-sparse")
exec.Command("latte", "install", "suite-sparse")
}
