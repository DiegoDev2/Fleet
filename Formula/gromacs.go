package main

// Gromacs - Versatile package for molecular dynamics calculations
// Homepage: https://www.gromacs.org/

import (
	"fmt"
	
	"os/exec"
)

func installGromacs() {
	// Método 1: Descargar y extraer .tar.gz
	gromacs_tar_url := "https://ftp.gromacs.org/pub/gromacs/gromacs-2024.3.tar.gz"
	gromacs_cmd_tar := exec.Command("curl", "-L", gromacs_tar_url, "-o", "package.tar.gz")
	err := gromacs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gromacs_zip_url := "https://ftp.gromacs.org/pub/gromacs/gromacs-2024.3.zip"
	gromacs_cmd_zip := exec.Command("curl", "-L", gromacs_zip_url, "-o", "package.zip")
	err = gromacs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gromacs_bin_url := "https://ftp.gromacs.org/pub/gromacs/gromacs-2024.3.bin"
	gromacs_cmd_bin := exec.Command("curl", "-L", gromacs_bin_url, "-o", "binary.bin")
	err = gromacs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gromacs_src_url := "https://ftp.gromacs.org/pub/gromacs/gromacs-2024.3.src.tar.gz"
	gromacs_cmd_src := exec.Command("curl", "-L", gromacs_src_url, "-o", "source.tar.gz")
	err = gromacs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gromacs_cmd_direct := exec.Command("./binary")
	err = gromacs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
