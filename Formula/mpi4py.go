package main

// Mpi4py - Python bindings for MPI
// Homepage: https://mpi4py.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installMpi4py() {
	// Método 1: Descargar y extraer .tar.gz
	mpi4py_tar_url := "https://github.com/mpi4py/mpi4py/releases/download/4.0.0/mpi4py-4.0.0.tar.gz"
	mpi4py_cmd_tar := exec.Command("curl", "-L", mpi4py_tar_url, "-o", "package.tar.gz")
	err := mpi4py_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpi4py_zip_url := "https://github.com/mpi4py/mpi4py/releases/download/4.0.0/mpi4py-4.0.0.zip"
	mpi4py_cmd_zip := exec.Command("curl", "-L", mpi4py_zip_url, "-o", "package.zip")
	err = mpi4py_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpi4py_bin_url := "https://github.com/mpi4py/mpi4py/releases/download/4.0.0/mpi4py-4.0.0.bin"
	mpi4py_cmd_bin := exec.Command("curl", "-L", mpi4py_bin_url, "-o", "binary.bin")
	err = mpi4py_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpi4py_src_url := "https://github.com/mpi4py/mpi4py/releases/download/4.0.0/mpi4py-4.0.0.src.tar.gz"
	mpi4py_cmd_src := exec.Command("curl", "-L", mpi4py_src_url, "-o", "source.tar.gz")
	err = mpi4py_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpi4py_cmd_direct := exec.Command("./binary")
	err = mpi4py_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
