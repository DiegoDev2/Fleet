package main

// Adios2 - Next generation of ADIOS developed in the Exascale Computing Program
// Homepage: https://adios2.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installAdios2() {
	// Método 1: Descargar y extraer .tar.gz
	adios2_tar_url := "https://github.com/ornladios/ADIOS2/archive/refs/tags/v2.10.1.tar.gz"
	adios2_cmd_tar := exec.Command("curl", "-L", adios2_tar_url, "-o", "package.tar.gz")
	err := adios2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adios2_zip_url := "https://github.com/ornladios/ADIOS2/archive/refs/tags/v2.10.1.zip"
	adios2_cmd_zip := exec.Command("curl", "-L", adios2_zip_url, "-o", "package.zip")
	err = adios2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adios2_bin_url := "https://github.com/ornladios/ADIOS2/archive/refs/tags/v2.10.1.bin"
	adios2_cmd_bin := exec.Command("curl", "-L", adios2_bin_url, "-o", "binary.bin")
	err = adios2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adios2_src_url := "https://github.com/ornladios/ADIOS2/archive/refs/tags/v2.10.1.src.tar.gz"
	adios2_cmd_src := exec.Command("curl", "-L", adios2_src_url, "-o", "source.tar.gz")
	err = adios2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adios2_cmd_direct := exec.Command("./binary")
	err = adios2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: nlohmann-json")
exec.Command("latte", "install", "nlohmann-json")
	fmt.Println("Instalando dependencia: c-blosc")
exec.Command("latte", "install", "c-blosc")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: libfabric")
exec.Command("latte", "install", "libfabric")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
	fmt.Println("Instalando dependencia: mpi4py")
exec.Command("latte", "install", "mpi4py")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
	fmt.Println("Instalando dependencia: pugixml")
exec.Command("latte", "install", "pugixml")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: yaml-cpp")
exec.Command("latte", "install", "yaml-cpp")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
