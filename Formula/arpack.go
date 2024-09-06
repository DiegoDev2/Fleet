package main

// Arpack - Routines to solve large scale eigenvalue problems
// Homepage: https://github.com/opencollab/arpack-ng

import (
	"fmt"
	
	"os/exec"
)

func installArpack() {
	// Método 1: Descargar y extraer .tar.gz
	arpack_tar_url := "https://github.com/opencollab/arpack-ng/archive/refs/tags/3.9.1.tar.gz"
	arpack_cmd_tar := exec.Command("curl", "-L", arpack_tar_url, "-o", "package.tar.gz")
	err := arpack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arpack_zip_url := "https://github.com/opencollab/arpack-ng/archive/refs/tags/3.9.1.zip"
	arpack_cmd_zip := exec.Command("curl", "-L", arpack_zip_url, "-o", "package.zip")
	err = arpack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arpack_bin_url := "https://github.com/opencollab/arpack-ng/archive/refs/tags/3.9.1.bin"
	arpack_cmd_bin := exec.Command("curl", "-L", arpack_bin_url, "-o", "binary.bin")
	err = arpack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arpack_src_url := "https://github.com/opencollab/arpack-ng/archive/refs/tags/3.9.1.src.tar.gz"
	arpack_cmd_src := exec.Command("curl", "-L", arpack_src_url, "-o", "source.tar.gz")
	err = arpack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arpack_cmd_direct := exec.Command("./binary")
	err = arpack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
