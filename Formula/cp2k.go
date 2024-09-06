package main

// Cp2k - Quantum chemistry and solid state physics software package
// Homepage: https://www.cp2k.org/

import (
	"fmt"
	
	"os/exec"
)

func installCp2k() {
	// Método 1: Descargar y extraer .tar.gz
	cp2k_tar_url := "https://github.com/cp2k/cp2k/releases/download/v2024.2/cp2k-2024.2.tar.bz2"
	cp2k_cmd_tar := exec.Command("curl", "-L", cp2k_tar_url, "-o", "package.tar.gz")
	err := cp2k_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cp2k_zip_url := "https://github.com/cp2k/cp2k/releases/download/v2024.2/cp2k-2024.2.tar.bz2"
	cp2k_cmd_zip := exec.Command("curl", "-L", cp2k_zip_url, "-o", "package.zip")
	err = cp2k_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cp2k_bin_url := "https://github.com/cp2k/cp2k/releases/download/v2024.2/cp2k-2024.2.tar.bz2"
	cp2k_cmd_bin := exec.Command("curl", "-L", cp2k_bin_url, "-o", "binary.bin")
	err = cp2k_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cp2k_src_url := "https://github.com/cp2k/cp2k/releases/download/v2024.2/cp2k-2024.2.tar.bz2"
	cp2k_cmd_src := exec.Command("curl", "-L", cp2k_src_url, "-o", "source.tar.gz")
	err = cp2k_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cp2k_cmd_direct := exec.Command("./binary")
	err = cp2k_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: libxc")
	exec.Command("latte", "install", "libxc").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: scalapack")
	exec.Command("latte", "install", "scalapack").Run()
}
