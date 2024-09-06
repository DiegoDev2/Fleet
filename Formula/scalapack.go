package main

// Scalapack - High-performance linear algebra for distributed memory machines
// Homepage: https://netlib.org/scalapack/

import (
	"fmt"
	
	"os/exec"
)

func installScalapack() {
	// Método 1: Descargar y extraer .tar.gz
	scalapack_tar_url := "https://netlib.org/scalapack/scalapack-2.2.0.tgz"
	scalapack_cmd_tar := exec.Command("curl", "-L", scalapack_tar_url, "-o", "package.tar.gz")
	err := scalapack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scalapack_zip_url := "https://netlib.org/scalapack/scalapack-2.2.0.tgz"
	scalapack_cmd_zip := exec.Command("curl", "-L", scalapack_zip_url, "-o", "package.zip")
	err = scalapack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scalapack_bin_url := "https://netlib.org/scalapack/scalapack-2.2.0.tgz"
	scalapack_cmd_bin := exec.Command("curl", "-L", scalapack_bin_url, "-o", "binary.bin")
	err = scalapack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scalapack_src_url := "https://netlib.org/scalapack/scalapack-2.2.0.tgz"
	scalapack_cmd_src := exec.Command("curl", "-L", scalapack_src_url, "-o", "source.tar.gz")
	err = scalapack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scalapack_cmd_direct := exec.Command("./binary")
	err = scalapack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
