package main

// Spack - Package manager that builds multiple versions and configurations of software
// Homepage: https://spack.io

import (
	"fmt"
	
	"os/exec"
)

func installSpack() {
	// Método 1: Descargar y extraer .tar.gz
	spack_tar_url := "https://github.com/spack/spack/archive/refs/tags/v0.22.1.tar.gz"
	spack_cmd_tar := exec.Command("curl", "-L", spack_tar_url, "-o", "package.tar.gz")
	err := spack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spack_zip_url := "https://github.com/spack/spack/archive/refs/tags/v0.22.1.zip"
	spack_cmd_zip := exec.Command("curl", "-L", spack_zip_url, "-o", "package.zip")
	err = spack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spack_bin_url := "https://github.com/spack/spack/archive/refs/tags/v0.22.1.bin"
	spack_cmd_bin := exec.Command("curl", "-L", spack_bin_url, "-o", "binary.bin")
	err = spack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spack_src_url := "https://github.com/spack/spack/archive/refs/tags/v0.22.1.src.tar.gz"
	spack_cmd_src := exec.Command("curl", "-L", spack_src_url, "-o", "source.tar.gz")
	err = spack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spack_cmd_direct := exec.Command("./binary")
	err = spack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
