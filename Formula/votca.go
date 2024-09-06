package main

// Votca - Versatile Object-oriented Toolkit for Coarse-graining Applications
// Homepage: https://www.votca.org/

import (
	"fmt"
	
	"os/exec"
)

func installVotca() {
	// Método 1: Descargar y extraer .tar.gz
	votca_tar_url := "https://github.com/votca/votca/archive/refs/tags/v2024.1.tar.gz"
	votca_cmd_tar := exec.Command("curl", "-L", votca_tar_url, "-o", "package.tar.gz")
	err := votca_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	votca_zip_url := "https://github.com/votca/votca/archive/refs/tags/v2024.1.zip"
	votca_cmd_zip := exec.Command("curl", "-L", votca_zip_url, "-o", "package.zip")
	err = votca_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	votca_bin_url := "https://github.com/votca/votca/archive/refs/tags/v2024.1.bin"
	votca_cmd_bin := exec.Command("curl", "-L", votca_bin_url, "-o", "binary.bin")
	err = votca_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	votca_src_url := "https://github.com/votca/votca/archive/refs/tags/v2024.1.src.tar.gz"
	votca_cmd_src := exec.Command("curl", "-L", votca_src_url, "-o", "source.tar.gz")
	err = votca_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	votca_cmd_direct := exec.Command("./binary")
	err = votca_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: libecpint")
	exec.Command("latte", "install", "libecpint").Run()
	fmt.Println("Instalando dependencia: libint")
	exec.Command("latte", "install", "libint").Run()
	fmt.Println("Instalando dependencia: libxc")
	exec.Command("latte", "install", "libxc").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: libaec")
	exec.Command("latte", "install", "libaec").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
}
